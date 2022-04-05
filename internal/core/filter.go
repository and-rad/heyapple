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

package core

import (
	"reflect"
	"strings"
)

var (
	FoodType = reflect.TypeOf(Food{})
	FoodMin  = reflect.ValueOf(Food{})
	FoodMax  = reflect.ValueOf(Food{
		KCal:     900,
		Fat:      100,
		FatSat:   83,
		FatO3:    54,
		FatO6:    70,
		Carbs:    100,
		Sugar:    100,
		Fructose: 56,
		Glucose:  36,
		Sucrose:  100,
		Protein:  89,
		Fiber:    71,
		Salt:     100,
	})

	foodFieldsByTag = getFieldsByTag(FoodType)
)

var (
	RecipeType     = reflect.TypeOf(Recipe{})
	recFieldsByTag = getFieldsByTag(RecipeType)
)

type FloatRange [2]float32
type IntRange [2]int
type Filter map[string]interface{}

func (f Filter) MatchFood(food Food) bool {
	if len(f) == 0 {
		return true
	}

	value := reflect.ValueOf(food)
	for k, v := range f {
		if k == "flags" {
			if flag, ok := v.(int); ok {
				if flag&food.Flags != 0 {
					continue
				}
			}
			return false
		}

		field := value.FieldByName(foodFieldsByTag[k])
		val := float32(field.Float())
		if r, ok := v.(FloatRange); ok {
			if val < r[0] || r[1] < val {
				return false
			}
		} else if val != v {
			return false
		}
	}

	return true
}

func (f Filter) MatchRecipe(rec Recipe) bool {
	if len(f) == 0 {
		return true
	}

	value := reflect.ValueOf(rec)
	for k, v := range f {
		field := value.FieldByName(recFieldsByTag[k])

		if field.Kind() == reflect.String {
			val := field.String()
			if name, ok := v.(string); ok {
				if !strings.Contains(val, name) {
					return false
				}
			}
		} else if field.Kind() == reflect.Int {
			val := int(field.Int())
			if r, ok := v.(IntRange); ok {
				if val < r[0] || r[1] < val {
					return false
				}
			} else if val != v {
				return false
			}
		} else if field.Kind() == reflect.Float32 {
			val := float32(field.Float())
			if r, ok := v.(FloatRange); ok {
				if val < r[0] || r[1] < val {
					return false
				}
			} else if val != v {
				return false
			}
		}
	}

	return true
}

func getFieldsByTag(t reflect.Type) map[string]string {
	tags := map[string]string{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tags[field.Tag.Get("json")] = field.Name
	}
	return tags
}
