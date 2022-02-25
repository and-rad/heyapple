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
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

const (
	testStorageDir = "/tmp/heyappletest/memory/config"
)

func TestNewDBWithBackup(t *testing.T) {
	os.Setenv(envStorageInterval, "200ms")
	defer os.Unsetenv(envStorageInterval)

	for idx, data := range []struct {
		dir  string
		file string
	}{
		{ //00// no write permission
			dir:  "/opt/tmp/heyapple",
			file: "",
		},
		{ //01// job scheduler successful
			dir:  testStorageDir,
			file: backup0,
		},
	} {
		os.Setenv(envStorageDir, data.dir)
		defer os.Unsetenv(envStorageDir)
		defer os.RemoveAll(data.dir)

		db := NewDBWithBackup(mock.NewLog())
		time.Sleep(time.Millisecond * 500)
		db.Close()

		file, _ := ioutil.ReadFile(filepath.Join(data.dir, storageFile+".json"))
		if contents := string(file); contents != data.file {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, contents, data.file)
		}
	}
}

func TestDB_Execute(t *testing.T) {
	for idx, data := range []struct {
		db  *DB
		cmd *app.CreateFood

		err error
	}{
		{ //00//
			db:  NewDB(mock.NewLog()),
			cmd: &app.CreateFood{},
		},
	} {
		err := data.db.Execute(data.cmd)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if f, _ := data.db.Food(data.cmd.ID); err == nil && f.ID != data.cmd.ID {
			t.Errorf("test case %d: id mismatch \nhave: %v\nwant: %v", idx, f.ID, data.cmd.ID)
		}
	}
}

func TestDB_Fetch(t *testing.T) {
	for idx, data := range []struct {
		db  *DB
		cmd *app.GetFood

		err error
	}{
		{ //00// empty database
			db:  NewDB(mock.NewLog()),
			cmd: &app.GetFood{Item: core.Food{ID: 1}},
			err: app.ErrNotFound,
		},
		{ //01// food item doesn't exist
			db:  &DB{food: map[int]core.Food{2: {ID: 2}}},
			cmd: &app.GetFood{Item: core.Food{ID: 1}},
			err: app.ErrNotFound,
		},
		{ //02// success
			db:  &DB{food: map[int]core.Food{2: {ID: 2}}},
			cmd: &app.GetFood{Item: core.Food{ID: 2}},
		},
	} {
		err := data.db.Fetch(data.cmd)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if f, err := data.db.Food(data.cmd.Item.ID); err == nil && data.cmd.Item != f {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.cmd.Item, f)
		}
	}
}

func TestDB_NewFood(t *testing.T) {
	for idx, data := range []struct {
		db *DB

		food core.Food
		err  error
	}{
		{ //00// empty database
			db:   NewDB(mock.NewLog()),
			food: core.Food{ID: 1},
		},
		{ //01// increment id
			db:   &DB{food: map[int]core.Food{2: {ID: 2}, 3: {ID: 3}}, foodID: 3},
			food: core.Food{ID: 4},
		},
	} {
		id, err := data.db.NewFood()

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if f, _ := data.db.Food(id); f != data.food {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, f, data.food)
		}
	}
}

func TestDB_Foods(t *testing.T) {
	for idx, data := range []struct {
		db *DB

		foods []core.Food
		err   error
	}{
		{ //00// empty database
			db:    NewDB(mock.NewLog()),
			foods: []core.Food{},
		},
		{ //01// all items returned
			db:    &DB{food: map[int]core.Food{1: mock.Food1, 2: mock.Food2}},
			foods: []core.Food{mock.Food1, mock.Food2},
		},
	} {
		foods, err := data.db.Foods()

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(foods, data.foods) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, foods, data.foods)
		}
	}
}

func TestDB_Food(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int

		food core.Food
		err  error
	}{
		{ //00// empty database
			db:  NewDB(mock.NewLog()),
			err: app.ErrNotFound,
		},
		{ //01// item doesn't exist
			db:  &DB{food: map[int]core.Food{1: mock.Food1}},
			id:  2,
			err: app.ErrNotFound,
		},
		{ //02// success
			db:   &DB{food: map[int]core.Food{1: mock.Food1, 2: mock.Food2}},
			id:   2,
			food: mock.Food2,
		},
	} {
		food, err := data.db.Food(data.id)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if food != data.food {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, food, data.food)
		}
	}
}

