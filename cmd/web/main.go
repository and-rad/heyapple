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
	"heyapple/pkg/storage"
	"heyapple/pkg/storage/memory"
	"heyapple/web"
	"io/fs"
	"net/http"
	"os"
	"strconv"

	"github.com/and-rad/scs/v2"
	"github.com/julienschmidt/httprouter"
)

func main() {
	log := app.NewLog(os.Stdout)
	translator := l10n.NewTranslator()
	notifier := email.NewNotifier(translator)
	db := memory.NewDB().WithBackup(log).WithDefaults(storage.Defaults())
	sessions := scs.New()

	env := &handler.Environment{
		DB:      db,
		Log:     log,
		Msg:     notifier,
		L10n:    translator,
		Session: sessions,
	}

	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(mw.Options)
	router.GET("/", chain(handler.Home(env), mw.Anon(env, "/app")))
	router.GET("/app", chain(handler.App(env), mw.Auth(env, "/auth")))
	router.GET("/auth", chain(handler.Login(env), mw.Anon(env, "/app")))
	router.GET("/legal", handler.Legal(env))
	router.GET("/privacy", handler.Privacy(env))
	router.GET("/terms", handler.Terms(env))
	router.GET("/confirm/:token", handler.Confirm(env))
	router.GET("/reset/:token", handler.Reset(env))

	router.POST("/auth/local", auth.LocalLogin(env))
	router.POST("/auth/reset", auth.ResetRequest(env))
	router.PUT("/auth/reset", auth.ResetConfirm(env))
	router.DELETE("/auth/local", auth.LocalLogout(env))

	router.GET("/api/v1/l10n/:lang", chain(api.L10n(env), mw.JSON()))
	router.POST("/api/v1/user", chain(api.NewUser(env), mw.JSON()))

	router.GET("/api/v1/foods", chain(api.Foods(env), mw.JSON()))
	router.GET("/api/v1/food/:id", chain(api.Food(env), mw.JSON()))
	router.POST("/api/v1/food", chain(api.NewFood(env), mw.JSON()))
	router.PUT("/api/v1/food/:id", chain(api.SaveFood(env), mw.JSON()))

	router.GET("/api/v1/recipes", chain(api.Recipes(env), mw.JSON()))
	router.GET("/api/v1/recipe/:id", chain(api.Recipe(env), mw.JSON()))
	router.GET("/api/v1/recipe/:id/owner", chain(api.RecipeOwner(env), mw.JSON()))
	router.POST("/api/v1/recipe", chain(api.NewRecipe(env), mw.JSON()))
	router.POST("/api/v1/recipe/:id/ingredient/:ing", chain(api.SaveIngredient(env), mw.JSON()))
	router.PUT("/api/v1/recipe/:id", chain(api.SaveRecipe(env), mw.JSON()))
	router.PUT("/api/v1/recipe/:id/ingredient/:ing", chain(api.SaveIngredient(env), mw.JSON()))

	router.GET("/api/v1/diary", chain(api.Diary(env), mw.JSON()))
	router.GET("/api/v1/diary/:year", chain(api.Diary(env), mw.JSON()))
	router.GET("/api/v1/diary/:year/:month", chain(api.Diary(env), mw.JSON()))
	router.GET("/api/v1/diary/:year/:month/:day", chain(api.Diary(env), mw.JSON()))
	router.GET("/api/v1/diary/:year/:month/:day/entries", chain(api.DiaryEntries(env), mw.JSON()))
	router.POST("/api/v1/diary/:date", chain(api.SaveDiaryEntry(env), mw.JSON()))
	router.PUT("/api/v1/diary/:date", chain(api.SaveDiaryEntry(env), mw.JSON()))

	if dir := os.Getenv("HEYAPPLE_DATA_DIR"); dir != "" {
		router.ServeFiles("/data/*filepath", http.Dir(dir))
	}

	if sub, err := fs.Sub(web.Static, "static"); err == nil {
		router.NotFound = http.FileServer(http.FS(sub))
	}

	log.Log("######################")
	log.Log("# Starting Hey Apple #")
	log.Log("######################")

	log.Error(http.ListenAndServe(address(),
		env.Session.LoadAndSave(mw.CSRF(env, router))))
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
