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
	"heyapple/pkg/core"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Foods returns a JSON-formatted list of all the food
// items stored in the database.
func Foods(db core.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cmd := &app.GetFoods{}
		if err := cmd.Fetch(db); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			sendResponse(cmd.Items, w)
		}
	}
}
