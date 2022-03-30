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
	"reflect"
	"testing"
	"time"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/core"
	"github.com/and-rad/heyapple/internal/mock"
)

func TestShoppingList_Fetch(t *testing.T) {
	for idx, data := range []struct {
		db    *mock.DB
		query *app.ShoppingList

		items []core.ShopItem
		err   error
	}{
		{ //00// no id provided
			db:    mock.NewDB(),
			query: &app.ShoppingList{},
			err:   app.ErrMissing,
		},
		{ //01// no date provided
			query: &app.ShoppingList{ID: 1},
			db:    mock.NewDB(),
			err:   app.ErrMissing,
		},
		{ //02// missing name
			query: &app.ShoppingList{ID: 12, Date: []time.Time{mock.Day1}},
			db:    mock.NewDB(),
			err:   app.ErrMissing,
		},
		{ //03// invalid shopping list name
			query: &app.ShoppingList{ID: 1, Name: "Dinner", Date: []time.Time{mock.Day1}},
			db:    mock.NewDB(),
			err:   app.ErrNotFound,
		},
		{ //04// connection failed
			db:    mock.NewDB().WithError(mock.ErrDOS),
			query: &app.ShoppingList{ID: 1, Name: "diary", Date: []time.Time{time.Now()}},
			err:   mock.ErrDOS,
		},
		{ //05// empty db
			query: &app.ShoppingList{ID: 1, Name: "diary", Date: []time.Time{time.Now()}},
			db:    mock.NewDB(),
			items: []core.ShopItem{},
		},
		{ //06// success
			query: &app.ShoppingList{ID: 1, Name: "diary", Date: []time.Time{mock.Day1}},
			db:    mock.NewDB().WithDate(mock.Day1).WithShoppingList(mock.List1()...),
			items: mock.List1(),
		},
		{ //07// return empty slice when DB didn't find anything
			query: &app.ShoppingList{ID: 12, Name: "diary", Date: []time.Time{time.Now()}},
			db:    mock.NewDB().WithError(app.ErrNotFound),
			items: []core.ShopItem{},
		},
	} {
		err := data.query.Fetch(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.query.Items, data.items) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.query.Items, data.items)
		}
	}
}

func TestSaveShoppingListDone_Execute(t *testing.T) {
	for idx, data := range []struct {
		db  *mock.DB
		cmd *app.SaveShoppingListDone

		items []core.ShopItem
		err   error
	}{
		{ //00// no id provided
			db:  mock.NewDB(),
			cmd: &app.SaveShoppingListDone{},
			err: app.ErrMissing,
		},
		{ //01// no data
			cmd: &app.SaveShoppingListDone{ID: 12},
			db:  mock.NewDB(),
			err: app.ErrMissing,
		},
		{ //02// missing name
			cmd: &app.SaveShoppingListDone{ID: 12, Items: map[int]bool{1: true}},
			db:  mock.NewDB(),
			err: app.ErrMissing,
		},
		{ //03// connection failed
			db:  mock.NewDB().WithError(nil, mock.ErrDOS).WithFoods(mock.Food1),
			cmd: &app.SaveShoppingListDone{ID: 1, Name: "diary", Items: map[int]bool{1: true}},
			err: mock.ErrDOS,
		},
		{ //04// ignore non-existent food
			db:  mock.NewDB(),
			cmd: &app.SaveShoppingListDone{ID: 1, Name: "diary", Items: map[int]bool{1: true}},
		},
		{ //05// success
			cmd:   &app.SaveShoppingListDone{ID: 1, Name: "diary", Items: map[int]bool{1: true}},
			db:    mock.NewDB().WithShoppingList(mock.List1()...).WithFoods(mock.Food1),
			items: func() []core.ShopItem { l := mock.List1(); l[0].Done = true; return l }(),
		},
		{ //06// invalid shopping list name
			db:  mock.NewDB().WithFoods(mock.Food1),
			cmd: &app.SaveShoppingListDone{ID: 1, Name: "Dinner", Items: map[int]bool{1: true}},
			err: app.ErrNotFound,
		},
	} {
		err := data.cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.db.ShopList, data.items) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.ShopList, data.items)
		}
	}
}
