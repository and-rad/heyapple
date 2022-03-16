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

package app_test

import (
	"heyapple/internal/mock"
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"reflect"
	"testing"
	"time"
)

func TestAddDiaryEntries_Execute(t *testing.T) {
	for idx, data := range []struct {
		cmd *app.AddDiaryEntries
		db  *mock.DB

		entries []core.DiaryEntry
		err     error
	}{
		{ //00// don't return error for amount=0
			db:      mock.NewDB(),
			cmd:     &app.AddDiaryEntries{},
			entries: []core.DiaryEntry{},
		},
		{ //01// connection failure
			db:      mock.NewDB().WithError(mock.ErrDOS),
			cmd:     &app.AddDiaryEntries{Date: mock.Day1, Food: []core.DiaryEntry{mock.Entry1()}},
			entries: []core.DiaryEntry{},
			err:     mock.ErrDOS,
		},
		{ //02// delayed connection failure
			db:      mock.NewDB().WithError(nil, mock.ErrDOS),
			cmd:     &app.AddDiaryEntries{ID: 1, Date: mock.Day1, Food: []core.DiaryEntry{mock.Entry1()}},
			entries: []core.DiaryEntry{},
			err:     mock.ErrDOS,
		},
		{ //03// delayed connection failure
			db:      mock.NewDB().WithEntries(mock.Entry1()).WithError(nil, mock.ErrDOS),
			cmd:     &app.AddDiaryEntries{ID: 1, Date: mock.Day1, Food: []core.DiaryEntry{mock.Entry1New()}},
			entries: []core.DiaryEntry{mock.Entry1()},
			err:     mock.ErrDOS,
		},
		{ //04// ignore "not found" errors
			db:      mock.NewDB().WithError(app.ErrNotFound),
			cmd:     &app.AddDiaryEntries{Date: mock.Day1, Food: []core.DiaryEntry{mock.Entry1()}},
			entries: []core.DiaryEntry{mock.Entry1()},
		},
		{ //05// success, add single food
			db: mock.NewDB(),
			cmd: &app.AddDiaryEntries{
				ID:   1,
				Date: mock.Day1,
				Food: []core.DiaryEntry{mock.Entry1()},
			},
			entries: []core.DiaryEntry{mock.Entry1()},
		},
		{ //06// success, add to existing single food
			db: mock.NewDB().WithEntries(mock.Entry1()),
			cmd: &app.AddDiaryEntries{
				ID:   1,
				Date: mock.Day1,
				Food: []core.DiaryEntry{mock.Entry1New()},
			},
			entries: []core.DiaryEntry{
				func() core.DiaryEntry {
					e := mock.Entry1New()
					e.Date = mock.Entry1().Date
					e.Food.Amount += mock.Entry1().Food.Amount
					return e
				}(),
			},
		},
		{ //07// success, complex operations
			db: mock.NewDB().WithEntries(
				core.DiaryEntry{Date: mock.Date1, Food: core.Ingredient{ID: 1, Amount: 100}},
				core.DiaryEntry{Date: mock.Date1, Food: core.Ingredient{ID: 2, Amount: 50}},
				core.DiaryEntry{Date: mock.Date2, Food: core.Ingredient{ID: 1, Amount: 30}, Recipe: "Rec1"},
				core.DiaryEntry{Date: mock.Date2, Food: core.Ingredient{ID: 2, Amount: 40}, Recipe: "Rec1"},
				core.DiaryEntry{Date: mock.Date2, Food: core.Ingredient{ID: 3, Amount: 70}, Recipe: "Rec1"},
				core.DiaryEntry{Date: mock.Date3, Food: core.Ingredient{ID: 1, Amount: 70}, Recipe: "Rec2"},
			),
			cmd: &app.AddDiaryEntries{ID: 1, Date: mock.Day1, Food: []core.DiaryEntry{
				{Date: mock.Date1, Food: core.Ingredient{ID: 1, Amount: 0}},                   // ignore, amount is 0
				{Date: mock.Date1p2, Food: core.Ingredient{ID: 2, Amount: 100}},               // add to existing
				{Date: mock.Date2, Food: core.Ingredient{ID: 1, Amount: 100}},                 // add as new entry
				{Date: mock.Date2, Food: core.Ingredient{ID: 1, Amount: 100}, Recipe: "Rec1"}, // add to existing
				{Date: mock.Date3, Food: core.Ingredient{ID: 1, Amount: 100}},                 // ignore, belongs to another day
			}},
			entries: []core.DiaryEntry{
				{Date: mock.Date1, Food: core.Ingredient{ID: 1, Amount: 100}},
				{Date: mock.Date1, Food: core.Ingredient{ID: 2, Amount: 150}},
				{Date: mock.Date2, Food: core.Ingredient{ID: 1, Amount: 130}, Recipe: "Rec1"},
				{Date: mock.Date2, Food: core.Ingredient{ID: 2, Amount: 40}, Recipe: "Rec1"},
				{Date: mock.Date2, Food: core.Ingredient{ID: 3, Amount: 70}, Recipe: "Rec1"},
				{Date: mock.Date3, Food: core.Ingredient{ID: 1, Amount: 70}, Recipe: "Rec2"},
				{Date: mock.Date2, Food: core.Ingredient{ID: 1, Amount: 100}},
			},
		},
	} {
		err := data.cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.db.Entries, data.entries) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.Entries, data.entries)
		}
	}
}

