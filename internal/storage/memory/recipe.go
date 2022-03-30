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
	"heyapple/internal/app"
	"heyapple/internal/core"
	"sort"
)

func (db *DB) NewRecipe(name string) (int, error) {
	db.recID++
	rec := core.NewRecipe(db.recID)
	rec.Name = name
	db.recipes[rec.ID] = rec
	return rec.ID, nil
}

func (db *DB) SetRecipe(rec core.Recipe) error {
	if _, ok := db.recipes[rec.ID]; ok {
		db.recipes[rec.ID] = rec
		return nil
	}
	return app.ErrNotFound
}

func (db *DB) SetRecipeAccess(user, rec, perms int) error {
	if _, ok := db.users[user]; !ok {
		return app.ErrNotFound
	}
	if _, ok := db.recipes[rec]; !ok {
		return app.ErrNotFound
	}
	if _, ok := db.userRec[user]; !ok {
		db.userRec[user] = make(map[int]int)
	}
	if _, ok := db.recUser[rec]; !ok {
		db.recUser[rec] = make(map[int]int)
	}
	db.userRec[user][rec] = perms
	db.recUser[rec][user] = perms
	return nil
}

func (db *DB) Recipe(id int) (core.Recipe, error) {
	if r, ok := db.recipes[id]; ok {
		rec := r
		rec.Items = append([]core.Ingredient{}, r.Items...)
		return rec, nil
	}
	return core.Recipe{}, app.ErrNotFound
}

func (db *DB) RecipeAccess(user, rec int) (int, error) {
	combined := app.PermNone
	if acc, ok := db.userRec[0]; ok {
		combined |= acc[rec]
	}
	if acc, ok := db.userRec[user]; ok {
		combined |= acc[rec]
	}
	return combined, nil
}

func (db *DB) Recipes(uid int, f core.Filter) ([]core.Recipe, error) {
	ids := map[int]struct{}{}
	recs := []core.Recipe{}

	for id, perm := range db.userRec[0] {
		r := db.recipes[id]
		if perm != app.PermNone && f.MatchRecipe(r) {
			ids[id] = struct{}{}
			rec := r
			rec.Items = append([]core.Ingredient{}, r.Items...)
			recs = append(recs, rec)
		}
	}

	for id, perm := range db.userRec[uid] {
		if _, ok := ids[id]; !ok {
			r := db.recipes[id]
			if perm != app.PermNone && f.MatchRecipe(r) {
				rec := r
				rec.Items = append([]core.Ingredient{}, r.Items...)
				recs = append(recs, rec)
			}
		}
	}

	sort.Slice(recs, func(i, j int) bool {
		return recs[i].ID < recs[j].ID
	})

	return recs, nil
}
