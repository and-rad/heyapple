////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2021-2024 The HeyApple Authors.
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

package app

import (
	"regexp"
)

// Validator provides validation for different types of user input.
type Validator struct {
	patternEmail    string
	patternPassword string
	patternLanguage string
}

// NewValidator returns a new Validator.
func NewValidator() Validator {
	return Validator{
		patternEmail:    `^[^@]+@[^@]+$`,
		patternPassword: `^\S{10}.*$`,
		patternLanguage: `^[a-z]+(-|_)*[a-zA-Z0-9]+$`,
	}
}

// MatchEmail returns true when s is a valid e-mail address.
func (v Validator) MatchEmail(s string) bool {
	ok, _ := regexp.MatchString(v.patternEmail, s)
	return ok
}

// MatchPass returns true when s is a valid password.
func (v Validator) MatchPass(s string) bool {
	ok, _ := regexp.MatchString(v.patternPassword, s)
	return ok
}

// MatchLang returns true when s is a valid IETF BCP 47 language tag.
func (v Validator) MatchLang(s string) bool {
	ok, _ := regexp.MatchString(v.patternLanguage, s)
	return ok
}
