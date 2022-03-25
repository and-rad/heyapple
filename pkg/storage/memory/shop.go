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
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"sort"
	"time"
)

func (db *DB) ShoppingList(id int, date ...time.Time) ([]core.ShopItem, error) {
	days, ok := db.entries[id]
	if !ok {
		return nil, app.ErrNotFound
	}

	items := map[int]core.ShopItem{}
	for _, dd := range date {
		day, ok := days[dd.Truncate(time.Hour*24)]
		if !ok {
			continue
		}
		for _, entry := range day {
			id := entry.Food.ID
			item, ok := items[id]
			if !ok {
				item.ID = id
				// TODO these ned to be implemented
				item.Price = [2]float32{}
				item.Aisle = 0
				item.Done = false
			}
			item.Amount += entry.Food.Amount
			items[id] = item
		}
	}

	result := make([]core.ShopItem, 0, len(items))
	for _, i := range items {
		result = append(result, i)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result, nil
}
