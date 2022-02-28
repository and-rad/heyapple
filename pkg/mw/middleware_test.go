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

package mw_test

import (
	"heyapple/internal/mock"
	"heyapple/pkg/handler"
	"heyapple/pkg/mw"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/and-rad/scs/v2"
)

func TestJSON(t *testing.T) {
	for idx, data := range []struct {
		header http.Header
		next   mock.Handler
		status int
	}{
		{ //00//
			next: mock.Handler{Body: "okay"},
			header: http.Header{
				"Access-Control-Allow-Origin": {"*"},
				"Content-Type":                {"application/json; charset=utf-8"},
			},
			status: http.StatusOK,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		mw.JSON()(data.next.Handle())(res, req, nil)

		if header := res.Result().Header; !reflect.DeepEqual(header, data.header) {
			t.Errorf("test case %d: header mismatch \nhave: %v \nwant: %v", idx, header, data.header)
		}

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.next.Body {
			t.Errorf("test case %d: body mismatch \nhave: %v \nwant: %v", idx, body, data.next.Body)
		}
	}
}

func TestOptions(t *testing.T) {
	for idx, data := range []struct {
		inhead  http.Header
		outhead http.Header
		status  int
	}{
		{ //00// no headers sent
			outhead: http.Header{},
			status:  http.StatusNoContent,
		},
		{ //01// no relevant headers sent
			inhead:  http.Header{"X-whazzup": {"whazzzzuuuuup"}},
			outhead: http.Header{},
			status:  http.StatusNoContent,
		},
		{ //02// CORS preflight header
			inhead: http.Header{
				"Access-Control-Request-Method": {"POST"},
			},
			outhead: http.Header{
				"Access-Control-Allow-Headers": {"Content-Type"},
				"Access-Control-Allow-Methods": {"DELETE, GET, OPTIONS, POST, PUT"},
				"Access-Control-Allow-Origin":  {"*"},
				"Access-Control-Max-Age":       {"86400"},
			},
			status: http.StatusNoContent,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		req.Header = data.inhead
		res := httptest.NewRecorder()
		mw.Options(res, req)

		if header := res.Result().Header; !reflect.DeepEqual(header, data.outhead) {
			t.Errorf("test case %d: header mismatch \nhave: %v \nwant: %v", idx, header, data.outhead)
		}

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}
	}
}

func TestAnon(t *testing.T) {
	for idx, data := range []struct {
		setCookie bool

		target string
		status int
	}{
		{ //00// anonymous user
			setCookie: false,
			status:    200,
		},
		{ //01// authenticated user
			setCookie: true,
			target:    "/overhere",
			status:    303,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{Session: scs.New()}

		if data.setCookie {
			if ctx, err := env.Session.Load(req.Context(), "abc"); err == nil {
				req = req.WithContext(ctx)
				env.Session.Put(req.Context(), "id", "hi")
			}
		}

		mw.Anon(env, data.target)((mock.Handler{}).Handle())(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if loc, err := res.Result().Location(); err == nil && loc.Path != data.target {
			t.Errorf("test case %d: location mismatch \nhave: %v\nwant: %v", idx, loc.Path, data.target)
		}
	}
}

func TestAuth(t *testing.T) {
	for idx, data := range []struct {
		setCookie bool

		target string
		status int
	}{
		{ //00// anonymous user
			setCookie: false,
			target:    "/somewhere",
			status:    303,
		},
		{ //01// authenticated
			setCookie: true,
			status:    200,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{Session: scs.New()}

		if data.setCookie {
			if ctx, err := env.Session.Load(req.Context(), "abc"); err == nil {
				req = req.WithContext(ctx)
				env.Session.Put(req.Context(), "id", "hi")
			}
		}

		mw.Auth(env, data.target)((mock.Handler{}).Handle())(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if loc, err := res.Result().Location(); err == nil && loc.Path != data.target {
			t.Errorf("test case %d: location mismatch \nhave: %v\nwant: %v", idx, loc.Path, data.target)
		}
	}
}

func TestCSRF(t *testing.T) {
	for idx, data := range []struct {
		env    map[string]string
		method string

		msg    string
		status int
	}{
		{ //00// missing env var, warn about it
			env:    map[string]string{"HEYAPPLE_CSRF_KEY": "1234abcd"},
			msg:    mw.ErrNoSecure.Error(),
			status: 200,
		},
		{ //01// missing env var, warn about it
			env:    map[string]string{"HEYAPPLE_CSRF_SECURE": "true"},
			msg:    mw.ErrNoKey.Error(),
			status: 200,
		},
		{ //02// insecure env var, warn about it
			env: map[string]string{
				"HEYAPPLE_CSRF_KEY":    "1234abcd",
				"HEYAPPLE_CSRF_SECURE": "false",
			},
			msg:    mw.ErrNotSecure.Error(),
			status: 200,
		},
		{ //03// always allow GET
			env: map[string]string{
				"HEYAPPLE_CSRF_KEY":    "1234abcd",
				"HEYAPPLE_CSRF_SECURE": "true",
			},
			method: http.MethodGet,
			status: 200,
		},
		{ //04// redirect POST if no token is present
			env: map[string]string{
				"HEYAPPLE_CSRF_KEY":    "1234abcd",
				"HEYAPPLE_CSRF_SECURE": "true",
			},
			method: http.MethodPost,
			status: 403,
		},
	} {
		os.Clearenv()
		for k, v := range data.env {
			os.Setenv(k, v)
			defer os.Unsetenv(k)
		}

		req := httptest.NewRequest(data.method, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{Log: mock.NewLog()}

		mw.CSRF(env, mock.Handler{}).ServeHTTP(res, req)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if msg := env.Log.(*mock.Log).Warning; msg != data.msg {
			t.Errorf("test case %d: log mismatch \nhave: %v\nwant: %v", idx, msg, data.msg)
		}
	}
}
