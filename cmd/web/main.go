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
	router.GET("/app", chain(handler.App(db), mw.Auth(sm, "/login")))
	router.GET("/login", chain(handler.Login(db), mw.Anon(sm, "/app")))
	router.GET("/confirm/:token", handler.Confirm(db))

	router.POST("/auth/local", auth.LocalLogin(sm, db))
	router.DELETE("/auth/local", auth.LocalLogout(sm))

	router.POST("/api/v1/user", chain(api.NewUser(out, nf, db), mw.JSON()))

	router.GET("/api/v1/foods", chain(api.Foods(db), mw.JSON()))
	router.GET("/api/v1/food/:id", chain(api.Food(db), mw.JSON()))
	router.POST("/api/v1/food", chain(api.NewFood(sm, db), mw.JSON()))
	router.PUT("/api/v1/food/:id", chain(api.SaveFood(sm, db), mw.JSON()))

	router.POST("/api/v1/recipe", chain(api.NewRecipe(sm, db), mw.JSON()))
	router.PUT("/api/v1/recipe/:id", chain(api.SaveRecipe(sm, db), mw.JSON()))

	if dir := os.Getenv("HEYAPPLE_DATA_DIR"); dir != "" {
		router.ServeFiles("/data/*filepath", http.Dir(dir))
	}

	if sub, err := fs.Sub(web.Static, "static"); err == nil {
		router.ServeFiles("/static/*filepath", http.FS(sub))
	}

	log.Fatal(http.ListenAndServe(address(), sm.LoadAndSave(router)))
}

// address builds the ListenAndServe address from the app config.
func address() string {
	host := os.Getenv("HEYAPPLE_WEB_HOST")
	port, err := strconv.Atoi(os.Getenv("HEYAPPLE_WEB_PORT"))
	if err != nil {
		port = 8080
	}
	return fmt.Sprintf("%s:%d", host, port)
}

// chain is a convenience function to chain middleware in
// a more readable manner.
func chain(h httprouter.Handle, m ...mw.Func) httprouter.Handle {
	handler := h
	for i := len(m) - 1; i >= 0; i-- {
		handler = m[i](handler)
	}
	return handler
}
