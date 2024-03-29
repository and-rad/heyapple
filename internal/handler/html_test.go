////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2021-2024 The HeyApple Authors.
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
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/and-rad/heyapple/internal/handler"
	"github.com/and-rad/heyapple/internal/mock"
	"github.com/and-rad/heyapple/internal/web"

	"github.com/and-rad/scs/v2"
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
