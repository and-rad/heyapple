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

package handler_test

import (
	"heyapple/internal/mock"
	"heyapple/pkg/app"
	"heyapple/pkg/handler"
	"heyapple/web"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestHome(t *testing.T) {
	home := web.Home

	for idx, data := range []struct {
		tmpl string

		out    string
		status int
	}{
		{ //00// function does not exist
			tmpl:   `{{ .Foo "Bar" }}`,
			status: 500,
		},
		{ //01// success
			tmpl:   `{{ if true }}hi{{ end }}`,
			status: 200,
			out:    "hi",
		},
	} {
		web.Home = template.Must(template.New("home.html").Parse(data.tmpl))
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		handler.Home(mock.NewDB())(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, body, data.out)
		}
	}

	web.Home = home
}

func TestLogin(t *testing.T) {
	login := web.Login

	for idx, data := range []struct {
		tmpl string

		out    string
		status int
	}{
		{ //00// function does not exist
			tmpl:   `{{ .Foo "Bar" }}`,
			status: 500,
		},
		{ //01// success
			tmpl:   `{{ if true }}hi{{ end }}`,
			status: 200,
			out:    "hi",
		},
	} {
		web.Login = template.Must(template.New("login.html").Parse(data.tmpl))
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		handler.Login(mock.NewDB())(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, body, data.out)
		}
	}

	web.Login = login
}

func TestApp(t *testing.T) {
	app := web.App

	for idx, data := range []struct {
		tmpl string

		out    string
		status int
	}{
		{ //00// function does not exist
			tmpl:   `{{ .Foo "Bar" }}`,
			status: 500,
		},
		{ //01// success
			tmpl:   `{{ if true }}hi{{ end }}`,
			status: 200,
			out:    "hi",
		},
	} {
		web.App = template.Must(template.New("app.html").Parse(data.tmpl))
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		handler.App(mock.NewDB())(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, body, data.out)
		}
	}

	web.App = app
}

func TestConfirm(t *testing.T) {
	tmp := web.Confirm

	for idx, data := range []struct {
		db   *mock.DB
		ps   httprouter.Params
		tmpl string

		out    string
		status int
	}{
		{ //00// token missing, bad request
			tmpl:   `{{ .status }}`,
			out:    "400",
			status: 200,
		},
		{ //01// token doesn't exist
			db:     mock.NewDB(),
			ps:     httprouter.Params{{Key: "token", Value: "abcd"}},
			tmpl:   `{{ .status }}`,
			out:    "404",
			status: 200,
		},
		{ //02// database failure
			db:     mock.NewDB().WithError(mock.ErrDOS),
			ps:     httprouter.Params{{Key: "token", Value: "abcd"}},
			tmpl:   `{{ .status }}`,
			out:    "500",
			status: 200,
		},
		{ //03// function does not exist
			tmpl:   `{{ .Foo "Bar" }}`,
			status: 500,
		},
		{ //04// token successfully processed
			db:     mock.NewDB().WithToken(app.Token{ID: 1}).WithUser(app.User{ID: 1}),
			ps:     httprouter.Params{{Key: "token", Value: "abcd"}},
			tmpl:   `{{ .status }}`,
			out:    "200",
			status: 200,
		},
		{ //05// success
			tmpl:   `{{ if true }}hi{{ end }}`,
			status: 200,
			out:    "hi",
		},
	} {
		web.Confirm = template.Must(template.New("confirm.html").Parse(data.tmpl))
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		handler.Confirm(data.db)(res, req, data.ps)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, body, data.out)
		}
	}

	web.Confirm = tmp
}
