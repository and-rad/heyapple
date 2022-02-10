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
	"fmt"
	"heyapple/internal/mock"
	"heyapple/pkg/api/v1"
	"heyapple/pkg/core"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFoods(t *testing.T) {
	for idx, data := range []struct {
		db     *mock.DB
		out    string
		status int
	}{
		{ //00// connection failure
			db:     mock.NewDB().WithError(mock.ErrDOS),
			status: http.StatusInternalServerError,
		},
		{ //01// empty set
			db:     mock.NewDB(),
			status: http.StatusOK,
			out:    "[]",
		},
		{ //02// success
			db:     mock.NewDB().WithFoods([]core.Food{mock.Food1, mock.Food2}),
			out:    fmt.Sprintf(`[%s,%s]`, mock.Food1Json, mock.Food2Json),
			status: http.StatusOK,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		api.Foods(data.db)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, body, data.out)
		}
	}
}

func TestNewFood(t *testing.T) {
	for idx, data := range []struct {
		db     *mock.DB
		out    string
		status int
	}{
		{ //00// connection failure
			db:     mock.NewDB().WithError(mock.ErrDOS),
			status: http.StatusInternalServerError,
		},
		{ //01// success
			db:     mock.NewDB().WithID(23),
			status: http.StatusCreated,
			out:    "23",
		},
	} {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		api.NewFood(data.db)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, body, data.out)
		}
	}
}
