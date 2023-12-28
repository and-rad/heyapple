////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2021-2024 The HeyApple Authors.
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

package mem

import (
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/core"
	"github.com/and-rad/heyapple/internal/mock"
)

func TestDB_NewDiaryEntries(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int
		in []core.DiaryEntry

		entries entryMap
		days    dayMap
		err     error
	}{
		{ //00// create new entry
			id: 1,
			db: &DB{
				food:    map[int]core.Food{1: mock.Food1, 2: mock.Food2},
				entries: entryMap{},
				days:    dayMap{},
			},
			in:      []core.DiaryEntry{mock.Entry1()},
			entries: entryMap{1: {mock.Day1: {mock.Entry1()}}},
			days: dayMap{1: {2022: {3: {{
				Date:       "2022-03-12",
				KCal:       mock.Food2.KCal * 1.5,
				Fat:        mock.Food2.Fat * 1.5,
				Carbs:      mock.Food2.Carbs * 1.5,
				Sugar:      mock.Food2.Sugar * 1.5,
				Fructose:   mock.Food2.Fructose * 1.5,
				Glucose:    mock.Food2.Glucose * 1.5,
				Sucrose:    mock.Food2.Sucrose * 1.5,
				Protein:    mock.Food2.Protein * 1.5,
				Fiber:      mock.Food2.Fiber * 1.5,
				Iron:       mock.Food2.Iron * 1.5,
				Zinc:       mock.Food2.Zinc * 1.5,
				Magnesium:  mock.Food2.Magnesium * 1.5,
				Chlorine:   mock.Food2.Chlorine * 1.5,
				Sodium:     mock.Food2.Sodium * 1.5,
				Calcium:    mock.Food2.Calcium * 1.5,
				Potassium:  mock.Food2.Potassium * 1.5,
				Phosphorus: mock.Food2.Phosphorus * 1.5,
				Copper:     mock.Food2.Copper * 1.5,
				Iodine:     mock.Food2.Iodine * 1.5,
				Manganse:   mock.Food2.Manganse * 1.5,
				VitA:       mock.Food2.VitA * 1.5,
				VitB1:      mock.Food2.VitB1 * 1.5,
				VitB2:      mock.Food2.VitB2 * 1.5,
				VitB6:      mock.Food2.VitB6 * 1.5,
				VitC:       mock.Food2.VitC * 1.5,
				VitE:       mock.Food2.VitE * 1.5,
			}}}}},
		},
		{ //01// make sure entries are sorted
			id: 1,
			db: &DB{
				food:    map[int]core.Food{1: mock.Food1, 2: mock.Food2},
				entries: entryMap{},
				days:    dayMap{},
			},
			in:      []core.DiaryEntry{mock.Entry2(), mock.Entry1()},
			entries: entryMap{1: {mock.Day1: {mock.Entry1(), mock.Entry2()}}},
			days:    dayMap{1: {2022: {3: {mock.Diary220312()}}}},
		},
	} {
		err := data.db.NewDiaryEntries(data.id, data.in...)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.db.entries, data.entries) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.entries, data.entries)
		}

		if !reflect.DeepEqual(data.db.days, data.days) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.days, data.days)
		}
	}
}

func TestDB_DelDiaryEntries(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int
		in []core.DiaryEntry

		entries entryMap
		days    dayMap
		err     error
	}{
		{ //00// diary doesn't exist
			id:      1,
			db:      NewDB(),
			entries: entryMap{},
			days:    dayMap{},
			err:     app.ErrNotFound,
		},
		{ //01// empty diary, nothing to do
			id:      1,
			db:      &DB{entries: entryMap{1: {}}},
			in:      []core.DiaryEntry{mock.Entry1()},
			entries: entryMap{1: {}},
		},
		{ //02// success, simple case
			id: 1,
			db: &DB{
				entries: entryMap{1: {mock.Day1: {mock.Entry1(), mock.Entry2()}}},
				days:    dayMap{1: {2022: {3: {mock.Diary220312()}}}},
			},
			in:      []core.DiaryEntry{mock.Entry2()},
			entries: entryMap{1: {mock.Day1: {mock.Entry1()}}},
			days:    dayMap{1: {2022: {3: {}}}},
		},
		{ //03// success, complex case
			id: 1,
			db: &DB{
				entries: entryMap{1: {
					mock.Day1: {mock.Entry1(), mock.Entry2()},
					mock.Day2: {mock.Entry3(), mock.Entry4()},
				}},
				days: dayMap{1: {2022: {3: {
					mock.Diary220312(),
					mock.Diary220313(),
				}}}},
			},
			in: []core.DiaryEntry{
				mock.Entry3(),
				mock.Entry4(),
				mock.Entry2(),
			},
			entries: entryMap{1: {
				mock.Day1: {mock.Entry1()},
				mock.Day2: {},
			}},
			days: dayMap{1: {2022: {3: {}}}},
		},
	} {
		err := data.db.DelDiaryEntries(data.id, data.in...)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.db.entries, data.entries) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.entries, data.entries)
		}

		if !reflect.DeepEqual(data.db.days, data.days) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.days, data.days)
		}
	}
}

