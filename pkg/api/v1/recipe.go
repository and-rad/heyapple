package api

import (
	"heyapple/pkg/app"
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
//   400 - Missing form data
//   500 - Internal server error
// Example input:
//   name=Pie
// Example output:
//   42
func NewRecipe(db app.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cmd := &app.CreateRecipe{Name: r.FormValue("name")}
		if cmd.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else if err := db.Execute(cmd); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
			sendResponse(cmd.ID, w)
		}
	}
}

// SaveRecipe updates a recipe in the database identified
// by the {id} parameter. The ingredients are passed in
// the request body. The response body will always be
// empty, success or failure is communicated by the status
// codes.
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
//   404 - Recipe doesn't exist
//   500 - Internal server error
// Example input:
//   size=12&item=1&item=4&amount=150&amount=255
func SaveRecipe(db app.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		recID, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		r.ParseForm()
		cmd := &app.SaveRecipe{ID: recID}
		if size, err := strconv.Atoi(r.PostForm.Get("size")); err == nil {
			cmd.Size = size
		}

		ids := r.PostForm["item"]
		amounts := r.PostForm["amount"]
		if len(ids) != len(amounts) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		cmd.Items = make(map[int]float32)
		for k := range ids {
			if id, err := strconv.Atoi(ids[k]); err == nil {
				if amount, err := strconv.ParseFloat(amounts[k], 32); err == nil {
					cmd.Items[id] = float32(amount)
				}
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
