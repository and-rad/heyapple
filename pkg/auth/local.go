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

package auth

import (
	"heyapple/pkg/app"
	"heyapple/pkg/handler"
	"net/http"

	"github.com/and-rad/scs/v2"
	"github.com/julienschmidt/httprouter"
)

// LocaLogin handles login for users with local accounts
// as opposed to users who authenticate with SSO services
// like OAuth. The response body is always empty.
//
// Endpoint:
//   /auth/local
// Methods:
//   POST
// Possible status codes:
//   200 - Login successful
//   400 - Malformed or missing form data
//   401 - Unsuccessful login attempt
//   500 - Internal server error
// Example input:
//   email=user@example.com&pass=topsecret
func LocalLogin(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		query := &app.Authenticate{
			Email: r.FormValue("email"),
			Pass:  r.FormValue("pass"),
		}

		if query.Email == "" || query.Pass == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else if err := env.DB.Fetch(query); err == app.ErrCredentials {
			w.WriteHeader(http.StatusUnauthorized)
		} else if err == app.ErrNotFound {
			w.WriteHeader(http.StatusUnauthorized)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			env.Session.Put(r.Context(), "id", query.ID)
			env.Session.Put(r.Context(), "lang", query.Lang)
		}
	}
}

// LocaLogout handles logout for users with local accounts
// as opposed to users who authenticate with SSO services
// like OAuth. The response body is always empty.
//
// Endpoint:
//   /auth/local
// Methods:
//   DELETE
// Possible status codes:
//   200 - Logout successful
//   404 - Session not found, user is not logged in
//   500 - Internal server error
func LocalLogout(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := env.Session.Destroy(r.Context()); err == scs.ErrNoSession {
			w.WriteHeader(http.StatusNotFound)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

// ResetRequest creates a new password reset request and
// sends a notification to the user with instructions
// on how to complete the request. The response body is
// always empty.
//
// Endpoint:
//   /auth/reset
// Methods:
//   POST
// Possible status codes:
//   200 - Request creation successful
//   400 - Malformed or missing form data
//   404 - User doesn't exist
//   500 - Internal server error
// Example input:
//   email=user@example.com
func ResetRequest(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cmd := &app.ResetPassword{Email: r.FormValue("email")}
		if cmd.Email == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := env.DB.Execute(cmd)
		if err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		data := app.NotificationData{
			"lang":  "en",
			"token": cmd.Token,
		}
		err = env.Msg.Send(cmd.Email, app.ResetNotification, data)
		if err != nil {
			env.Log.Error(err)
		}
	}
}

// ResetConfirm completes a password reset request. If
// successful, the associated token is deleted and the
// password is changed. The response body is always empty.
//
// Endpoint:
//   /auth/reset
// Methods:
//   PUT
// Possible status codes:
//   200 - Password reset successful
//   400 - Malformed or missing form data
//   404 - User or token doesn't exist
//   500 - Internal server error
// Example input:
//   email=user@example.com&pass=topsecret
func ResetConfirm(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cmd := &app.ChangePassword{
			Token: r.FormValue("token"),
			Pass:  r.FormValue("pass"),
		}

		if cmd.Token == "" || cmd.Pass == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else if err := env.DB.Execute(cmd); err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}
