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

package auth

import (
	"context"
	"net/http"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/handler"

	"github.com/and-rad/scs/v2"
	"github.com/julienschmidt/httprouter"
)

// Confirm completes the user registration by confirming
// the sign-up token. If successful, the associated token
// is deleted and the user is able to sign in. The
// response body is always empty.
//
// Endpoint:
//
//	/auth/confirm
//
// Methods:
//
//	PUT
//
// Possible status codes:
//
//	200 - Registration complete
//	400 - Malformed or missing form data
//	404 - User or token doesn't exist
//	500 - Internal server error
//
// Example input:
//
//	token=178a6ee3f1da299fed940aa2d7
func Confirm(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cmd := &app.Activate{Token: r.FormValue("token")}

		if cmd.Token == "" {
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

// LocaLogin handles login for users with local accounts
// as opposed to users who authenticate with SSO services
// like OAuth. The response body is always empty.
//
// Endpoint:
//
//	/auth/local
//
// Methods:
//
//	POST
//
// Possible status codes:
//
//	200 - Login successful
//	400 - Malformed or missing form data
//	401 - Unsuccessful login attempt
//	500 - Internal server error
//
// Example input:
//
//	email=user@example.com&pass=topsecret
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
			env.Session.Put(r.Context(), "perm", query.Perm)
		}
	}
}

// LocaLogout handles logout for users with local accounts
// as opposed to users who authenticate with SSO services
// like OAuth. The response body is always empty.
//
// Endpoint:
//
//	/auth/local
//
// Methods:
//
//	DELETE
//
// Possible status codes:
//
//	200 - Logout successful
//	404 - Session not found, user is not logged in
//	500 - Internal server error
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
// always empty. For security reasons, this is one of a handful
// of functions that return success status codes even when
// technically failing. This is done to make user enumeration
// more difficult.
//
// Endpoint:
//
//	/auth/reset
//
// Methods:
//
//	POST
//
// Possible status codes:
//
//	202 - Request was accepted
//	400 - Malformed or missing form data
//	500 - Internal server error
//
// Example input:
//
//	email=user@example.com
func ResetRequest(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cmd := &app.ResetPassword{Email: r.FormValue("email")}
		if cmd.Email == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := env.DB.Execute(cmd)
		if err == app.ErrNotFound {
			w.WriteHeader(http.StatusAccepted)
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
//
//	/auth/reset
//
// Methods:
//
//	PUT
//
// Possible status codes:
//
//	200 - Password reset successful
//	400 - Malformed or missing form data
//	404 - User or token doesn't exist
//	422 - New password is too weak
//	500 - Internal server error
//
// Example input:
//
//	email=user@example.com&pass=topsecret
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
		} else if err == app.ErrWeakPass {
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

// ChangePassword sets a new password for the logged-in user.
// Unlike ResetConfirm, it doesn't require a token, but it
// checks the requesting user's identity and expects the user's
// current password. If successful, the  password is changed.
// The response body is always empty.
//
// This is the route that should be called when users try to
// change their passwords from within their profile settings.
//
// Endpoint:
//
//	/auth/pass
//
// Methods:
//
//	PUT
//
// Possible status codes:
//
//	200 - Password change successful
//	400 - Malformed or missing form data
//	401 - Insufficient permission
//	404 - User doesn't exist
//	422 - New password is too weak
//	500 - Internal server error
//
// Example input:
//
//	passold=topsecret&passnew=topsecret123
func ChangePassword(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		old := r.FormValue("passold")
		new := r.FormValue("passnew")
		if old == "" || new == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		uid, ok := env.Session.Get(r.Context(), "id").(int)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		query := &app.Authorize{ID: uid, Pass: old}
		if err := env.DB.Fetch(query); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else if !query.Ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		if !query.Ok {
			return
		}

		cmd := &app.ChangePassword{ID: uid, Pass: new}
		if err := env.DB.Execute(cmd); err == app.ErrWeakPass {
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

// logOut can be used to invalidate the session that
// belongs to the user identified by id. It should be
// called after making permission changes on that user
// as an extra security step.
func logOut(sm *scs.SessionManager, r *http.Request, id int) error {
	return sm.Iterate(r.Context(), func(ctx context.Context) error {
		if uid, ok := sm.Get(ctx, "id").(int); ok && uid == id {
			return sm.Destroy(ctx)
		}
		return nil
	})
}
