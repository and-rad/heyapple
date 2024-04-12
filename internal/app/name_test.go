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

package app

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/and-rad/heyapple/internal/core"
)

type database struct {
	user  User
	names []string
	err   []error
}

func (d *database) Execute(Command) error                                           { return nil }
func (d *database) Fetch(Query) error                                               { return nil }
func (d *database) NewUser(email, hash, token string) (int, error)                  { return 0, nil }
func (d *database) DelUser(id int) error                                            { return nil }
func (d *database) UserByEmail(email string) (User, error)                          { return User{}, nil }
func (d *database) UserPrefs(id int) (StoredPrefs, error)                           { return StoredPrefs{}, nil }
func (d *database) NewToken(id int, hash string, data interface{}) error            { return nil }
func (d *database) DelToken(string) error                                           { return nil }
func (d *database) Token(string) (Token, error)                                     { return Token{}, nil }
func (d *database) NewFood() (int, error)                                           { return 0, nil }
func (d *database) SetFood(core.Food) error                                         { return nil }
func (d *database) Food(id int) (core.Food, error)                                  { return core.Food{}, nil }
func (d *database) Foods(core.Filter) ([]core.Food, error)                          { return nil, nil }
func (d *database) FoodExists(id int) (bool, error)                                 { return false, nil }
func (d *database) NewRecipe(string) (int, error)                                   { return 0, nil }
func (d *database) SetRecipe(core.Recipe) error                                     { return nil }
func (d *database) SetRecipeAccess(user, rec, perms int) error                      { return nil }
func (d *database) SetRecipeInstructions(id int, text string) error                 { return nil }
func (d *database) DelRecipeInstructions(id int) error                              { return nil }
func (d *database) Recipe(id int) (core.Recipe, error)                              { return core.Recipe{}, nil }
func (d *database) Recipes(uid int, f core.Filter) ([]core.Recipe, error)           { return nil, nil }
func (d *database) RecipeAccess(user, rec int) (int, error)                         { return 0, nil }
func (d *database) RecipeInstructions(id int) (string, error)                       { return "", nil }
func (d *database) NewDiaryEntries(id int, entries ...core.DiaryEntry) error        { return nil }
func (d *database) SetDiaryEntries(id int, entries ...core.DiaryEntry) error        { return nil }
func (d *database) DelDiaryEntries(id int, entries ...core.DiaryEntry) error        { return nil }
func (d *database) DiaryEntries(id int, date time.Time) ([]core.DiaryEntry, error)  { return nil, nil }
func (d *database) DiaryDays(id, year, month, day int) ([]core.DiaryDay, error)     { return nil, nil }
func (d *database) SetShoppingListDone(id int, done map[int]bool) error             { return nil }
func (d *database) ShoppingList(id int, date ...time.Time) ([]core.ShopItem, error) { return nil, nil }
func (d *database) UserByName(name string) (User, error)                            { return User{}, nil }

func (d *database) UserByID(id int) (User, error) {
	if err := d.popError(); err != nil {
		return User{}, err
	}
	if d.user.ID != id {
		return User{}, ErrNotFound
	}
	return d.user, nil
}

func (d *database) UserNames(prefix string) ([]string, error) {
	if err := d.popError(); err != nil {
		return nil, err
	}
	names := []string{}
	for _, n := range d.names {
		if strings.HasPrefix(n, prefix) {
			names = append(names, n)
		}
	}
	return names, nil
}

func (d *database) SetUser(u User) error {
	if err := d.popError(); err != nil {
		return err
	}
	if d.user.ID != u.ID {
		return ErrNotFound
	}
	d.user = u
	return nil
}

func (d *database) popError() (err error) {
	if len(d.err) > 0 {
		err = d.err[0]
		d.err = d.err[1:]
	}
	return err
}

var errdos = errors.New("dos")

func TestChangeName(t *testing.T) {
	for idx, data := range []struct {
		db *database
		id int

		err error
	}{
		{ //00// invalid user id
			db:  &database{},
			err: ErrNotFound,
		},
		{ //01// user doesn't exist
			id:  2,
			db:  &database{},
			err: ErrNotFound,
		},
		{ //02// database error
			id:  1,
			db:  &database{user: User{}, err: []error{errdos}},
			err: errdos,
		},
		{ //03// success
			id: 1,
			db: &database{user: User{ID: 1, Name: "AnnoyingOrange"}},
		},
	} {
		old := data.db.user.Name

		cmd := &ChangeName{ID: data.id}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if data.err == err && err == nil {
			if data.db.user.Name == old {
				t.Errorf("test case %d: name change fail \nhave: %v\nwant: %v", idx, data.db.user.Name, cmd.Name)
			}
		}
	}
}

func TestChangeNameCollision(t *testing.T) {
	adj := adjectives
	noun := nouns
	defer func() {
		adjectives = adj
		nouns = noun
	}()

	for idx, data := range []struct {
		db    *database
		adjs  []string
		nouns []string

		name string
		err  error
	}{
		{ //00// database error
			db: &database{
				user:  User{ID: 1},
				names: []string{"GoodApple", "BadApple", "GoodBanana", "BadBanana"},
				err:   []error{nil, errdos},
			},
			adjs:  []string{"Good", "Bad"},
			nouns: []string{"Apple", "Banana"},
			err:   errdos,
		},
		{ //01// name exists once
			db: &database{
				user:  User{ID: 1},
				names: []string{"GoodApple", "GoodBanana", "BadBanana"},
			},
			adjs:  []string{"Good"},
			nouns: []string{"Apple"},
			name:  "GoodApple1",
		},
		{ //02// name exists multiple times
			db: &database{
				user:  User{ID: 1},
				names: []string{"GoodApple2", "GoodApple5", "GoodBanana", "BadBanana"},
			},
			adjs:  []string{"Good"},
			nouns: []string{"Apple"},
			name:  "GoodApple6",
		},
		{ //03// name exists multiple times, but some are technically invalid
			db: &database{
				user:  User{ID: 1},
				names: []string{"GoodApple", "BadApple", "BadApple2", "BadAppleXY", "GoodBanana", "BadBanana"},
			},
			adjs:  []string{"Bad"},
			nouns: []string{"Apple"},
			name:  "BadApple3",
		},
	} {
		adjectives = data.adjs
		nouns = data.nouns

		cmd := &ChangeName{ID: 1}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if cmd.Name != data.name {
			t.Errorf("test case %d: name mismatch \nhave: %v\nwant: %v", idx, cmd.Name, data.name)
		}
	}
}
