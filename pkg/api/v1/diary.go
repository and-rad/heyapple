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

package api

import (
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"heyapple/pkg/handler"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// SaveDiaryEntry edits food in the diary on the day
// identified by {date}. Food id, amount, and recipe
// name are passed in the request body. The response
// body will always be empty, success or failure is
// communicated by the status codes.
//
// Food items that do not fit the date are just ignored
// without returning an error. POST adds the given
// amount to the amount of an existing item while PUT
// replaces it.
//
// Endpoint:
//   /api/v1/diary/{date}
// Methods:
//   POST, PUT
// Possible status codes:
//   204 - Update successful
//   400 - Malformed or missing form data
//   401 - Insufficient permission
//   404 - Diary doesn't exist
//   500 - Internal server error
// Example input:
//   recipe=Pie&id=1&amount=125&date=2022-03-12T12:45:00
func SaveDiaryEntry(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		day, err := time.Parse("2006-01-02", ps.ByName("date"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, ok := env.Session.Get(r.Context(), "id").(int)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r.ParseForm()
		ids := r.Form["id"]
		amounts := r.Form["amount"]
		dates := r.Form["date"]
		if !(len(ids) == len(amounts) && len(ids) == len(dates)) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		recipe := r.FormValue("recipe")
		entries := []core.DiaryEntry{}
		for i := range ids {
			id, err := strconv.Atoi(ids[i])
			if err != nil {
				continue
			}
			amount, err := strconv.ParseFloat(amounts[i], 32)
			if err != nil {
				continue
			}
			date, err := time.Parse(time.RFC3339, dates[i])
			if err != nil {
				continue
			}
			entries = append(entries, core.DiaryEntry{
				Date:   date,
				Recipe: recipe,
				Food:   core.Ingredient{ID: id, Amount: float32(amount)},
			})
		}

		if r.Method == http.MethodPost {
			err = env.DB.Execute(&app.AddDiaryEntries{ID: id, Date: day, Food: entries})
		} else {
			err = env.DB.Execute(&app.SaveDiaryEntries{ID: id, Date: day, Food: entries})
		}

		if err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
