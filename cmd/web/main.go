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

package main

import (
	"heyapple/pkg/api/v1"
	"heyapple/pkg/middleware"
	"heyapple/pkg/storage/memory"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := memory.NewDB()

	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(middleware.Options)
	router.GET("/api/v1/foods", api.Foods(db))

	handler := middleware.Headers(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
