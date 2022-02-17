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
	"encoding/json"
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const (
	backupFile  = "db"
	storageFile = "db"
)

type backup struct {
	db         *DB
	lastBackup string
}

type backupData struct {
	Users  map[int]app.User  `json:"users"`
	Food   map[int]core.Food `json:"food"`
	UserID int               `json:"userid"`
	FoodID int               `json:"foodid"`
}

func (b *backup) Run() {
	if err := b.save(); err != nil {
		b.db.log.Error(err)
	}

	if now := time.Now().Format("2006-01-02"); now != b.lastBackup {
		if err := b.backUp(); err != nil {
			b.db.log.Error(err)
		}
	}
}

func (b *backup) load() {
	b.db.mtx.Lock()
	defer b.db.mtx.Unlock()

	path := filepath.Join(getConfig().storageDir, storageFile+".json")
	if data, err := ioutil.ReadFile(path); err == nil {
		var db backupData
		if err = json.Unmarshal(data, &db); err == nil {
			b.db.userID = db.UserID
			b.db.users = db.Users
			for k, v := range b.db.users {
				b.db.emails[v.Email] = k
			}

			b.db.food = db.Food
			b.db.foodID = db.FoodID
		}
	}
}

func (b *backup) save() error {
	b.db.mtx.RLock()
	defer b.db.mtx.RUnlock()

	dir := getConfig().storageDir
	if err := os.MkdirAll(dir, 0750); err != nil {
		return err
	}

	data, _ := json.Marshal(backupData{
		Users:  b.db.users,
		UserID: b.db.userID,
		Food:   b.db.food,
		FoodID: b.db.foodID,
	})

	path := filepath.Join(dir, storageFile+".json")
	return os.WriteFile(path, data, 0644)
}

func (b *backup) backUp() error {
	b.db.mtx.RLock()
	defer b.db.mtx.RUnlock()

	dir := getConfig().backupDir
	if err := os.MkdirAll(dir, 0750); err != nil {
		return err
	}

	data, _ := json.Marshal(backupData{
		Users:  b.db.users,
		UserID: b.db.userID,
		Food:   b.db.food,
		FoodID: b.db.foodID,
	})

	now := time.Now().Format("2006-01-02")
	path := filepath.Join(dir, backupFile+now+".json")
	err := os.WriteFile(path, data, 0644)

	if err == nil {
		b.lastBackup = now
	}

	return err
}