func TestDB_SetDiaryEntries(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int
		in []core.DiaryEntry

		entries entryMap
		days    dayMap
		err     error
	}{
		{ //00// diary doesn't exist
			id:      1,
			db:      NewDB(),
			entries: entryMap{},
			days:    dayMap{},
			err:     app.ErrNotFound,
		},
		{ //01// empty diary, nothing to do
			id:      1,
			db:      &DB{entries: entryMap{1: {}}},
			in:      []core.DiaryEntry{mock.Entry1()},
			entries: entryMap{1: {}},
		},
		{ //02// success, simple case
			id: 1,
			db: &DB{
				food:    map[int]core.Food{1: mock.Food1, 2: mock.Food2},
				entries: entryMap{1: {mock.Day1: {mock.Entry1(), mock.Entry2()}}},
				days:    dayMap{1: {2022: {3: {mock.Diary220312()}}}},
			},
			in:      []core.DiaryEntry{mock.Entry1New()},
			entries: entryMap{1: {mock.Day1: {mock.Entry1New(), mock.Entry2()}}},
			days: dayMap{1: {2022: {3: {{
				Date:       "2022-03-12",
				KCal:       mock.Food1.KCal*0.5 + mock.Food2.KCal*0.75,
				Fat:        mock.Food1.Fat*0.5 + mock.Food2.Fat*0.75,
				Carbs:      mock.Food1.Carbs*0.5 + mock.Food2.Carbs*0.75,
				Sugar:      mock.Food1.Sugar*0.5 + mock.Food2.Sugar*0.75,
				Fructose:   mock.Food1.Fructose*0.5 + mock.Food2.Fructose*0.75,
				Glucose:    mock.Food1.Glucose*0.5 + mock.Food2.Glucose*0.75,
				Sucrose:    mock.Food1.Sucrose*0.5 + mock.Food2.Sucrose*0.75,
				Protein:    mock.Food1.Protein*0.5 + mock.Food2.Protein*0.75,
				Fiber:      mock.Food1.Fiber*0.5 + mock.Food2.Fiber*0.75,
				Iron:       mock.Food1.Iron*0.5 + mock.Food2.Iron*0.75,
				Zinc:       mock.Food1.Zinc*0.5 + mock.Food2.Zinc*0.75,
				Magnesium:  mock.Food1.Magnesium*0.5 + mock.Food2.Magnesium*0.75,
				Chlorine:   mock.Food1.Chlorine*0.5 + mock.Food2.Chlorine*0.75,
				Sodium:     mock.Food1.Sodium*0.5 + mock.Food2.Sodium*0.75,
				Calcium:    mock.Food1.Calcium*0.5 + mock.Food2.Calcium*0.75,
				Potassium:  mock.Food1.Potassium*0.5 + mock.Food2.Potassium*0.75,
				Phosphorus: mock.Food1.Phosphorus*0.5 + mock.Food2.Phosphorus*0.75,
				Copper:     mock.Food1.Copper*0.5 + mock.Food2.Copper*0.75,
				Iodine:     mock.Food1.Iodine*0.5 + mock.Food2.Iodine*0.75,
				Manganse:   mock.Food1.Manganse*0.5 + mock.Food2.Manganse*0.75,
				VitA:       mock.Food1.VitA*0.5 + mock.Food2.VitA*0.75,
				VitB1:      mock.Food1.VitB1*0.5 + mock.Food2.VitB1*0.75,
				VitB2:      mock.Food1.VitB2*0.5 + mock.Food2.VitB2*0.75,
				VitB6:      mock.Food1.VitB6*0.5 + mock.Food2.VitB6*0.75,
				VitC:       mock.Food1.VitC*0.5 + mock.Food2.VitC*0.75,
				VitE:       mock.Food1.VitE*0.5 + mock.Food2.VitE*0.75,
			}}}}},
		},
		{ //03// success, complex case
			id: 1,
			db: &DB{
				food: map[int]core.Food{
					1: mock.Food1,
					2: mock.Food2,
				},
				entries: entryMap{1: {
					mock.Day1: {mock.Entry1(), mock.Entry2()},
					mock.Day2: {mock.Entry3(), mock.Entry4()},
				}},
				days: dayMap{1: {2022: {3: {
					mock.Diary220312(),
					mock.Diary220313(),
				}}}},
			},
			in: []core.DiaryEntry{
				mock.Entry3New(),
				mock.Entry4New(),
				mock.Entry2New(),
			},
			entries: entryMap{1: {
				mock.Day1: {mock.Entry1(), mock.Entry2New()},
				mock.Day2: {mock.Entry3New(), mock.Entry4New()},
			}},
			days: dayMap{1: {2022: {3: {
				{
					Date:       "2022-03-12",
					KCal:       mock.Food1.KCal*0.9 + mock.Food2.KCal*1.5,
					Fat:        mock.Food1.Fat*0.9 + mock.Food2.Fat*1.5,
					Carbs:      mock.Food1.Carbs*0.9 + mock.Food2.Carbs*1.5,
					Sugar:      mock.Food1.Sugar*0.9 + mock.Food2.Sugar*1.5,
					Fructose:   mock.Food1.Fructose*0.9 + mock.Food2.Fructose*1.5,
					Glucose:    mock.Food1.Glucose*0.9 + mock.Food2.Glucose*1.5,
					Sucrose:    mock.Food1.Sucrose*0.9 + mock.Food2.Sucrose*1.5,
					Protein:    mock.Food1.Protein*0.9 + mock.Food2.Protein*1.5,
					Fiber:      mock.Food1.Fiber*0.9 + mock.Food2.Fiber*1.5,
					Iron:       mock.Food1.Iron*0.9 + mock.Food2.Iron*1.5,
					Zinc:       mock.Food1.Zinc*0.9 + mock.Food2.Zinc*1.5,
					Magnesium:  mock.Food1.Magnesium*0.9 + mock.Food2.Magnesium*1.5,
					Chlorine:   mock.Food1.Chlorine*0.9 + mock.Food2.Chlorine*1.5,
					Sodium:     mock.Food1.Sodium*0.9 + mock.Food2.Sodium*1.5,
					Calcium:    mock.Food1.Calcium*0.9 + mock.Food2.Calcium*1.5,
					Potassium:  mock.Food1.Potassium*0.9 + mock.Food2.Potassium*1.5,
					Phosphorus: mock.Food1.Phosphorus*0.9 + mock.Food2.Phosphorus*1.5,
					Copper:     mock.Food1.Copper*0.9 + mock.Food2.Copper*1.5,
					Iodine:     mock.Food1.Iodine*0.9 + mock.Food2.Iodine*1.5,
					Manganse:   mock.Food1.Manganse*0.9 + mock.Food2.Manganse*1.5,
					VitA:       mock.Food1.VitA*0.9 + mock.Food2.VitA*1.5,
					VitB1:      mock.Food1.VitB1*0.9 + mock.Food2.VitB1*1.5,
					VitB2:      mock.Food1.VitB2*0.9 + mock.Food2.VitB2*1.5,
					VitB6:      mock.Food1.VitB6*0.9 + mock.Food2.VitB6*1.5,
					VitC:       mock.Food1.VitC*0.9 + mock.Food2.VitC*1.5,
					VitE:       mock.Food1.VitE*0.9 + mock.Food2.VitE*1.5,
				},
				{
					Date:       "2022-03-13",
					KCal:       mock.Food1.KCal*5.2 + mock.Food2.KCal*1.5,
					Fat:        mock.Food1.Fat*5.2 + mock.Food2.Fat*1.5,
					Carbs:      mock.Food1.Carbs*5.2 + mock.Food2.Carbs*1.5,
					Sugar:      mock.Food1.Sugar*5.2 + mock.Food2.Sugar*1.5,
					Fructose:   mock.Food1.Fructose*5.2 + mock.Food2.Fructose*1.5,
					Glucose:    mock.Food1.Glucose*5.2 + mock.Food2.Glucose*1.5,
					Sucrose:    mock.Food1.Sucrose*5.2 + mock.Food2.Sucrose*1.5,
					Protein:    mock.Food1.Protein*5.2 + mock.Food2.Protein*1.5,
					Fiber:      mock.Food1.Fiber*5.2 + mock.Food2.Fiber*1.5,
					Iron:       mock.Food1.Iron*5.2 + mock.Food2.Iron*1.5,
					Zinc:       mock.Food1.Zinc*5.2 + mock.Food2.Zinc*1.5,
					Magnesium:  mock.Food1.Magnesium*5.2 + mock.Food2.Magnesium*1.5,
					Chlorine:   mock.Food1.Chlorine*5.2 + mock.Food2.Chlorine*1.5,
					Sodium:     mock.Food1.Sodium*5.2 + mock.Food2.Sodium*1.5,
					Calcium:    mock.Food1.Calcium*5.2 + mock.Food2.Calcium*1.5,
					Potassium:  mock.Food1.Potassium*5.2 + mock.Food2.Potassium*1.5,
					Phosphorus: mock.Food1.Phosphorus*5.2 + mock.Food2.Phosphorus*1.5,
					Copper:     mock.Food1.Copper*5.2 + mock.Food2.Copper*1.5,
					Iodine:     mock.Food1.Iodine*5.2 + mock.Food2.Iodine*1.5,
					Manganse:   mock.Food1.Manganse*5.2 + mock.Food2.Manganse*1.5,
					VitA:       mock.Food1.VitA*5.2 + mock.Food2.VitA*1.5,
					VitB1:      mock.Food1.VitB1*5.2 + mock.Food2.VitB1*1.5,
					VitB2:      mock.Food1.VitB2*5.2 + mock.Food2.VitB2*1.5,
					VitB6:      mock.Food1.VitB6*5.2 + mock.Food2.VitB6*1.5,
					VitC:       mock.Food1.VitC*5.2 + mock.Food2.VitC*1.5,
					VitE:       mock.Food1.VitE*5.2 + mock.Food2.VitE*1.5,
				},
			}}}},
		},
	} {
		err := data.db.SetDiaryEntries(data.id, data.in...)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.db.entries, data.entries) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.entries, data.entries)
		}

		if !reflect.DeepEqual(data.db.days, data.days) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.days, data.days)
		}
	}
}

