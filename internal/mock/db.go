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
	FoodItem  core.Food
	FoodItems []core.Food

	Err error
	ID  uint32
}

func NewDB() *DB {
	return &DB{
		FoodItems: []core.Food{},
	}
}

func (db *DB) WithError(e error) *DB {
	db.Err = e
	return db
}

func (db *DB) WithID(id uint32) *DB {
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

func (db *DB) Execute(c app.Command) error {
	return c.Execute(db)
}

func (db *DB) Fetch(q app.Query) error {
	return q.Fetch(db)
}

func (db *DB) Food(id uint32) (core.Food, error) {
	if db.Err != nil {
		return core.Food{}, db.Err
	}
	if db.FoodItem.ID != id {
		return core.Food{}, app.ErrNotFound
	}
	return db.FoodItem, nil
}

func (db *DB) Foods() ([]core.Food, error) {
	if db.Err != nil {
		return nil, db.Err
	}
	return db.FoodItems, nil
}

func (db *DB) NewFood() (uint32, error) {
	if db.Err != nil {
		return 0, db.Err
	}
	return db.ID, nil
}

func (db *DB) SetFood(food core.Food) error {
	if db.Err != nil {
		return db.Err
	}
	if db.FoodItem.ID != food.ID {
		return app.ErrNotFound
	}
	db.FoodItem = food
	return nil
}