func TestDB_SetFood(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		food core.Food

		err error
	}{
		{ //00// empty database
			db:  NewDB(mock.NewLog()),
			err: app.ErrNotFound,
		},
		{ //01// food item doesn't exist
			db:   &DB{food: map[int]core.Food{1: mock.Food1}},
			food: mock.Food2,
			err:  app.ErrNotFound,
		},
		{ //02// success
			db:   &DB{food: map[int]core.Food{1: mock.Food1, 2: mock.Food2}},
			food: core.Food{ID: 2, KCal: 200, Fat: 30, Sugar: 45.6},
		},
	} {
		err := data.db.SetFood(data.food)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if f, err := data.db.Food(data.food.ID); err == nil && f != data.food {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, f, data.food)
		}
	}
}

func TestDB_UserByName(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		name string

		user app.User
		err  error
	}{
		{ //00// empty database
			db:  NewDB(mock.NewLog()),
			err: app.ErrNotFound,
		},
		{ //01// user doesn't exist
			db: &DB{
				users:  map[int]app.User{1: {ID: 1, Email: "a@a.a"}},
				emails: map[string]int{"a@a.a": 1},
			},
			name: "b@b.b",
			err:  app.ErrNotFound,
		},
		{ //02// success
			db: &DB{
				users: map[int]app.User{
					1: {ID: 1, Email: "a@a.a"},
					2: {ID: 2, Email: "b@b.b"},
				},
				emails: map[string]int{
					"a@a.a": 1,
					"b@b.b": 2,
				},
			},
			name: "b@b.b",
			user: app.User{ID: 2, Email: "b@b.b"},
		},
	} {
		user, err := data.db.UserByName(data.name)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if user != data.user {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, user, data.user)
		}
	}
}

func TestDB_NewUser(t *testing.T) {
	for idx, data := range []struct {
		db    *DB
		name  string
		hash  string
		token string

		user app.User
		err  error
	}{
		{ //00// empty database
			db:   NewDB(mock.NewLog()),
			name: "a@a.a",
			hash: "djwrifkgh",
			user: app.User{ID: 1, Email: "a@a.a", Pass: "djwrifkgh", Lang: getConfig().defaultLang},
		},
		{ //01// username already exists
			db: &DB{
				emails: map[string]int{"a@a.a": 1},
				users:  map[int]app.User{1: {ID: 1, Email: "a@a.a", Pass: "qpwoeirutz"}},
				userID: 1,
			},
			name: "a@a.a",
			hash: "djwrifkgh",
			user: app.User{ID: 1, Email: "a@a.a", Pass: "qpwoeirutz"},
			err:  app.ErrExists,
		},
		{ //02// success
			db: &DB{
				emails: map[string]int{"a@a.a": 1},
				users:  map[int]app.User{1: {ID: 1, Email: "a@a.a", Pass: "qpwoeirutz"}},
				tokens: map[string]app.Token{},
				userID: 1,
			},
			name:  "b@b.b",
			hash:  "djwrifkgh",
			token: "aabbccdd",
			user:  app.User{ID: 2, Email: "b@b.b", Pass: "djwrifkgh", Lang: getConfig().defaultLang},
		},
	} {
		id, err := data.db.NewUser(data.name, data.hash, data.token)

		if err == nil && id != data.db.userID {
			t.Errorf("test case %d: id mismatch \nhave: %v\nwant: %v", idx, id, data.db.userID)
		}

		if err == nil && id != data.db.tokens[data.token].ID {
			t.Errorf("test case %d: token mismatch \nhave: %v\nwant: %v", idx, data.db.tokens[data.token].ID, id)
		}

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if u, _ := data.db.UserByName(data.name); u != data.user {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, u, data.user)
		}
	}
}

