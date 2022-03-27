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

package mock

import "fmt"

type Translator struct {
	Map         map[string]interface{}
	Lang        string
	DefaultLang string
}

func NewTranslator() *Translator {
	return &Translator{
		Map:         make(map[string]interface{}),
		Lang:        "en",
		DefaultLang: "en",
	}
}

func (t *Translator) Translate(input interface{}, lang string) string {
	if s, ok := input.(string); ok {
		if tr, ok := t.Map[s]; ok {
			return tr.(string)
		}
		return s
	} else if e, ok := input.(error); ok {
		if tr, ok := t.Map[e.Error()]; ok {
			return tr.(string)
		}
		return e.Error()
	}

	return fmt.Sprintf("%v", input)
}

func (t *Translator) Default() string {
	if t.DefaultLang == "" {
		return "en"
	}
	return t.DefaultLang
}

func (t *Translator) Get(lang string) map[string]interface{} {
	if lang == t.Lang {
		return t.Map
	}
	return map[string]interface{}{}
}
