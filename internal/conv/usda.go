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

package conv

import (
	"encoding/json"
	"math"

	"github.com/and-rad/heyapple/internal/core"
)

var (
	usdaCats = map[string]core.Category{
		"Dairy and Egg Products":            core.CatDairy,
		"Spices and Herbs":                  core.CatSpice,
		"Baby Foods":                        core.CatBaby,
		"Fats and Oils":                     core.CatFat,
		"Poultry Products":                  core.CatPoultry,
		"Soups, Sauces, and Gravies":        core.CatSoup,
		"Sausages and Luncheon Meats":       core.CatSausage,
		"Breakfast Cereals":                 core.CatCereal,
		"Fruits and Fruit Juices":           core.CatFruit,
		"Pork Products":                     core.CatPork,
		"Vegetables and Vegetable Products": core.CatVegetable,
		"Nut and Seed Products":             core.CatNut,
		"Beef Products":                     core.CatBeef,
		"Beverages":                         core.CatDrink,
		"Finfish and Shellfish Products":    core.CatFish,
		"Legumes and Legume Products":       core.CatLegume,
		"Lamb, Veal, and Game Products":     core.CatLamb,
		"Baked Products":                    core.CatBaked,
		"Sweets":                            core.CatSweets,
		"Cereal Grains and Pasta":           core.CatGrains,
		"Fast Foods":                        core.CatFastFood,
		"Meals, Entrees, and Side Dishes":   core.CatMeals,
		"Snacks":                            core.CatSnacks,
		"Restaurant Foods":                  core.CatRestaurant,
		"Alcoholic Beverages":               core.CatAlcohol,
	}

	usdaAisles = map[core.Category]core.Aisle{
		core.CatDairy:     core.AisleDairy,
		core.CatSpice:     core.AisleSpice,
		core.CatBaby:      core.AisleBaby,
		core.CatFish:      core.AisleFish,
		core.CatFruit:     core.AisleProduce,
		core.CatVegetable: core.AisleProduce,
		core.CatPoultry:   core.AisleMeat,
		core.CatPork:      core.AisleMeat,
		core.CatBeef:      core.AisleMeat,
		core.CatLegume:    core.AislePasta,
		core.CatLamb:      core.AisleMeat,
		core.CatBaked:     core.AisleBread,
		core.CatSweets:    core.AisleSnacks,
		core.CatGrains:    core.AislePasta,
		core.CatSnacks:    core.AisleSnacks,
		core.CatDrink:     core.AisleDrink,
		core.CatAlcohol:   core.AisleDrink,
		core.CatFat:       core.AisleOil,
	}
)

type usdaNutrientData struct {
	Name string `json:"name"`
	Unit string `json:"unitName"`
	ID   int    `json:"id"`
}

type usdaNutrient struct {
	Data   usdaNutrientData `json:"nutrient"`
	Amount float32          `json:"amount"`
}

type usdaCategory struct {
	Desc string `json:"description"`
}

type usdaFood struct {
	Desc      string         `json:"description"`
	Category  usdaCategory   `json:"foodCategory"`
	Nutrients []usdaNutrient `json:"foodNutrients"`
}

type usda struct {
	Foods []usdaFood `json:"SRLegacyFoods"`
}

func FromUSDA(data []byte) ([]Food, error) {
	usda := usda{}
	if err := json.Unmarshal(data, &usda); err != nil {
		return nil, ErrFromJSON
	}

	foods := []Food{}
	for i, inFood := range usda.Foods {
		category := usdaCats[inFood.Category.Desc]
		outFood := Food{
			Name:  inFood.Desc,
			Aisle: usdaAisles[category],
			Food: core.Food{
				ID:       i + 1,
				Category: category,
			},
		}

		for _, n := range inFood.Nutrients {
			switch n.Data.ID {
			case 1003:
				outFood.Protein = usdaAmount(n, "g")
			case 1004:
				outFood.Fat = usdaAmount(n, "g")
			case 1005:
				outFood.Carbs = usdaAmount(n, "g")
			case 1008:
				outFood.KCal = usdaAmount(n, "kcal")
			case 1010:
				outFood.Sucrose = usdaAmount(n, "g")
			case 1011:
				outFood.Glucose = usdaAmount(n, "g")
			case 1012:
				outFood.Fructose = usdaAmount(n, "g")
			case 1079:
				outFood.Fiber = usdaAmount(n, "g")
			case 1087:
				outFood.Calcium = usdaAmount(n, "mg")
			case 1089:
				outFood.Iron = usdaAmount(n, "mg")
			case 1090:
				outFood.Magnesium = usdaAmount(n, "mg")
			case 1091:
				outFood.Phosphorus = usdaAmount(n, "mg")
			case 1092:
				outFood.Potassium = usdaAmount(n, "mg")
			case 1093:
				outFood.Sodium = usdaAmount(n, "mg")
				outFood.Salt = outFood.Sodium * 0.0025
			case 1095:
				outFood.Zinc = usdaAmount(n, "mg")
			case 1098:
				outFood.Copper = usdaAmount(n, "mg")
			case 1101:
				outFood.Manganse = usdaAmount(n, "mg")
			case 1103:
				outFood.Selenium = usdaAmount(n, "mg")
			case 1106:
				outFood.VitA = usdaAmount(n, "mg")
			case 1109:
				outFood.VitE = usdaAmount(n, "mg")
			case 1114:
				outFood.VitD = usdaAmount(n, "mg")
			case 1162:
				outFood.VitC = usdaAmount(n, "mg")
			case 1165:
				outFood.VitB1 = usdaAmount(n, "mg")
			case 1166:
				outFood.VitB2 = usdaAmount(n, "mg")
			case 1167:
				outFood.VitB3 = usdaAmount(n, "mg")
			case 1170:
				outFood.VitB5 = usdaAmount(n, "mg")
			case 1175:
				outFood.VitB6 = usdaAmount(n, "mg")
			case 1190:
				outFood.VitB9 = usdaAmount(n, "mg")
			case 1178:
				outFood.VitB12 = usdaAmount(n, "mg")
			case 1184, 1185:
				outFood.VitK += usdaAmount(n, "mg")
			case 1258:
				outFood.FatSat = usdaAmount(n, "g")
			case 1272, 1278, 1280, 1404:
				outFood.FatO3 += usdaAmount(n, "g")
			case 1313, 1316, 1321, 1406, 1408:
				outFood.FatO6 += usdaAmount(n, "g")
			case 2000:
				outFood.Sugar = usdaAmount(n, "g")
			}
		}

		foods = append(foods, outFood)
	}

	return foods, nil
}

func usdaAmount(n usdaNutrient, base string) float32 {
	if n.Data.Unit == base {
		return n.Amount
	}

	var i, j int
	for k, v := range []string{"g", "mg", "µg"} {
		if base == v {
			i = k
		}
		if n.Data.Unit == v {
			j = k
		}
	}

	return n.Amount * float32(math.Pow(1000, float64(i-j)))
}
