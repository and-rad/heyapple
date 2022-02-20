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
	"net/url"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
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
			db:     mock.NewDB().WithFoods(mock.Food1, mock.Food2),
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

func TestFood(t *testing.T) {
	for idx, data := range []struct {
		db     *mock.DB
		params httprouter.Params

		out    string
		status int
	}{
		{ //00// missing id
			db:     mock.NewDB(),
			status: http.StatusBadRequest,
		},
		{ //01// missing id
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "name", Value: "12"}},
			status: http.StatusBadRequest,
		},
		{ //02// wrong data type for id
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "id", Value: "someone"}},
			status: http.StatusBadRequest,
		},
		{ //03// connection failure
			db:     mock.NewDB().WithError(mock.ErrDOS),
			params: httprouter.Params{{Key: "id", Value: "42"}},
			status: http.StatusInternalServerError,
		},
		{ //04// item doesn't exist
			db:     mock.NewDB().WithFood(mock.Food1),
			params: httprouter.Params{{Key: "id", Value: "42"}},
			status: http.StatusNotFound,
		},
		{ //05// success
			db:     mock.NewDB().WithFood(mock.Food1),
			params: httprouter.Params{{Key: "id", Value: "1"}},
			status: http.StatusOK,
			out:    mock.Food1Json,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		api.Food(data.db)(res, req, data.params)

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

func TestSaveFood(t *testing.T) {
	for idx, data := range []struct {
		db     *mock.DB
		params httprouter.Params
		in     url.Values

		food   core.Food
		status int
	}{
		{ //00// missing mandatory form values
			db:     mock.NewDB(),
			status: http.StatusBadRequest,
		},
		{ //01// missing mandatory form values
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "name", Value: "12"}},
			status: http.StatusBadRequest,
		},
		{ //02// wrong data type for id
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "id", Value: "someone"}},
			status: http.StatusBadRequest,
		},
		{ //03// connection failure
			db:     mock.NewDB().WithError(mock.ErrDOS),
			params: httprouter.Params{{Key: "id", Value: "42"}},
			status: http.StatusInternalServerError,
		},
		{ //04// item doesn't exist
			db:     mock.NewDB().WithFood(mock.Food1),
			params: httprouter.Params{{Key: "id", Value: "42"}},
			in:     url.Values{"kcal": {"360"}},
			status: http.StatusNotFound,
			food:   mock.Food1,
		},
		{ //05// success
			db:     mock.NewDB().WithFood(mock.Food1),
			params: httprouter.Params{{Key: "id", Value: "1"}},
			in:     url.Values{"kcal": {"360"}},
			status: http.StatusNoContent,
			food:   func() core.Food { f := mock.Food1; f.KCal = 360; return f }(),
		},
		{ //06// ignore invalid values
			db:     mock.NewDB().WithFood(mock.Food1),
			params: httprouter.Params{{Key: "id", Value: "1"}},
			in:     url.Values{"prot": {"34.5"}, "fat": {"alot"}, "power": {"9000"}},
			status: http.StatusNoContent,
			food:   func() core.Food { f := mock.Food1; f.Protein = 34.5; return f }(),
		},
	} {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(data.in.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()
		api.SaveFood(data.db)(res, req, data.params)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if data.db.FoodItem != data.food {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, data.db.FoodItem, data.food)
		}
	}
}
