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
	"github.com/and-rad/heyapple/internal/core"
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

// RecipeInstructions is a query that fetches preparation
// instructions for a given recipe. If successful, the text
// is stored in the query.
type RecipeInstructions struct {
	RecID        int    `json:"id"`
	Instructions string `json:"inst"`
}

func (q *RecipeInstructions) Fetch(db DB) error {
	if q.RecID == 0 {
		return ErrNotFound
	}

	if inst, err := db.RecipeInstructions(q.RecID); err != nil {
		return err
	} else {
		q.Instructions = inst
	}

	return nil
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

// Recipe is a query to retrieve a single recipe from
// the food database. The item's ID is expected to be set
// before the query is executed.
type Recipe struct {
	Item core.Recipe
}

func (q *Recipe) Fetch(db DB) error {
	if q.Item.ID == 0 {
		return ErrNotFound
	}

	if rec, err := db.Recipe(q.Item.ID); err != nil {
		return err
	} else {
		q.Item = rec
	}

	return nil
}

// SaveIngredient is a command to make changes to a recipe's
// list of ingredients. If Amount is 0, the ingredient will
// be removed from the recipe, otherwise it will be added
// or updated. If Replace is true, it will replace the current
// amount, otherwise it will add to it. In any case, the
// recipe's cached nutrients will be recalculated after the
// operation.
type SaveIngredient struct {
	RecipeID     int
	IngredientID int
	Amount       float32
	Replace      bool
}

func (c *SaveIngredient) Execute(db DB) error {
	rec, err := db.Recipe(c.RecipeID)
	if err != nil {
		return err
	}

	rec = core.Recipe{
		ID:       rec.ID,
		Name:     rec.Name,
		Size:     rec.Size,
		PrepTime: rec.PrepTime,
		CookTime: rec.CookTime,
		MiscTime: rec.MiscTime,
		Items:    rec.Items,
	}

	for k, v := range rec.Items {
		if v.ID == c.IngredientID {
			if !c.Replace {
				c.Amount += v.Amount
			}
			rec.Items = append(rec.Items[:k], rec.Items[k+1:]...)
			break
		}
	}

	if c.Amount > 0 {
		if ok, _ := db.FoodExists(c.IngredientID); ok {
			rec.Items = append(rec.Items, core.Ingredient{
				ID:     c.IngredientID,
				Amount: c.Amount},
			)
		}
	}

	for _, item := range rec.Items {
		if f, err := db.Food(item.ID); err == nil {
			amount := item.Amount * 0.01
			rec.Flags |= f.Flags

			rec.KCal += amount * f.KCal
			rec.Fat += amount * f.Fat
			rec.FatSat += amount * f.FatSat
			rec.FatO3 += amount * f.FatO3
			rec.FatO6 += amount * f.FatO6
			rec.Carbs += amount * f.Carbs
			rec.Sugar += amount * f.Sugar
			rec.Fructose += amount * f.Fructose
			rec.Glucose += amount * f.Glucose
			rec.Sucrose += amount * f.Sucrose
			rec.Fiber += amount * f.Fiber
			rec.Protein += amount * f.Protein
			rec.Salt += amount * f.Salt

			rec.Potassium += amount * f.Potassium
			rec.Chlorine += amount * f.Chlorine
			rec.Sodium += amount * f.Sodium
			rec.Calcium += amount * f.Calcium
			rec.Phosphorus += amount * f.Phosphorus
			rec.Magnesium += amount * f.Magnesium
			rec.Iron += amount * f.Iron
			rec.Zinc += amount * f.Zinc
			rec.Manganse += amount * f.Manganse
			rec.Copper += amount * f.Copper
			rec.Iodine += amount * f.Iodine
			rec.Chromium += amount * f.Chromium
			rec.Molybdenum += amount * f.Molybdenum
			rec.Selenium += amount * f.Selenium

			rec.VitA += amount * f.VitA
			rec.VitB1 += amount * f.VitB1
			rec.VitB2 += amount * f.VitB2
			rec.VitB3 += amount * f.VitB3
			rec.VitB5 += amount * f.VitB5
			rec.VitB6 += amount * f.VitB6
			rec.VitB7 += amount * f.VitB7
			rec.VitB9 += amount * f.VitB9
			rec.VitB12 += amount * f.VitB12
			rec.VitC += amount * f.VitC
			rec.VitD += amount * f.VitD
			rec.VitE += amount * f.VitE
			rec.VitK += amount * f.VitK
		}
	}

	if rec.Flags&core.FlagAnimal != 0 {
		rec.Flags &^= core.FlagVegan
	}

	if rec.Flags&core.FlagMeat != 0 {
		rec.Flags &^= core.FlagVegetarian
	}

	return db.SetRecipe(rec)
}
