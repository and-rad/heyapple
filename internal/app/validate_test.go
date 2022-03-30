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

package app_test

import (
	"testing"

	"github.com/and-rad/heyapple/internal/app"
)

func TestMatchEmail(t *testing.T) {
	testdata := []struct {
		input string
		valid bool
	}{
		{ //00//
			"", false,
		},
		{ //01//
			"atsignmissing", false,
		},
		{ //02//
			"@@@.com", false,
		},
		{ //03//
			"a@b.c", true,
		},
		{ //04//
			"user@localhost", true,
		},
		{ //05//
			"user@", false,
		},
		{ //06//
			"@localhost", false,
		},
	}

	for idx, data := range testdata {
		ok := app.NewValidator().MatchEmail(data.input)

		if ok != data.valid {
			t.Errorf("test case %d: validity mismatch \nexpected: %v \nactual  : %v", idx, data.valid, ok)
		}
	}
}

func TestMatchPassword(t *testing.T) {
	testdata := []struct {
		input string
		valid bool
	}{
		{ //00//
			"", false,
		},
		{ //01//
			"tooshort", false,
		},
		{ //02//
			"dont do white space", false,
		},
		{ //03//
			"longenoughbutkindabad", true,
		},
		{ //04//
			"§Correct.Horse!Battery$Staple#", true,
		},
	}

	for idx, data := range testdata {
		ok := app.NewValidator().MatchPass(data.input)

		if ok != data.valid {
			t.Errorf("test case %d: validity mismatch \nexpected: %v \nactual  : %v", idx, data.valid, ok)
		}
	}
}

func TestMatchLanguage(t *testing.T) {
	testdata := []struct {
		input string
		valid bool
	}{
		{ //00//
			"", false,
		},
		{ //01//
			"de", true,
		},
		{ //02//
			"en-us", true,
		},
		{ //03//
			"es-419", true,
		},
		{ //04//
			"es_ES", true,
		},
		{ //05//
			"<script>inject()</script>", false,
		},
		{ //06//
			"§$%&/().,=?", false,
		},
	}

	for idx, data := range testdata {
		ok := app.NewValidator().MatchLang(data.input)

		if ok != data.valid {
			t.Errorf("test case %d: validity mismatch \nhave: %v \nwant: %v", idx, ok, data.valid)
		}
	}
}
