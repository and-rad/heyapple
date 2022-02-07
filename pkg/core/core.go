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

// Package core defines core structs and interfaces.
package core

import "time"

// Food represents an edible object. All nutrients are
// stored per 100 base units. The base unit is either gram
// or milliliter. The actual unit of measurement for
// nutrients is grams.
type Food struct {
	ID uint32 `json:"id"`

	KCal    float32 `json:"cal"`
	Fat     float32 `json:"fat"`
	Carbs   float32 `json:"carbs"`
	Protein float32 `json:"protein"`
}

// Ingredient represents a single ingredient in a recipe.
// It's a combination of a food item and how much of that
// food is used, measured in grams.
type Ingredient struct {
	ID     uint32  `json:"id"`
	Amount float32 `json:"amount"`
}

// Recipe is a collection of ingredients.
type Recipe struct {
	Items []Ingredient `json:"items"`
	ID    uint32       `json:"id"`
}

// Meal is a collection of food items that were consumed at
// a specific point in time. The food can be provided by a
// recipe, but this is not required.
type Meal struct {
	Date   time.Time    `json:"date"`
	Items  []Ingredient `json:"items"`
	Recipe uint32       `json:"recipe"`
}

// DB defines an interface that provides access to the
// food data for reading and writing.
type DB interface {
	Food(uint32) (Food, error)
	Foods() ([]Food, error)
	NewFood() (uint32, error)
	SetFood(Food) error
}

// A Command encapsulates a single action that changes the
// underlying data. It can carry input and output parameters.
type Command interface {
	Execute(db DB) error
}

// A Query encapsulates a single read action on the underlying
// data. It should not make any changes to the data. It can
// carry input and output parameters.
type Query interface {
	Fetch(db DB) error
}
