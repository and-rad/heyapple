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

package app

import (
	"heyapple/pkg/core"
)

// CreateRecipe is a command to create a new named recipe
// in the food database. If successful, the new item id is
// stored in the command.
type CreateRecipe struct {
	Name string
	ID   int
}

func (c *CreateRecipe) Execute(db DB) error {
	if id, err := db.NewRecipe(c.Name); err != nil {
		return err
	} else {
		c.ID = id
	}

	return nil
}

// SaveRecipe is a command that upates a recipe's number
// of servings or its list of ingredients.
type SaveRecipe struct {
	Items map[int]float32
	Size  int
	ID    int
}

func (c *SaveRecipe) Execute(db DB) error {
	if c.ID == 0 {
		return ErrNotFound
	}

	rec, err := db.Recipe(c.ID)
	if err != nil {
		return err
	}

	if c.Size > 0 {
		rec.Size = c.Size
	}

	rec.Items = []core.Ingredient{}
	for k, v := range c.Items {
		if ok, err := db.FoodExists(k); ok && err == nil {
			rec.Items = append(rec.Items, core.Ingredient{
				ID:     k,
				Amount: v,
			})
		}
	}

	return db.SetRecipe(rec)
}