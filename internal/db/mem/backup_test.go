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
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/core"
	"github.com/and-rad/heyapple/internal/mock"
)

var (
	backup0 = `{"users":{},"prefs":{},"tokens":{},"food":{},"recipes":{},"inst":{},"recaccess":{},"entries":{},"aisles":{},"prices":{},"done":{},"userid":0,"foodid":0,"recid":0}`
	backup1 = fmt.Sprintf(`{"users":{"1":%s},"prefs":{"1":%s},"tokens":{"abcd":{"id":1}},"food":{"1":%s,"2":%s},"recipes":{"1":%s},"inst":{"1":"%s"},"recaccess":{"1":[{"res":1,"perms":3840}]},"entries":{"1":[%s,%s,%s]},"aisles":{"0":{"1":1,"2":12}},"prices":{"1":{"2":[1.5,2.5]}},"done":{"1":{"1":true}},"userid":1,"foodid":2,"recid":1}`,
		mock.User1Json, mock.StoredPrefs1Json, mock.Food1Json, mock.Food2Json, mock.Recipe1Json, mock.Instructions1, mock.Entry1Json, mock.Entry2Json, mock.Entry3Json,
	)

	database1 = &DB{
		users:        map[int]app.User{mock.User1.ID: mock.User1},
		prefs:        map[int]app.StoredPrefs{mock.User1.ID: mock.StoredPrefs1},
		emails:       map[string]int{mock.User1.Email: mock.User1.ID},
		tokens:       map[string]app.Token{"abcd": {ID: mock.User1.ID}},
		food:         map[int]core.Food{mock.Food1.ID: mock.Food1, mock.Food2.ID: mock.Food2},
		recipes:      map[int]core.Recipe{mock.Recipe1().ID: mock.Recipe1()},
		instructions: map[int]string{mock.Recipe1().ID: mock.Instructions1},
		recUser:      map[int]map[int]int{1: {1: app.PermCreate | app.PermEdit | app.PermRead | app.PermDelete}},
		userRec:      map[int]map[int]int{1: {1: app.PermCreate | app.PermEdit | app.PermRead | app.PermDelete}},
		entries:      entryMap{mock.User1.ID: {mock.Date1.Truncate(time.Hour * 24): {mock.Entry1(), mock.Entry2()}, mock.Day2: {mock.Entry3()}}},
		days:         dayMap{mock.User1.ID: {2022: {3: {mock.Diary220312(), mock.Diary220313()}}}},
		aisles:       aisleMap{0: {mock.Food1.ID: 1, mock.Food2.ID: 12}},
		prices:       priceMap{mock.User1.ID: {mock.Food2.ID: [2]float32{1.5, 2.5}}},
		done:         doneMap{mock.User1.ID: {mock.Food1.ID: true}},
		userID:       1,
		foodID:       2,
		recID:        1,
	}
)

func Test_backup_Run(t *testing.T) {
	for idx, data := range []struct {
		db  *DB
		dir string

		file string
	}{
		{ //00// no directory permission
			dir: func() string {
				if runtime.GOOS == "windows" {
					return "C:/Program Files/Temp/HeyApple"
				} else {
					return "/opt/tmp/heyapple"
				}
			}(),
			db: NewDB(),
		},
		{ //01// empty database
			dir:  testStorageDir,
			db:   NewDB(),
			file: backup0,
		},
		{ //02// save filled database
			dir:  testStorageDir,
			db:   database1,
			file: backup1,
		},
	} {
		os.Setenv(envStorageDir, data.dir)
		os.Setenv(envBackupDir, data.dir)
		defer os.Unsetenv(envStorageDir)
		defer os.Unsetenv(envBackupDir)
		defer os.RemoveAll(data.dir)

		(&backup{db: data.db, log: mock.NewLog()}).Run()

		file, _ := os.ReadFile(filepath.Join(data.dir, storageFile+".json"))
		if contents := string(file); contents != data.file {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, contents, data.file)
		}

		now := time.Now().Format("2006-01-02")
		file, _ = os.ReadFile(filepath.Join(data.dir, backupFile+now+".json"))
		if contents := string(file); contents != data.file {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, contents, data.file)
		}
	}
}

