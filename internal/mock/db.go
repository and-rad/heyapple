package mock

import (
	"errors"
	"heyapple/pkg/core"
)

// Error definitions
var (
	ErrNotFound = errors.New("not found")
	ErrDOS      = errors.New("dos")
)

var (
	nilFood = core.Food{}
)

/*
type Food struct {
	ID      uint32
	KCal    float32
	Fat     float32
	Carbs   float32
	Protein float32
}
*/

type DB struct {
	FoodInfo []core.Food

	LastFoodID uint32

	FailFood bool
}

func (db *DB) Food(id uint32) (core.Food, error) {
	if db.FailFood {
		return nilFood, ErrDOS
	}

	for _, f := range db.FoodInfo {
		if f.ID == id {
			return f, nil
		}
	}

	return nilFood, ErrNotFound
}

func (db *DB) Foods() ([]core.Food, error) {
	if db.FailFood {
		return []core.Food{}, ErrDOS
	}
	return db.FoodInfo, nil
}

func (db *DB) NewFood() (uint32, error) {
	if db.FailFood {
		return 0, ErrDOS
	}

	db.LastFoodID++
	db.FoodInfo = append(db.FoodInfo, core.Food{ID: db.LastFoodID})

	return db.LastFoodID, nil
}

func (db *DB) SetFood(food core.Food) error {
	if db.FailFood {
		return ErrDOS
	}

	for i, f := range db.FoodInfo {
		if f.ID == food.ID {
			db.FoodInfo[i] = food
			return nil
		}
	}

	return ErrNotFound
}
