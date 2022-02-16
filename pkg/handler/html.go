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
