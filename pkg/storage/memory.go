package storage

import (
	"heyapple/pkg/core"
	"sync"
)

type DB struct {
	mtx  sync.RWMutex
	food map[uint32]core.Food
}

func NewDB() *DB {
	return &DB{
		food: make(map[uint32]core.Food),
	}
}

func (db *DB) Food(uint32) (core.Food, error) {
	panic("not imlemented")
}

func (db *DB) Foods() ([]core.Food, error) {
	panic("not implemented")
}

func (db *DB) NewFood() (uint32, error) {
	panic("not implemented")
}

func (db *DB) SetFood(core.Food) error {
	panic("not implemented")
}

func (db *DB) Set(c core.Command) error {
	db.mtx.Lock()
	defer db.mtx.Unlock()
	return c.Execute(db)
}

func (db *DB) Get(q core.Query) error {
	db.mtx.RLock()
	defer db.mtx.RUnlock()
	return q.Fetch(db)
}
