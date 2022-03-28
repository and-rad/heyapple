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

func TestNewTranslator(t *testing.T) {
	for idx, data := range []struct {
		tr   translator
		lang string
		key  interface{}

		out string
	}{
		{ //00//
			tr: translator{data: translations{
				"de": {"yes": "Ja!", "no": "Nee!"},
			}},
			key:  nil,
			lang: "de",
			out:  "",
		},
		{ //01//
			tr: translator{data: translations{
				"de": {"yes": "Ja!", "no": "Nee!"},
			}},
			key:  "",
			lang: "en",
			out:  "",
		},
		{ //02//
			tr: translator{data: translations{
				"de": {"yes": "Ja!", "no": "Nee!"},
				"en": {"yes": "Yes!", "no": "Nope!"},
			}},
			key:  "yes",
			lang: "en",
			out:  "Yes!",
		},
		{ //03//
			tr: translator{data: translations{
				"de": {"yes": "Ja!", "no": "Nee!"},
				"en": {"yes": "Yes!", "no": "Nope!"},
			}},
			key:  "yes",
			lang: "de",
			out:  "Ja!",
		},
		{ //04// fall back to English
			tr: translator{data: translations{
				"en": {"yes": "Yes!", "no": "Nope!"},
			}},
			key:  "yes",
			lang: "fr",
			out:  "Yes!",
		},
		{ //05//
			tr: translator{data: translations{
				"en": {"err": translation{"err": "Error"}},
			}},
			key:  errors.New("sql: wrong table"),
			lang: "en",
			out:  "Error",
		},
		{ //06//
			tr: translator{data: translations{
				"en": {"err": translation{"err": "Error"}},
			}},
			key:  errors.New("no"),
			lang: "en",
			out:  "Error",
		},
		{ //07//
			tr: translator{data: translations{
				"de": {"err": translation{"err": "Fehler"}},
			}},
			key:  errors.New("no"),
			lang: "de",
			out:  "Fehler",
		},
		{ //08//
			tr: translator{debug: true, data: translations{
				"en": {"err": translation{"err": "Error"}},
			}},
			key:  errors.New("sql: wrong table"),
			lang: "en",
			out:  "sql: wrong table",
		},
		{ //09//
			tr: translator{debug: true, data: translations{
				"en": {"yes": "Yes!", "no": "Nope!"},
			}},
			key:  errors.New("no"),
			lang: "en",
			out:  "Nope!",
		},
		{ //10//
			tr: translator{debug: true, data: translations{
				"de": {"yes": "Ja!", "no": "Nee!"},
			}},
			key:  errors.New("no"),
			lang: "de",
			out:  "Nee!",
		},
		{ //11// parse list of preferences
			tr: translator{data: translations{
				"de": {"yes": "Ja!", "no": "Nee!"},
				"en": {"yes": "Yes!", "no": "Nope!"},
			}},
			key:  "yes",
			lang: "de-DE,de;q=0.8,en;q=0.2",
			out:  "Ja!",
		},
		{ //12// nested keys
			tr: translator{data: translations{
				"de": {"yes": translation{"maybe": "Vielleicht", "definitely": "Auf jeden!"}},
			}},
			lang: "de",
			key:  "yes.maybe",
			out:  "Vielleicht",
		},
		{ //13// fall back to English if lang exists but key isn't found...
			tr: translator{data: translations{
				"de": {"yes": "Ja!"},
				"en": {"yes": "Yes!", "no": "Nope!"},
			}},
			key:  "no",
			lang: "de",
			out:  "Nope!",
		},
		{ //14// ...but don't do it in debug mode
			tr: translator{debug: true, data: translations{
				"de": {"yes": "Ja!"},
				"en": {"yes": "Yes!", "no": "Nope!"},
			}},
			key:  "no",
			lang: "de",
			out:  "no",
		},
	} {
		val := data.tr.Translate(data.key, data.lang)

		if val != data.out {
			t.Errorf("test case %d: value mismatch \nexpected: %s \nactual  : %s", idx, data.out, val)
		}
	}
}

