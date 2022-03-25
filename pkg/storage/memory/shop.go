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
			fid := entry.Food.ID
			item, ok := items[fid]
			if !ok {
				item.ID = fid
				item.Price = db.prices[id][fid]
				item.Aisle = db.aisle(id, fid)
				item.Done = db.done[id][fid]
			}
			item.Amount += entry.Food.Amount
			items[fid] = item
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

func (db *DB) aisle(uid, fid int) core.Aisle {
	if aisles, ok := db.aisles[uid]; ok {
		if aisle, ok := aisles[fid]; ok {
			return aisle
		}
	}
	return db.aisles[0][fid]
}
