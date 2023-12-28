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
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/and-rad/heyapple/internal/api/v1"
	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/handler"
	"github.com/and-rad/heyapple/internal/mock"

	"github.com/and-rad/scs/v2"
	"github.com/julienschmidt/httprouter"
)

func TestSaveListDone(t *testing.T) {
	for idx, data := range []struct {
		db        *mock.DB
		params    httprouter.Params
		values    url.Values
		setCookie bool

		status int
	}{
		{ //00// no session
			db:     mock.NewDB(),
			status: http.StatusUnauthorized,
		},
		{ //01// missing input data
			db:        mock.NewDB(),
			params:    httprouter.Params{{Key: "name", Value: "diary"}},
			values:    url.Values{},
			setCookie: true,
			status:    http.StatusBadRequest,
		},
		{ //02// number of input values is off
			db:        mock.NewDB(),
			params:    httprouter.Params{},
			values:    url.Values{"id": {"1", "2"}, "done": {"true"}},
			setCookie: true,
			status:    http.StatusBadRequest,
		},
		{ //03// invalid data type
			db:        mock.NewDB(),
			params:    httprouter.Params{},
			values:    url.Values{"id": {"one"}, "done": {"true"}},
			setCookie: true,
			status:    http.StatusBadRequest,
		},
		{ //04// invalid data type
			db:        mock.NewDB(),
			params:    httprouter.Params{},
			values:    url.Values{"id": {"1"}, "done": {"hellyeah"}},
			setCookie: true,
			status:    http.StatusBadRequest,
		},
		{ //05// connection failure
			db:        mock.NewDB().WithError(nil, mock.ErrDOS).WithFoods(mock.Food1),
			params:    httprouter.Params{{Key: "name", Value: "diary"}},
			values:    url.Values{"id": {"1"}, "done": {"true"}},
			setCookie: true,
			status:    http.StatusInternalServerError,
		},
		{ //06// success
			db:        mock.NewDB().WithFoods(mock.Food1),
			params:    httprouter.Params{{Key: "name", Value: "diary"}},
			values:    url.Values{"id": {"1"}, "done": {"true"}},
			setCookie: true,
			status:    http.StatusNoContent,
		},
	} {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(data.values.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()
		env := &handler.Environment{DB: data.db, Session: scs.New()}

		if data.setCookie {
			if ctx, err := env.Session.Load(req.Context(), "abc"); err == nil {
				req = req.WithContext(ctx)
				env.Session.Put(req.Context(), "id", 1)
			}
		}

		api.SaveListDone(env)(res, req, data.params)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}
	}
}

func TestShoppingList(t *testing.T) {
	for idx, data := range []struct {
		db        *mock.DB
		params    httprouter.Params
		in        url.Values
		setCookie bool

		out    string
		status int
	}{
		{ //00// no session
			db:     mock.NewDB(),
			status: http.StatusUnauthorized,
		},
		{ //01// missing date
			db:        mock.NewDB(),
			params:    httprouter.Params{{Key: "name", Value: "diary"}},
			setCookie: true,
			status:    http.StatusBadRequest,
		},
		{ //02// connection failure
			db:        mock.NewDB().WithError(mock.ErrDOS),
			params:    httprouter.Params{{Key: "name", Value: "diary"}},
			in:        url.Values{"date": {"2022-03-01"}},
			setCookie: true,
			status:    http.StatusInternalServerError,
		},
		{ //03// shopping list doesn't exist
			db:        mock.NewDB().WithError(app.ErrNotFound),
			params:    httprouter.Params{{Key: "name", Value: "Dinner"}},
			in:        url.Values{"date": {"2022-03-01"}},
			setCookie: true,
			status:    http.StatusNotFound,
		},
		{ //04// success
			db:        mock.NewDB().WithDate(mock.Day1).WithShoppingList(mock.List1()...),
			params:    httprouter.Params{{Key: "name", Value: "diary"}},
			in:        url.Values{"date": {mock.Date1Date}},
			setCookie: true,
			out:       mock.List1Json,
			status:    http.StatusOK,
		},
	} {
		queryURL := "/"
		if len(data.in) != 0 {
			queryURL = "/?" + data.in.Encode()
		}

		req := httptest.NewRequest(http.MethodGet, queryURL, strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{DB: data.db, Session: scs.New()}

		if data.setCookie {
			if ctx, err := env.Session.Load(req.Context(), "abc"); err == nil {
				req = req.WithContext(ctx)
				env.Session.Put(req.Context(), "id", 1)
			}
		}

		api.ShoppingList(env)(res, req, data.params)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, body, data.out)
		}
	}
}

func TestShoppingLists(t *testing.T) {
	for idx, data := range []struct {
		db        *mock.DB
		setCookie bool

		out    string
		status int
	}{
		{ //00// no session
			db:     mock.NewDB(),
			status: http.StatusUnauthorized,
		},
		{ //01// empty database
			db:        mock.NewDB(),
			setCookie: true,
			out:       `[]`,
			status:    http.StatusOK,
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

		api.ShoppingLists(env)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, body, data.out)
		}
	}
}
