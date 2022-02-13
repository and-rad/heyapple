package handler_test

import (
	"heyapple/internal/mock"
	"heyapple/pkg/handler"
	"heyapple/web"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
