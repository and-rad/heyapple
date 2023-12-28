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

package mem

import (
	"reflect"
	"sync"
	"testing"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/core"
	"github.com/and-rad/heyapple/internal/mock"
)

func TestDB_NewRecipe(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		name string

		rec core.Recipe
		err error
	}{
		{ //00// empty database, nameless recipe
			db:  NewDB(),
			rec: core.NewRecipe(1),
		},
		{ //01// increment id
			db: &DB{
				recipes: map[int]core.Recipe{2: {ID: 2}, 3: {ID: 3}},
				recID:   3,
			},
			name: "Pie",
			rec:  func() core.Recipe { r := core.NewRecipe(4); r.Name = "Pie"; return r }(),
		},
		{ //02// duplicate names allowed
			db: &DB{
				recipes: map[int]core.Recipe{1: {ID: 1, Name: "Pie"}},
				recID:   1,
			},
			name: "Pie",
			rec:  func() core.Recipe { r := core.NewRecipe(2); r.Name = "Pie"; return r }(),
		},
	} {
		id, err := data.db.NewRecipe(data.name)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if r, _ := data.db.Recipe(id); !reflect.DeepEqual(r, data.rec) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, r, data.rec)
		}
	}
}

func TestDB_SetRecipe(t *testing.T) {
	for idx, data := range []struct {
		db  *DB
		rec core.Recipe

		err error
	}{
		{ //00// empty database
			db:  NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// recipe doesn't exist
			db:  &DB{recipes: map[int]core.Recipe{1: mock.Recipe1()}},
			rec: mock.Recipe2(),
			err: app.ErrNotFound,
		},
		{ //02// success
			db: &DB{
				food:    map[int]core.Food{1: mock.Food1},
				recipes: map[int]core.Recipe{2: mock.Recipe2()},
			},
			rec: core.Recipe{
				ID:      2,
				Name:    mock.Recipe2().Name,
				Size:    5,
				Items:   []core.Ingredient{{ID: 1, Amount: 600}},
				KCal:    mock.Food1.KCal * 6,
				Fat:     mock.Food1.Fat * 6,
				Carbs:   mock.Food1.Carbs * 6,
				Protein: mock.Food1.Protein * 6,
			},
		},
	} {
		err := data.db.SetRecipe(data.rec)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if r, err := data.db.Recipe(data.rec.ID); err == nil && !reflect.DeepEqual(r, data.rec) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, r, data.rec)
		}
	}
}

func TestDB_Recipe(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int

		rec core.Recipe
		err error
	}{
		{ //00// empty database
			db:  NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// recipe doesn't exist
			db:  &DB{recipes: map[int]core.Recipe{1: mock.Recipe1()}},
			id:  2,
			err: app.ErrNotFound,
		},
		{ //02// success
			db:  &DB{recipes: map[int]core.Recipe{2: mock.Recipe2()}},
			id:  2,
			rec: mock.Recipe2(),
		},
	} {
		rec, err := data.db.Recipe(data.id)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(rec, data.rec) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, rec, data.rec)
		}
	}
}

func TestDB_Recipe_Race(t *testing.T) {
	db := NewDB()
	db.recipes[1] = mock.Recipe1()

	rec1, _ := db.Recipe(1)
	rec2, _ := db.Recipe(1)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		rec1.Name = "Rec 1"
		rec1.Items[0].Amount = 12
		wg.Done()
	}()

	go func() {
		rec2.Name = "Rec 2"
		rec2.Items[0].Amount = 120
		wg.Done()
	}()

	wg.Wait()
}

func TestDB_RecipeAccess(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		user int
		rec  int

		perms int
		err   error
	}{
		{ //00// empty database
			db:    NewDB(),
			perms: app.PermNone,
		},
		{ //01// user doesn't exist
			db:    &DB{userRec: map[int]map[int]int{1: {5: app.PermRead}}},
			user:  2,
			rec:   5,
			perms: app.PermNone,
		},
		{ //02// success
			db:    &DB{userRec: map[int]map[int]int{1: {5: app.PermRead}}},
			user:  1,
			rec:   5,
			perms: app.PermRead,
		},
		{ //03// success for public recipe
			db:    &DB{userRec: map[int]map[int]int{0: {5: app.PermRead}}},
			user:  1,
			rec:   5,
			perms: app.PermRead,
		},
		{ //04// public recipe doesn't block custom access
			db: &DB{userRec: map[int]map[int]int{
				0: {2: app.PermRead},
				1: {5: app.PermCreate},
			}},
			user:  1,
			rec:   5,
			perms: app.PermCreate,
		},
		{ //05// permissions are combined
			db: &DB{userRec: map[int]map[int]int{
				0: {5: app.PermRead},
				1: {5: app.PermCreate},
			}},
			user:  1,
			rec:   5,
			perms: app.PermRead | app.PermCreate,
		},
	} {
		perms, err := data.db.RecipeAccess(data.user, data.rec)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if perms != data.perms {
			t.Errorf("test case %d: permission mismatch \nhave: %v\nwant: %v", idx, perms, data.perms)
		}
	}
}

