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
	"time"
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
