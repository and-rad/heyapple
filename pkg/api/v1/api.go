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

// Package api implements a REST API to interact with the app over
// HTTP requests. This is v1 of the API.
//
// All objects returned from the functions in this package are valid
// JSON values, including strings, numbers, objects, arrays, and null.
// Look at the documentation for individual functions to learn about
// API endpoints, possible status codes, and example output.
package api

import (
	"encoding/json"
	"net/http"
)

func sendResponse(data interface{}, w http.ResponseWriter) {
	if data, err := json.Marshal(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(data)
	}
}
