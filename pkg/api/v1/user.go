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
	"heyapple/pkg/app"
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
func NewUser(log app.Logger, nf app.Notifier, db app.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cmd := &app.CreateUser{
			Email: r.FormValue("email"),
			Pass:  r.FormValue("pass"),
		}

		if cmd.Email == "" || cmd.Pass == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := db.Execute(cmd)
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
		if err = nf.Send(cmd.Email, app.RegisterNotification, data); err != nil {
			log.Error(err)
		}
	}
}
