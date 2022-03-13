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

// AddDiaryEntries is a command to add several food
// items to the diary identified by ID and the day
// identified by Date. If any of the food already exists,
// the amount is added on top, otherwise a new entry
// is created.
//
// Any food with an amount of 0 or a date that does not
// actually fall in the same day with Date is silently
// dropped.
type AddDiaryEntries struct {
	Date time.Time
	Food []core.DiaryEntry
	ID   int
}

func (c *AddDiaryEntries) Execute(db DB) error {
	date := c.Date.Truncate(time.Hour * 24)
	clean := make([]core.DiaryEntry, 0, len(c.Food))
	for _, f := range c.Food {
		if f.Food.Amount > 0 && f.Day() == date {
			clean = append(clean, f)
		}
	}

	if len(clean) == 0 {
		return nil
	}

	entries, err := db.DiaryEntries(c.ID, date)
	if err != nil {
		return err
	}

	entriesToAdd := []core.DiaryEntry{}
	entriesToSet := []core.DiaryEntry{}

	for _, food := range clean {
		exists := false
		for _, entry := range entries {
			if food.Equal(entry) {
				exists = true
				food.Date = entry.Date
				food.Food.Amount += entry.Food.Amount
				entriesToSet = append(entriesToSet, food)
				break
			}
		}
		if !exists {
			entriesToAdd = append(entriesToAdd, food)
		}
	}

	if len(entriesToAdd) != 0 {
		if err = db.NewDiaryEntries(c.ID, entriesToAdd...); err != nil {
			return err
		}
	}

	if len(entriesToSet) != 0 {
		if err = db.SetDiaryEntries(c.ID, entriesToSet...); err != nil {
			return err
		}
	}

	return nil
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
