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
	"heyapple/internal/app"
	"heyapple/internal/handler"
	"heyapple/internal/mock"
	"heyapple/internal/web"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/and-rad/scs/v2"
	"github.com/julienschmidt/httprouter"
)

var (
	funcs = template.FuncMap{"l10n": func(interface{}) string { return "" }}
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
		{ //02// translate string
			tmpl:   `{{ l10n "msg.hi" }}`,
			status: 200,
			out:    "Hi!",
		},
	} {
		web.Home = template.Must(template.New("home.html").Funcs(funcs).Parse(data.tmpl))
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{
			DB:      mock.NewDB(),
			Session: scs.New(),
			L10n:    &mock.Translator{Map: map[string]interface{}{"msg.hi": "Hi!"}},
		}

		handler.Home(env)(res, req, nil)

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
		{ //02// translate string
			tmpl:   `{{ l10n "msg.hi" }}`,
			status: 200,
			out:    "Hi!",
		},
	} {
		web.Login = template.Must(template.New("login.html").Funcs(funcs).Parse(data.tmpl))
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{
			DB:      mock.NewDB(),
			Session: scs.New(),
			L10n:    &mock.Translator{Map: map[string]interface{}{"msg.hi": "Hi!"}},
		}

		handler.Login(env)(res, req, nil)

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
		{ //02// translate string
			tmpl:   `{{ l10n "msg.hi" }}`,
			status: 200,
			out:    "Hi!",
		},
	} {
		web.App = template.Must(template.New("app.html").Funcs(funcs).Parse(data.tmpl))
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{
			DB:      mock.NewDB(),
			Session: scs.New(),
			L10n:    &mock.Translator{Map: map[string]interface{}{"msg.hi": "Hi!"}},
		}

		handler.App(env)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, body, data.out)
		}
	}

	web.App = app
}

func TestLegal(t *testing.T) {
	tmp := web.Legal

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
		{ //02// translate string
			tmpl:   `{{ l10n "msg.hi" }}`,
			status: 200,
			out:    "Hi!",
		},
	} {
		web.Legal = template.Must(template.New("legal.html").Funcs(funcs).Parse(data.tmpl))
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{
			DB:      mock.NewDB(),
			Session: scs.New(),
			L10n:    &mock.Translator{Map: map[string]interface{}{"msg.hi": "Hi!"}},
		}

		handler.Legal(env)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, body, data.out)
		}
	}

	web.Legal = tmp
}

func TestPrivacy(t *testing.T) {
	tmp := web.Privacy

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
		{ //02// translate string
			tmpl:   `{{ l10n "msg.hi" }}`,
			status: 200,
			out:    "Hi!",
		},
	} {
		web.Privacy = template.Must(template.New("privacy.html").Funcs(funcs).Parse(data.tmpl))
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{
			DB:      mock.NewDB(),
			Session: scs.New(),
			L10n:    &mock.Translator{Map: map[string]interface{}{"msg.hi": "Hi!"}},
		}

		handler.Privacy(env)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, body, data.out)
		}
	}

	web.Home = tmp
}

func TestTerms(t *testing.T) {
	tmp := web.Terms

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
		{ //02// translate string
			tmpl:   `{{ l10n "msg.hi" }}`,
			status: 200,
			out:    "Hi!",
		},
	} {
		web.Terms = template.Must(template.New("terms.html").Funcs(funcs).Parse(data.tmpl))
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{
			DB:      mock.NewDB(),
			Session: scs.New(),
			L10n:    &mock.Translator{Map: map[string]interface{}{"msg.hi": "Hi!"}},
		}

		handler.Terms(env)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, body, data.out)
		}
	}

	web.Home = tmp
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
		{ //06// translate string
			tmpl:   `{{ l10n "msg.hi" }}`,
			status: 200,
			out:    "Hi!",
		},
	} {
		web.Confirm = template.Must(template.New("confirm.html").Funcs(funcs).Parse(data.tmpl))
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{
			DB:      data.db,
			Session: scs.New(),
			L10n:    &mock.Translator{Map: map[string]interface{}{"msg.hi": "Hi!"}},
		}

		handler.Confirm(env)(res, req, data.ps)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, body, data.out)
		}
	}

	web.Confirm = tmp
}

func TestReset(t *testing.T) {
	tmp := web.Reset

	for idx, data := range []struct {
		db   *mock.DB
		ps   httprouter.Params
		tmpl string

		out    string
		status int
	}{
		{ //00// token missing
			tmpl:   `{{ .token }}`,
			status: 200,
			out:    "",
		},
		{ //01// function does not exist
			tmpl:   `{{ .Foo "Bar" }}`,
			status: 500,
		},
		{ //02// success
			tmpl:   `{{ .token }}`,
			ps:     httprouter.Params{{Key: "token", Value: "abcd"}},
			status: 200,
			out:    "abcd",
		},
		{ //03// translate string
			tmpl:   `{{ l10n "msg.hi" }}`,
			status: 200,
			out:    "Hi!",
		},
	} {
		web.Reset = template.Must(template.New("reset.html").Funcs(funcs).Parse(data.tmpl))
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{
			DB:      data.db,
			Session: scs.New(),
			L10n:    &mock.Translator{Map: map[string]interface{}{"msg.hi": "Hi!"}},
		}

		handler.Reset(env)(res, req, data.ps)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, body, data.out)
		}
	}

	web.Reset = tmp
}

func TestPing(t *testing.T) {
	for idx, data := range []struct {
		status int
	}{
		{ //00//
			status: http.StatusOK,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()

		handler.Ping(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}
	}
}
