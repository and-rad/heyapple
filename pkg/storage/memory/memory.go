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
	"bytes"
	"encoding/json"
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"heyapple/pkg/job"
	"io/fs"
	"sort"
	"sync"
)

type DB struct {
	jobs *job.Scheduler

	users   map[int]app.User
	emails  map[string]int
	tokens  map[string]app.Token
	food    map[int]core.Food
	recipes map[int]core.Recipe
	recMeta map[int]core.RecipeMeta

	userRec map[int]map[int]int
	recUser map[int]map[int]int

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
		recMeta: make(map[int]core.RecipeMeta),
		userRec: make(map[int]map[int]int),
		recUser: make(map[int]map[int]int),
	}
}

func (db *DB) WithDefaults(fs fs.FS) *DB {
	if len(db.food) == 0 {
		file, err := fs.Open("food.json")
		if err != nil {
			return db
		}

		defer file.Close()

		var buf bytes.Buffer
		if _, err := buf.ReadFrom(file); err != nil {
			return db
		}

		food := []core.Food{}
		if err := json.Unmarshal(buf.Bytes(), &food); err != nil {
			return db
		}

		db.foodID = len(food)
		for _, f := range food {
			db.food[f.ID] = f
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

func (db *DB) NewUser(name, hash, token string) (int, error) {
	if _, ok := db.emails[name]; ok {
		return 0, app.ErrExists
	}

	db.userID++
	db.emails[name] = db.userID
	db.tokens[token] = app.Token{ID: db.userID}
	db.users[db.userID] = app.User{
		ID:    db.userID,
		Email: name,
		Pass:  hash,
		Lang:  getConfig().defaultLang,
	}

	return db.userID, nil
}

func (db *DB) SetUser(user app.User) error {
	if u, ok := db.users[user.ID]; ok {
		delete(db.emails, u.Email)
		db.users[user.ID] = user
		db.emails[user.Email] = user.ID
		return nil
	}
	return app.ErrNotFound
}

func (db *DB) UserByName(name string) (app.User, error) {
	if id, ok := db.emails[name]; ok {
		if user, ok := db.users[id]; ok {
			return user, nil
		}
	}
	return app.User{}, app.ErrNotFound
}

func (db *DB) UserByID(id int) (app.User, error) {
	if user, ok := db.users[id]; ok {
		return user, nil
	}
	return app.User{}, app.ErrNotFound
}

func (db *DB) Token(hash string) (app.Token, error) {
	if token, ok := db.tokens[hash]; ok {
		return token, nil
	}
	return app.Token{}, app.ErrNotFound
}

func (db *DB) NewToken(id int, hash string, data interface{}) error {
	db.tokens[hash] = app.Token{ID: id, Data: data}
	return nil
}

func (db *DB) DeleteToken(hash string) error {
	delete(db.tokens, hash)
	return nil
}

func (db *DB) Food(id int) (core.Food, error) {
	if food, ok := db.food[id]; ok {
		return food, nil
	}
	return core.Food{}, app.ErrNotFound
}

func (db *DB) Foods(filter core.Filter) ([]core.Food, error) {
	foods := []core.Food{}
	for _, f := range db.food {
		if filter.MatchFood(f) {
			foods = append(foods, f)
		}
	}

	sort.Slice(foods, func(i, j int) bool {
		return foods[i].ID < foods[j].ID
	})

	return foods, nil
}

func (db *DB) FoodExists(id int) (bool, error) {
	_, ok := db.food[id]
	return ok, nil
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

func (db *DB) NewRecipe(name string) (int, error) {
	db.recID++
	db.recipes[db.recID] = core.Recipe{ID: db.recID, Size: 1}
	db.recMeta[db.recID] = core.RecipeMeta{ID: db.recID, Name: name}
	return db.recID, nil
}

func (db *DB) SetRecipe(rec core.Recipe) error {
	if _, ok := db.recipes[rec.ID]; !ok {
		return app.ErrNotFound
	}

	meta := db.recMeta[rec.ID]
	meta.KCal = 0
	meta.Fat = 0
	meta.Carbs = 0
	meta.Protein = 0
	for _, v := range rec.Items {
		food := db.food[v.ID]
		amount := v.Amount * 0.01
		meta.KCal += food.KCal * amount
		meta.Fat += food.Fat * amount
		meta.Carbs += food.Carbs * amount
		meta.Protein += food.Protein * amount
	}

	db.recipes[rec.ID] = rec
	db.recMeta[rec.ID] = meta

	return nil
}

func (db *DB) SetRecipeAccess(user, rec, perms int) error {
	if _, ok := db.users[user]; !ok {
		return app.ErrNotFound
	}
	if _, ok := db.recipes[rec]; !ok {
		return app.ErrNotFound
	}
	if _, ok := db.userRec[user]; !ok {
		db.userRec[user] = make(map[int]int)
	}
	if _, ok := db.recUser[rec]; !ok {
		db.recUser[rec] = make(map[int]int)
	}
	db.userRec[user][rec] = perms
	db.recUser[rec][user] = perms
	return nil
}

func (db *DB) Recipe(id int) (core.Recipe, error) {
	if rec, ok := db.recipes[id]; ok {
		return rec, nil
	}
	return core.Recipe{}, app.ErrNotFound
}

func (db *DB) RecipeMeta(id int) (core.RecipeMeta, error) {
	if meta, ok := db.recMeta[id]; ok {
		return meta, nil
	}
	return core.RecipeMeta{}, app.ErrNotFound
}

func (db *DB) RecipeAccess(user, rec int) (int, error) {
	if acc, ok := db.userRec[user]; ok {
		return acc[rec], nil
	}
	return app.PermNone, nil
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
