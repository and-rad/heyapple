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

package core_test

import (
	"heyapple/pkg/core"
	"reflect"
	"testing"
)

func TestNewRecipe(t *testing.T) {
	for idx, data := range []struct {
		id  int
		rec core.Recipe
	}{
		{ //00//
			id:  1,
			rec: core.Recipe{ID: 1, Size: 1, Items: []core.Ingredient{}},
		},
		{ //01//
			id:  9000,
			rec: core.Recipe{ID: 9000, Size: 1, Items: []core.Ingredient{}},
		},
	} {
		if rec := core.NewRecipe(data.id); !reflect.DeepEqual(rec, data.rec) {
			t.Errorf("test case %d: recipe mismatch \nhave: %v\nwant: %v", idx, rec, data.rec)
		}
	}
}
