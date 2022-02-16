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
	"heyapple/internal/mock"
	"heyapple/pkg/app"
	"heyapple/pkg/auth"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/and-rad/scs/v2"
)

func TestLocalLogin(t *testing.T) {
	for idx, data := range []struct {
		db *mock.DB
		in url.Values

		status int
	}{
		{ //00// no data
			db:     mock.NewDB(),
			status: http.StatusUnauthorized,
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
			db:     mock.NewDB().WithUser(app.User{ID: 1, Email: "a@a.a", Pass: "$2a$10$ADm2JBRbt8UvB0uI7NNFBupOdTq7XKae6Dvc7NfVCnw89rPZr3.zK"}),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusOK,
		},
	} {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data.in.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()

		auth.LocalLogin(scs.New(), data.db)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}
	}
}

func TestLocalLogout(t *testing.T) {
	for idx, data := range []struct {
		status int
	}{
		{ //00//
			status: http.StatusNotFound,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()

		auth.LocalLogout(scs.New())(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}
	}
}
