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
	"net/http"
	"strconv"

	"github.com/and-rad/scs/v2"
	"github.com/julienschmidt/httprouter"
)

// Foods returns a JSON-formatted list of all the food
// items stored in the database. The function body is
// empty when errors occur and will always be an array
// on success, even when there are no food items in the
// database.
//
// Endpoint:
//   /api/v1/foods
// Methods:
//   GET
// Possible status codes:
//   200 - OK
//   500 - Internal server error
// Example output:
//   [
//     { "id": 1, "kcal": 230, ... },
//     { "id": 6, "kcal": 522, ... },
//     ...
//   ]
func Foods(db app.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		query := &app.GetFoods{}
		if err := db.Fetch(query); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			sendResponse(query.Items, w)
		}
	}
}

// Food returns a JSON-formatted food item identified
// by {id}. The function body is empty when errors occur
// and will be a single JSON object on success.
//
// Endpoint:
//   /api/v1/food/{id}
// Methods:
//   GET
// Possible status codes:
//   200 - OK
//   400 - Missing or malformed id
//   404 - Food item doesn't exist
//   500 - Internal server error
// Example output:
//     { "id": 1, "kcal": 230, ... }
func Food(db app.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		query := &app.GetFood{}
		if id, err := strconv.Atoi(ps.ByName("id")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			query.Item.ID = id
		}

		if err := db.Fetch(query); err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			sendResponse(query.Item, w)
		}
	}
}

// NewFood creates a new food item and returns the item's
// ID on success. The response body will be empty if any
// errors occur.
//
// Endpoint:
//   /api/v1/food
// Methods:
//   POST
// Possible status codes:
//   201 - Creation successful
//   401 - Insufficient permissions
//   500 - Internal server error
// Example output:
//   42
func NewFood(sm *scs.SessionManager, db app.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var hasPermission bool
		if id, ok := sm.Get(r.Context(), "id").(int); ok {
			if u, err := db.UserByID(id); err == nil {
				hasPermission = (u.Perm&app.PermCreateFood == app.PermCreateFood)
			}
		}

		if !hasPermission {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		cmd := &app.CreateFood{}
		if err := db.Execute(cmd); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
			sendResponse(cmd.ID, w)
		}
	}
}

// SaveFood updates a food item in the database identified
// by the {id} parameter. The nutrient values are passed in
// the request body. The response body will always be
// empty, success or failure is communicated by the status
// codes.
//
// Invalid form data does not trigger an error and will
// just be dropped silently. As long as data is valid and
// corresponds to an existing nutrient, it's parsed and
// stored.
//
// Endpoint:
//   /api/v1/food/{id}
// Methods:
//   PUT
// Possible status codes:
//   204 - Update successful
//   400 - Malformed form data
//   401 - Insufficient permissions
//   404 - Food item doesn't exist
//   500 - Internal server error
// Example input:
//   kcal=123&fat=4.5&vitb1=0.06
func SaveFood(sm *scs.SessionManager, db app.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var hasPermission bool
		if id, ok := sm.Get(r.Context(), "id").(int); ok {
			if u, err := db.UserByID(id); err == nil {
				hasPermission = (u.Perm&app.PermEditFood == app.PermEditFood)
			}
		}

		if !hasPermission {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		cmd := &app.SaveFood{}
		if id, err := strconv.Atoi(ps.ByName("id")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			cmd.ID = id
		}

		r.ParseForm()
		cmd.Data = make(map[string]float32)
		for k, v := range r.PostForm {
			if val, err := strconv.ParseFloat(v[0], 32); err == nil {
				cmd.Data[k] = float32(val)
			}
		}

		if err := db.Execute(cmd); err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
