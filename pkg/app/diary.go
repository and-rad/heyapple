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
// actually fall on the same day with Date is silently
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
	if err != nil && err != ErrNotFound {
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

// SaveDiaryEntries is a command to update existing
// food items in the diary identified by ID with the
// given amount. The old amount is always replaced
// with the new one.
//
// Any food with an amount of 0 will be removed from
// the database. Any food with a date that does not
// fall on the same day with Date is silently
// dropped.
type SaveDiaryEntries struct {
	Date time.Time
	Food []core.DiaryEntry
	ID   int
}

func (c *SaveDiaryEntries) Execute(db DB) error {
	date := c.Date.Truncate(time.Hour * 24)
	clean := make([]core.DiaryEntry, 0, len(c.Food))
	for _, f := range c.Food {
		if f.Day() == date {
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

	entriesToSet := []core.DiaryEntry{}
	entriesToDel := []core.DiaryEntry{}

	for _, food := range clean {
		for _, entry := range entries {
			if !food.Equal(entry) {
				continue
			}
			if food.Food.Amount > 0 {
				food.Date = entry.Date
				entriesToSet = append(entriesToSet, food)
			} else {
				entriesToDel = append(entriesToDel, food)
			}
			break
		}
	}

	if len(entriesToSet) != 0 {
		if err = db.SetDiaryEntries(c.ID, entriesToSet...); err != nil {
			return err
		}
	}

	if len(entriesToDel) != 0 {
		if err = db.DelDiaryEntries(c.ID, entriesToDel...); err != nil {
			return err
		}
	}

	return nil
}

type DiaryEntries struct {
	Date    time.Time
	Entries []core.DiaryEntry
	ID      int
}

func (q *DiaryEntries) Fetch(db DB) error {
	if q.ID == 0 {
		return ErrNotFound
	}

	date := q.Date.Truncate(time.Hour * 24)
	if entries, err := db.DiaryEntries(q.ID, date); err != nil {
		return err
	} else {
		q.Entries = entries
	}

	return nil
}

type DiaryDays struct {
	Days  []core.DiaryDay
	ID    int
	Year  int
	Month int
}

func (q *DiaryDays) Fetch(db DB) error {
	if q.ID == 0 {
		return ErrNotFound
	}

	if days, err := db.DiaryDays(q.ID, q.Year, q.Month); err != nil {
		return err
	} else {
		q.Days = days
	}

	return nil
}