func TestSaveDiaryEntries_Execute(t *testing.T) {
	for idx, data := range []struct {
		cmd *app.SaveDiaryEntries
		db  *mock.DB

		entries []core.DiaryEntry
		err     error
	}{
		{ //00// nothing to do
			db:      mock.NewDB(),
			cmd:     &app.SaveDiaryEntries{},
			entries: []core.DiaryEntry{},
		},
		{ //01// connection failure
			db:      mock.NewDB().WithError(mock.ErrDOS),
			cmd:     &app.SaveDiaryEntries{Date: mock.Day1, Food: []core.DiaryEntry{mock.Entry1()}},
			entries: []core.DiaryEntry{},
			err:     mock.ErrDOS,
		},
		{ //02// delayed connection failure
			db:      mock.NewDB().WithEntries(mock.Entry1()).WithError(nil, mock.ErrDOS),
			cmd:     &app.SaveDiaryEntries{ID: 1, Date: mock.Day1, Food: []core.DiaryEntry{mock.Entry1()}},
			entries: []core.DiaryEntry{mock.Entry1()},
			err:     mock.ErrDOS,
		},
		{ //03// delayed connection failure
			db: mock.NewDB().WithEntries(mock.Entry1()).WithError(nil, mock.ErrDOS),
			cmd: &app.SaveDiaryEntries{ID: 1, Date: mock.Day1, Food: []core.DiaryEntry{
				func() core.DiaryEntry { e := mock.Entry1(); e.Food.Amount = 0; return e }(),
			}},
			entries: []core.DiaryEntry{mock.Entry1()},
			err:     mock.ErrDOS,
		},
		{ //04// success, save single food
			db: mock.NewDB().WithEntries(mock.Entry1()),
			cmd: &app.SaveDiaryEntries{
				ID:   1,
				Date: mock.Day1,
				Food: []core.DiaryEntry{mock.Entry1New()},
			},
			entries: []core.DiaryEntry{
				func() core.DiaryEntry {
					e := mock.Entry1New()
					e.Date = mock.Entry1().Date
					return e
				}(),
			},
		},
		{ //05// success, delete single food
			db: mock.NewDB().WithEntries(mock.Entry1()),
			cmd: &app.SaveDiaryEntries{ID: 1, Date: mock.Day1, Food: []core.DiaryEntry{
				func() core.DiaryEntry { e := mock.Entry1(); e.Food.Amount = 0; return e }(),
			}},
			entries: []core.DiaryEntry{},
		},
		{ //06// success, complex operations
			db: mock.NewDB().WithEntries(
				core.DiaryEntry{Date: mock.Date1, Food: core.Ingredient{ID: 1, Amount: 100}},
				core.DiaryEntry{Date: mock.Date1, Food: core.Ingredient{ID: 2, Amount: 50}},
				core.DiaryEntry{Date: mock.Date2, Food: core.Ingredient{ID: 1, Amount: 30}, Recipe: "Rec1"},
				core.DiaryEntry{Date: mock.Date2, Food: core.Ingredient{ID: 2, Amount: 40}, Recipe: "Rec1"},
				core.DiaryEntry{Date: mock.Date2, Food: core.Ingredient{ID: 3, Amount: 70}, Recipe: "Rec1"},
				core.DiaryEntry{Date: mock.Date3, Food: core.Ingredient{ID: 1, Amount: 80}, Recipe: "Rec2"},
			),
			cmd: &app.SaveDiaryEntries{ID: 1, Date: mock.Day1, Food: []core.DiaryEntry{
				{Date: mock.Date1, Food: core.Ingredient{ID: 1, Amount: 0}},                   // remove
				{Date: mock.Date1p2, Food: core.Ingredient{ID: 2, Amount: 100}},               // save
				{Date: mock.Date2, Food: core.Ingredient{ID: 1, Amount: 100}},                 // ignore, doesn't exist
				{Date: mock.Date2, Food: core.Ingredient{ID: 1, Amount: 100}, Recipe: "Rec1"}, // save
				{Date: mock.Date3, Food: core.Ingredient{ID: 1, Amount: 100}},                 // ignore, belongs to another day
			}},
			entries: []core.DiaryEntry{
				{Date: mock.Date1, Food: core.Ingredient{ID: 2, Amount: 100}},
				{Date: mock.Date2, Food: core.Ingredient{ID: 1, Amount: 100}, Recipe: "Rec1"},
				{Date: mock.Date2, Food: core.Ingredient{ID: 2, Amount: 40}, Recipe: "Rec1"},
				{Date: mock.Date2, Food: core.Ingredient{ID: 3, Amount: 70}, Recipe: "Rec1"},
				{Date: mock.Date3, Food: core.Ingredient{ID: 1, Amount: 80}, Recipe: "Rec2"},
			},
		},
	} {
		err := data.cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.db.Entries, data.entries) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.Entries, data.entries)
		}
	}
}

