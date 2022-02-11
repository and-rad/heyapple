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
	"sync"
)

type DB struct {
	mtx  sync.RWMutex
	food map[uint32]core.Food
}

func NewDB() *DB {
	return &DB{
		food: make(map[uint32]core.Food),
	}
}

func (db *DB) Food(id uint32) (core.Food, error) {
	if food, ok := db.food[id]; ok {
		return food, nil
	}
	return core.Food{}, app.ErrNotFound
}

func (db *DB) Foods() ([]core.Food, error) {
	panic("not implemented")
}

func (db *DB) NewFood() (uint32, error) {
	panic("not implemented")
}

func (db *DB) SetFood(core.Food) error {
	panic("not implemented")
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
