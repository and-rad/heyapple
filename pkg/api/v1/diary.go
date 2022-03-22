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
	"fmt"
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"heyapple/pkg/handler"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// SaveDiaryEntry edits food in the diary on the day
// identified by {date}. Time, food id, amount, and
// recipe name are passed in the request body. The response
// body will always be empty, success or failure is
// communicated by the status codes.
//
// Food items that do not fit the date are just ignored
// without returning an error. POST adds the given
// amount to the amount of an existing item while PUT
// replaces it.
//
// Date is expected to be in YYYY-MM-DD format, the time
// in simplified 24-hour hh:mm format. Any food item
// that does not conform to this will be ignored.
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

		fmt.Printf("%v\n", r.Form)

		ids := r.Form["id"]
		amounts := r.Form["amount"]
		times := r.Form["time"]
		recs := r.Form["recipe"]
		if !(len(ids) == len(amounts) && len(ids) == len(times) && len(ids) == len(recs)) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

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
			format := ps.ByName("date") + "T" + times[i] + ":00Z"
			date, err := time.Parse(time.RFC3339, format)
			if err != nil {
				continue
			}
			entries = append(entries, core.DiaryEntry{
				Date:   date,
				Recipe: recs[i],
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

// Diary returns a JSON-formatted list of objects that
// summarize days from a user's diary. If no query
// parameters are provided, the entire diary is returned,
// otherwise only those dates that match {year} and {month}.
// The function body is empty when errors occur and will
// always be an array on success, even when there are
// no entries in the database.
//
// Endpoint:
//   /api/v1/diary/{year}/{month}
// Methods:
//   GET
// Possible status codes:
//   200 - OK
//   400 - Malformed query parameters
//   401 - Insufficient permission
//   400 - Diary doesn't exist
//   500 - Internal server error
// Example output:
//   [
//     { "date": "2022-03-12", "kcal": 1987, ... },
//     { "date": "2022-03-13", "kcal": 2150, ... },
//     ...
//   ]
func Diary(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		query := &app.DiaryDays{}

		if param := ps.ByName("year"); param != "" {
			if year, err := strconv.Atoi(param); err == nil {
				query.Year = year
			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		if param := ps.ByName("month"); param != "" {
			if month, err := strconv.Atoi(param); err == nil {
				query.Month = month
			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		if param := ps.ByName("day"); param != "" {
			if day, err := strconv.Atoi(param); err == nil {
				query.Day = day
			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		if id, ok := env.Session.Get(r.Context(), "id").(int); ok {
			query.ID = id
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err := env.DB.Fetch(query); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			sendResponse(query.Days, w)
		}
	}
}

// DiaryEntries returns a JSON-formatted list of food
// that is recorded in a user's diary. The date params
// are not optional. The function body is empty when
// errors occur and will always be an array on success,
// even when there are no entries in the database.
//
// Endpoint:
//   /api/v1/diary/{year}/{month}/{day}/entries
// Methods:
//   GET
// Possible status codes:
//   200 - OK
//   400 - Malformed query parameters
//   401 - Insufficient permission
//   400 - Diary doesn't exist
//   500 - Internal server error
// Example output:
//   [
//     "date": "2022-03-12T08:43:00Z", "food":{"id": 2, "amount": 150}, ...
//     "date": "2022-03-12T13:05:00Z", "food":{"id": 1, "amount": 240}, ...
//     ...
//   ]
func DiaryEntries(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		year, err := strconv.Atoi(ps.ByName("year"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		month, err := strconv.Atoi(ps.ByName("month"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		day, err := strconv.Atoi(ps.ByName("day"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		query := &app.DiaryEntries{
			Date: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC),
		}

		if id, ok := env.Session.Get(r.Context(), "id").(int); ok {
			query.ID = id
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err := env.DB.Fetch(query); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			sendResponse(query.Entries, w)
		}
	}
}
