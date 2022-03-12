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

package app

import (
	"heyapple/pkg/core"
	"time"
)

// AddDiaryEntry is a command to add a food item to the
// diary identified by ID at the time given by Date. If
// the same food exists at the same time, the amount is
// added on top, otherwise a new entry is created.
type AddDiaryEntry struct {
	Date   time.Time
	Recipe string
	ID     int
	Food   core.Ingredient
}

func (c *AddDiaryEntry) Execute(db DB) error {
	if c.Food.Amount == 0 {
		return nil
	}

	entry, err := db.DiaryEntry(c.ID, c.Food.ID, c.Date)
	if err == ErrNotFound {
		entry.Date = c.Date
		entry.Recipe = c.Recipe
		entry.Food = c.Food
		return db.NewDiaryEntries(c.ID, entry)
	}

	if err == nil {
		entry.Food.Amount += c.Food.Amount
		return db.SetDiaryEntries(c.ID, entry)
	}

	return err
}

// SaveDiaryEntry is a command to update an existing
// food item in the diary identified by ID with the
// given amount. The old amount is always replaced
// with the new one. If the new amount is 0, the item
// will be removed from the diary.
type SaveDiaryEntry struct {
	Date time.Time
	Food core.Ingredient
	ID   int
}

func (c *SaveDiaryEntry) Execute(db DB) error {
	if c.Food.Amount == 0 {
		entry := core.DiaryEntry{Date: c.Date, Food: c.Food}
		return db.DelDiaryEntries(c.ID, entry)
	}

	entry, err := db.DiaryEntry(c.ID, c.Food.ID, c.Date)
	if err != nil {
		return err
	}

	entry.Food.Amount = c.Food.Amount
	return db.SetDiaryEntries(c.ID, entry)
}
