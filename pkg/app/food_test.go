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

func TestCreateFood_Execute(t *testing.T) {
	for idx, data := range []struct {
		db  *mock.DB
		err error
	}{
		{ //00//
			db:  mock.NewDB().WithError(mock.ErrDOS),
			err: mock.ErrDOS,
		},
		{ //01//
			db: mock.NewDB().WithID(42),
		},
	} {
		cmd := &app.CreateFood{}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v \nwant: %v", idx, err, data.err)
		}

		if cmd.ID != data.db.ID {
			t.Errorf("test case %d: id mismatch \nhave: %v \nwant: %v", idx, cmd.ID, data.db.ID)
		}
	}
}

func TestSaveFood_Execute(t *testing.T) {
	for idx, data := range []struct {
		data map[string]float32
		id   uint32
		db   *mock.DB

		food core.Food
		err  error
	}{
		{ //00// connection failed
			db:  mock.NewDB().WithError(mock.ErrDOS),
			err: mock.ErrDOS,
		},
		{ //01// id not found
			id:   2,
			db:   mock.NewDB().WithFood(mock.Food1),
			food: mock.Food1,
			err:  mock.ErrNotFound,
		},
		{ //02// empty data, no changes
			id:   1,
			db:   mock.NewDB().WithFood(mock.Food1),
			food: mock.Food1,
		},
		{ //03// change some values
			id:   1,
			db:   mock.NewDB().WithFood(mock.Food1),
			data: map[string]float32{"kcal": 120, "carb": 33.3},
			food: func() core.Food { f := mock.Food1; f.KCal = 120; f.Carbs = 33.3; return f }(),
		},
		{ //04// change all values
			id: 1,
			db: mock.NewDB().WithFood(mock.Food1),
			data: map[string]float32{
				"kcal": 10, "fat": 11, "fatsat": 12, "fato3": 13, "fato6": 14, "carb": 15, "sug": 16, "prot": 17, "fib": 18,
				"pot": 19, "chl": 20, "sod": 21, "calc": 22, "phos": 23, "mag": 24, "iron": 25, "zinc": 26, "mang": 27, "cop": 28, "iod": 29, "chr": 30, "mol": 31, "sel": 32,
				"vita": 33, "vitb1": 34, "vitb2": 35, "vitb3": 36, "vitb5": 37, "vitb6": 38, "vitb7": 39, "vitb9": 40, "vitb12": 41, "vitc": 42, "vitd": 43, "vite": 44, "vitk": 45,
			},
			food: core.Food{
				ID: 1, KCal: 10, Fat: 11, FatSat: 12, FatO3: 13, FatO6: 14, Carbs: 15, Sugar: 16, Protein: 17, Fiber: 18,
				Potassium: 19, Chlorine: 20, Sodium: 21, Calcium: 22, Phosphorus: 23, Magnesium: 24, Iron: 25, Zinc: 26, Manganse: 27, Copper: 28, Iodine: 29, Chromium: 30, Molybdenum: 31, Selenium: 32,
				VitA: 33, VitB1: 34, VitB2: 35, VitB3: 36, VitB5: 37, VitB6: 38, VitB7: 39, VitB9: 40, VitB12: 41, VitC: 42, VitD: 43, VitE: 44, VitK: 45,
			},
		},
	} {
		cmd := &app.SaveFood{ID: data.id, Data: data.data}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v \nwant: %v", idx, err, data.err)
		}

		if data.food != data.db.FoodItem {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, data.db.FoodItem, data.food)
		}
	}
}

func TestGetFood_Fetch(t *testing.T) {
	for idx, data := range []struct {
		id uint32
		db *mock.DB

		food core.Food
		err  error
	}{
		{ //00// connection failed
			db:  mock.NewDB().WithError(mock.ErrDOS),
			err: mock.ErrDOS,
		},
		{ //01// empty db
			id:   1,
			db:   mock.NewDB(),
			err:  mock.ErrNotFound,
			food: core.Food{ID: 1},
		},
		{ //02// id not found
			id:   1,
			db:   mock.NewDB().WithFood(mock.Food2),
			err:  mock.ErrNotFound,
			food: core.Food{ID: 1},
		},
		{ //03// success
			id:   1,
			db:   mock.NewDB().WithFood(mock.Food1),
			food: mock.Food1,
		},
	} {
		qry := &app.GetFood{Item: core.Food{ID: data.id}}
		err := qry.Fetch(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v \nwant: %v", idx, err, data.err)
		}

		if qry.Item != data.food {
			t.Errorf("test case %d: id mismatch \nhave: %v \nwant: %v", idx, qry.Item, data.food)
		}
	}
}

func TestGetFoods_Fetch(t *testing.T) {
	for idx, data := range []struct {
		db    *mock.DB
		foods []core.Food
		err   error
	}{
		{ //00// connection failed
			db:  mock.NewDB().WithError(mock.ErrDOS),
			err: mock.ErrDOS,
		},
		{ //01// empty db
			db:    mock.NewDB(),
			foods: []core.Food{},
		},
		{ //02// success
			db:    mock.NewDB().WithFoods([]core.Food{mock.Food1, mock.Food2}),
			foods: []core.Food{mock.Food1, mock.Food2},
		},
	} {
		qry := &app.GetFoods{}
		err := qry.Fetch(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(qry.Items, data.foods) {
			t.Errorf("test case %d: data mismatch \nhave: %#v\nwant: %#v", idx, qry.Items, data.foods)
		}
	}
}