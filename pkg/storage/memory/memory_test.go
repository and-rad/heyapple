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
		db   *DB
		name string
		hash string

		user app.User
		err  error
	}{
		{ //00// empty database
			db:   NewDB(mock.NewLog()),
			name: "a@a.a",
			hash: "djwrifkgh",
			user: app.User{ID: 1, Email: "a@a.a", Pass: "djwrifkgh"},
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
				userID: 1,
			},
			name: "b@b.b",
			hash: "djwrifkgh",
			user: app.User{ID: 2, Email: "b@b.b", Pass: "djwrifkgh"},
		},
	} {
		id, err := data.db.NewUser(data.name, data.hash)

		if err == nil && id != data.db.userID {
			t.Errorf("test case %d: id mismatch \nhave: %v\nwant: %v", idx, id, data.db.userID)
		}

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if u, _ := data.db.UserByName(data.name); u != data.user {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, u, data.user)
		}
	}
}