func TestDB_SetUser(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		user app.User

		err error
	}{
		{ //00// empty database
			db:  NewDB(mock.NewLog()),
			err: app.ErrNotFound,
		},
		{ //01// user doesn't exist
			db:   &DB{users: map[int]app.User{2: {ID: 2}}},
			user: mock.User1,
			err:  app.ErrNotFound,
		},
		{ //02// success
			db: &DB{
				users:  map[int]app.User{1: mock.User1},
				emails: map[string]int{mock.User1.Email: 1},
			},
			user: app.User{ID: 1, Email: "b@b.b", Pass: "kdjfhghr", Perm: app.PermLogin},
		},
	} {
		err := data.db.SetUser(data.user)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if u, err := data.db.UserByID(data.user.ID); err == nil && u != data.user {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, u, data.user)
		}

		if err == nil && data.db.emails[data.user.Email] != data.user.ID {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.emails, data.user.Email)
		}
	}
}

func TestDB_UserByID(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int

		user app.User
		err  error
	}{
		{ //00// empty database
			db:  NewDB(mock.NewLog()),
			err: app.ErrNotFound,
		},
		{ //01// user doesn't exist
			db: &DB{
				users: map[int]app.User{1: {ID: 1, Email: "a@a.a"}},
			},
			id:  2,
			err: app.ErrNotFound,
		},
		{ //02// success
			db: &DB{
				users: map[int]app.User{
					1: {ID: 1, Email: "a@a.a"},
					2: {ID: 2, Email: "b@b.b"},
				},
			},
			id:   2,
			user: app.User{ID: 2, Email: "b@b.b"},
		},
	} {
		user, err := data.db.UserByID(data.id)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if user != data.user {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, user, data.user)
		}
	}
}

func TestDB_Token(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		hash string

		token app.Token
		err   error
	}{
		{ //00// empty database
			db:  NewDB(mock.NewLog()),
			err: app.ErrNotFound,
		},
		{ //01// token not found
			db:   &DB{tokens: map[string]app.Token{"abcd": {ID: 1}}},
			hash: "bbbb",
			err:  app.ErrNotFound,
		},
		{ //02// success
			db:    &DB{tokens: map[string]app.Token{"abcd": {ID: 1, Data: "hi"}}},
			hash:  "abcd",
			token: app.Token{ID: 1, Data: "hi"},
		},
	} {
		token, err := data.db.Token(data.hash)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if token != data.token {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, token, data.token)
		}
	}
}

func TestDB_DeleteToken(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		hash string

		tokens map[string]app.Token
		err    error
	}{
		{ //00// empty database, no error
			db:     NewDB(mock.NewLog()),
			tokens: make(map[string]app.Token),
		},
		{ //01// token deleted
			db: &DB{tokens: map[string]app.Token{
				"abcd": {ID: 1, Data: "hi"},
				"efef": {ID: 2, Data: 9000},
			}},
			hash: "abcd",
			tokens: map[string]app.Token{
				"efef": {ID: 2, Data: 9000},
			},
		},
	} {
		err := data.db.DeleteToken(data.hash)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.db.tokens, data.tokens) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.tokens, data.tokens)
		}
	}
}

func TestDB_NewRecipe(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		name string

		rec  core.Recipe
		meta core.RecipeMeta
		err  error
	}{
		{ //00// empty database, nameless recipe
			db:   NewDB(mock.NewLog()),
			rec:  core.Recipe{ID: 1, Size: 1},
			meta: core.RecipeMeta{ID: 1},
		},
		{ //01// increment id
			db: &DB{
				recipes: map[int]core.Recipe{2: {ID: 2}, 3: {ID: 3}},
				recMeta: map[int]core.RecipeMeta{2: {ID: 2, Name: "Food"}, 3: {ID: 3, Name: "Drink"}},
				recID:   3,
			},
			name: "Pie",
			rec:  core.Recipe{ID: 4, Size: 1},
			meta: core.RecipeMeta{ID: 4, Name: "Pie"},
		},
		{ //02// duplicate names allowed
			db: &DB{
				recipes: map[int]core.Recipe{1: {ID: 1}},
				recMeta: map[int]core.RecipeMeta{1: {ID: 1, Name: "Apple Pie"}},
				recID:   1,
			},
			name: "Apple Pie",
			rec:  core.Recipe{ID: 2, Size: 1},
			meta: core.RecipeMeta{ID: 2, Name: "Apple Pie"},
		},
	} {
		id, err := data.db.NewRecipe(data.name)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if r, _ := data.db.Recipe(id); !reflect.DeepEqual(r, data.rec) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, r, data.rec)
		}

		if m, _ := data.db.RecipeMeta(id); m != data.meta {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, m, data.meta)
		}
	}
}

