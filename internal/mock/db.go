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

// Package mock supplies various stubs and mock objects that can be
// used for unit testing and development deployments.
// nolint
package mock

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/core"
)

// Error definitions
var (
	ErrDOS = errors.New("dos")
)

type Access struct {
	User     int
	Resource int
	Perms    int
}

type Instructions struct {
	Recipe       int
	Instructions string
}

type DB struct {
	User         app.User
	Prefs        app.StoredPrefs
	Tok          app.Token
	FoodItem     core.Food
	FoodItems    []core.Food
	RecipeItem   core.Recipe
	RecipeItems  []core.Recipe
	Entries      []core.DiaryEntry
	Days         []core.DiaryDay
	ShopList     []core.ShopItem
	Access       Access
	Instructions Instructions

	Err  []error
	ID   int
	Name string

	Filter core.Filter

	date time.Time
}

func NewDB() *DB {
	return &DB{
		FoodItems:   []core.Food{},
		RecipeItems: []core.Recipe{},
		Entries:     []core.DiaryEntry{},
		Days:        []core.DiaryDay{},
	}
}

func (db *DB) WithError(e ...error) *DB {
	db.Err = e
	return db
}

func (db *DB) WithID(id int) *DB {
	db.ID = id
	return db
}

func (db *DB) WithFood(food core.Food) *DB {
	db.FoodItem = food
	return db
}

func (db *DB) WithFoods(foods ...core.Food) *DB {
	db.FoodItems = foods
	return db
}

func (db *DB) WithUser(user app.User) *DB {
	db.User = user
	return db
}

func (db *DB) WithPrefs(prefs app.StoredPrefs) *DB {
	db.Prefs = prefs
	return db
}

func (db *DB) WithToken(tok app.Token) *DB {
	db.Tok = tok
	return db
}

func (db *DB) WithName(name string) *DB {
	db.Name = name
	return db
}

func (db *DB) WithDate(date time.Time) *DB {
	db.date = date.Truncate(time.Hour * 24)
	return db
}

func (db *DB) WithRecipe(rec core.Recipe) *DB {
	db.RecipeItem = rec
	return db
}

func (db *DB) WithRecipes(recs ...core.Recipe) *DB {
	db.RecipeItems = recs
	return db
}

func (db *DB) WithAccess(a Access) *DB {
	db.Access = a
	return db
}

func (db *DB) WithInstructions(id int, text string) *DB {
	db.Instructions = Instructions{Recipe: id, Instructions: text}
	return db
}

func (db *DB) WithEntries(ents ...core.DiaryEntry) *DB {
	db.Entries = ents
	return db
}

func (db *DB) WithDays(days ...core.DiaryDay) *DB {
	db.Days = days
	return db
}

func (db *DB) WithShoppingList(items ...core.ShopItem) *DB {
	db.ShopList = items
	return db
}

func (db *DB) Execute(c app.Command) error {
	return c.Execute(db)
}

func (db *DB) Fetch(q app.Query) error {
	return q.Fetch(db)
}

func (db *DB) NewUser(email, hash, token string) (int, error) {
	if err := db.popError(); err != nil {
		return 0, err
	}
	db.User = app.User{ID: db.ID, Email: email, Pass: hash}
	db.Tok = app.Token{ID: db.ID}
	return db.ID, nil
}

func (db *DB) SetUser(user app.User) error {
	if err := db.popError(); err != nil {
		return err
	}
	if db.User.ID != user.ID {
		return app.ErrNotFound
	}
	db.User = user
	return nil
}

func (db *DB) DelUser(id int) error {
	if err := db.popError(); err != nil {
		return err
	}
	if db.User.ID != id {
		return nil
	}
	db.User = app.User{}
	return nil
}

func (db *DB) UserByID(id int) (app.User, error) {
	if err := db.popError(); err != nil {
		return app.User{}, err
	}
	if db.User.ID != id {
		return app.User{}, app.ErrNotFound
	}
	return db.User, nil
}

