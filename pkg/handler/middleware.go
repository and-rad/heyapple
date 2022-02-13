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

// Package web defines handlers for static web content as well as
// various pieces of middleware that conform to the http.Handler and the
// httprouter.Handle interfaces.
package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// JSON is a middleware to identify response bodies as JSON data.
func JSON(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next(w, r, ps)
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
