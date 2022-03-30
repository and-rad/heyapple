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

package auth_test

import (
	"heyapple/internal/app"
	"heyapple/internal/auth"
	"heyapple/internal/handler"
	"heyapple/internal/mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/and-rad/scs/v2"
	"github.com/and-rad/scs/v2/memstore"
)

func TestLocalLogin(t *testing.T) {
	for idx, data := range []struct {
		db *mock.DB
		in url.Values

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
		{ //02// empty DB
			db:     mock.NewDB(),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusUnauthorized,
		},
		{ //03// user doesn't exist
			db:     mock.NewDB().WithUser(app.User{ID: 1, Email: "b@b.b", Pass: "ahdjhehlk"}),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusUnauthorized,
		},
		{ //04// wrong password
			db:     mock.NewDB().WithUser(app.User{ID: 1, Email: "a@a.a", Pass: "ahdjhehlk"}),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusUnauthorized,
		},
		{ //05// success
			db:     mock.NewDB().WithUser(app.User{ID: 1, Perm: app.PermLogin, Email: "a@a.a", Pass: "$2a$10$ADm2JBRbt8UvB0uI7NNFBupOdTq7XKae6Dvc7NfVCnw89rPZr3.zK"}),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusOK,
		},
	} {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data.in.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()
		env := &handler.Environment{DB: data.db, Session: scs.New()}

		auth.LocalLogin(env)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}
	}
}

func TestLocalLogout(t *testing.T) {
	for idx, data := range []struct {
		setSession bool
		store      scs.Store
		status     int
	}{
		{ //00// no session
			setSession: false,
			status:     http.StatusNotFound,
		},
		{ //01// fail for internal reasons
			setSession: true,
			store:      mock.NewSessionStore().WithFailDestroy(),
			status:     http.StatusInternalServerError,
		},
		{ //02// success
			setSession: true,
			store:      memstore.New(),
			status:     http.StatusOK,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{Session: scs.New()}

		if data.setSession {
			env.Session.Store = data.store
			if ctx, err := env.Session.Load(req.Context(), ""); err == nil {
				req = req.WithContext(ctx)
			}
		}

		auth.LocalLogout(env)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}
	}
}

func TestResetRequest(t *testing.T) {
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
			in:     url.Values{"email": {"a@a.a"}},
			status: http.StatusInternalServerError,
		},
		{ //02// notification failure
			db:     mock.NewDB().WithUser(mock.User1),
			nf:     mock.NewNotifier().WithError(mock.ErrDOS),
			in:     url.Values{"email": {"a@a.a"}},
			status: http.StatusOK,
			err:    mock.ErrDOS.Error(),
		},
		{ //03// user doesn't exist
			db:     mock.NewDB(),
			in:     url.Values{"email": {"a@a.a"}},
			status: http.StatusNotFound,
		},
		{ //04// success
			db:     mock.NewDB().WithUser(mock.User1),
			nf:     mock.NewNotifier(),
			in:     url.Values{"email": {"a@a.a"}},
			status: http.StatusOK,
		},
	} {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data.in.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()
		env := &handler.Environment{DB: data.db, Msg: data.nf, Log: mock.NewLog()}

		auth.ResetRequest(env)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		err := env.Log.(*mock.Log).Err
		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if data.nf != nil && err == "" {
			if data.nf.Msg != app.ResetNotification {
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

func TestResetConfirm(t *testing.T) {
	for idx, data := range []struct {
		db *mock.DB
		in url.Values

		status int
	}{
		{ //00// no data
			db:     mock.NewDB(),
			status: http.StatusBadRequest,
		},
		{ //01// connection failure
			db:     mock.NewDB().WithError(mock.ErrDOS),
			in:     url.Values{"token": {"abcd"}, "pass": {"password123"}},
			status: http.StatusInternalServerError,
		},
		{ //02// empty DB
			db:     mock.NewDB(),
			in:     url.Values{"token": {"abcd"}, "pass": {"password123"}},
			status: http.StatusNotFound,
		},
		{ //03// success
			db:     mock.NewDB().WithUser(app.User{ID: 1}).WithToken(app.Token{ID: 1, Data: "reset"}),
			in:     url.Values{"token": {"abcd"}, "pass": {"password123"}},
			status: http.StatusOK,
		},
	} {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data.in.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()
		env := &handler.Environment{DB: data.db}

		auth.ResetConfirm(env)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}
	}
}