func Test_translator_Get(t *testing.T) {
	for idx, data := range []struct {
		langs translations
		in    string
		out   map[string]interface{}
	}{
		{ //00// nothing in, nothing out
			out: translation{},
		},
		{ //01// language doesn't exist
			langs: translations{
				"de": {"key": "Wert"},
			},
			in:  "zh",
			out: translation{},
		},
		{ //02// language doesn't exist, fall back to English
			langs: translations{
				"de": {"key": "Wert"},
				"en": {"key": "value"},
			},
			in:  "zh",
			out: translation{"key": "value"},
		},
		{ //03// success
			langs: translations{
				"de": {"key": "Wert"},
				"en": {"key": "value"},
			},
			in:  "de",
			out: translation{"key": "Wert"},
		},
		{ //04// success with preferences
			langs: translations{
				"de": {"key": "Wert"},
				"en": {"key": "value"},
			},
			in:  "de-DE,de;q=0.9,en;q=0.4",
			out: translation{"key": "Wert"},
		},
	} {
		out := (&translator{data: data.langs}).Get(data.in)

		if !reflect.DeepEqual(out, data.out) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, out, data.out)
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
		{ //01//
			env:  "en",
			lang: "en",
		},
		{ //02//
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

func Test_translator_match(t *testing.T) {
	for idx, data := range []struct {
		langs translations
		in    string
		out   string
	}{
		{ //00// no languages defined
			out: "",
		},
		{ //01// empty input -> default language
			langs: translations{
				"de": translation{"key": "Wert"},
				"en": translation{"key": "value"},
			},
			out: "en",
		},
		{ //02// exact match
			langs: translations{
				"de": translation{"key": "Wert"},
				"en": translation{"key": "value"},
			},
			in:  "de",
			out: "de",
		},
		{ //03// match base language
			langs: translations{
				"de": translation{"key": "Wert"},
				"en": translation{"key": "value"},
			},
			in:  "de-CH",
			out: "de",
		},
		{ //04// match list of preferences
			langs: translations{
				"de": translation{"key": "Wert"},
				"en": translation{"key": "value"},
			},
			in:  "es-419,es;q=0.9,en;q=0.5",
			out: "en",
		},
	} {
		if out := (&translator{data: data.langs}).match(data.in); out != data.out {
			t.Errorf("test case %d: language mismatch \nhave: %v\nwant: %v", idx, out, data.out)
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
		out translations
	}{
		{ //00// folder doesn't exist in file system
			fs:  fstest.MapFS{},
			out: translations{},
		},
		{ //01// empty file
			fs:  fstest.MapFS{"l10n/en.json": {Data: []byte(`{}`)}},
			out: translations{"en": {}},
		},
		{ //02// single file
			fs:  fstest.MapFS{"l10n/en.json": {Data: []byte(`{"err":"Error"}`)}},
			out: translations{"en": {"err": "Error"}},
		},
		{ //03// multiple files
			fs: fstest.MapFS{
				"l10n/en.json": {Data: []byte(`{"err":"Error","hi":"Hi!"}`)},
				"l10n/de.json": {Data: []byte(`{"err":"Fehler","hi":"Hallo!"}`)},
			},
			out: translations{
				"en": {"err": "Error", "hi": "Hi!"},
				"de": {"err": "Fehler", "hi": "Hallo!"},
			},
		},
		{ //04// nested data
			fs: fstest.MapFS{
				"l10n/en.json": {Data: []byte(`{"err":{"bad":"Bad","very":"Very Bad"}}`)},
			},
			out: translations{
				"en": {"err": translation{"bad": "Bad", "very": "Very Bad"}},
			},
		},
	} {
		tr := loadTranslations(data.fs)

		if !reflect.DeepEqual(tr, data.out) {
			t.Errorf("test case %d: translator mismatch \nhave: %#v\nwant: %#v", idx, tr, data.out)
		}
	}
}

func Test_mergeTranslations(t *testing.T) {
	for idx, data := range []struct {
		base translations
		more translations
		out  translations
	}{
		{ //00//
			out: nil,
		},
		{ //01//
			base: translations{"en": {"key": "value"}},
			out:  translations{"en": {"key": "value"}},
		},
		{ //02//
			base: translations{
				"en": {"key1": "value 1", "key2": "value 2"},
				"de": {"key1": "Wert 1", "key3": "Wert 3"},
			},
			out: translations{
				"en": {"key1": "value 1", "key2": "value 2"},
				"de": {"key1": "Wert 1", "key3": "Wert 3"},
			},
		},
		{ //03//
			base: translations{
				"en": {"key1": "value 1", "key2": "value 2"},
				"de": {"key1": "Wert 1", "key3": "Wert 3"},
			},
			more: translations{
				"en": {"key1": "new value 1", "key3": "value 3"},
			},
			out: translations{
				"en": {"key1": "new value 1", "key2": "value 2", "key3": "value 3"},
				"de": {"key1": "Wert 1", "key3": "Wert 3"},
			},
		},
		{ //04//
			base: translations{
				"en": {"key1": "value 1", "key2": "value 2"},
				"de": {"key1": "Wert 1", "key3": "Wert 3"},
			},
			more: translations{
				"en": {"key1": "new value 1", "key3": "value 3"},
				"es": {"key1": "hola"},
				"de": {"key2": "Wert 2"},
			},
			out: translations{
				"en": {"key1": "new value 1", "key2": "value 2", "key3": "value 3"},
				"de": {"key1": "Wert 1", "key2": "Wert 2", "key3": "Wert 3"},
				"es": {"key1": "hola"},
			},
		},
		{ //05// don't merge empty strings
			base: translations{
				"en": {"key1": "value 1", "key2": "value 2"},
				"de": {"key1": "Wert 1", "key3": "Wert 3"},
			},
			more: translations{
				"en": {"key1": "", "key2": "new value 2"},
			},
			out: translations{
				"en": {"key1": "value 1", "key2": "new value 2"},
				"de": {"key1": "Wert 1", "key3": "Wert 3"},
			},
		},
		{ //06// merge nested translations
			base: translations{
				"en": {"key": translation{"key1": "value 1", "key2": "value 2"}},
			},
			more: translations{
				"en": {"key": translation{"key1": "new value 1"}},
			},
			out: translations{
				"en": {"key": translation{"key1": "new value 1", "key2": "value 2"}},
			},
		},
	} {
		out := mergeTranslations(data.base, data.more)

		if !reflect.DeepEqual(out, data.out) {
			t.Errorf("test case %d: translator mismatch \nhave: %v \nwant: %v", idx, out, data.out)
		}
	}
}
