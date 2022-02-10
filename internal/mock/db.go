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
	ErrNotFound = errors.New("not found")
	ErrDOS      = errors.New("dos")
)

var (
	nilFood = core.Food{}
)

type DB struct {
	FoodInfo []core.Food

	LastFoodID uint32

	FailFood bool
}

func NewDB() *DB {
	return &DB{
		FoodInfo: []core.Food{},
	}
}

// Fail is a convenience function for method chaining
// configuration parameters
func (db *DB) Fail(fail bool) *DB {
	db.FailFood = fail
	return db
}

func (db *DB) Prefill() *DB {
	db.FoodInfo = []core.Food{
		Food1,
		Food2,
	}

	db.LastFoodID = 2

	return db
}

func (db *DB) Execute(c app.Command) error {
	return c.Execute(db)
}

func (db *DB) Fetch(q app.Query) error {
	return q.Fetch(db)
}

func (db *DB) Food(id uint32) (core.Food, error) {
	if db.FailFood {
		return nilFood, ErrDOS
	}

	for _, f := range db.FoodInfo {
		if f.ID == id {
			return f, nil
		}
	}

	return nilFood, ErrNotFound
}

func (db *DB) Foods() ([]core.Food, error) {
	if db.FailFood {
		return []core.Food{}, ErrDOS
	}
	return db.FoodInfo, nil
}

func (db *DB) NewFood() (uint32, error) {
	if db.FailFood {
		return 0, ErrDOS
	}

	db.LastFoodID++
	db.FoodInfo = append(db.FoodInfo, core.Food{ID: db.LastFoodID})

	return db.LastFoodID, nil
}

func (db *DB) SetFood(food core.Food) error {
	if db.FailFood {
		return ErrDOS
	}

	for i, f := range db.FoodInfo {
		if f.ID == food.ID {
			db.FoodInfo[i] = food
			return nil
		}
	}

	return ErrNotFound
}
