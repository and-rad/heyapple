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

type ShoppingList struct {
	Name  string
	Date  []time.Time
	Items []core.ShopItem
	ID    int
}

func (q *ShoppingList) Fetch(db DB) error {
	if q.ID == 0 {
		return ErrMissing
	}

	if len(q.Date) == 0 {
		return ErrMissing
	}

	if q.Name == "" {
		return ErrMissing
	}

	if q.Name != "diary" {
		// TODO implement custom shopping lists
		return ErrNotFound
	}

	items, err := db.ShoppingList(q.ID, q.Date...)
	if err == ErrNotFound {
		q.Items = []core.ShopItem{}
	} else if err != nil {
		return err
	} else {
		q.Items = items
	}

	return nil
}
