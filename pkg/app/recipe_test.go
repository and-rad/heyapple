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
	"heyapple/internal/mock"
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"reflect"
	"testing"
)

func TestCreateRecipe_Execute(t *testing.T) {
	for idx, data := range []struct {
		db   *mock.DB
		name string

		err error
	}{
		{ //00//
			db:  mock.NewDB().WithError(mock.ErrDOS),
			err: mock.ErrDOS,
		},
		{ //01//
			db:   mock.NewDB().WithID(42),
			name: "Apple Pie",
		},
	} {
		cmd := &app.CreateRecipe{Name: data.name}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if cmd.ID != data.db.ID {
			t.Errorf("test case %d: id mismatch \nhave: %v\nwant: %v", idx, cmd.ID, data.db.ID)
		}

		if cmd.Name != data.db.Name {
			t.Errorf("test case %d: name mismatch \nhave: %v\nwant: %v", idx, data.db.Name, cmd.Name)
		}
	}
}

func TestSaveRecipe_Execute(t *testing.T) {
	for idx, data := range []struct {
		data map[int]float32
		size int
		id   int
		db   *mock.DB

		rec core.Recipe
		err error
	}{
		{ //00// id missing or invalid
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// connection failed
			id:  1,
			db:  mock.NewDB().WithError(mock.ErrDOS),
			err: mock.ErrDOS,
		},
		{ //02// id not found
			id:  2,
			db:  mock.NewDB().WithRecipe(mock.Recipe1),
			rec: mock.Recipe1,
			err: app.ErrNotFound,
		},
		{ //03// ignore non-existent food
			id:   1,
			size: 12,
			data: map[int]float32{1: 200, 2: 100, 3: 450},
			db:   mock.NewDB().WithRecipe(mock.Recipe1).WithFoods(mock.Food1, mock.Food2),
			rec: core.Recipe{
				ID:    1,
				Size:  12,
				Items: []core.Ingredient{{ID: 1, Amount: 200}, {ID: 2, Amount: 100}},
			},
		},
	} {
		cmd := &app.SaveRecipe{ID: data.id, Items: data.data, Size: data.size}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v \nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.rec, data.db.RecipeItem) {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, data.db.RecipeItem, data.rec)
		}
	}
}
