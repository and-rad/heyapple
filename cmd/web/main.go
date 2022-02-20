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
	"fmt"
	"heyapple/pkg/api/v1"
	"heyapple/pkg/app"
	"heyapple/pkg/auth"
	"heyapple/pkg/email"
	"heyapple/pkg/handler"
	"heyapple/pkg/l10n"
	"heyapple/pkg/mw"
	"heyapple/pkg/storage/memory"
	"heyapple/web"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/and-rad/scs/v2"
	"github.com/julienschmidt/httprouter"
)

func main() {
	out := app.NewLog(os.Stdout)
	out.Log("######################")
	out.Log("# Starting Hey Apple #")
	out.Log("######################")

	db := memory.NewDBWithBackup(out)
	tr := l10n.NewTranslator()
	nf := email.NewNotifier(tr)
	sm := scs.New()

	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(mw.Options)
	router.GET("/", handler.Home(db))
	router.GET("/app", mw.Auth(sm, "/login", handler.App(db)))
	router.GET("/login", mw.Anon(sm, "/app", handler.Login(db)))
	router.GET("/confirm/:token", handler.Confirm(db))

	router.POST("/auth/local", auth.LocalLogin(sm, db))
	router.DELETE("/auth/local", auth.LocalLogout(sm))

	router.POST("/api/v1/user", mw.JSON(api.NewUser(out, nf, db)))

	router.GET("/api/v1/foods", mw.JSON(api.Foods(db)))
	router.GET("/api/v1/food/:id", mw.JSON(api.Food(db)))
	router.POST("/api/v1/food", mw.JSON(api.NewFood(db)))
	router.POST("/api/v1/recipe", mw.JSON(api.NewRecipe(db)))
	router.PUT("/api/v1/food/:id", mw.JSON(api.SaveFood(db)))

	if dir := os.Getenv("HEYAPPLE_DATA_DIR"); dir != "" {
		router.ServeFiles("/data/*filepath", http.Dir(dir))
	}

	if sub, err := fs.Sub(web.Static, "static"); err == nil {
		router.ServeFiles("/static/*filepath", http.FS(sub))
	}

	log.Fatal(http.ListenAndServe(address(), sm.LoadAndSave(router)))
}

func address() string {
	host := os.Getenv("HEYAPPLE_WEB_HOST")
	port, err := strconv.Atoi(os.Getenv("HEYAPPLE_WEB_PORT"))
	if err != nil {
		port = 8080
	}
	return fmt.Sprintf("%s:%d", host, port)
}