func (db *DB) UserNames(prefix string) ([]string, error) {
	if err := db.popError(); err != nil {
		return nil, err
	}
	if db.User.Name == "" {
		return []string{}, nil
	}
	return []string{db.User.Name}, nil
}

func (db *DB) UserByEmail(email string) (app.User, error) {
	if err := db.popError(); err != nil {
		return app.User{}, err
	}
	if db.User.Email != email {
		return app.User{}, app.ErrNotFound
	}
	return db.User, nil
}

func (db *DB) UserByName(name string) (app.User, error) {
	if err := db.popError(); err != nil {
		return app.User{}, err
	}
	if db.User.Name != name {
		return app.User{}, app.ErrNotFound
	}
	return db.User, nil
}

func (db *DB) SetUserPrefs(id int, prefs app.StoredPrefs) error {
	if err := db.popError(); err != nil {
		return err
	}
	if db.User.ID != id {
		return app.ErrNotFound
	}
	db.Prefs = prefs
	return nil
}

func (db *DB) UserPrefs(id int) (app.StoredPrefs, error) {
	if err := db.popError(); err != nil {
		return app.StoredPrefs{}, err
	}
	if db.User.ID != id {
		return app.StoredPrefs{}, app.ErrNotFound
	}
	prefs := db.Prefs
	return prefs, nil
}

func (db *DB) Token(string) (app.Token, error) {
	if err := db.popError(); err != nil {
		return app.Token{}, err
	}
	if db.Tok.ID == 0 {
		return app.Token{}, app.ErrNotFound
	}
	return db.Tok, nil
}

func (db *DB) NewToken(id int, hash string, data interface{}) error {
	if err := db.popError(); err != nil {
		return err
	}
	db.Tok = app.Token{ID: id, Data: data}
	return nil
}

func (db *DB) DelToken(string) error {
	db.Tok = app.Token{}
	return nil
}

func (db *DB) Food(id int) (core.Food, error) {
	if err := db.popError(); err != nil {
		return core.Food{}, err
	}
	if db.FoodItem.ID != id {
		for _, f := range db.FoodItems {
			if f.ID == id {
				return f, nil
			}
		}
		return core.Food{}, app.ErrNotFound
	}
	return db.FoodItem, nil
}

func (db *DB) Foods(f core.Filter) ([]core.Food, error) {
	if err := db.popError(); err != nil {
		return nil, err
	}
	db.Filter = f
	return db.FoodItems, nil
}

func (db *DB) NewFood() (int, error) {
	if err := db.popError(); err != nil {
		return 0, err
	}
	return db.ID, nil
}

func (db *DB) SetFood(food core.Food) error {
	if err := db.popError(); err != nil {
		return err
	}
	if db.FoodItem.ID != food.ID {
		return app.ErrNotFound
	}
	db.FoodItem = food
	return nil
}

func (db *DB) FoodExists(id int) (bool, error) {
	if err := db.popError(); err != nil {
		return false, err
	}
	for _, f := range db.FoodItems {
		if f.ID == id {
			return true, nil
		}
	}
	return false, nil
}

func (db *DB) NewRecipe(name string) (int, error) {
	if err := db.popError(); err != nil {
		return 0, err
	}
	db.Name = name
	return db.ID, nil
}

func (db *DB) SetRecipe(rec core.Recipe) error {
	if err := db.popError(); err != nil {
		return err
	}
	if db.RecipeItem.ID != rec.ID {
		return app.ErrNotFound
	}
	db.RecipeItem = rec
	return nil
}

func (db *DB) SetRecipeAccess(user, rec, perms int) error {
	if err := db.popError(); err != nil {
		return err
	}
	db.Access = Access{User: user, Resource: rec, Perms: perms}
	return nil
}

func (db *DB) SetRecipeInstructions(id int, text string) error {
	if err := db.popError(); err != nil {
		return err
	}

	if db.Instructions.Recipe != id {
		return app.ErrNotFound
	}

	db.Instructions.Instructions = text
	return nil
}

func (db *DB) DelRecipeInstructions(id int) error {
	if err := db.popError(); err != nil {
		return err
	}

	if db.Instructions.Recipe != id {
		return nil
	}

	db.Instructions = Instructions{}
	return nil
}

