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
	"heyapple/internal/web"
	"io/fs"
	"os"
	"strings"

	"golang.org/x/text/language"
)

type translation map[string]interface{}
type translations map[string]translation

type translator struct {
	data  translations
	debug bool
}

// NewTranslator returns an implementation of the app.Translator interface.
func NewTranslator() *translator {
	conf := getConfig()
	base := loadTranslations(web.L10n)
	user := loadTranslations(os.DirFS(conf.dataDir))

	return &translator{
		data:  mergeTranslations(base, user),
		debug: conf.debugMode,
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

	for _, l := range []string{lang, t.Default()} {
		if data, ok := t.data[t.match(l)]; ok {
			parts := strings.Split(key, ".")
			for _, p := range parts {
				if val, ok := data[p].(string); ok {
					return val
				}
				if val, ok := data[p].(translation); ok {
					data = val
				}
			}

			if t.debug {
				return key
			}
		}
	}

	return key
}

// Get implements the app.Translator interface.
func (t *translator) Get(lang string) map[string]interface{} {
	if data, ok := t.data[t.match(lang)]; ok {
		return data
	}
	return translation{}
}

// Default implements the app.Translator interface.
func (t *translator) Default() string {
	return getConfig().defaultLang
}

func (t *translator) match(lang string) string {
	if _, ok := t.data[lang]; ok {
		return lang
	}

	if tags, _, err := language.ParseAcceptLanguage(lang); err == nil {
		for _, tag := range tags {
			code := tag.String()
			if _, ok := t.data[code]; ok {
				return code
			}
			code = strings.Split(code, "-")[0]
			if _, ok := t.data[code]; ok {
				return code
			}
		}
	}

	fallback := t.Default()
	if _, ok := t.data[fallback]; ok {
		return fallback
	}

	return lang
}

func loadTranslations(dir fs.FS) translations {
	out := translations{}

	if files, err := fs.ReadDir(dir, "l10n"); err == nil {
		for _, f := range files {
			data, err := fs.ReadFile(dir, "l10n/"+f.Name())
			if err != nil {
				panic(fmt.Sprintf("error reading language file: %s", err))
			}

			l10n := translation{}
			if err = json.Unmarshal(data, &l10n); err != nil {
				panic(fmt.Sprintf("error parsing localization data: %s", err))
			}

			out[strings.TrimSuffix(f.Name(), ".json")] = parseMap(l10n)
		}
	}

	return out
}

func mergeTranslations(base translations, custom translations) translations {
	for lang, newTrn := range custom {
		if oldTrn, ok := base[lang]; ok {
			base[lang] = mergeTranslation(oldTrn, newTrn)
		} else {
			base[lang] = newTrn
		}
	}
	return base
}

func mergeTranslation(base translation, custom translation) translation {
	for newKey, newVal := range custom {
		oldVal, ok := base[newKey]
		if !ok {
			base[newKey] = newVal
			continue
		}
		if newVal == "" {
			continue
		}
		if newStr, ok := newVal.(string); ok {
			if _, ok := oldVal.(string); ok {
				base[newKey] = newStr
				continue
			}
		}

		if newTrn, ok := newVal.(translation); ok {
			if oldTrn, ok := oldVal.(translation); ok {
				base[newKey] = mergeTranslation(oldTrn, newTrn)
			}
		}
	}

	return base
}

func parseMap(tr translation) translation {
	result := translation{}
	for k, v := range tr {
		switch t := v.(type) {
		case map[string]interface{}:
			result[k] = parseMap(translation(t))
		default:
			result[k] = t
		}
	}
	return result
}
