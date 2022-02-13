package handler

import (
	"heyapple/pkg/app"
	"heyapple/web"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Home(db app.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := web.Home.ExecuteTemplate(w, "home.html", struct{}{}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func Login(db app.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := web.Login.ExecuteTemplate(w, "login.html", struct{}{}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func App(db app.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := web.App.ExecuteTemplate(w, "app.html", struct{}{}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