func TestDiaryEntries_Fetch(t *testing.T) {
	for idx, data := range []struct {
		db   *mock.DB
		id   int
		date time.Time

		entries []core.DiaryEntry
		err     error
	}{
		{ //00// connection failed
			db:  mock.NewDB().WithError(mock.ErrDOS),
			id:  1,
			err: mock.ErrDOS,
		},
		{ //01// no id provided
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //02// diary doesn't exist
			db:      mock.NewDB().WithError(app.ErrNotFound),
			id:      2,
			entries: []core.DiaryEntry{},
		},
		{ //03// empty db
			db:      mock.NewDB(),
			id:      1,
			entries: []core.DiaryEntry{},
		},
		{ //04// success
			db:      mock.NewDB().WithEntries(mock.Entry1(), mock.Entry2()),
			id:      1,
			date:    mock.Date1,
			entries: []core.DiaryEntry{mock.Entry1(), mock.Entry2()},
		},
		{ //05// no entries for specified date
			db:      mock.NewDB().WithEntries(mock.Entry1(), mock.Entry2()),
			id:      1,
			date:    mock.Day2,
			entries: []core.DiaryEntry{},
		},
	} {
		qry := &app.DiaryEntries{ID: data.id, Date: data.date}
		err := qry.Fetch(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(qry.Entries, data.entries) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, qry.Entries, data.entries)
		}
	}
}

func TestDiaryDays_Fetch(t *testing.T) {
	for idx, data := range []struct {
		db    *mock.DB
		id    int
		year  int
		month int

		days []core.DiaryDay
		err  error
	}{
		{ //00// connection failed
			db:  mock.NewDB().WithError(mock.ErrDOS),
			id:  1,
			err: mock.ErrDOS,
		},
		{ //01// no id provided
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //02// diary doesn't exist
			db:   mock.NewDB().WithError(app.ErrNotFound),
			id:   2,
			days: []core.DiaryDay{},
		},
		{ //03// empty db
			db:   mock.NewDB(),
			id:   1,
			days: []core.DiaryDay{},
		},
		{ //04// success
			db:   mock.NewDB().WithDays(mock.Diary210101(), mock.Diary210102(), mock.Diary220301()),
			id:   1,
			year: 2021,
			days: []core.DiaryDay{mock.Diary210101(), mock.Diary210102()},
		},
		{ //05// no entries for specified date
			db:   mock.NewDB().WithDays(mock.Diary210101(), mock.Diary210102(), mock.Diary220301()),
			id:   1,
			year: 2020,
			days: []core.DiaryDay{},
		},
		{ //06// success for specific month
			db:    mock.NewDB().WithDays(mock.Diary210101(), mock.Diary210102(), mock.Diary210201()),
			id:    1,
			year:  2021,
			month: 2,
			days:  []core.DiaryDay{mock.Diary210201()},
		},
		{ //07// no entries for specified month
			db:    mock.NewDB().WithDays(mock.Diary210101(), mock.Diary210102(), mock.Diary210201()),
			id:    1,
			year:  2021,
			month: 3,
			days:  []core.DiaryDay{},
		},
	} {
		qry := &app.DiaryDays{ID: data.id, Year: data.year, Month: data.month}
		err := qry.Fetch(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(qry.Days, data.days) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, qry.Days, data.days)
		}
	}
}