func TestDB_DiaryEntries(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		id   int
		date time.Time

		entries []core.DiaryEntry
		err     error
	}{
		{ //00// diary doesn't exist
			id:  1,
			db:  NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// day doesn't exist
			id:   1,
			date: mock.Date1,
			db:   &DB{entries: entryMap{1: {}}},
			err:  app.ErrNotFound,
		},
		{ //02// success with fuzzy date match
			id:      1,
			date:    mock.Date1,
			db:      &DB{entries: entryMap{1: {mock.Day1: {mock.Entry1(), mock.Entry2()}}}},
			entries: []core.DiaryEntry{mock.Entry1(), mock.Entry2()},
		},
	} {
		entries, err := data.db.DiaryEntries(data.id, data.date)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(entries, data.entries) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, entries, data.entries)
		}
	}
}

func TestDB_DiaryEntries_Race(t *testing.T) {
	db := NewDB()
	db.entries = entryMap{1: {mock.Day1: {mock.Entry1()}}}

	e1, _ := db.DiaryEntries(1, mock.Day1)
	e2, _ := db.DiaryEntries(1, mock.Day1)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		e1[0].Recipe = "Rec 1"
		wg.Done()
	}()

	go func() {
		e2[0].Recipe = "Rec 1"
		wg.Done()
	}()

	wg.Wait()
}

