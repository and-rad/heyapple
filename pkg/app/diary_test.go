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
	"testing"
	"time"
)

func TestAddDiaryEntry_Execute(t *testing.T) {
	for idx, data := range []struct {
		cmd *app.AddDiaryEntry
		db  *mock.DB

		entry core.DiaryEntry
		err   error
	}{
		{ //00// don't return error for amount=0
			cmd: &app.AddDiaryEntry{},
			db:  mock.NewDB(),
		},
		{ //01// connection failure
			cmd: &app.AddDiaryEntry{Food: core.Ingredient{ID: 2, Amount: 150}},
			db:  mock.NewDB().WithError(mock.ErrDOS),
			err: mock.ErrDOS,
		},
		{ //02// add a new entry
			db: mock.NewDB(),
			cmd: &app.AddDiaryEntry{
				ID:   1,
				Food: core.Ingredient{ID: 2, Amount: 150},
				Date: mock.Date1,
			},
			entry: core.DiaryEntry{
				Date: mock.Date1,
				Food: core.Ingredient{ID: 2, Amount: 150},
			},
		},
		{ //03// update an entry with fuzzy date matching
			db: mock.NewDB().WithEntries(mock.Entry1()),
			cmd: &app.AddDiaryEntry{
				ID:   1,
				Food: core.Ingredient{ID: 2, Amount: 75},
				Date: mock.Date1p2,
			},
			entry: func() core.DiaryEntry {
				e := mock.Entry1()
				e.Food.Amount += 75
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
