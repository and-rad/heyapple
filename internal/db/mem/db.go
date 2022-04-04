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

// package mem provides an implementation of the app.DB interface
// that lives entirely in the application's runtime memory. It can be
// saved to and loaded from JSON files for persistence.
package mem

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"sync"
	"time"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/core"
	"github.com/and-rad/heyapple/internal/job"
)

// entryMap   uid     day           entries
type entryMap map[int]map[time.Time][]core.DiaryEntry

// dayMap   uid     yr      mon     entries
type dayMap map[int]map[int]map[int][]core.DiaryDay

// aisleMap   uid     food    aisle
type aisleMap map[int]map[int]core.Aisle

// priceMap   uid     food    price range
type priceMap map[int]map[int][2]float32

// doneMap   uid     food    done
type doneMap map[int]map[int]bool

type DB struct {
	jobs *job.Scheduler

	users   map[int]app.User
	emails  map[string]int
	tokens  map[string]app.Token
	food    map[int]core.Food
	recipes map[int]core.Recipe
	entries entryMap
	days    dayMap

	userRec map[int]map[int]int
	recUser map[int]map[int]int

	aisles aisleMap
	prices priceMap
	done   doneMap

	userID int
	foodID int
	recID  int

	mtx sync.RWMutex
}

func NewDB() *DB {
	return &DB{
		users:   make(map[int]app.User),
		tokens:  make(map[string]app.Token),
		emails:  make(map[string]int),
		food:    make(map[int]core.Food),
		recipes: make(map[int]core.Recipe),
		userRec: make(map[int]map[int]int),
		recUser: make(map[int]map[int]int),
		entries: make(entryMap),
		days:    make(dayMap),
		aisles:  make(aisleMap),
		prices:  make(priceMap),
		done:    make(doneMap),
	}
}

func (db *DB) WithDefaults(fs fs.FS) *DB {
	if len(db.users) == 0 {
		users := []app.User{}
		data := loadDefault(fs, "user.json")
		if err := json.Unmarshal(data, &users); err != nil {
			return db
		}

		db.userID = len(users)
		for _, u := range users {
			db.users[u.ID] = u
			db.emails[u.Email] = u.ID
		}
	}

	if len(db.food) == 0 {
		food := []struct {
			core.Food
			Aisle core.Aisle
		}{}
		data := loadDefault(fs, "food.json")
		if err := json.Unmarshal(data, &food); err != nil {
			return db
		}

		if num := len(food); num > 0 {
			db.aisles[0] = map[int]core.Aisle{}
		}

		for _, f := range food {
			db.food[f.ID] = f.Food
			db.aisles[0][f.ID] = f.Aisle
			if db.foodID < f.ID {
				db.foodID = f.ID
			}
		}
	}

	if len(db.recipes) == 0 {
		recs := []core.Recipe{}
		data := loadDefault(fs, "recipe.json")
		if err := json.Unmarshal(data, &recs); err != nil {
			return db
		}

		db.recID = len(recs)
		db.userRec[0] = map[int]int{}
		for _, r := range recs {
			db.recipes[r.ID] = r
			db.userRec[0][r.ID] = app.PermRead
		}
	}

	return db
}

func (db *DB) WithBackup(log app.Logger) *DB {
	backupper := &backup{db: db, log: log}
	backupper.load()

	db.jobs = job.NewScheduler(
		getConfig().storageInterval,
		backupper,
	)

	go db.jobs.Run()
	return db
}

func (db *DB) Close() error {
	if db.jobs != nil {
		db.jobs.Stop()
	}
	return nil
}

func (db *DB) Execute(c app.Command) error {
	db.mtx.Lock()
	defer db.mtx.Unlock()
	return c.Execute(db)
}

func (db *DB) Fetch(q app.Query) error {
	db.mtx.RLock()
	defer db.mtx.RUnlock()
	return q.Fetch(db)
}

func loadDefault(fs fs.FS, name string) []byte {
	file, err := fs.Open(name)
	if err != nil {
		return nil
	}

	defer file.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(file); err != nil {
		return nil
	}

	return buf.Bytes()
}
