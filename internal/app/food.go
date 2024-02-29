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
	"reflect"

	"github.com/and-rad/heyapple/internal/core"
)

// CreateFood is a command to create a new food item in the
// food database. If successful, the new item id is stored
// in the command.
type CreateFood struct {
	Food core.Food
}

func (c *CreateFood) Execute(db DB) error {
	id, err := db.NewFood()
	if err != nil {
		return err
	}

	c.Food.ID = id
	return nil
}

// SaveFood is a command that changes the specified values
// of a food item identified by ID.
type SaveFood struct {
	Data map[string]float32
	ID   int
}

func (c *SaveFood) Execute(db DB) error {
	if c.ID == 0 {
		return ErrNotFound
	}

	food, err := db.Food(c.ID)
	if err != nil {
		return err
	}

	foodType := reflect.TypeOf(food)
	foodVal := reflect.ValueOf(&food).Elem()
	for i := 0; i < foodType.NumField(); i++ {
		field := foodType.Field(i)
		if field.Type.Kind() != reflect.Float32 {
			continue
		}

		tag := field.Tag.Get("json")
		if v, ok := c.Data[tag]; ok {
			foodVal.Field(i).SetFloat(float64(v))
		}
	}

	return db.SetFood(food)
}

// GetFood is a query to retrieve a single food item from
// the food database. The item's ID is expected to be set
// before the query is executed.
type GetFood struct {
	Item core.Food
}

func (q *GetFood) Fetch(db DB) error {
	if q.Item.ID == 0 {
		return ErrNotFound
	}

	food, err := db.Food(q.Item.ID)
	if err != nil {
		return err
	}

	q.Item = food
	return nil
}

// GetFoods is a query to retrieve all food items from
// the food database.
type GetFoods struct {
	Filter core.Filter
	Items  []core.Food
}

func (q *GetFoods) Fetch(db DB) error {
	foods, err := db.Foods(q.Filter)
	if err != nil {
		return err
	}

	q.Items = foods
	return nil
}
