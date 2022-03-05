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
		{ //00// fail exact value
			filter: core.Filter{"kcal": float32(123)},
			food:   mock.Food1,
			ok:     false,
		},
		{ //00// match exact value
			filter: core.Filter{"kcal": mock.Food1.KCal},
			food:   mock.Food1,
			ok:     true,
		},
		{ //00// fail range
			filter: core.Filter{"fib": core.FloatRange{4, 9}},
			food:   mock.Food1,
			ok:     false,
		},
		{ //00// match range
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
