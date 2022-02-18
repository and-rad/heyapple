////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2021-2022 The HeyApple Authors.
//
// Use of this source code is governed by the GNU Affero General
// Public License as published by the Free Software Foundation,
// either version 3 of the License, or any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.
//
////////////////////////////////////////////////////////////////////////

package l10n

import (
	"errors"
	"io/fs"
	"os"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestTranslate(t *testing.T) {
	testdata := []struct {
		key   interface{}
		lang  string
		debug bool

		out string
	}{
		{ //00//
			key:  nil,
			lang: "de",
			out:  "",
		},
		{ //01//
			key:  "",
			lang: "en",
			out:  "",
		},
		{ //02//
			key:  "yes",
			lang: "en",
			out:  "Yes!",
		},
		{ //03//
			key:  "yes",
			lang: "de",
			out:  "Ja!",
		},
		{ //04//
			key:  "yes",
			lang: "fr",
			out:  "yes",
		},
		{ //05//
			key:  errors.New("sql: wrong table"),
			lang: "en",
			out:  "Error",
		},
		{ //06//
			key:  errors.New("no"),
			lang: "en",
			out:  "Error",
		},
		{ //07//
			key:  errors.New("no"),
			lang: "de",
			out:  "Fehler",
		},
		{ //08//
			key:   errors.New("sql: wrong table"),
			lang:  "en",
			debug: true,
			out:   "sql: wrong table",
		},
		{ //09//
			key:   errors.New("no"),
			lang:  "en",
			debug: true,
			out:   "Nope!",
		},
		{ //10//
			key:   errors.New("no"),
			lang:  "de",
			debug: true,
			out:   "Nee!",
		},
		{ //11//
			key:  "doowop",
			lang: "de",
			out:  "doowop",
		},
		{ //12//
			key:  "doowop",
			lang: "fr",
			out:  "doowop",
		},
	}

	translator := &translator{data: map[string]map[string]string{
		"de": {"yes": "Ja!", "no": "Nee!", "err.err": "Fehler"},
		"en": {"yes": "Yes!", "no": "Nope!", "err.err": "Error"},
	}}

	for idx, data := range testdata {
		translator.debug = data.debug
		val := translator.Translate(data.key, data.lang)

		if val != data.out {
			t.Errorf("test case %d: value mismatch \nexpected: %s \nactual  : %s", idx, data.out, val)
		}
	}
}

func Test_translator_Get(t *testing.T) {
	testdata := []struct {
		langs map[string]map[string]string
		in    string
		out   map[string]string
	}{
		{ //00// nothing in, nothing out
			out: map[string]string{},
		},
		{ //01// language doesn't exist
			langs: map[string]map[string]string{
				"de": {"key": "Wert"},
			},
			in:  "zh",
			out: map[string]string{},
		},
		{ //02// language doesn't exist, fall back to English
			langs: map[string]map[string]string{
				"de": {"key": "Wert"},
				"en": {"key": "value"},
			},
			in:  "zh",
			out: map[string]string{"key": "value"},
		},
		{ //03// success
			langs: map[string]map[string]string{
				"de": {"key": "Wert"},
				"en": {"key": "value"},
			},
			in:  "de",
			out: map[string]string{"key": "Wert"},
		},
	}

	for idx, data := range testdata {
		out := (&translator{data: data.langs}).Get(data.in)

		if !reflect.DeepEqual(out, data.out) {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, out, data.out)
		}
	}
}

func Test_translator_Default(t *testing.T) {
	for idx, data := range []struct {
		env  string
		lang string
	}{
		{ //00//
			env:  "",
			lang: "en",
		},
		{ //00//
			env:  "en",
			lang: "en",
		},
		{ //00//
			env:  "de",
			lang: "de",
		},
	} {
		os.Setenv(envDefault, data.env)
		defer os.Unsetenv(envDefault)

		if lang := NewTranslator().Default(); lang != data.lang {
			t.Errorf("test case %d: language mismatch \nhave: %v\nwant: %v", idx, lang, data.lang)
		}
	}
}

func Test_loadTranslations_NoFile(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("expected panic did not occur")
		}
	}()

	loadTranslations(fstest.MapFS{"l10n/en.json": {Mode: fs.ModeDir}})
}

func Test_loadTranslations_NoJSON(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("expected panic did not occur")
		}
	}()

	fs := fstest.MapFS{"l10n/en.json": {Data: []byte(`{"err":}`)}}

	loadTranslations(fs)
}

func Test_loadTranslations(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("unexpected panic: %v", err)
		}
	}()

	for idx, data := range []struct {
		fs  fs.FS
		out translation
	}{
		{ //00// folder doesn't exist in file system
			fs:  fstest.MapFS{},
			out: translation{},
		},
		{ //01// empty file
			fs:  fstest.MapFS{"l10n/en.json": {Data: []byte(`{}`)}},
			out: translation{"en": {}},
		},
		{ //02// single file
			fs:  fstest.MapFS{"l10n/en.json": {Data: []byte(`{"err":"Error"}`)}},
			out: translation{"en": {"err": "Error"}},
		},
		{ //03// multiple files
			fs: fstest.MapFS{
				"l10n/en.json": {Data: []byte(`{"err":"Error","hi":"Hi!"}`)},
				"l10n/de.json": {Data: []byte(`{"err":"Fehler","hi":"Hallo!"}`)},
			},
			out: translation{
				"en": {"err": "Error", "hi": "Hi!"},
				"de": {"err": "Fehler", "hi": "Hallo!"},
			},
		},
	} {
		tr := loadTranslations(data.fs)

		if !reflect.DeepEqual(tr, data.out) {
			t.Errorf("test case %d: translator mismatch \nhave: %v \nwant: %v", idx, tr, data.out)
		}
	}
}
