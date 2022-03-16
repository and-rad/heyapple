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
	"fmt"
	"heyapple/internal/mock"
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"testing"
	"testing/fstest"
	"time"
)

const (
	testStorageDir = "/tmp/heyappletest/memory/config"
)

func TestDB_WithBackup(t *testing.T) {
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

		db := NewDB().WithBackup(mock.NewLog())
		time.Sleep(time.Millisecond * 500)
		db.Close()

		file, _ := ioutil.ReadFile(filepath.Join(data.dir, storageFile+".json"))
		if contents := string(file); contents != data.file {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, contents, data.file)
		}
	}
}

func TestDB_WithDefaults(t *testing.T) {
	for idx, data := range []struct {
		fs fs.FS
		db *DB
	}{
		{ //00// file not found
			fs: fstest.MapFS{},
			db: NewDB(),
		},
		{ //01// not a file
			fs: fstest.MapFS{
				"user.json": {Mode: fs.ModeDir},
			},
			db: NewDB(),
		},
		{ //02// invalid JSON
			fs: fstest.MapFS{
				"user.json": {Data: []byte(`{"email":}`)},
			},
			db: NewDB(),
		},
		{ //03// success
			fs: fstest.MapFS{
				"user.json": {Data: []byte(fmt.Sprintf(`[%s]`, mock.User1Json))},
			},
			db: func() *DB {
				db := NewDB()
				db.users = map[int]app.User{1: mock.User1}
				db.emails = map[string]int{mock.User1.Email: 1}
				db.userID = 1
				return db
			}(),
		},
		{ //04// not a file
			fs: fstest.MapFS{
				"user.json": {Data: []byte(`[]`)},
				"food.json": {Mode: fs.ModeDir},
			},
			db: NewDB(),
		},
		{ //05// invalid JSON
			fs: fstest.MapFS{
				"user.json": {Data: []byte(`[]`)},
				"food.json": {Data: []byte(`{"err":}`)},
			},
			db: NewDB(),
		},
		{ //06// success
			fs: fstest.MapFS{
				"user.json": {Data: []byte(`[]`)},
				"food.json": {Data: []byte(fmt.Sprintf(`[%s]`, mock.Food1Json))},
			},
			db: func() *DB {
				db := NewDB()
				db.food = map[int]core.Food{1: mock.Food1}
				db.foodID = 1
				return db
			}(),
		},
		{ //07// not a file
			fs: fstest.MapFS{
				"user.json":   {Data: []byte(`[]`)},
				"food.json":   {Data: []byte(`[]`)},
				"recipe.json": {Mode: fs.ModeDir},
			},
			db: NewDB(),
		},
		{ //08// invalid JSON
			fs: fstest.MapFS{
				"user.json":   {Data: []byte(`[]`)},
				"food.json":   {Data: []byte(`[]`)},
				"recipe.json": {Data: []byte(`{"err":}`)},
			},
			db: NewDB(),
		},
		{ //09// success
			fs: fstest.MapFS{
				"user.json":   {Data: []byte(`[]`)},
				"food.json":   {Data: []byte(`[]`)},
				"recipe.json": {Data: []byte(fmt.Sprintf(`[%s]`, mock.Recipe1Json))},
			},
			db: func() *DB {
				db := NewDB()
				db.recID = 1
				db.recipes = map[int]core.Recipe{1: mock.Recipe1()}
				db.userRec = map[int]map[int]int{0: {1: app.PermRead}}
				return db
			}(),
		},
	} {
		db := NewDB().WithDefaults(data.fs)

		if !reflect.DeepEqual(db, data.db) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, db, data.db)
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
			db:  NewDB(),
			cmd: &app.CreateFood{},
		},
	} {
		err := data.db.Execute(data.cmd)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if f, _ := data.db.Food(data.cmd.Food.ID); err == nil && f != data.cmd.Food {
			t.Errorf("test case %d: id mismatch \nhave: %v\nwant: %v", idx, f, data.cmd.Food)
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
			db:  NewDB(),
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
			db:   NewDB(),
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
		db     *DB
		filter core.Filter

		foods []core.Food
		err   error
	}{
		{ //00// empty database
			db:    NewDB(),
			foods: []core.Food{},
		},
		{ //01// all items returned
			db:    &DB{food: map[int]core.Food{1: mock.Food1, 2: mock.Food2}},
			foods: []core.Food{mock.Food1, mock.Food2},
		},
		{ //02// filter by range
			db:     &DB{food: map[int]core.Food{1: mock.Food1, 2: mock.Food2}},
			filter: core.Filter{"kcal": core.FloatRange{20, 60}},
			foods:  []core.Food{mock.Food1},
		},
		{ //03// filter by exact value
			db:     &DB{food: map[int]core.Food{1: mock.Food1, 2: mock.Food2}},
			filter: core.Filter{"prot": float32(1)},
			foods:  []core.Food{mock.Food2},
		},
	} {
		foods, err := data.db.Foods(data.filter)

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
			db:  NewDB(),
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
			db:  NewDB(),
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
			db:  NewDB(),
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
			db:   NewDB(),
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
			db:  NewDB(),
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
			db:  NewDB(),
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
			db:  NewDB(),
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
			db:     NewDB(),
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

func TestDB_FoodExists(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int

		ok  bool
		err error
	}{
		{ //00// empty database
			db: NewDB(),
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

func TestDB_NewToken(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		id   int
		hash string
		data interface{}

		token app.Token
		err   error
	}{
		{ //00//
			db:    NewDB(),
			token: app.Token{},
		},
		{ //00//
			db:    NewDB(),
			id:    12,
			hash:  "abcd",
			data:  true,
			token: app.Token{ID: 12, Data: true},
		},
	} {
		err := data.db.NewToken(data.id, data.hash, data.data)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if tok, _ := data.db.Token(data.hash); tok != data.token {
			t.Errorf("test case %d: token mismatch \nhave: %v\nwant: %v", idx, tok, data.token)
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

func TestDB_NewDiaryEntries(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int
		in []core.DiaryEntry

		out entryMap
		err error
	}{
		{ //00// create new entry
			id:  1,
			db:  NewDB(),
			in:  []core.DiaryEntry{mock.Entry1()},
			out: entryMap{1: {mock.Day1: {mock.Entry1()}}},
		},
		{ //01// make sure entries are sorted
			id:  1,
			db:  NewDB(),
			in:  []core.DiaryEntry{mock.Entry2(), mock.Entry1()},
			out: entryMap{1: {mock.Day1: {mock.Entry1(), mock.Entry2()}}},
		},
	} {
		err := data.db.NewDiaryEntries(data.id, data.in...)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.db.entries, data.out) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.entries, data.out)
		}
	}
}

func TestDB_DelDiaryEntries(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int
		in []core.DiaryEntry

		out entryMap
		err error
	}{
		{ //00// diary doesn't exist
			id:  1,
			db:  NewDB(),
			out: entryMap{},
			err: app.ErrNotFound,
		},
		{ //01// empty diary, nothing to do
			id:  1,
			db:  &DB{entries: entryMap{1: {}}},
			in:  []core.DiaryEntry{mock.Entry1()},
			out: entryMap{1: {}},
		},
		{ //02// success, simple case
			id:  1,
			db:  &DB{entries: entryMap{1: {mock.Day1: {mock.Entry1(), mock.Entry2()}}}},
			in:  []core.DiaryEntry{mock.Entry2()},
			out: entryMap{1: {mock.Day1: {mock.Entry1()}}},
		},
		{ //03// success, complex case
			id: 1,
			db: &DB{entries: entryMap{1: {
				mock.Day1: {mock.Entry1(), mock.Entry2()},
				mock.Day2: {mock.Entry3(), mock.Entry4()},
			}}},
			in: []core.DiaryEntry{
				mock.Entry3(),
				mock.Entry4(),
				mock.Entry2(),
			},
			out: entryMap{1: {
				mock.Day1: {mock.Entry1()},
				mock.Day2: {},
			}},
		},
	} {
		err := data.db.DelDiaryEntries(data.id, data.in...)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.db.entries, data.out) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.entries, data.out)
		}
	}
}

