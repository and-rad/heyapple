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

// Package memory provides an implementation of the app.DB interface
// that lives entirely in the application's runtime memory. It can be
// saved to and loaded from JSON files for persistence.
package memory

import (
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"heyapple/pkg/job"
	"sort"
	"sync"
)

type DB struct {
	log  app.Logger
	jobs *job.Scheduler

	users  map[int]app.User
	emails map[string]int
	food   map[int]core.Food

	foodID int

	mtx sync.RWMutex
}

func NewDB(log app.Logger) *DB {
	return &DB{
		log:  log,
		food: make(map[int]core.Food),
	}
}

func NewDBWithBackup(log app.Logger) *DB {
	backupper := &backup{db: NewDB(log)}
	backupper.load()

	db := backupper.db
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

func (db *DB) UserByName(name string) (app.User, error) {
	if id, ok := db.emails[name]; ok {
		if user, ok := db.users[id]; ok {
			return user, nil
		}
	}
	return app.User{}, app.ErrNotFound
}

func (db *DB) Food(id int) (core.Food, error) {
	if food, ok := db.food[id]; ok {
		return food, nil
	}
	return core.Food{}, app.ErrNotFound
}

func (db *DB) Foods() ([]core.Food, error) {
	foods := []core.Food{}
	for _, f := range db.food {
		foods = append(foods, f)
	}

	sort.Slice(foods, func(i, j int) bool {
		return foods[i].ID < foods[j].ID
	})

	return foods, nil
}

func (db *DB) NewFood() (int, error) {
	db.foodID++
	db.food[db.foodID] = core.Food{ID: db.foodID}
	return db.foodID, nil
}

func (db *DB) SetFood(food core.Food) error {
	if _, ok := db.food[food.ID]; ok {
		db.food[food.ID] = food
		return nil
	}
	return app.ErrNotFound
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
