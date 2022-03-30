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
	"strconv"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/core"
)

func sendResponse(data interface{}, w http.ResponseWriter) {
	if data, err := json.Marshal(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(data)
	}
}

// asInt checks if param exists in the request's form and
// if so, returns it as int or int range if there are
// several values found.
func asInt(param string, r *http.Request) (interface{}, error) {
	v, ok := r.Form[param]
	if !ok {
		return nil, app.ErrNotFound
	}

	v1, err := strconv.Atoi(v[0])
	if err != nil {
		return nil, err
	}

	v2 := v1
	if len(v) > 1 {
		if v2, err = strconv.Atoi(v[1]); err != nil {
			return nil, err
		}
	}

	if v1 != v2 {
		return core.IntRange{v1, v2}, nil
	}

	return v1, nil
}

// asFloat checks if param exists in the request's form and
// if so, returns it as float32 or float32 range if there are
// several values found.
func asFloat(param string, r *http.Request) (interface{}, error) {
	v, ok := r.Form[param]
	if !ok {
		return nil, app.ErrNotFound
	}

	v1, err := strconv.ParseFloat(v[0], 32)
	if err != nil {
		return nil, err
	}

	v2 := v1
	if len(v) > 1 {
		if v2, err = strconv.ParseFloat(v[1], 32); err != nil {
			return nil, err
		}
	}

	if v1 != v2 {
		return core.FloatRange{float32(v1), float32(v2)}, nil
	}

	return float32(v1), nil
}