func (db *DB) Recipe(id int) (core.Recipe, error) {
	if err := db.popError(); err != nil {
		return core.Recipe{}, err
	}
	if db.RecipeItem.ID != id {
		return core.Recipe{}, app.ErrNotFound
	}
	return db.RecipeItem, nil
}

func (db *DB) RecipeAccess(user, rec int) (int, error) {
	if err := db.popError(); err != nil {
		return 0, err
	}
	if db.Access.User == user && db.Access.Resource == rec {
		return db.Access.Perms, nil
	}
	return 0, nil
}

func (db *DB) RecipeInstructions(id int) (string, error) {
	if err := db.popError(); err != nil {
		return "", err
	}
	if db.Instructions.Recipe == id {
		return db.Instructions.Instructions, nil
	}
	return "", nil
}

func (db *DB) Recipes(uid int, f core.Filter) ([]core.Recipe, error) {
	if err := db.popError(); err != nil {
		return nil, err
	}
	db.Filter = f
	return db.RecipeItems, nil
}

func (db *DB) NewDiaryEntries(id int, entries ...core.DiaryEntry) error {
	if err := db.popError(); err != nil {
		return err
	}
	db.Entries = append(db.Entries, entries...)
	return nil
}

func (db *DB) SetDiaryEntries(id int, entries ...core.DiaryEntry) error {
	if err := db.popError(); err != nil {
		return err
	}
	for i, old := range db.Entries {
		for _, new := range entries {
			if old.Equal(new) {
				db.Entries[i] = new
				break
			}
		}
	}
	return nil
}

func (db *DB) DelDiaryEntries(id int, entries ...core.DiaryEntry) error {
	if err := db.popError(); err != nil {
		return err
	}

	tmp := []core.DiaryEntry{}
	for _, old := range db.Entries {
		found := false
		for _, new := range entries {
			if new.Equal(old) {
				found = true
				break
			}
		}
		if !found {
			tmp = append(tmp, old)
		}
	}

	db.Entries = tmp
	return nil
}

func (db *DB) DiaryEntries(id int, date time.Time) ([]core.DiaryEntry, error) {
	if err := db.popError(); err != nil {
		return nil, err
	}
	day := date.Truncate(time.Hour * 24)
	entries := []core.DiaryEntry{}
	for _, e := range db.Entries {
		if e.Day() == day {
			entries = append(entries, e)
		}
	}
	return entries, nil
}

func (db *DB) DiaryDays(id, year, month, day int) ([]core.DiaryDay, error) {
	if err := db.popError(); err != nil {
		return nil, err
	}

	if year < 1 && month < 1 && day < 1 {
		return db.Days, nil
	}

	comp := strconv.FormatInt(int64(year), 10)
	if month > 0 {
		comp += fmt.Sprintf("-%02d", month)
	}
	if day > 0 {
		comp += fmt.Sprintf("-%02d", day)
	}

	days := []core.DiaryDay{}
	for _, d := range db.Days {
		if strings.HasPrefix(d.Date, comp) {
			days = append(days, d)
		}
	}

	return days, nil
}

func (db *DB) SetShoppingListDone(id int, done map[int]bool) error {
	if err := db.popError(); err != nil {
		return err
	}
	for i, l := range db.ShopList {
		if v, ok := done[l.ID]; ok {
			db.ShopList[i].Done = v
		}
	}
	return nil
}

func (db *DB) ShoppingList(id int, date ...time.Time) ([]core.ShopItem, error) {
	if err := db.popError(); err != nil {
		return nil, err
	}
	if len(date) == 0 {
		return []core.ShopItem{}, nil
	}
	day := date[0].Truncate(time.Hour * 24)
	if day != db.date {
		return []core.ShopItem{}, nil
	}
	result := append([]core.ShopItem{}, db.ShopList...)
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result, nil
}

func (db *DB) popError() error {
	var err error
	if len(db.Err) > 0 {
		err = db.Err[0]
		db.Err = db.Err[1:]
	}
	return err
}
