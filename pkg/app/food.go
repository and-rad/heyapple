// Package app contains the core functionality of the aplication. Commands
// and queries that operate on food, recipes, and meals can be found here.
package app

import (
	"heyapple/pkg/core"
	"reflect"
)

// CreateFood is a command to create a new food item in the
// food database. If successful, the new item id is stored
// in the command.
type CreateFood struct {
	ID uint32
}

func (c *CreateFood) Execute(db core.DB) error {
	if id, err := db.NewFood(); err != nil {
		return err
	} else {
		c.ID = id
	}

	return nil
}

// SaveFood is a command that changes the specified values
// of a food item identified by ID.
type SaveFood struct {
	Data map[string]float32
	ID   uint32
}

func (c *SaveFood) Execute(db core.DB) error {
	food, err := db.Food(c.ID)
	if err != nil {
		return err
	}

	foodType := reflect.TypeOf(food)
	foodVal := reflect.ValueOf(&food).Elem()
	for i := 0; i < foodType.NumField(); i++ {
		tag := foodType.Field(i).Tag.Get("json")
		if v, ok := c.Data[tag]; ok {
			foodVal.Field(i).SetFloat(float64(v))
		}
	}

	return db.SetFood(food)
}

// GetFood is a query to retrieve a single food item from
// the food database. The item's ID is expected to be set
// before the query is executed.
type GetFood struct {
	Item core.Food
}

func (q *GetFood) Fetch(db core.DB) error {
	if food, err := db.Food(q.Item.ID); err != nil {
		return err
	} else {
		q.Item = food
	}

	return nil
}

// GetFoods is a query to retrieve all food items from
// the food database.
type GetFoods struct {
	Items []core.Food
}

func (q *GetFoods) Fetch(db core.DB) error {
	if foods, err := db.Foods(); err != nil {
		return err
	} else {
		q.Items = foods
	}

	return nil
}
