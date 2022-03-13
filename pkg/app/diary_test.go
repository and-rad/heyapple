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

func TestSaveDiaryEntry_Execute(t *testing.T) {
	for idx, data := range []struct {
		cmd *app.SaveDiaryEntry
		db  *mock.DB

		entry core.DiaryEntry
		err   error
	}{
		{ //00// nothing in, nothing out
			cmd: &app.SaveDiaryEntry{},
			db:  mock.NewDB(),
		},
		{ //01// delete if amount=0
			db:    mock.NewDB().WithEntries(mock.Entry1()),
			cmd:   &app.SaveDiaryEntry{ID: 1, Food: core.Ingredient{ID: 2}},
			entry: core.DiaryEntry{},
		},
		{ //02// item doesn't exist
			db: mock.NewDB().WithEntries(mock.Entry1()),
			cmd: &app.SaveDiaryEntry{
				ID:   1,
				Food: core.Ingredient{ID: 2, Amount: 123},
				Date: time.Now(),
			},
			err: app.ErrNotFound,
		},
		{ //03// update an entry with fuzzy date matching
			db: mock.NewDB().WithEntries(mock.Entry1()),
			cmd: &app.SaveDiaryEntry{
				ID:   1,
				Food: core.Ingredient{ID: 2, Amount: 75},
				Date: mock.Entry1().Date.Add(time.Minute * 2),
			},
			entry: func() core.DiaryEntry {
				e := mock.Entry1()
				e.Food.Amount = 75
				return e
			}(),
		},
	} {
		cmd := data.cmd
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if entry, _ := data.db.DiaryEntry(cmd.ID, cmd.Food.ID, cmd.Date); entry != data.entry {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, entry, data.entry)
		}
	}
}

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
		{ //04// success, add single food
			db: mock.NewDB(),
			cmd: &app.AddDiaryEntries{
				ID:   1,
				Date: mock.Day1,
				Food: []core.DiaryEntry{mock.Entry1()},
			},
			entries: []core.DiaryEntry{mock.Entry1()},
		},
		{ //05// success, add to existing single food
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
		{ //06// success, complex operations
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
		cmd := data.cmd
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.db.Entries, data.entries) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.Entries, data.entries)
		}
	}
}
