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
	"encoding/json"
	"fmt"
	"heyapple/web"
	"io/fs"
	"strings"
)

type translation map[string]map[string]string

type translator struct {
	data  translation
	debug bool
}

// NewTranslator returns an implementation of the app.Translator interface.
func NewTranslator() *translator {
	return &translator{
		data:  loadTranslations(web.L10n),
		debug: getConfig().debugMode,
	}
}

// Translate implements the app.Translator interface.
func (t *translator) Translate(input interface{}, lang string) string {
	var key string

	switch it := input.(type) {
	case string:
		key = it
	case error:
		if t.debug {
			key = it.Error()
		} else {
			key = "err.err"
		}
	}

	if data, ok := t.data[lang]; ok {
		if val, ok := data[key]; ok {
			return val
		}
	}

	return key
}

// Get implements the app.Translator interface.
func (t *translator) Get(lang string) map[string]string {
	if data, ok := t.data[lang]; ok {
		return data
	}

	if data, ok := t.data["en"]; ok {
		return data
	}

	return map[string]string{}
}

// Default implements the app.Translator interface.
func (t *translator) Default() string {
	return getConfig().defaultLang
}

func loadTranslations(dir fs.FS) translation {
	out := translation{}

	if files, err := fs.ReadDir(dir, "l10n"); err == nil {
		for _, f := range files {
			data, err := fs.ReadFile(dir, "l10n/"+f.Name())
			if err != nil {
				panic(fmt.Sprintf("error reading language file: %s", err))
			}

			l10n := map[string]string{}
			if err = json.Unmarshal(data, &l10n); err != nil {
				panic(fmt.Sprintf("error parsing localization data: %s", err))
			}

			out[strings.TrimSuffix(f.Name(), ".json")] = l10n
		}
	}

	return out
}
