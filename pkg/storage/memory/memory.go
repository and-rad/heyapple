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

package memory

import (
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

func (db *DB) Food(uint32) (core.Food, error) {
	panic("not imlemented")
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

func (db *DB) Set(c core.Command) error {
	db.mtx.Lock()
	defer db.mtx.Unlock()
	return c.Execute(db)
}

func (db *DB) Get(q core.Query) error {
	db.mtx.RLock()
	defer db.mtx.RUnlock()
	return q.Fetch(db)
}
