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

package core_test

import (
	"heyapple/internal/mock"
	"heyapple/pkg/core"
	"reflect"
	"testing"
)

func TestFoodMin(t *testing.T) {
	values := map[string]interface{}{
		"kcal":   0.0,
		"fat":    0.0,
		"fatsat": 0.0,
		"fato3":  0.0,
		"fato6":  0.0,
		"carb":   0.0,
		"sug":    0.0,
		"prot":   0.0,
		"fib":    0.0,
		"salt":   0.0,
	}

	for i := 0; i < core.FoodType.NumField(); i++ {
		field := core.FoodType.Field(i)
		tag := field.Tag.Get("json")

		var val interface{}
		switch field.Type.Kind() {
		case reflect.Float32, reflect.Float64:
			val = core.FoodMin.Field(i).Float()
		}

		if v, ok := values[tag]; ok && v != val {
			t.Errorf("test case %d: value mismatch \nhave: %v\nwant: %v", i, val, v)
		}
	}
}

func TestFoodMax(t *testing.T) {
	values := map[string]interface{}{
		"kcal":   900.0,
		"fat":    100.0,
		"fatsat": 83.0,
		"fato3":  54.0,
		"fato6":  70.0,
		"carb":   100.0,
		"sug":    100.0,
		"prot":   89.0,
		"fib":    71.0,
		"salt":   100.0,
	}

	for i := 0; i < core.FoodType.NumField(); i++ {
		field := core.FoodType.Field(i)
		tag := field.Tag.Get("json")

		var val interface{}
		switch field.Type.Kind() {
		case reflect.Float32, reflect.Float64:
			val = core.FoodMax.Field(i).Float()
		}

		if v, ok := values[tag]; ok && v != val {
			t.Errorf("test case %d: value mismatch \nhave: %v\nwant: %v", i, val, v)
		}
	}
}

func TestFilter_MatchFood(t *testing.T) {
	for idx, data := range []struct {
		filter core.Filter
		food   core.Food

		ok bool
	}{
		{ //00// empty filter, always true
			food: mock.Food1,
			ok:   true,
		},
		{ //01// fail exact value
			filter: core.Filter{"kcal": float32(123)},
			food:   mock.Food1,
			ok:     false,
		},
		{ //02// match exact value
			filter: core.Filter{"kcal": mock.Food1.KCal},
			food:   mock.Food1,
			ok:     true,
		},
		{ //03// fail range
			filter: core.Filter{"fib": core.FloatRange{4, 9}},
			food:   mock.Food1,
			ok:     false,
		},
		{ //04// match range
			filter: core.Filter{"fib": core.FloatRange{0, 100}},
			food:   mock.Food1,
			ok:     true,
		},
	} {
		if ok := data.filter.MatchFood(data.food); ok != data.ok {
			t.Errorf("test case %d: result mismatch \nhave: %v\nwant: %v", idx, ok, data.ok)
		}
	}
}

func TestFilter_MatchRecipe(t *testing.T) {
	for idx, data := range []struct {
		filter core.Filter
		rec    core.Recipe

		ok bool
	}{
		{ //00// empty filter, always true
			rec: mock.Recipe1(),
			ok:  true,
		},
		{ //01// fail exact value
			filter: core.Filter{"name": "Brownie"},
			rec:    mock.Recipe1(),
			ok:     false,
		},
		{ //02// match exact value
			filter: core.Filter{"name": mock.Recipe1().Name},
			rec:    mock.Recipe1(),
			ok:     true,
		},
		{ //03// fail exact value
			filter: core.Filter{"size": 4},
			rec:    mock.Recipe1(),
			ok:     false,
		},
		{ //04// fail int range
			filter: core.Filter{"size": core.IntRange{4, 6}},
			rec:    mock.Recipe1(),
			ok:     false,
		},
		{ //05// match int range
			filter: core.Filter{"size": core.IntRange{1, 6}},
			rec:    mock.Recipe1(),
			ok:     true,
		},
		{ //06// fail exact value
			filter: core.Filter{"kcal": 4.4},
			rec:    mock.Recipe1(),
			ok:     false,
		},
		{ //07// fail float range
			filter: core.Filter{"kcal": core.FloatRange{150, 180}},
			rec:    mock.Recipe1(),
			ok:     false,
		},
		{ //08// match float range
			filter: core.Filter{"kcal": core.FloatRange{40, 150}},
			rec:    mock.Recipe1(),
			ok:     true,
		},
	} {
		if ok := data.filter.MatchRecipe(data.rec); ok != data.ok {
			t.Errorf("test case %d: result mismatch \nhave: %v\nwant: %v", idx, ok, data.ok)
		}
	}
}
