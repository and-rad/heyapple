////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2021-2024 The HeyApple Authors.
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
	"net/http"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/handler"

	"github.com/julienschmidt/httprouter"
)

// NewUser creates a new user and sends out a registration
// notification on success. The response body is always
// empty. For security reasons, this is one of a handful
// of functions that return success status codes even when
// technically failing. This is done to make user enumeration
// more difficult.
//
// Endpoint:
//
//	/api/v1/user
//
// Methods:
//
//	POST
//
// Possible status codes:
//
//	202 - Creation request accepted
//	400 - Malformed or missing form data
//	422 - The provided password is not strong enough
//	500 - Internal server error
//
// Example input:
//
//	email=user@example.com&pass=topsecret
func NewUser(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cmd := &app.CreateUser{
			Email: r.FormValue("email"),
			Pass:  r.FormValue("pass"),
		}

		if cmd.Email == "" || cmd.Pass == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := env.DB.Execute(cmd)
		if err == app.ErrExists {
			w.WriteHeader(http.StatusAccepted)
			return
		}
		if err == app.ErrWeakPass {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)

		data := app.NotificationData{
			"lang":  "en",
			"token": cmd.Token,
		}
		err = env.Msg.Send(cmd.Email, app.RegisterNotification, data)
		if err != nil {
			env.Log.Error(err)
		}
	}
}

// L10n returns all localized strings for the language
// {lang}. The response body will be a JSON-formatted
// collection of key-value pairs.
//
// Endpoint:
//
//	/api/v1/l10n/{lang}
//
// Methods:
//
//	GET
//
// Possible status codes:
//
//	200 - OK
//
// Example output:
//
//	{
//	  "app.name": "HeyApple",
//	  "msg.hello": "What's up?",
//	  ...
//	}
func L10n(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tr := env.L10n.Get(ps.ByName("lang"))
		w.WriteHeader(http.StatusOK)
		sendResponse(tr, w)
	}
}

// Preferences returns the account and app settings
// for the session user.
//
// Endpoint:
//
//	/api/v1/prefs
//
// Methods:
//
//	GET
//
// Possible status codes:
//
//	200 - OK
//	404 - User doesn't exist
//	500 - Internal server error
//
// Example output:
//
//	{
//	  "account": { "email": "user@example.com", ... },
//	  "rdi": { "kcal": 2000, "fat": 60, ... },
//	  ...
//	}
func Preferences(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		uid, ok := env.Session.Get(r.Context(), "id").(int)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		query := &app.Preferences{ID: uid}
		if err := env.DB.Fetch(query); err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			sendResponse(query.Prefs, w)
		}
	}
}
