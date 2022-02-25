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
			sm.Put(r.Context(), "lang", query.Lang)
		}
	}
}

func LocalLogout(sm *scs.SessionManager) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := sm.Destroy(r.Context()); err == scs.ErrNoSession {
			w.WriteHeader(http.StatusNotFound)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}
