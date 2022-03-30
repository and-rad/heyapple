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

package api_test

import (
	"encoding/json"
	"heyapple/internal/api/v1"
	"heyapple/internal/app"
	"heyapple/internal/handler"
	"heyapple/internal/mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

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
		{ //01// connection failure
			db:     mock.NewDB().WithError(mock.ErrDOS),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusInternalServerError,
		},
		{ //02// notification failure
			db:     mock.NewDB(),
			nf:     mock.NewNotifier().WithError(mock.ErrDOS),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusCreated,
			err:    mock.ErrDOS.Error(),
		},
		{ //03// username exists
			db:     mock.NewDB().WithError(app.ErrExists),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusConflict,
		},
		{ //04// success
			db:     mock.NewDB(),
			nf:     mock.NewNotifier(),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusCreated,
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
