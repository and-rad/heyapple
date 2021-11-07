package api

import (
	"encoding/json"
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Foods returns a JSON-formatted list of all the food
// items stored in the database.
func Foods(db core.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cmd := &app.GetFoods{}
		if err := cmd.Fetch(db); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else if data, err := json.Marshal(cmd.Items); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write(data)
		}
	}
}
