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

package api

import (
	"heyapple/internal/app"
	"heyapple/internal/handler"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NewUser creates a new user and sends out a registration
// notification on success. The response body is always
// empty.
//
// Endpoint:
//   /api/v1/user
// Methods:
//   POST
// Possible status codes:
//   201 - Creation successful
//   400 - Malformed or missing form data
//   500 - Internal server error
// Example input:
//   email=user@example.com&pass=topsecret
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
			w.WriteHeader(http.StatusConflict)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

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
//   /api/v1/l10n/{lang}
// Methods:
//   GET
// Possible status codes:
//   200 - OK
// Example output:
//   {
//     "app.name": "HeyApple",
//     "msg.hello": "What's up?",
//     ...
//   }
func L10n(env *handler.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tr := env.L10n.Get(ps.ByName("lang"))
		w.WriteHeader(http.StatusOK)
		sendResponse(tr, w)
	}
}