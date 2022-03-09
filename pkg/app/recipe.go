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
// in the food database. If successful, the new item is
// stored in the command.
type CreateRecipe struct {
	Name   string
	Recipe core.Recipe
}

func (c *CreateRecipe) Execute(db DB) error {
	if id, err := db.NewRecipe(c.Name); err != nil {
		return err
	} else {
		c.Recipe = core.NewRecipe(id)
		c.Recipe.Name = c.Name
	}

	return nil
}

// SaveRecipe is a command that upates a recipe's data,
// including the name and the number of servings but
// never the list of ingredients.
type SaveRecipe struct {
	Data map[string]interface{}
	ID   int
}

func (c *SaveRecipe) Execute(db DB) error {
	if c.ID == 0 {
		return ErrNotFound
	}

	rec, err := db.Recipe(c.ID)
	if err != nil {
		return err
	}

	if name, ok := c.Data["name"].(string); ok && name != "" {
		rec.Name = name
	}

	if size, ok := c.Data["size"].(int); ok && size > 0 {
		rec.Size = size
	}

	if time, ok := c.Data["preptime"].(int); ok && time >= 0 {
		rec.PrepTime = time
	}

	if time, ok := c.Data["cooktime"].(int); ok && time >= 0 {
		rec.CookTime = time
	}

	if time, ok := c.Data["misctime"].(int); ok && time >= 0 {
		rec.MiscTime = time
	}

	return db.SetRecipe(rec)
}

// SaveRecipeAccess is a command that changes a user's
// access rights for a given recipe.
type SaveRecipeAccess struct {
	UserID     int
	RecID      int
	Permission int
}

func (c *SaveRecipeAccess) Execute(db DB) error {
	if c.UserID == 0 || c.RecID == 0 {
		return ErrNotFound
	}
	if c.Permission < 0x00000100 || 0x0000ff00 < c.Permission {
		return ErrPermission
	}
	return db.SetRecipeAccess(c.UserID, c.RecID, c.Permission)
}

// RecipeAccess is a query that fetches a user's access
// rights for a given recipe. If successful, the permission
// value is stored in the query.
type RecipeAccess struct {
	UserID     int
	RecID      int
	Permission int
}

func (q *RecipeAccess) Fetch(db DB) error {
	if q.UserID == 0 || q.RecID == 0 {
		return ErrNotFound
	}

	if perm, err := db.RecipeAccess(q.UserID, q.RecID); err != nil {
		return err
	} else {
		q.Permission = perm
	}

	return nil
}

func (q *RecipeAccess) HasPerms(perms int) bool {
	return q.Permission&perms == perms
}

// Recipes is a query to retrieve all recipes from
// the food database.
type Recipes struct {
	Filter core.Filter
	Items  []core.Recipe
	UserID int
}

func (q *Recipes) Fetch(db DB) error {
	if q.UserID == 0 {
		return ErrNotFound
	}

	if recs, err := db.Recipes(q.UserID, q.Filter); err != nil {
		return err
	} else {
		q.Items = recs
	}

	return nil
}
