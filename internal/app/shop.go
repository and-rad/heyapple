////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2021-2024 The HeyApple Authors.
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
	"time"

	"github.com/and-rad/heyapple/internal/core"
)

type SaveShoppingListDone struct {
	Items map[int]bool
	Name  string
	ID    int
}

func (c *SaveShoppingListDone) Execute(db DB) error {
	if c.ID == 0 {
		return ErrMissing
	}

	if len(c.Items) == 0 {
		return ErrMissing
	}

	if c.Name == "" {
		return ErrMissing
	}

	clean := make(map[int]bool, len(c.Items))
	for k, v := range c.Items {
		if ok, _ := db.FoodExists(k); ok {
			clean[k] = v
		}
	}

	if len(clean) == 0 {
		return nil
	}

	if c.Name != "diary" {
		// TODO implement custom shopping lists
		return ErrNotFound
	}

	return db.SetShoppingListDone(c.ID, clean)
}

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
