package auth

import (
	"heyapple/pkg/app"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func LocalLogin(db app.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	}
}

func LocalLogout(db app.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	}
}