func TestDB_SetRecipeAccess(t *testing.T) {
	for idx, data := range []struct {
		db    *DB
		user  int
		rec   int
		perms int

		err error
	}{
		{ //00// empty database
			db:  NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// recipe doesn't exist
			db: &DB{
				users:   map[int]app.User{1: mock.User1},
				userRec: map[int]map[int]int{},
				recUser: map[int]map[int]int{},
			},
			user: 1,
			rec:  2,
			err:  app.ErrNotFound,
		},
		{ //02// set new permission
			db: &DB{
				users:   map[int]app.User{1: mock.User1},
				recipes: map[int]core.Recipe{2: mock.Recipe2()},
				userRec: map[int]map[int]int{},
				recUser: map[int]map[int]int{},
			},
			user:  1,
			rec:   2,
			perms: app.PermRead,
		},
		{ //03// update existing permissions
			db: &DB{
				users:   map[int]app.User{1: mock.User1},
				recipes: map[int]core.Recipe{2: mock.Recipe2()},
				userRec: map[int]map[int]int{1: {2: app.PermNone}},
				recUser: map[int]map[int]int{2: {1: app.PermNone}},
			},
			user:  1,
			rec:   2,
			perms: app.PermRead | app.PermEdit,
		},
	} {
		err := data.db.SetRecipeAccess(data.user, data.rec, data.perms)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if perms, _ := data.db.RecipeAccess(data.user, data.rec); perms != data.perms {
			t.Errorf("test case %d: permission mismatch \nhave: %v\nwant: %v", idx, perms, data.perms)
		}
	}
}

func TestDB_Recipes(t *testing.T) {
	for idx, data := range []struct {
		db     *DB
		filter core.Filter
		uid    int

		recs []core.Recipe
		err  error
	}{
		{ //00// empty database
			db:   NewDB(),
			recs: []core.Recipe{},
		},
		{ //01// missing permissions
			db:   &DB{recipes: map[int]core.Recipe{1: mock.Recipe1(), 2: mock.Recipe2()}},
			uid:  1,
			recs: []core.Recipe{},
		},
		{ //02// denied by explicit permission
			db: &DB{
				recipes: map[int]core.Recipe{1: mock.Recipe1(), 2: mock.Recipe2()},
				userRec: map[int]map[int]int{1: {1: app.PermNone}},
			},
			uid:  1,
			recs: []core.Recipe{},
		},
		{ //03// no filter, all items with access returned
			db: &DB{
				recipes: map[int]core.Recipe{1: mock.Recipe1(), 2: mock.Recipe2()},
				userRec: map[int]map[int]int{1: {1: app.PermOwner}},
			},
			uid:  1,
			recs: []core.Recipe{mock.Recipe1()},
		},
		{ //04// all items returned
			db: &DB{
				recipes: map[int]core.Recipe{1: mock.Recipe1(), 2: mock.Recipe2()},
				userRec: map[int]map[int]int{1: {1: app.PermOwner, 2: app.PermRead}},
			},
			uid:  1,
			recs: []core.Recipe{mock.Recipe1(), mock.Recipe2()},
		},
		{ //05// filter by range
			db: &DB{
				recipes: map[int]core.Recipe{1: mock.Recipe1(), 2: mock.Recipe2()},
				userRec: map[int]map[int]int{1: {1: app.PermOwner, 2: app.PermRead}},
			},
			uid:    1,
			filter: core.Filter{"kcal": core.FloatRange{60, 180}},
			recs:   []core.Recipe{mock.Recipe1()},
		},
		{ //06// filter by exact value
			db: &DB{
				recipes: map[int]core.Recipe{1: mock.Recipe1(), 2: mock.Recipe2()},
				userRec: map[int]map[int]int{1: {1: app.PermOwner, 2: app.PermRead}},
			},
			uid:    1,
			filter: core.Filter{"size": mock.Recipe2().Size},
			recs:   []core.Recipe{mock.Recipe2()},
		},
		{ //07// include public recipes
			db: &DB{
				recipes: map[int]core.Recipe{1: mock.Recipe1(), 2: mock.Recipe2()},
				userRec: map[int]map[int]int{0: {1: app.PermRead}, 1: {2: app.PermOwner}},
			},
			uid:  1,
			recs: []core.Recipe{mock.Recipe1(), mock.Recipe2()},
		},
		{ //08// avoid duplicates
			db: &DB{
				recipes: map[int]core.Recipe{1: mock.Recipe1(), 2: mock.Recipe2()},
				userRec: map[int]map[int]int{0: {2: app.PermRead}, 1: {2: app.PermOwner}},
			},
			uid:  1,
			recs: []core.Recipe{mock.Recipe2()},
		},
	} {
		recs, err := data.db.Recipes(data.uid, data.filter)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(recs, data.recs) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, recs, data.recs)
		}
	}
}

func TestDB_Recipes_Race(t *testing.T) {
	db := NewDB()
	db.recipes[1] = mock.Recipe1()
	db.userRec[1] = map[int]int{1: app.PermRead}

	rec1, _ := db.Recipes(1, nil)
	rec2, _ := db.Recipes(1, nil)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		rec1[0].Name = "Rec 1"
		rec1[0].Items[0].Amount = 12
		wg.Done()
	}()

	go func() {
		rec2[0].Name = "Rec 2"
		rec2[0].Items[0].Amount = 120
		wg.Done()
	}()

	wg.Wait()
}
