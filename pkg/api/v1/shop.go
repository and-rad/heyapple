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
	"heyapple/pkg/handler"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// SaveListDone enables or disables items on the shopping
// list identified by {name}. Disabled here represents the
// "already bought" state. The response body will always
// be empty, success or failure is communicated by the
// status codes.
//
// Endpoint:
//   /api/v1/list/{name}/done
// Methods:
//   Put
// Possible status codes:
//   204 - Update successful
//   400 - Malformed or missing input data
//   401 - Insufficient permission
//   500 - Internal server error
// Example input:
//   id=1&done=true&id=23&done=false
func SaveListDone(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cmd := &app.SaveShoppingListDone{
			ID:    env.Session.GetInt(r.Context(), "id"),
			Name:  ps.ByName("name"),
			Items: map[int]bool{},
		}

		if cmd.ID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r.ParseForm()
		ids := r.Form["id"]
		done := r.Form["done"]
		if len(ids) != len(done) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for i := range ids {
			id, err := strconv.Atoi(ids[i])
			if err != nil {
				continue
			}
			ok, err := strconv.ParseBool(done[i])
			if err != nil {
				continue
			}
			cmd.Items[id] = ok
		}

		if err := env.DB.Execute(cmd); err == app.ErrMissing {
			w.WriteHeader(http.StatusBadRequest)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

// ShoppingLists returns a JSON-formatted list with
// the names of a user's custom shopping lists.
//
// Endpoint:
//   /api/v1/lists
// Methods:
//   GET
// Possible status codes:
//   200 - OK
//   401 - Insufficient permission
//   500 - Internal server error
// Example output:
//   [ "Dinner", "Cleaning Stuff", ... ]
func ShoppingLists(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if id := env.Session.GetInt(r.Context(), "id"); id == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
		sendResponse([]string{}, w)
	}
}

// ShoppingList returns a JSON-formatted list of shopping
// items generated from a user's diary. It expects at least
// one valid date in the request body in the yyyy-mm-dd
// format. The response body is empty when errors occur
// and will always be an array on success, even when there
// are no entries in the database.
//
// Endpoint:
//   /api/v1/list/{name}
// Methods:
//   GET
// Possible status codes:
//   200 - OK
//   400 - Malformed or missing query parameters
//   401 - Insufficient permission
//   404 - Shopping list doesn't exist
//   500 - Internal server error
// Example input:
//   date=2022-03-12&date=2022-03-13
// Example output:
//   [
//     "id": "1", "amount": 150, "aisle": 12, "done": false, ...
//     "id": "4", "amount": 620, "aisle": 9, "done": true, ...
//     ...
//   ]
func ShoppingList(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		query := &app.ShoppingList{
			Name: ps.ByName("name"),
			ID:   env.Session.GetInt(r.Context(), "id"),
		}

		if query.ID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r.ParseForm()
		for _, date := range r.Form["date"] {
			if day, err := time.Parse("2006-01-02", date); err == nil {
				query.Date = append(query.Date, day)
			}
		}

		if err := env.DB.Fetch(query); err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err == app.ErrMissing {
			w.WriteHeader(http.StatusBadRequest)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			sendResponse(query.Items, w)
		}
	}
}
