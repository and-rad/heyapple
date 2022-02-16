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
	"heyapple/pkg/app"
	"heyapple/pkg/auth"
	"heyapple/pkg/handler"
	"heyapple/pkg/storage/memory"
	"heyapple/web"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/and-rad/scs/v2"
	"github.com/julienschmidt/httprouter"
)

func main() {
	out := app.NewLog(os.Stdout)
	out.Log("######################")
	out.Log("# Starting Hey Apple #")
	out.Log("######################")

	db := memory.NewDBWithBackup(out)
	sm := scs.New()

	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(handler.Options)
	router.GET("/", handler.Home(db))
	router.GET("/app", handler.Auth(sm, "/login", handler.App(db)))
	router.GET("/login", handler.Anon(sm, "/app", handler.Login(db)))

	router.POST("/auth/local", auth.LocalLogin(sm, db))
	router.DELETE("/auth/local", auth.LocalLogout(db))

	router.GET("/api/v1/foods", handler.JSON(api.Foods(db)))
	router.GET("/api/v1/food/:id", handler.JSON(api.Food(db)))
	router.POST("/api/v1/food", handler.JSON(api.NewFood(db)))
	router.PUT("/api/v1/food/:id", handler.JSON(api.SaveFood(db)))

	if dir := os.Getenv("HEYAPPLE_DATA_DIR"); dir != "" {
		router.ServeFiles("/data/*filepath", http.Dir(dir))
	}

	if sub, err := fs.Sub(web.Static, "static"); err == nil {
		router.ServeFiles("/static/*filepath", http.FS(sub))
	}

	log.Fatal(http.ListenAndServe(":8080", sm.LoadAndSave(router)))
}
