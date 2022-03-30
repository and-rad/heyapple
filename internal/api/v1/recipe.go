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
	"heyapple/internal/app"
	"heyapple/internal/core"
	"heyapple/internal/handler"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// NewRecipe creates a new named recipe and returns its
// ID on success. The response body will be empty if any
// errors occur.
//
// Endpoint:
//   /api/v1/recipe
// Methods:
//   POST
// Possible status codes:
//   201 - Creation successful
//   202 - Partial success
//   400 - Missing form data
//   401 - Insufficient permission
//   500 - Internal server error
// Example input:
//   name=Pie
// Example output:
//   { "id": 1, "size": 0, "items": [] }
func NewRecipe(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cmd := &app.CreateRecipe{Name: r.FormValue("name")}
		if cmd.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else if uid, ok := env.Session.Get(r.Context(), "id").(int); !ok {
			w.WriteHeader(http.StatusUnauthorized)
		} else if err := env.DB.Execute(cmd); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else if err := env.DB.Execute(&app.SaveRecipeAccess{
			UserID:     uid,
			RecID:      cmd.Recipe.ID,
			Permission: app.PermOwner,
		}); err != nil {
			w.WriteHeader(http.StatusAccepted)
			sendResponse(cmd.Recipe, w)
		} else {
			w.WriteHeader(http.StatusCreated)
			sendResponse(cmd.Recipe, w)
		}
	}
}

// SaveRecipe updates a recipe in the database identified
// by the {id} parameter. The data is passed in the request
// body. The response body will always be empty, success
// or failure is communicated by the status codes.
//
// This endpoint does not update the list of ingredients!
//
// Invalid form data does not trigger an error and will
// just be dropped silently. As long as data is valid and
// corresponds to an existing food item, it's parsed and
// stored.
//
// Endpoint:
//   /api/v1/recipe/{id}
// Methods:
//   PUT
// Possible status codes:
//   204 - Update successful
//   400 - Malformed or missing form data
//   401 - Insufficient permission
//   404 - Recipe doesn't exist
//   500 - Internal server error
// Example input:
//   size=12&name=Butterscotch
func SaveRecipe(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		recID, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var hasPermission bool
		if id, ok := env.Session.Get(r.Context(), "id").(int); ok {
			query := &app.RecipeAccess{UserID: id, RecID: recID}
			if env.DB.Fetch(query) == nil {
				hasPermission = query.HasPerms(app.PermEdit)
			}
		}

		if !hasPermission {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		cmd := &app.SaveRecipe{ID: recID, Data: map[string]interface{}{}}
		if name := r.FormValue("name"); name != "" {
			cmd.Data["name"] = name
		}
		if size, err := strconv.Atoi(r.FormValue("size")); err == nil {
			cmd.Data["size"] = size
		}
		if time, err := strconv.Atoi(r.FormValue("preptime")); err == nil {
			cmd.Data["preptime"] = time
		}
		if time, err := strconv.Atoi(r.FormValue("cooktime")); err == nil {
			cmd.Data["cooktime"] = time
		}
		if time, err := strconv.Atoi(r.FormValue("misctime")); err == nil {
			cmd.Data["misctime"] = time
		}

		if err := env.DB.Execute(cmd); err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

// Recipes returns a JSON-formatted list of recipes
// stored in the database. If no query parameters are
// transmitted, it returns all recipes that the session
// user has access to. The function body is empty when
// errors occur and will always be an array on success,
// even when there are no recipes in the database.
//
// If a query parameter is included more than once, its
// first two values are interpreted as the minimum and
// the maximum valid amount.
//
// Endpoint:
//   /api/v1/recipes
// Methods:
//   GET
// Possible status codes:
//   200 - OK
//   500 - Internal server error
// Example input:
//   name=Pie&kcal=32&kcal=120&prot=12
// Example output:
//   [
//     { "id": 1, "name": "Pie", "size": 6, ... },
//     { "id": 6, "name": "Omelette", "size": 1, ... },
//     ...
//   ]
func Recipes(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		query := &app.Recipes{Filter: getRecipeFilter(r)}
		if id, ok := env.Session.Get(r.Context(), "id").(int); ok {
			query.UserID = id
		}

		if err := env.DB.Fetch(query); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			sendResponse(query.Items, w)
		}
	}
}

