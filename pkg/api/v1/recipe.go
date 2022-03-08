package api

import (
	"heyapple/pkg/app"
	"heyapple/pkg/handler"
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