func TestDB_SetRecipe(t *testing.T) {
	for idx, data := range []struct {
		db  *DB
		rec core.Recipe

		meta core.RecipeMeta
		err  error
	}{
		{ //00// empty database
			db:  NewDB(mock.NewLog()),
			err: app.ErrNotFound,
		},
		{ //01// recipe doesn't exist
			db:  &DB{recipes: map[int]core.Recipe{1: mock.Recipe1}},
			rec: mock.Recipe2,
			err: app.ErrNotFound,
		},
		{ //02// success
			db: &DB{
				food:    map[int]core.Food{1: mock.Food1},
				recipes: map[int]core.Recipe{2: mock.Recipe2},
				recMeta: map[int]core.RecipeMeta{2: mock.RecMeta2},
			},
			rec: core.Recipe{ID: 2, Size: 5, Items: []core.Ingredient{{ID: 1, Amount: 600}}},
			meta: func() core.RecipeMeta {
				m := mock.RecMeta2
				m.KCal = mock.Food1.KCal * 6
				m.Fat = mock.Food1.Fat * 6
				m.Carbs = mock.Food1.Carbs * 6
				m.Protein = mock.Food1.Protein * 6
				return m
			}(),
		},
	} {
		err := data.db.SetRecipe(data.rec)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if r, err := data.db.Recipe(data.rec.ID); err == nil && !reflect.DeepEqual(r, data.rec) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, r, data.rec)
		}

		if m, _ := data.db.RecipeMeta(data.rec.ID); m != data.meta {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, m, data.meta)
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
			db:  NewDB(mock.NewLog()),
			err: app.ErrNotFound,
		},
		{ //01// recipe doesn't exist
			db:  &DB{recipes: map[int]core.Recipe{1: mock.Recipe1}},
			id:  2,
			err: app.ErrNotFound,
		},
		{ //02// success
			db:  &DB{recipes: map[int]core.Recipe{2: mock.Recipe2}},
			id:  2,
			rec: mock.Recipe2,
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

func TestDB_RecipeMeta(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int

		meta core.RecipeMeta
		err  error
	}{
		{ //00// empty database
			db:  NewDB(mock.NewLog()),
			err: app.ErrNotFound,
		},
		{ //01// recipe doesn't exist
			db:  &DB{recMeta: map[int]core.RecipeMeta{1: mock.RecMeta1}},
			id:  2,
			err: app.ErrNotFound,
		},
		{ //02// success
			db:   &DB{recMeta: map[int]core.RecipeMeta{1: mock.RecMeta1, 2: mock.RecMeta2}},
			id:   2,
			meta: mock.RecMeta2,
		},
	} {
		meta, err := data.db.RecipeMeta(data.id)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if meta != data.meta {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, meta, data.meta)
		}
	}
}

func TestDB_FoodExists(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int

		ok  bool
		err error
	}{
		{ //00// empty database
			db: NewDB(mock.NewLog()),
			ok: false,
		},
		{ //01// item doesn't exist
			db: &DB{food: map[int]core.Food{1: mock.Food1}},
			id: 2,
			ok: false,
		},
		{ //02// success
			db: &DB{food: map[int]core.Food{1: mock.Food1, 2: mock.Food2}},
			id: 2,
			ok: true,
		},
	} {
		ok, err := data.db.FoodExists(data.id)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if ok != data.ok {
			t.Errorf("test case %d: result mismatch \nhave: %v\nwant: %v", idx, ok, data.ok)
		}
	}
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
			db:    NewDB(mock.NewLog()),
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
			db:  NewDB(mock.NewLog()),
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
				recipes: map[int]core.Recipe{2: mock.Recipe2},
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
				recipes: map[int]core.Recipe{2: mock.Recipe2},
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