func TestDB_SetDiaryEntries(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int
		in []core.DiaryEntry

		out entryMap
		err error
	}{
		{ //00// diary doesn't exist
			id:  1,
			db:  NewDB(),
			out: entryMap{},
			err: app.ErrNotFound,
		},
		{ //01// empty diary, nothing to do
			id:  1,
			db:  &DB{entries: entryMap{1: {}}},
			in:  []core.DiaryEntry{mock.Entry1()},
			out: entryMap{1: {}},
		},
		{ //02// success, simple case
			id:  1,
			db:  &DB{entries: entryMap{1: {mock.Day1: {mock.Entry1(), mock.Entry2()}}}},
			in:  []core.DiaryEntry{mock.Entry1New()},
			out: entryMap{1: {mock.Day1: {mock.Entry1New(), mock.Entry2()}}},
		},
		{ //03// success, complex case
			id: 1,
			db: &DB{entries: entryMap{1: {
				mock.Day1: {mock.Entry1(), mock.Entry2()},
				mock.Day2: {mock.Entry3(), mock.Entry4()},
			}}},
			in: []core.DiaryEntry{
				mock.Entry3New(),
				mock.Entry4New(),
				mock.Entry2New(),
			},
			out: entryMap{1: {
				mock.Day1: {mock.Entry1(), mock.Entry2New()},
				mock.Day2: {mock.Entry3New(), mock.Entry4New()},
			}},
		},
	} {
		err := data.db.SetDiaryEntries(data.id, data.in...)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.db.entries, data.out) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.entries, data.out)
		}
	}
}

