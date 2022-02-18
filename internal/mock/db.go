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

// Package mock supplies various stubs and mock objects that can be
// used for unit testing and development deployments.
// nolint
package mock

import (
	"errors"
	"heyapple/pkg/app"
	"heyapple/pkg/core"
)

// Error definitions
var (
	ErrDOS = errors.New("dos")
)

type DB struct {
	User      app.User
	Tok       app.Token
	FoodItem  core.Food
	FoodItems []core.Food

	Err []error
	ID  int
}

func NewDB() *DB {
	return &DB{
		FoodItems: []core.Food{},
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

func (db *DB) WithFoods(foods []core.Food) *DB {
	db.FoodItems = foods
	return db
}

func (db *DB) WithUser(user app.User) *DB {
	db.User = user
	return db
}

func (db *DB) WithToken(tok app.Token) *DB {
	db.Tok = tok
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

func (db *DB) UserByID(id int) (app.User, error) {
	if err := db.popError(); err != nil {
		return app.User{}, err
	}
	if db.User.ID != id {
		return app.User{}, app.ErrNotFound
	}
	return db.User, nil
}

func (db *DB) UserByName(name string) (app.User, error) {
	if err := db.popError(); err != nil {
		return app.User{}, err
	}
	if db.User.Email != name {
		return app.User{}, app.ErrNotFound
	}
	return db.User, nil
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

func (db *DB) DeleteToken(string) error {
	db.Tok = app.Token{}
	return nil
}

func (db *DB) Food(id int) (core.Food, error) {
	if err := db.popError(); err != nil {
		return core.Food{}, err
	}
	if db.FoodItem.ID != id {
		return core.Food{}, app.ErrNotFound
	}
	return db.FoodItem, nil
}

func (db *DB) Foods() ([]core.Food, error) {
	if err := db.popError(); err != nil {
		return nil, err
	}
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

func (db *DB) popError() error {
	var err error
	if len(db.Err) > 0 {
		err = db.Err[0]
		db.Err = db.Err[1:]
	}
	return err
}
