package core

import "reflect"

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

	foodFieldsByTag = getFoodFieldsByTag()
)

type FloatRange [2]float32
type Filter map[string]interface{}

func (f Filter) MatchFood(food Food) bool {
	if len(f) == 0 {
		return true
	}

	for k, v := range f {
		val := float32(reflect.ValueOf(food).
			FieldByName(foodFieldsByTag[k]).
			Float(),
		)
		if r, ok := v.(FloatRange); ok {
			if val < r[0] || r[1] < val {
				return false
			}
		} else {
			if val != v {
				return false
			}
		}
	}
	return true
}

func getFoodFieldsByTag() map[string]string {
	tags := map[string]string{}
	for i := 0; i < FoodType.NumField(); i++ {
		field := FoodType.Field(i)
		tags[field.Tag.Get("json")] = field.Name
	}
	return tags
}