func Test_backup_load(t *testing.T) {
	os.Setenv(envStorageDir, testStorageDir)
	defer os.Unsetenv(envStorageDir)

	for idx, data := range []struct {
		file string
		db   *DB
	}{
		{ //00// empty file
			db: NewDB(),
		},
		{ //01// invalid JSON
			file: `{"food":`,
			db:   NewDB(),
		},
		{ //02// success
			file: backup1,
			db:   database1,
		},
	} {
		dir := os.Getenv(envStorageDir)
		os.MkdirAll(dir, 0755)
		defer os.RemoveAll(dir)

		file := filepath.Join(dir, storageFile+".json")
		err := os.WriteFile(file, []byte(data.file), 0644)
		if err != nil {
			t.Error(err)
		}

		db := NewDB()
		(&backup{db: db}).load()

		if !reflect.DeepEqual(db, data.db) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, db, data.db)
		}
	}
}

func Test_backup_save(t *testing.T) {
	for idx, data := range []struct {
		db  *DB
		dir string

		file string
		err  error
	}{
		{ //00// no directory permission
			dir: func() string {
				if runtime.GOOS == "windows" {
					return "C:/Program Files/Temp/HeyApple"
				} else {
					return "/opt/tmp/heyapple"
				}
			}(),
			db:  NewDB(),
			err: &fs.PathError{},
		},
		{ //01// empty database
			dir:  testStorageDir,
			db:   NewDB(),
			file: backup0,
		},
		{ //02// save filled database
			dir:  testStorageDir,
			db:   database1,
			file: backup1,
		},
	} {
		os.Setenv(envStorageDir, data.dir)
		defer os.Unsetenv(envStorageDir)
		defer os.RemoveAll(data.dir)

		err := (&backup{db: data.db}).save()

		if reflect.TypeOf(err) != reflect.TypeOf(data.err) {
			t.Errorf("test case %d: error mismatch \nhave: %#v\nwant: %#v", idx, err, data.err)
		}

		if err == nil {
			file, err := os.ReadFile(filepath.Join(data.dir, storageFile+".json"))
			if err != nil {
				t.Errorf("test case %d: couldn't read storage file", idx)
			}

			if contents := string(file); contents != data.file {
				t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, contents, data.file)
			}
		}
	}
}

func Test_backup_backUp(t *testing.T) {
	for idx, data := range []struct {
		db  *DB
		dir string

		file string
		err  error
	}{
		{ //00// no directory permission
			dir: func() string {
				if runtime.GOOS == "windows" {
					return "C:/Program Files/Temp/HeyApple"
				} else {
					return "/opt/tmp/heyapple"
				}
			}(),
			db:  NewDB(),
			err: &fs.PathError{},
		},
		{ //01// empty database
			dir:  testStorageDir,
			db:   NewDB(),
			file: backup0,
		},
		{ //02// save filled database
			dir:  testStorageDir,
			db:   database1,
			file: backup1,
		},
	} {
		os.Setenv(envBackupDir, data.dir)
		defer os.Unsetenv(envBackupDir)
		defer os.RemoveAll(data.dir)

		backup := &backup{db: data.db}
		err := backup.backUp()

		if reflect.TypeOf(err) != reflect.TypeOf(data.err) {
			t.Errorf("test case %d: error mismatch \nhave: %#v\nwant: %#v", idx, err, data.err)
		}

		if err == nil {
			now := time.Now().Format("2006-01-02")
			file, err := os.ReadFile(filepath.Join(data.dir, backupFile+now+".json"))
			if err != nil {
				t.Errorf("test case %d: couldn't read storage file", idx)
			}

			if contents := string(file); contents != data.file {
				t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, contents, data.file)
			}

			if backup.lastBackup != now {
				t.Errorf("test case %d: date mismatch \nhave: %v\nwant: %v", idx, backup.lastBackup, now)
			}
		}
	}
}
