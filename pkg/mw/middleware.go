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

// Package mw defines various pieces of middleware that conform to the
// http.Handler and the httprouter.Handle interfaces.
package mw

import (
	"errors"
	"heyapple/pkg/handler"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/csrf"
	"github.com/julienschmidt/httprouter"
)

var (
	ErrNoKey     = errors.New("HEYAPPLE_CSRF_KEY environment variable not found. Using insecure default value. Do NOT run this in a production environment!")
	ErrNoSecure  = errors.New("HEYAPPLE_CSRF_SECURE environment variable not found. CSRF cookies will be sent over insecure connections. Do NOT run this in a production environment!")
	ErrNotSecure = errors.New("HEYAPPLE_CSRF_SECURE is set to 'false'. CSRF cookies will be sent over insecure connections. Do NOT run this in a production environment!")
)

type Func func(httprouter.Handle) httprouter.Handle

// Anon is a middleware that prevents logged-in users from accessing a resource.
func Anon(env *handler.Environment, target string) Func {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			if env.Session.Get(r.Context(), "id") == nil {
				next(w, r, ps)
			} else {
				http.Redirect(w, r, target, http.StatusSeeOther)
			}
		}
	}
}

// Auth is a middleware that grants only logged-in users access to a resource.
func Auth(env *handler.Environment, target string) Func {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			if env.Session.Get(r.Context(), "id") != nil {
				next(w, r, ps)
			} else {
				http.Redirect(w, r, target, http.StatusSeeOther)
			}
		}
	}
}

// JSON is a middleware to identify response bodies as JSON data.
func JSON() Func {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			next(w, r, ps)
		}
	}
}

// Options is a middleware that tells clients about CORS capabilities.
func Options(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	// check for CORS headers
	if r.Header.Get("Access-Control-Request-Method") != "" {
		header.Set("Access-Control-Allow-Headers", "Content-Type")
		header.Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Max-Age", "86400")
	}

	w.WriteHeader(http.StatusNoContent)
}

func CSRF(env *handler.Environment, next http.Handler) http.Handler {
	key := os.Getenv("HEYAPPLE_CSRF_KEY")
	if key == "" {
		env.Log.Warn(ErrNoKey)
		key = "1f2e3d4c5b6a7a8b9c1d2f"
	}

	var secure bool
	if secEnv, ok := os.LookupEnv("HEYAPPLE_CSRF_SECURE"); !ok {
		env.Log.Warn(ErrNoSecure)
	} else if b, err := strconv.ParseBool(secEnv); err != nil || !b {
		env.Log.Warn(ErrNotSecure)
	} else {
		secure = true
	}

	protect := csrf.Protect([]byte(key),
		csrf.CookieName("_csrf"),
		csrf.Secure(secure),
	)

	return protect(next)
}
