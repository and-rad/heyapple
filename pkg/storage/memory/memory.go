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
	"time"
)

// entryMap       uid     day         entries
type entryMap map[int]map[time.Time][]core.DiaryEntry

type DB struct {
	jobs *job.Scheduler

	users   map[int]app.User
	emails  map[string]int
	tokens  map[string]app.Token
	food    map[int]core.Food
	recipes map[int]core.Recipe
	entries entryMap

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
		entries: make(entryMap),
		userRec: make(map[int]map[int]int),
		recUser: make(map[int]map[int]int),
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
		food := []core.Food{}
		data := loadDefault(fs, "food.json")
		if err := json.Unmarshal(data, &food); err != nil {
			return db
		}

		db.foodID = len(food)
		for _, f := range food {
			db.food[f.ID] = f
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
	rec := core.NewRecipe(db.recID)
	rec.Name = name
	db.recipes[rec.ID] = rec
	return rec.ID, nil
}

func (db *DB) SetRecipe(rec core.Recipe) error {
	if _, ok := db.recipes[rec.ID]; ok {
		db.recipes[rec.ID] = rec
		return nil
	}
	return app.ErrNotFound
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
	if r, ok := db.recipes[id]; ok {
		rec := r
		rec.Items = append([]core.Ingredient{}, r.Items...)
		return rec, nil
	}
	return core.Recipe{}, app.ErrNotFound
}

func (db *DB) RecipeAccess(user, rec int) (int, error) {
	combined := app.PermNone
	if acc, ok := db.userRec[0]; ok {
		combined |= acc[rec]
	}
	if acc, ok := db.userRec[user]; ok {
		combined |= acc[rec]
	}
	return combined, nil
}

func (db *DB) Recipes(uid int, f core.Filter) ([]core.Recipe, error) {
	ids := map[int]struct{}{}
	recs := []core.Recipe{}

	for id, perm := range db.userRec[0] {
		r := db.recipes[id]
		if perm != app.PermNone && f.MatchRecipe(r) {
			ids[id] = struct{}{}
			rec := r
			rec.Items = append([]core.Ingredient{}, r.Items...)
			recs = append(recs, rec)
		}
	}

	for id, perm := range db.userRec[uid] {
		if _, ok := ids[id]; !ok {
			r := db.recipes[id]
			if perm != app.PermNone && f.MatchRecipe(r) {
				rec := r
				rec.Items = append([]core.Ingredient{}, r.Items...)
				recs = append(recs, rec)
			}
		}
	}

	sort.Slice(recs, func(i, j int) bool {
		return recs[i].ID < recs[j].ID
	})

	return recs, nil
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

func (db *DB) NewDiaryEntries(id int, entries ...core.DiaryEntry) error {
	days, ok := db.entries[id]
	if !ok {
		db.entries[id] = map[time.Time][]core.DiaryEntry{}
		days = db.entries[id]
	}

	dirty := map[time.Time]struct{}{}
	for _, e := range entries {
		date := e.Date.Truncate(time.Hour * 24)
		if _, ok := days[date]; !ok {
			days[date] = []core.DiaryEntry{e}
			continue
		}
		days[date] = append(days[date], e)
		dirty[date] = struct{}{}
	}

	for date := range dirty {
		sort.Slice(days[date], func(i, j int) bool {
			return days[date][i].Date.Before(days[date][j].Date)
		})
	}

	return nil
}

func (db *DB) SetDiaryEntries(id int, entries ...core.DiaryEntry) error {
	days, ok := db.entries[id]
	if !ok {
		return app.ErrNotFound
	}

	for _, entry := range entries {
		date := entry.Day()
		if _, ok := days[date]; !ok {
			continue
		}

		for i, e := range days[date] {
			if e.Equal(entry) {
				days[date][i] = entry
				break
			}
		}
	}

	return nil
}

func (db *DB) DelDiaryEntries(id int, entries ...core.DiaryEntry) error {
	days, ok := db.entries[id]
	if !ok {
		return app.ErrNotFound
	}

	for _, entry := range entries {
		date := entry.Day()
		if _, ok := days[date]; !ok {
			continue
		}

		end := 0
		for _, e := range days[date] {
			if !e.Equal(entry) {
				days[date][end] = e
				end++
			}
		}

		days[date] = days[date][:end]
	}

	return nil
}

func (db *DB) DiaryEntry(diary, food int, date time.Time) (core.DiaryEntry, error) {
	days, ok := db.entries[diary]
	if !ok {
		return core.DiaryEntry{}, app.ErrNotFound
	}

	day, ok := days[date.Truncate(time.Hour*24)]
	if !ok {
		return core.DiaryEntry{}, app.ErrNotFound
	}

	test := core.DiaryEntry{
		Food: core.Ingredient{ID: food},
		Date: date,
	}

	for _, e := range day {
		if e.Equal(test) {
			return e, nil
		}
	}

	return core.DiaryEntry{}, app.ErrNotFound
}

func (db *DB) DiaryEntries(id int, date time.Time) ([]core.DiaryEntry, error) {
	if days, ok := db.entries[id]; ok {
		if day, ok := days[date.Truncate(time.Hour*24)]; ok {
			return append(make([]core.DiaryEntry, 0, len(day)), day...), nil
		}
	}
	return nil, app.ErrNotFound
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
