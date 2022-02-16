package auth

import (
	"heyapple/pkg/app"
	"net/http"

	"github.com/and-rad/scs/v2"
	"github.com/julienschmidt/httprouter"
)

func LocalLogin(sm *scs.SessionManager, db app.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		query := &app.Authenticate{
			Email: r.FormValue("email"),
			Pass:  r.FormValue("pass"),
		}

		if query.Email == "" || query.Pass == "" {
			w.WriteHeader(http.StatusUnauthorized)
		} else if err := db.Fetch(query); err == app.ErrCredentials {
			w.WriteHeader(http.StatusUnauthorized)
		} else if err == app.ErrNotFound {
			w.WriteHeader(http.StatusUnauthorized)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			sm.Put(r.Context(), "id", query.ID)
		}
	}
}

func LocalLogout(db app.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	}
}
