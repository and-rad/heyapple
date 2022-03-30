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

package mem

import (
	"sort"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/core"
)

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
