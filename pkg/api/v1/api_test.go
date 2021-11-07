package api_test

import "heyapple/pkg/core"

var (
	food1 = core.Food{ID: 1, KCal: 2, Fat: 3, Carbs: 4, Protein: 5}
	food2 = core.Food{ID: 2, KCal: 3, Fat: 4, Carbs: 5, Protein: 6}

	food1json = `{"id":1,"cal":2,"fat":3,"carbs":4,"protein":5}`
	food2json = `{"id":2,"cal":3,"fat":4,"carbs":5,"protein":6}`
)
