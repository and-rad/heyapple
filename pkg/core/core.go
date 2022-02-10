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

import (
	"time"
)

// Food represents an edible object. All nutrients are
// stored per 100 base units. The base unit is either gram
// or milliliter. The actual unit of measurement for
// nutrients is grams.
type Food struct {
	ID    uint32 `json:"id"`
	Brand uint32 `json:"brand"`

	KCal    float32 `json:"kcal"`
	Fat     float32 `json:"fat"`
	FatSat  float32 `json:"fatsat"`
	FatO3   float32 `json:"fato3"`
	FatO6   float32 `json:"fato6"`
	Carbs   float32 `json:"carb"`
	Sugar   float32 `json:"sug"`
	Protein float32 `json:"prot"`
	Fiber   float32 `json:"fib"`

	Potassium  float32 `json:"pot"`
	Chlorine   float32 `json:"chl"`
	Sodium     float32 `json:"sod"`
	Calcium    float32 `json:"calc"`
	Phosphorus float32 `json:"phos"`
	Magnesium  float32 `json:"mag"`
	Iron       float32 `json:"iron"`
	Zinc       float32 `json:"zinc"`
	Manganse   float32 `json:"mang"`
	Copper     float32 `json:"cop"`
	Iodine     float32 `json:"iod"`
	Chromium   float32 `json:"chr"`
	Molybdenum float32 `json:"mol"`
	Selenium   float32 `json:"sel"`

	VitA   float32 `json:"vita"`
	VitB1  float32 `json:"vitb1"`
	VitB2  float32 `json:"vitb2"`
	VitB3  float32 `json:"vitb3"`
	VitB5  float32 `json:"vitb5"`
	VitB6  float32 `json:"vitb6"`
	VitB7  float32 `json:"vitb7"`
	VitB9  float32 `json:"vitb9"`
	VitB12 float32 `json:"vitb12"`
	VitC   float32 `json:"vitc"`
	VitD   float32 `json:"vitd"`
	VitE   float32 `json:"vite"`
	VitK   float32 `json:"vitk"`
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
