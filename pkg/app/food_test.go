package app_test

import (
	"heyapple/internal/mock"
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"reflect"
	"testing"
)

func TestCreateFood_Execute(t *testing.T) {
	for idx, data := range []struct {
		indb  *mock.DB
		outdb *mock.DB
		err   error
	}{
		{ //00//
			indb:  &mock.DB{FailFood: true},
			outdb: &mock.DB{FailFood: true},
			err:   mock.ErrDOS,
		},
		{ //01//
			indb: &mock.DB{},
			outdb: &mock.DB{
				FoodInfo:   []core.Food{{ID: 1}},
				LastFoodID: 1,
			},
		},
		{ //02//
			indb: &mock.DB{
				FoodInfo:   []core.Food{{ID: 1}, {ID: 2}},
				LastFoodID: 8,
			},
			outdb: &mock.DB{
				FoodInfo:   []core.Food{{ID: 1}, {ID: 2}, {ID: 9}},
				LastFoodID: 9,
			},
		},
	} {
		cmd := &app.CreateFood{}
		err := cmd.Execute(data.indb)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v \nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.indb, data.outdb) {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, data.indb, data.outdb)
		}

		if cmd.ID != data.outdb.LastFoodID {
			t.Errorf("test case %d: id mismatch \nhave: %v \nwant: %v", idx, cmd.ID, data.outdb.LastFoodID)
		}
	}
}

func TestSaveFood_Execute(t *testing.T) {
	for idx, data := range []struct {
		err   error
		indb  *mock.DB
		outdb *mock.DB
		data  map[string]float32
		id    uint32
	}{
		{ //00// connection failed
			indb:  &mock.DB{FailFood: true},
			outdb: &mock.DB{FailFood: true},
			err:   mock.ErrDOS,
		},
		{ //01// empty DB
			indb:  &mock.DB{},
			outdb: &mock.DB{},
			err:   mock.ErrNotFound,
		},
		{ //02// id not found
			indb:  &mock.DB{FoodInfo: []core.Food{food1}},
			outdb: &mock.DB{FoodInfo: []core.Food{food1}},
			id:    2,
			err:   mock.ErrNotFound,
		},
		{ //03// empty data, no changes
			indb:  &mock.DB{FoodInfo: []core.Food{food1}},
			outdb: &mock.DB{FoodInfo: []core.Food{food1}},
			id:    1,
		},
		{ //04// change some values
			indb: &mock.DB{FoodInfo: []core.Food{
				{ID: 1, KCal: 2, Fat: 3, Carbs: 4, Protein: 5},
			}},
			outdb: &mock.DB{FoodInfo: []core.Food{
				{ID: 1, KCal: 120, Fat: 3, Carbs: 33.3, Protein: 5},
			}},
			data: map[string]float32{"cal": 120, "carbs": 33.3},
			id:   1,
		},
		{ //05// change all values
			indb: &mock.DB{FoodInfo: []core.Food{
				{ID: 1, KCal: 2, Fat: 3, Carbs: 4, Protein: 5},
			}},
			outdb: &mock.DB{FoodInfo: []core.Food{
				{ID: 1, KCal: 100, Fat: 30, Carbs: 20, Protein: 50},
			}},
			data: map[string]float32{
				"cal":     100,
				"fat":     30,
				"carbs":   20,
				"protein": 50,
			},
			id: 1,
		},
	} {
		cmd := &app.SaveFood{ID: data.id, Data: data.data}
		err := cmd.Execute(data.indb)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v \nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.indb, data.outdb) {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, data.indb, data.outdb)
		}
	}
}

func TestGetFood_Fetch(t *testing.T) {
	for idx, data := range []struct {
		db  *mock.DB
		err error
		out core.Food
	}{
		{ //00// connection failed
			db:  &mock.DB{FailFood: true},
			err: mock.ErrDOS,
		},
		{ //01// empty db
			db:  &mock.DB{},
			out: core.Food{ID: 1},
			err: mock.ErrNotFound,
		},
		{ //02// id not found
			db:  &mock.DB{FoodInfo: []core.Food{{ID: 1}, {ID: 2}}},
			out: core.Food{ID: 3},
			err: mock.ErrNotFound,
		},
		{ //03// success
			db:  &mock.DB{FoodInfo: []core.Food{food1}},
			out: food1,
		},
	} {
		qry := &app.GetFood{Item: core.Food{ID: data.out.ID}}
		err := qry.Fetch(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v \nwant: %v", idx, err, data.err)
		}

		if qry.Item != data.out {
			t.Errorf("test case %d: id mismatch \nhave: %v \nwant: %v", idx, qry.Item, data.out)
		}
	}
}

func TestGetFoods_Fetch(t *testing.T) {
	for idx, data := range []struct {
		db  *mock.DB
		err error
		out []core.Food
	}{
		{ //00// connection failed
			db:  &mock.DB{FailFood: true},
			err: mock.ErrDOS,
		},
		{ //01// empty db
			db:  &mock.DB{FoodInfo: []core.Food{}},
			out: []core.Food{},
		},
		{ //02// success
			db:  &mock.DB{FoodInfo: []core.Food{food1, food2}},
			out: []core.Food{food1, food2},
		},
	} {
		qry := &app.GetFoods{}
		err := qry.Fetch(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v \nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(qry.Items, data.out) {
			t.Errorf("test case %d: id mismatch \nhave: %v \nwant: %v", idx, qry.Items, data.out)
		}
	}
}