func TestDB_DiaryEntries(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		id   int
		date time.Time

		entries []core.DiaryEntry
		err     error
	}{
		{ //00// diary doesn't exist
			id:  1,
			db:  NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// day doesn't exist
			id:   1,
			date: mock.Date1,
			db:   &DB{entries: entryMap{1: {}}},
			err:  app.ErrNotFound,
		},
		{ //02// success with fuzzy date match
			id:      1,
			date:    mock.Date1,
			db:      &DB{entries: entryMap{1: {mock.Day1: {mock.Entry1(), mock.Entry2()}}}},
			entries: []core.DiaryEntry{mock.Entry1(), mock.Entry2()},
		},
	} {
		entries, err := data.db.DiaryEntries(data.id, data.date)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(entries, data.entries) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, entries, data.entries)
		}
	}
}

func TestDB_DiaryEntries_Race(t *testing.T) {
	db := NewDB()
	db.entries = entryMap{1: {mock.Day1: {mock.Entry1()}}}

	e1, _ := db.DiaryEntries(1, mock.Day1)
	e2, _ := db.DiaryEntries(1, mock.Day1)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		e1[0].Recipe = "Rec 1"
		wg.Done()
	}()

	go func() {
		e2[0].Recipe = "Rec 1"
		wg.Done()
	}()

	wg.Wait()
}

func TestDB_DiaryDays(t *testing.T) {
	for idx, data := range []struct {
		db    *DB
		id    int
		year  int
		month int
		day   int

		days []core.DiaryDay
		err  error
	}{
		{ //00// diary doesn't exist
			id:  1,
			db:  NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// get everything
			id: 1,
			db: &DB{days: dayMap{1: {
				2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
				2022: {3: {mock.Diary220301(), mock.Diary220302()}},
			}}},
			days: []core.DiaryDay{mock.Diary210101(), mock.Diary210102(), mock.Diary210201(), mock.Diary210202(), mock.Diary220301(), mock.Diary220302()},
		},
		{ //02// no entries for specific year
			id:   1,
			year: 2020,
			db: &DB{days: dayMap{1: {
				2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
				2022: {3: {mock.Diary220301(), mock.Diary220302()}},
			}}},
			days: []core.DiaryDay{},
		},
		{ //03// get specific year
			id:   1,
			year: 2021,
			db: &DB{days: dayMap{1: {
				2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
				2022: {3: {mock.Diary220301(), mock.Diary220302()}},
			}}},
			days: []core.DiaryDay{mock.Diary210101(), mock.Diary210102(), mock.Diary210201(), mock.Diary210202()},
		},
		{ //04// no entries for specific month
			id:    1,
			year:  2021,
			month: 4,
			db: &DB{days: dayMap{1: {
				2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
				2022: {3: {mock.Diary220301(), mock.Diary220302()}},
			}}},
			days: []core.DiaryDay{},
		},
		{ //05// get specific month
			id:    1,
			year:  2021,
			month: 2,
			db: &DB{days: dayMap{1: {
				2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
				2022: {3: {mock.Diary220301(), mock.Diary220302()}},
			}}},
			days: []core.DiaryDay{mock.Diary210201(), mock.Diary210202()},
		},
		{ //06// no entries for specific day
			id:    1,
			year:  2021,
			month: 2,
			day:   12,
			db: &DB{days: dayMap{1: {
				2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
				2022: {3: {mock.Diary220301(), mock.Diary220302()}},
			}}},
			days: []core.DiaryDay{},
		},
		{ //07// get specific day
			id:    1,
			year:  2021,
			month: 2,
			day:   1,
			db: &DB{days: dayMap{1: {
				2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
				2022: {3: {mock.Diary220301(), mock.Diary220302()}},
			}}},
			days: []core.DiaryDay{mock.Diary210201()},
		},
	} {
		days, err := data.db.DiaryDays(data.id, data.year, data.month, data.day)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(days, data.days) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, days, data.days)
		}
	}
}

func TestDB_DiaryDays_Race(t *testing.T) {
	db := NewDB()
	db.days = dayMap{1: {
		2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
		2022: {3: {mock.Diary220301(), mock.Diary220302()}},
	}}

	d1, _ := db.DiaryDays(1, 2021, 1, 1)
	d2, _ := db.DiaryDays(1, 2021, 1, 1)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		d1[0].KCal = 123
		wg.Done()
	}()

	go func() {
		d2[0].KCal = 456
		wg.Done()
	}()

	wg.Wait()
}
