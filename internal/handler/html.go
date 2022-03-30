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

package handler

import (
	"heyapple/internal/app"
	"heyapple/internal/web"
	"html/template"
	"net/http"

	"github.com/and-rad/scs/v2"
	"github.com/gorilla/csrf"
	"github.com/julienschmidt/httprouter"
)

func Ping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func Home(env *Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data := map[string]interface{}{}
		lang, _ := sessionData(env.Session, r)
		l10n := func(in interface{}) string { return env.L10n.Translate(in, lang) }
		tpl := template.Must(web.Home.Clone()).Funcs(template.FuncMap{"l10n": l10n})
		if err := tpl.ExecuteTemplate(w, "home.html", data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func Login(env *Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data := map[string]interface{}{"csrf": csrf.Token(r)}
		lang, _ := sessionData(env.Session, r)
		l10n := func(in interface{}) string { return env.L10n.Translate(in, lang) }
		tpl := template.Must(web.Login.Clone()).Funcs(template.FuncMap{"l10n": l10n})
		if err := tpl.ExecuteTemplate(w, "login.html", data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func App(env *Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		lang, perm := sessionData(env.Session, r)
		data := map[string]interface{}{
			"csrf": csrf.Token(r),
			"perm": perm,
		}

		l10n := func(in interface{}) string { return env.L10n.Translate(in, lang) }
		tpl := template.Must(web.App.Clone()).Funcs(template.FuncMap{"l10n": l10n})
		if err := tpl.ExecuteTemplate(w, "app.html", data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func Legal(env *Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data := map[string]interface{}{}
		lang, _ := sessionData(env.Session, r)
		l10n := func(in interface{}) string { return env.L10n.Translate(in, lang) }
		tpl := template.Must(web.Legal.Clone()).Funcs(template.FuncMap{"l10n": l10n})
		if err := tpl.ExecuteTemplate(w, "legal.html", data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func Privacy(env *Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data := map[string]interface{}{}
		lang, _ := sessionData(env.Session, r)
		l10n := func(in interface{}) string { return env.L10n.Translate(in, lang) }
		tpl := template.Must(web.Privacy.Clone()).Funcs(template.FuncMap{"l10n": l10n})
		if err := tpl.ExecuteTemplate(w, "privacy.html", data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func Terms(env *Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data := map[string]interface{}{}
		lang, _ := sessionData(env.Session, r)
		l10n := func(in interface{}) string { return env.L10n.Translate(in, lang) }
		tpl := template.Must(web.Terms.Clone()).Funcs(template.FuncMap{"l10n": l10n})
		if err := tpl.ExecuteTemplate(w, "terms.html", data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func Confirm(env *Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data := map[string]interface{}{}
		if token := ps.ByName("token"); token == "" {
			data["status"] = http.StatusBadRequest
		} else if err := env.DB.Execute(&app.Activate{Token: token}); err == app.ErrNotFound {
			data["status"] = http.StatusNotFound
		} else if err != nil {
			data["status"] = http.StatusInternalServerError
		} else {
			data["status"] = http.StatusOK
		}

		lang, _ := sessionData(env.Session, r)
		l10n := func(in interface{}) string { return env.L10n.Translate(in, lang) }
		tpl := template.Must(web.Confirm.Clone()).Funcs(template.FuncMap{"l10n": l10n})
		if err := tpl.ExecuteTemplate(w, "confirm.html", data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func Reset(env *Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data := map[string]interface{}{
			"csrf":  csrf.Token(r),
			"token": ps.ByName("token"),
		}

		lang, _ := sessionData(env.Session, r)
		l10n := func(in interface{}) string { return env.L10n.Translate(in, lang) }
		tpl := template.Must(web.Reset.Clone()).Funcs(template.FuncMap{"l10n": l10n})
		if err := tpl.ExecuteTemplate(w, "reset.html", data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func sessionData(sm *scs.SessionManager, r *http.Request) (lang string, perm int) {
	if l, ok := sm.Get(r.Context(), "lang").(string); ok && l != "" {
		lang = l
	}
	if langs := r.Header.Get("Accept-Language"); langs != "" {
		lang = langs
	}
	if p, ok := sm.Get(r.Context(), "perm").(int); ok {
		perm = p
	}

	return
}