func TestDB_DiaryDays(t *testing.T) {
	for idx, data := range []struct {
		db    *DB
		id    int
		year  int
		month int
		day   int

		days []core.DiaryDay
		err  error
	}{
		{ //00// diary doesn't exist
			id:  1,
			db:  NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// get everything
			id: 1,
			db: &DB{days: dayMap{1: {
				2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
				2022: {3: {mock.Diary220301(), mock.Diary220302()}},
			}}},
			days: []core.DiaryDay{mock.Diary210101(), mock.Diary210102(), mock.Diary210201(), mock.Diary210202(), mock.Diary220301(), mock.Diary220302()},
		},
		{ //02// no entries for specific year
			id:   1,
			year: 2020,
			db: &DB{days: dayMap{1: {
				2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
				2022: {3: {mock.Diary220301(), mock.Diary220302()}},
			}}},
			days: []core.DiaryDay{},
		},
		{ //03// get specific year
			id:   1,
			year: 2021,
			db: &DB{days: dayMap{1: {
				2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
				2022: {3: {mock.Diary220301(), mock.Diary220302()}},
			}}},
			days: []core.DiaryDay{mock.Diary210101(), mock.Diary210102(), mock.Diary210201(), mock.Diary210202()},
		},
		{ //04// no entries for specific month
			id:    1,
			year:  2021,
			month: 4,
			db: &DB{days: dayMap{1: {
				2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
				2022: {3: {mock.Diary220301(), mock.Diary220302()}},
			}}},
			days: []core.DiaryDay{},
		},
		{ //05// get specific month
			id:    1,
			year:  2021,
			month: 2,
			db: &DB{days: dayMap{1: {
				2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
				2022: {3: {mock.Diary220301(), mock.Diary220302()}},
			}}},
			days: []core.DiaryDay{mock.Diary210201(), mock.Diary210202()},
		},
		{ //06// no entries for specific day
			id:    1,
			year:  2021,
			month: 2,
			day:   12,
			db: &DB{days: dayMap{1: {
				2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
				2022: {3: {mock.Diary220301(), mock.Diary220302()}},
			}}},
			days: []core.DiaryDay{},
		},
		{ //07// get specific day
			id:    1,
			year:  2021,
			month: 2,
			day:   1,
			db: &DB{days: dayMap{1: {
				2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
				2022: {3: {mock.Diary220301(), mock.Diary220302()}},
			}}},
			days: []core.DiaryDay{mock.Diary210201()},
		},
	} {
		days, err := data.db.DiaryDays(data.id, data.year, data.month, data.day)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(days, data.days) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, days, data.days)
		}
	}
}

