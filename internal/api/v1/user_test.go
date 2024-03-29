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

package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/and-rad/heyapple/internal/api/v1"
	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/handler"
	"github.com/and-rad/heyapple/internal/mock"
	"github.com/and-rad/scs/v2"

	"github.com/julienschmidt/httprouter"
)

func TestNewUser(t *testing.T) {
	for idx, data := range []struct {
		db *mock.DB
		nf *mock.Notifier
		in url.Values

		err    string
		status int
	}{
		{ //00// no data
			db:     mock.NewDB(),
			status: http.StatusBadRequest,
		},
		{ //01// weak password
			db:     mock.NewDB(),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"topsecret"}},
			status: http.StatusUnprocessableEntity,
		},
		{ //02// connection failure
			db:     mock.NewDB().WithError(mock.ErrDOS),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusInternalServerError,
		},
		{ //03// notification failure
			db:     mock.NewDB(),
			nf:     mock.NewNotifier().WithError(mock.ErrDOS),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusAccepted,
			err:    mock.ErrDOS.Error(),
		},
		{ //04// username exists
			db:     mock.NewDB().WithError(app.ErrExists),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusAccepted,
		},
		{ //05// success
			db:     mock.NewDB(),
			nf:     mock.NewNotifier(),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusAccepted,
		},
	} {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data.in.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()
		env := &handler.Environment{DB: data.db, Msg: data.nf, Log: mock.NewLog()}

		api.NewUser(env)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		err := env.Log.(*mock.Log).Err
		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if data.nf != nil && err == "" {
			if data.nf.Msg != app.RegisterNotification {
				t.Errorf("test case %d: message mismatch \nhave: %v\nwant: %v", idx, data.nf.Msg, app.RegisterNotification)
			}
			if data.nf.To != data.in["email"][0] {
				t.Errorf("test case %d: recipient mismatch \nhave: %v\nwant: %v", idx, data.nf.To, data.in["email"][0])
			}
			if data.nf.Data["lang"] != "en" {
				t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.nf.Data["lang"], "en")
			}
		}
	}
}

func TestL10n(t *testing.T) {
	for idx, data := range []struct {
		tr     *mock.Translator
		params httprouter.Params

		out    map[string]interface{}
		status int
	}{
		{ //00// empty params, return fallback
			tr: &mock.Translator{
				Map:  map[string]interface{}{"hi": "Hi!", "bye": "Bye!"},
				Lang: "en",
			},
			out:    map[string]interface{}{},
			status: 200,
		},
		{ //01// success
			tr: &mock.Translator{
				Map:  map[string]interface{}{"hi": "Hi!", "bye": "Bye!"},
				Lang: "en",
			},
			params: httprouter.Params{{Key: "lang", Value: "en"}},
			out:    map[string]interface{}{"hi": "Hi!", "bye": "Bye!"},
			status: 200,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{L10n: data.tr}
		api.L10n(env)(res, req, data.params)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		out := map[string]interface{}{}
		body := res.Body.Bytes()
		json.Unmarshal(body, &out)
		if !reflect.DeepEqual(out, data.out) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, out, data.out)
		}
	}
}

func TestPreferences(t *testing.T) {
	for idx, data := range []struct {
		db        *mock.DB
		setCookie bool

		out    string
		status int
	}{
		{ //00// missing session id
			db:     mock.NewDB(),
			status: http.StatusNotFound,
		},
		{ //01// user doesn't exist
			db:        mock.NewDB(),
			setCookie: true,
			status:    http.StatusNotFound,
		},
		{ //02// connection failure
			db:        mock.NewDB().WithUser(mock.User1).WithError(mock.ErrDOS),
			setCookie: true,
			status:    http.StatusInternalServerError,
		},
		{ //03// success
			db:        mock.NewDB().WithUser(mock.User1),
			setCookie: true,
			status:    http.StatusOK,
			out:       mock.Prefs1Json,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{DB: data.db, Session: scs.New()}

		if data.setCookie {
			if ctx, err := env.Session.Load(req.Context(), "abc"); err == nil {
				req = req.WithContext(ctx)
				env.Session.Put(req.Context(), "id", 1)
			}
		}

		api.Preferences(env)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, body, data.out)
		}
	}
}
