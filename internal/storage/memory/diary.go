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

package memory

import (
	"fmt"
	"sort"
	"time"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/core"
)

func (db *DB) NewDiaryEntries(id int, entries ...core.DiaryEntry) error {
	days, ok := db.entries[id]
	if !ok {
		db.entries[id] = map[time.Time][]core.DiaryEntry{}
		days = db.entries[id]
	}

	dirty := map[time.Time]struct{}{}
	unsorted := map[time.Time]struct{}{}
	for _, e := range entries {
		date := e.Date.Truncate(time.Hour * 24)
		if _, ok := days[date]; !ok {
			days[date] = []core.DiaryEntry{e}
		} else {
			days[date] = append(days[date], e)
			unsorted[date] = struct{}{}
		}
		dirty[date] = struct{}{}
	}

	for date := range unsorted {
		sort.Slice(days[date], func(i, j int) bool {
			return days[date][i].Date.Before(days[date][j].Date)
		})
	}

	for date := range dirty {
		db.refreshDiaryDay(id, date)
	}

	return nil
}

func (db *DB) SetDiaryEntries(id int, entries ...core.DiaryEntry) error {
	days, ok := db.entries[id]
	if !ok {
		return app.ErrNotFound
	}

	dirty := map[time.Time]struct{}{}
	for _, entry := range entries {
		date := entry.Day()
		if _, ok := days[date]; !ok {
			continue
		}

		for i, e := range days[date] {
			if e.Equal(entry) {
				days[date][i] = entry
				dirty[date] = struct{}{}
				break
			}
		}
	}

	for date := range dirty {
		db.refreshDiaryDay(id, date)
	}

	return nil
}

func (db *DB) DelDiaryEntries(id int, entries ...core.DiaryEntry) error {
	days, ok := db.entries[id]
	if !ok {
		return app.ErrNotFound
	}

	dirty := map[time.Time]struct{}{}
	for _, entry := range entries {
		date := entry.Day()
		if _, ok := days[date]; !ok {
			continue
		}

		end := 0
		for _, e := range days[date] {
			if !e.Equal(entry) {
				days[date][end] = e
				end++
			}
		}

		if end != len(days[date]) {
			days[date] = days[date][:end]
			dirty[date] = struct{}{}
		}
	}

	for date := range dirty {
		db.refreshDiaryDay(id, date)
	}

	return nil
}

func (db *DB) DiaryEntries(id int, date time.Time) ([]core.DiaryEntry, error) {
	if days, ok := db.entries[id]; ok {
		if day, ok := days[date.Truncate(time.Hour*24)]; ok {
			return append(make([]core.DiaryEntry, 0, len(day)), day...), nil
		}
	}
	return nil, app.ErrNotFound
}

func (db *DB) DiaryDays(id, year, month, day int) ([]core.DiaryDay, error) {
	years, ok := db.days[id]
	if !ok {
		return nil, app.ErrNotFound
	}

	result := []core.DiaryDay{}
	if year == 0 && month == 0 && day == 0 {
		for _, months := range years {
			for _, m := range months {
				result = append(result, m...)
			}
		}
	} else if months, ok := years[year]; ok {
		if month == 0 {
			for _, m := range months {
				result = append(result, m...)
			}
		} else if days, ok := months[month]; ok {
			if day == 0 {
				result = append(result, days...)
			} else {
				date := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
				for _, d := range days {
					if d.Date == date {
						return []core.DiaryDay{d}, nil
					}
				}
			}
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Date < result[j].Date
	})

	return result, nil
}

func (db *DB) refreshDiaryDay(id int, date time.Time) {
	if _, ok := db.days[id]; !ok {
		db.days[id] = map[int]map[int][]core.DiaryDay{}
	}

	day := core.DiaryDay{Date: date.Format("2006-01-02")}
	entries := db.entries[id][date]
	for _, e := range entries {
		if f, ok := db.food[e.Food.ID]; ok {
			amount := e.Food.Amount * 0.01
			day.KCal += f.KCal * amount
			day.Fat += f.Fat * amount
			day.Carbs += f.Carbs * amount
			day.Protein += f.Protein * amount
		}
	}

	year := date.Year()
	month := int(date.Month())

	if day.Empty() {
		months := db.days[id][year]
		for i, d := range months[month] {
			if d.Date == day.Date {
				months[month] = append(months[month][:i], months[month][i+1:]...)
				return
			}
		}
	}

	if _, ok := db.days[id][year]; !ok {
		db.days[id][year] = map[int][]core.DiaryDay{}
	}

	if _, ok := db.days[id][year][month]; !ok {
		db.days[id][year][month] = []core.DiaryDay{}
	}

	months := db.days[id][year]
	for i, d := range months[month] {
		if d.Date == day.Date {
			months[month][i] = day
			return
		}
	}

	months[month] = append(months[month], day)
	sort.Slice(months[month], func(i, j int) bool {
		return months[month][i].Date < months[month][j].Date
	})
}
