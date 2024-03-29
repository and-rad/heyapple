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

// Package core defines core structs and interfaces.
package core

import (
	"time"
)

const (
	FlagNone       = 0x00000000
	FlagVegan      = 0x00000001
	FlagVegetarian = 0x00000002
	FlagEgg        = 0x00000004
	FlagDairy      = 0x00000008
	FlagFish       = 0x00000010
	FlagPoultry    = 0x00000020
	FlagPork       = 0x00000040
	FlagBeef       = 0x00000080
	FlagShellfish  = 0x00000100
	FlagNut        = 0x00000200
	FlagPeanut     = 0x00000400
	FlagWheat      = 0x00000800
	FlagSoy        = 0x00001000

	FlagMeat   = FlagFish | FlagPork | FlagPoultry | FlagBeef | FlagShellfish
	FlagAnimal = FlagMeat | FlagEgg | FlagDairy
)

type Aisle int

const (
	AisleNone Aisle = iota
	AisleProduce
	AisleSpice
	AisleBread
	AisleDairy
	AisleMeat
	AisleFish
	AisleFrozen
	AislePasta
	AisleSauce
	AisleCanned
	AisleOil
	AisleBaking
	AisleSnacks
	AisleDeli
	AisleIntl
	AisleDrink
	AisleCleaning
	AisleHealth
	AisleBaby
	AisleVegan
)

type Category int

const (
	CatNone Category = iota
	CatDairy
	CatSpice
	CatBaby
	CatFat
	CatPoultry
	CatSoup
	CatSausage
	CatCereal
	CatFruit
	CatPork
	CatVegetable
	CatNut
	CatBeef
	CatDrink
	CatFish
	CatLegume
	CatLamb
	CatBaked
	CatSweets
	CatGrains
	CatFastFood
	CatMeals
	CatSnacks
	CatRestaurant
	CatAlcohol
	CatVegan
)

// Food represents an edible object. All nutrients are
// stored per 100 base units. The base unit is either gram
// or milliliter. The actual unit of measurement for
// macronutrients is grams while micronutrients are stored
// in milligrams for a smaller memory footprint.
type Food struct {
	ID int `json:"id"`

	Brand    int      `json:"brand"`
	Category Category `json:"cat"`
	Flags    int      `json:"flags"`

	KCal     float32 `json:"kcal"`
	Fat      float32 `json:"fat"`
	FatSat   float32 `json:"fatsat"`
	FatO3    float32 `json:"fato3"`
	FatO6    float32 `json:"fato6"`
	Carbs    float32 `json:"carb"`
	Sugar    float32 `json:"sug"`
	Fructose float32 `json:"fruc"`
	Glucose  float32 `json:"gluc"`
	Sucrose  float32 `json:"suc"`
	Fiber    float32 `json:"fib"`
	Protein  float32 `json:"prot"`
	Salt     float32 `json:"salt"`

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
	ID     int     `json:"id"`
	Amount float32 `json:"amount"`
}

// Recipe is a named collection of ingredients. Some data
// is entered manually, like the name, number of servings,
// and preparation instructions, while some is computed
// automatically, like nutrients and flags.
type Recipe struct {
	Name  string       `json:"name"`
	Items []Ingredient `json:"items"`

	ID       int `json:"id"`
	Size     int `json:"size"`
	Flags    int `json:"flags"`
	PrepTime int `json:"preptime"`
	CookTime int `json:"cooktime"`
	MiscTime int `json:"misctime"`

	KCal     float32 `json:"kcal"`
	Fat      float32 `json:"fat"`
	FatSat   float32 `json:"fatsat"`
	FatO3    float32 `json:"fato3"`
	FatO6    float32 `json:"fato6"`
	Carbs    float32 `json:"carb"`
	Sugar    float32 `json:"sug"`
	Fructose float32 `json:"fruc"`
	Glucose  float32 `json:"gluc"`
	Sucrose  float32 `json:"suc"`
	Fiber    float32 `json:"fib"`
	Protein  float32 `json:"prot"`
	Salt     float32 `json:"salt"`

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

func NewRecipe(id int) Recipe {
	return Recipe{ID: id, Size: 1, Items: []Ingredient{}}
}

// DiaryEntry represents a food item that was consumed at
// a specific point in time. The food can be provided by a
// recipe, but this is not required.
//
// The recipe is stored by name, not by id. The reason is
// persistence: Recipes can be renamed, repurposed, deleted.
// Storing just the id would be less useful to a user in
// these situations. The name provides the most useful
// information even after the original recipe is gone.
type DiaryEntry struct {
	Date   time.Time  `json:"date"`
	Recipe string     `json:"recipe"`
	Food   Ingredient `json:"food"`
}

func (d DiaryEntry) Equal(entry DiaryEntry) bool {
	if d.Food.ID != entry.Food.ID {
		return false
	}

	if d.Recipe != entry.Recipe {
		return false
	}

	if d.Date.After(entry.Date) {
		return d.Date.Sub(entry.Date) < time.Second*150
	}

	if d.Date.Before(entry.Date) {
		return entry.Date.Sub(d.Date) < time.Second*150
	}

	return true
}

func (d DiaryEntry) Day() time.Time {
	return d.Date.Truncate(time.Hour * 24)
}

// DiaryDay represents a summary of the diary entries
// for a given day. It acts as a cache of all the nutrient
// values taken from the entries and updated whenever
// the entries change.
type DiaryDay struct {
	Date string `json:"date"`

	KCal     float32 `json:"kcal"`
	Fat      float32 `json:"fat"`
	FatSat   float32 `json:"fatsat"`
	FatO3    float32 `json:"fato3"`
	FatO6    float32 `json:"fato6"`
	Carbs    float32 `json:"carb"`
	Sugar    float32 `json:"sug"`
	Fructose float32 `json:"fruc"`
	Glucose  float32 `json:"gluc"`
	Sucrose  float32 `json:"suc"`
	Fiber    float32 `json:"fib"`
	Protein  float32 `json:"prot"`
	Salt     float32 `json:"salt"`

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

func (d DiaryDay) Empty() bool {
	return d == (DiaryDay{Date: d.Date})
}

// ShopItem represents an item on the shopping list.
// It can be associated with a food item, in which
// case the Name field will usually be empty. If it is
// associated with non-food groceries, the ID will
// usually be 0.
type ShopItem struct {
	Name   string     `json:"name,omitempty"`
	Price  [2]float32 `json:"price,omitempty"`
	Done   bool       `json:"done"`
	Amount float32    `json:"amount"`
	Aisle  Aisle      `json:"aisle"`
	ID     int        `json:"id"`
}