// Recipe returns a JSON-formatted recipe identified
// by {id}. The function body is empty when errors occur
// and will be a single JSON object on success.
//
// Endpoint:
//   /api/v1/recipe/{id}
// Methods:
//   GET
// Possible status codes:
//   200 - OK
//   400 - Missing or malformed id
//   404 - Recipe doesn't exist
//   500 - Internal server error
// Example output:
//   { "id": 1, "name": "Pie", "size": 6, ... }
func Recipe(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		recID, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var hasPermission bool
		if id, ok := env.Session.Get(r.Context(), "id").(int); ok {
			query := &app.RecipeAccess{UserID: id, RecID: recID}
			if env.DB.Fetch(query) == nil {
				hasPermission = query.HasPerms(app.PermRead)
			}
		}

		if !hasPermission {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		query := &app.Recipe{Item: core.Recipe{ID: recID}}
		if err := env.DB.Fetch(query); err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			sendResponse(query.Item, w)
		}
	}
}

// RecipeOwner returns ownership information about the
// recipe identified by {id}. The function body is empty
// when errors occur and will be a single JSON object
// on success.
//
// Endpoint:
//   /api/v1/recipe/{id}/owner
// Methods:
//   GET
// Possible status codes:
//   200 - OK
//   400 - Missing or malformed id
//   404 - Recipe doesn't exist
//   500 - Internal server error
// Example output:
//   { "isowner": false, "owner": "NastyOrange" }
func RecipeOwner(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// TODO this function is structured a little odd because
		// we prepared it to be ready for adding another query
		// later once usernames are implemented.

		rid, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		uid := env.Session.GetInt(r.Context(), "id")
		data := map[string]interface{}{}

		query1 := &app.RecipeAccess{UserID: uid, RecID: rid}
		if err := env.DB.Fetch(query1); err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			data["isowner"] = query1.HasPerms(app.PermOwner)
		}

		// TODO replace this with an actual implementation
		if len(data) != 0 {
			data["owner"] = ""
		}

		if len(data) != 0 {
			w.WriteHeader(http.StatusOK)
			sendResponse(data, w)
		}
	}
}

// SaveIngredient adds the ingredient identified by {ing} to
// the recipe identified by {id}. The amount is passed in
// the request body. The response body will always be empty,
// success or failure is communicated by the status codes.
//
// If the ingredient does not exist in the database, it is
// just ignored without returning an error. POST adds the
// given amount to the amount of an existing ingredient while
// PUT replaces it.
//
// Endpoint:
//   /api/v1/recipe/{id}/ingredient/{ing}
// Methods:
//   POST, PUT
// Possible status codes:
//   204 - Update successful
//   400 - Malformed or missing form data
//   401 - Insufficient permission
//   404 - Recipe doesn't exist
//   500 - Internal server error
// Example input:
//   amount=125
func SaveIngredient(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		recID, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var hasPermission bool
		if id, ok := env.Session.Get(r.Context(), "id").(int); ok {
			query := &app.RecipeAccess{UserID: id, RecID: recID}
			if env.DB.Fetch(query) == nil {
				hasPermission = query.HasPerms(app.PermEdit)
			}
		}

		if !hasPermission {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ingID, err := strconv.Atoi(ps.ByName("ing"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		amount, err := strconv.ParseFloat(r.FormValue("amount"), 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		replace := r.Method == http.MethodPut
		if !replace && amount == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		cmd := &app.SaveIngredient{
			Amount:       float32(amount),
			RecipeID:     recID,
			IngredientID: ingID,
			Replace:      replace,
		}

		if err := env.DB.Execute(cmd); err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

func getRecipeFilter(r *http.Request) core.Filter {
	r.ParseForm()

	f := core.Filter{}
	if name := r.FormValue("name"); name != "" {
		f["name"] = name
		delete(r.Form, "name")
	}
	if size, err := asInt("size", r); err == nil {
		f["size"] = size
		delete(r.Form, "size")
	}
	if time, err := asInt("preptime", r); err == nil {
		f["preptime"] = time
		delete(r.Form, "preptime")
	}
	if time, err := asInt("cooktime", r); err == nil {
		f["cooktime"] = time
		delete(r.Form, "cooktime")
	}
	if time, err := asInt("misctime", r); err == nil {
		f["misctime"] = time
		delete(r.Form, "misctime")
	}

	for i := 0; i < core.RecipeType.NumField(); i++ {
		param := core.FoodType.Field(i).Tag.Get("json")
		if val, err := asFloat(param, r); err == nil {
			f[param] = val
		}
	}

	return f
}
