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
	"heyapple/internal/mock"
	"heyapple/pkg/api/v1"
	"heyapple/pkg/app"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
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
		{ //01// notification failure
			db:     mock.NewDB(),
			nf:     mock.NewNotifier().WithError(mock.ErrDOS),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusCreated,
			err:    mock.ErrDOS.Error(),
		},
		{ //01// success
			db:     mock.NewDB(),
			nf:     mock.NewNotifier(),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusCreated,
		},
	} {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data.in.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()
		log := mock.NewLog()

		api.NewUser(log, data.nf, data.db)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if log.Err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, log.Err, data.err)
		}

		if data.nf != nil && log.Err == "" {
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
