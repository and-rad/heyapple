package core

import (
	"reflect"
	"strings"
)

var (
	FoodType = reflect.TypeOf(Food{})
	FoodMin  = reflect.ValueOf(Food{})
	FoodMax  = reflect.ValueOf(Food{
		KCal:    900,
		Fat:     100,
		FatSat:  83,
		FatO3:   54,
		FatO6:   70,
		Carbs:   100,
		Sugar:   100,
		Protein: 89,
		Fiber:   71,
		Salt:    100,
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
