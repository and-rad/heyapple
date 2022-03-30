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
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"testing/fstest"
	"time"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/core"
	"github.com/and-rad/heyapple/internal/mock"
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
				db.aisles = aisleMap{0: {1: 0}}
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
