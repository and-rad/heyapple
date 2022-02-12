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
	"heyapple/pkg/core"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

var (
	backup0 = `{"food":{},"foodid":0}`
	backup1 = fmt.Sprintf(`{"food":{"1":%s,"2":%s},"foodid":2}`, mock.Food1Json, mock.Food2Json)

	database1 = &DB{log: mock.NewLog(), food: map[int]core.Food{1: mock.Food1, 2: mock.Food2}, foodID: 2}
)

func Test_backup_Run(t *testing.T) {
	for idx, data := range []struct {
		db  *DB
		dir string

		file string
	}{
		{ //00// no directory permission
			dir: "/opt/tmp/heyapple",
			db:  NewDB(mock.NewLog()),
		},
		{ //01// empty database
			dir:  testStorageDir,
			db:   NewDB(mock.NewLog()),
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

		(&backup{db: data.db}).Run()

		file, _ := ioutil.ReadFile(filepath.Join(data.dir, storageFile+".json"))
		if contents := string(file); contents != data.file {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, contents, data.file)
		}

		now := time.Now().Format("2006-01-02")
		file, _ = ioutil.ReadFile(filepath.Join(data.dir, backupFile+now+".json"))
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
			db: NewDB(mock.NewLog()),
		},
		{ //01// invalid JSON
			file: `{"food":`,
			db:   NewDB(mock.NewLog()),
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

		db := NewDB(mock.NewLog())
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
			dir: "/opt/tmp/heyapple",
			db:  NewDB(mock.NewLog()),
			err: &fs.PathError{},
		},
		{ //01// empty database
			dir:  testStorageDir,
			db:   NewDB(mock.NewLog()),
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
			file, err := ioutil.ReadFile(filepath.Join(data.dir, storageFile+".json"))
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
			dir: "/opt/tmp/heyapple",
			db:  NewDB(mock.NewLog()),
			err: &fs.PathError{},
		},
		{ //01// empty database
			dir:  testStorageDir,
			db:   NewDB(mock.NewLog()),
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
			file, err := ioutil.ReadFile(filepath.Join(data.dir, backupFile+now+".json"))
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
