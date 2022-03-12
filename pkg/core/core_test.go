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

package core_test

import (
	"heyapple/internal/mock"
	"heyapple/pkg/core"
	"reflect"
	"testing"
	"time"
)

func TestNewRecipe(t *testing.T) {
	for idx, data := range []struct {
		id  int
		rec core.Recipe
	}{
		{ //00//
			id:  1,
			rec: core.Recipe{ID: 1, Size: 1, Items: []core.Ingredient{}},
		},
		{ //01//
			id:  9000,
			rec: core.Recipe{ID: 9000, Size: 1, Items: []core.Ingredient{}},
		},
	} {
		if rec := core.NewRecipe(data.id); !reflect.DeepEqual(rec, data.rec) {
			t.Errorf("test case %d: recipe mismatch \nhave: %v\nwant: %v", idx, rec, data.rec)
		}
	}
}

func TestDiaryEntry_Equal(t *testing.T) {
	for idx, data := range []struct {
		a core.DiaryEntry
		b core.DiaryEntry

		ok bool
	}{
		{ //00//
			a:  mock.Entry1(),
			b:  mock.Entry2(),
			ok: false,
		},
		{ //01// identity
			a:  mock.Entry1(),
			b:  mock.Entry1(),
			ok: true,
		},
		{ //02// fuzzy date matching
			a: core.DiaryEntry{
				Date:   time.Now(),
				Recipe: "My Recipe",
				Food:   core.Ingredient{ID: 1, Amount: 230},
			},
			b: core.DiaryEntry{
				Date: time.Now().Add(time.Minute * 2),
				Food: core.Ingredient{ID: 1, Amount: 50},
			},
			ok: true,
		},
	} {
		if ok := data.a.Equal(data.b); ok != data.ok {
			t.Errorf("test case %d: equality mismatch \nhave: %v\nwant: %v", idx, ok, data.ok)
		}

		if ok := data.b.Equal(data.a); ok != data.ok {
			t.Errorf("test case %d: equality mismatch \nhave: %v\nwant: %v", idx, ok, data.ok)
		}
	}
}

func TestDiaryEntry_Day(t *testing.T) {
	for idx, data := range []struct {
		entry core.DiaryEntry
		day   string
	}{
		{ //00//
			entry: core.DiaryEntry{Date: time.Date(1987, 12, 6, 11, 45, 0, 0, time.UTC)},
			day:   "1987-12-06",
		},
		{ //01// midnight
			entry: core.DiaryEntry{Date: time.Date(1987, 12, 6, 0, 0, 0, 0, time.UTC)},
			day:   "1987-12-06",
		},
		{ //02// today
			entry: core.DiaryEntry{Date: time.Now().UTC()},
			day:   time.Now().UTC().Format("2006-01-02"),
		},
	} {
		if day := data.entry.Day().Format("2006-01-02"); day != data.day {
			t.Errorf("test case %d: date mismatch \nhave: %v\nwant: %v", idx, day, data.day)
		}
	}
}
