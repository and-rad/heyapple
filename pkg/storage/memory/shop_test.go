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

package memory

import (
	"heyapple/internal/mock"
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"reflect"
	"testing"
	"time"
)

func TestDB_ShoppingList(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		id   int
		date []time.Time

		list []core.ShopItem
		err  error
	}{
		{ //00// diary doesn't exist
			id:  1,
			db:  NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// day doesn't exist
			id:   1,
			date: []time.Time{mock.Date1},
			db:   &DB{entries: entryMap{1: {}}},
			list: []core.ShopItem{},
		},
		{ //02// success with fuzzy date match
			id:   1,
			date: []time.Time{mock.Date1},
			db:   &DB{entries: entryMap{1: {mock.Day1: {mock.Entry1(), mock.Entry2()}}}},
			list: mock.List1(),
		},
		{ //03// success with multiple days
			id:   1,
			date: []time.Time{mock.Day1, mock.Day2},
			db: &DB{entries: entryMap{1: {
				mock.Day1: {mock.Entry1(), mock.Entry2()},
				mock.Day2: {mock.Entry3(), mock.Entry4()},
			}}},
			list: mock.List12(),
		},
		{ //04// no date provided, empty list
			id: 1,
			db: &DB{entries: entryMap{1: {
				mock.Day1: {mock.Entry1(), mock.Entry2()},
				mock.Day2: {mock.Entry3(), mock.Entry4()},
			}}},
			list: []core.ShopItem{},
		},
	} {
		list, err := data.db.ShoppingList(data.id, data.date...)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(list, data.list) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, list, data.list)
		}
	}
}