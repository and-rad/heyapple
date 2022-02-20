package api

import (
	"heyapple/pkg/app"
	"net/http"

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