func TestDB_DiaryDays_Race(t *testing.T) {
	db := NewDB()
	db.days = dayMap{1: {
		2021: {1: {mock.Diary210101(), mock.Diary210102()}, 2: {mock.Diary210201(), mock.Diary210202()}},
		2022: {3: {mock.Diary220301(), mock.Diary220302()}},
	}}

	d1, _ := db.DiaryDays(1, 2021, 1, 1)
	d2, _ := db.DiaryDays(1, 2021, 1, 1)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		d1[0].KCal = 123
		wg.Done()
	}()

	go func() {
		d2[0].KCal = 456
		wg.Done()
	}()

	wg.Wait()
}

func TestDB_refreshDiaryDay(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		id   int
		date time.Time

		days dayMap
	}{
		{ //00// empty data, technically correct
			db:   &DB{food: map[int]core.Food{}, entries: entryMap{}, days: dayMap{}},
			days: dayMap{0: {1: {1: {core.DiaryDay{Date: "0001-01-01"}}}}},
		},
		{ //01// no entries, clear the day
			db: &DB{
				entries: entryMap{},
				days:    dayMap{1: {2022: {3: {mock.Diary220312()}}}},
			},
			id:   1,
			date: mock.Day1,
			days: dayMap{1: {2022: {3: {}}}},
		},
		{ //02// update existing day
			db: &DB{
				food:    map[int]core.Food{1: mock.Food1, 2: mock.Food2},
				entries: entryMap{1: {mock.Day1: {mock.Entry1(), mock.Entry2()}}},
				days:    dayMap{1: {2022: {3: {core.DiaryDay{Date: "2022-03-12"}}}}},
			},
			id:   1,
			date: mock.Day1,
			days: dayMap{1: {2022: {3: {mock.Diary220312()}}}},
		},
		{ //03// add new day
			db: &DB{
				food: map[int]core.Food{1: mock.Food1, 2: mock.Food2},
				entries: entryMap{1: {
					mock.Day1: {mock.Entry1(), mock.Entry2()},
					mock.Day2: {mock.Entry3()},
				}},
				days: dayMap{1: {2022: {3: {mock.Diary220312()}}}},
			},
			id:   1,
			date: mock.Day2,
			days: dayMap{1: {2022: {3: {mock.Diary220312(), mock.Diary220313()}}}},
		},
	} {
		data.db.refreshDiaryDay(data.id, data.date)

		if !reflect.DeepEqual(data.db.days, data.days) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.days, data.days)
		}
	}
}
