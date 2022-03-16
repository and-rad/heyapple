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
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"heyapple/pkg/handler"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/and-rad/scs/v2"
	"github.com/julienschmidt/httprouter"
)

func TestSaveDiaryEntry(t *testing.T) {
	for idx, data := range []struct {
		db        *mock.DB
		params    httprouter.Params
		in        url.Values
		setCookie bool
		method    string

		entries []core.DiaryEntry
		status  int
	}{
		{ //00// missing mandatory URL parts
			db:      mock.NewDB(),
			method:  http.MethodPost,
			entries: []core.DiaryEntry{},
			status:  http.StatusBadRequest,
		},
		{ //01// missing mandatory URL parts
			db:      mock.NewDB(),
			method:  http.MethodPost,
			params:  httprouter.Params{{Key: "date", Value: "12"}},
			entries: []core.DiaryEntry{},
			status:  http.StatusBadRequest,
		},
		{ //02// no session
			db:      mock.NewDB(),
			method:  http.MethodPost,
			params:  httprouter.Params{{Key: "date", Value: "2022-03-12"}},
			entries: []core.DiaryEntry{},
			status:  http.StatusUnauthorized,
		},
		{ //03// no data, nothing to do
			db:        mock.NewDB(),
			method:    http.MethodPost,
			params:    httprouter.Params{{Key: "date", Value: "2022-03-12"}},
			setCookie: true,
			entries:   []core.DiaryEntry{},
			status:    http.StatusNoContent,
		},
		{ //04// number of form params is off
			db:        mock.NewDB(),
			method:    http.MethodPost,
			params:    httprouter.Params{{Key: "date", Value: "2022-03-12"}},
			in:        url.Values{"id": {"1", "2"}, "amount": {"120"}},
			setCookie: true,
			entries:   []core.DiaryEntry{},
			status:    http.StatusBadRequest,
		},
		{ //05// invalid id param
			db:        mock.NewDB(),
			method:    http.MethodPost,
			params:    httprouter.Params{{Key: "date", Value: "2022-03-12"}},
			in:        url.Values{"id": {"hi"}, "amount": {"120"}, "time": {mock.Date1Time}},
			setCookie: true,
			entries:   []core.DiaryEntry{},
			status:    http.StatusNoContent,
		},
		{ //06// invalid amount param
			db:        mock.NewDB(),
			method:    http.MethodPost,
			params:    httprouter.Params{{Key: "date", Value: "2022-03-12"}},
			in:        url.Values{"id": {"1"}, "amount": {"lots"}, "time": {mock.Date1Time}},
			setCookie: true,
			entries:   []core.DiaryEntry{},
			status:    http.StatusNoContent,
		},
		{ //07// invalid time param
			db:        mock.NewDB(),
			method:    http.MethodPost,
			params:    httprouter.Params{{Key: "date", Value: "2022-03-12"}},
			in:        url.Values{"id": {"1"}, "amount": {"120"}, "time": {"now"}},
			setCookie: true,
			entries:   []core.DiaryEntry{},
			status:    http.StatusNoContent,
		},
		{ //08// connection failure
			db:        mock.NewDB().WithError(mock.ErrDOS),
			method:    http.MethodPost,
			params:    httprouter.Params{{Key: "date", Value: mock.Date1Date}},
			in:        url.Values{"id": {"1"}, "amount": {"120"}, "time": {mock.Date1Time}},
			setCookie: true,
			entries:   []core.DiaryEntry{},
			status:    http.StatusInternalServerError,
		},
		{ //09// diary id does not exist
			db:        mock.NewDB().WithError(app.ErrNotFound),
			method:    http.MethodPut,
			params:    httprouter.Params{{Key: "date", Value: mock.Date1Date}},
			in:        url.Values{"id": {"1"}, "amount": {"120"}, "time": {mock.Date1Time}},
			setCookie: true,
			entries:   []core.DiaryEntry{},
			status:    http.StatusNotFound,
		},
		{ //10// add entries
			db:     mock.NewDB(),
			method: http.MethodPost,
			params: httprouter.Params{{Key: "date", Value: mock.Date1Date}},
			in: url.Values{
				"id":     {"1", "2"},
				"amount": {"120", "210"},
				"time":   {mock.Date1Time, mock.Date1Time}},
			setCookie: true,
			entries: []core.DiaryEntry{
				{Date: mock.Date1, Food: core.Ingredient{ID: 1, Amount: 120}},
				{Date: mock.Date1, Food: core.Ingredient{ID: 2, Amount: 210}},
			},
			status: http.StatusNoContent,
		},
		{ //11// replace entries
			db: mock.NewDB().WithEntries(
				core.DiaryEntry{Date: mock.Date1, Food: core.Ingredient{ID: 1, Amount: 100}},
				core.DiaryEntry{Date: mock.Date1, Food: core.Ingredient{ID: 2, Amount: 50}},
				core.DiaryEntry{Date: mock.Date2, Food: core.Ingredient{ID: 1, Amount: 30}, Recipe: "Rec1"},
			),
			method: http.MethodPut,
			params: httprouter.Params{{Key: "date", Value: mock.Date1Date}},
			in: url.Values{
				"id":     {"1", "2"},
				"amount": {"120", "210"},
				"time":   {mock.Date1Time, mock.Date1Time}},
			setCookie: true,
			entries: []core.DiaryEntry{
				{Date: mock.Date1, Food: core.Ingredient{ID: 1, Amount: 120}},
				{Date: mock.Date1, Food: core.Ingredient{ID: 2, Amount: 210}},
				{Date: mock.Date2, Food: core.Ingredient{ID: 1, Amount: 30}, Recipe: "Rec1"},
			},
			status: http.StatusNoContent,
		},
	} {
		req := httptest.NewRequest(data.method, "/", strings.NewReader(data.in.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()
		env := &handler.Environment{DB: data.db, Session: scs.New()}

		if data.setCookie {
			if ctx, err := env.Session.Load(req.Context(), "abc"); err == nil {
				req = req.WithContext(ctx)
				env.Session.Put(req.Context(), "id", 1)
			}
		}

		api.SaveDiaryEntry(env)(res, req, data.params)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if !reflect.DeepEqual(data.db.Entries, data.entries) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.Entries, data.entries)
		}
	}
}

func TestDiary(t *testing.T) {
	for idx, data := range []struct {
		db        *mock.DB
		params    httprouter.Params
		setCookie bool

		out    string
		status int
	}{
		{ //00// no session
			db:     mock.NewDB(),
			status: http.StatusUnauthorized,
		},
		{ //01// connection failure
			db:        mock.NewDB().WithError(mock.ErrDOS),
			setCookie: true,
			status:    http.StatusInternalServerError,
		},
		{ //02// success
			db:        mock.NewDB().WithDays(mock.Diary210101(), mock.Diary210102(), mock.Diary220301()),
			setCookie: true,
			out:       fmt.Sprintf(`[%s,%s,%s]`, mock.Diary210101Json, mock.Diary210102Json, mock.Diary220301Json),
			status:    http.StatusOK,
		},
		{ //03// invalid year format
			db:        mock.NewDB().WithDays(mock.Diary210101(), mock.Diary210102(), mock.Diary220301()),
			params:    httprouter.Params{{Key: "year", Value: "thisone"}},
			setCookie: true,
			status:    http.StatusBadRequest,
		},
		{ //04// success for year
			db:        mock.NewDB().WithDays(mock.Diary210101(), mock.Diary210102(), mock.Diary220301()),
			params:    httprouter.Params{{Key: "year", Value: "2022"}},
			setCookie: true,
			out:       fmt.Sprintf(`[%s]`, mock.Diary220301Json),
			status:    http.StatusOK,
		},
		{ //05// invalid month format
			db:        mock.NewDB().WithDays(mock.Diary210101(), mock.Diary210102(), mock.Diary220301()),
			params:    httprouter.Params{{Key: "month", Value: "thisone"}},
			setCookie: true,
			status:    http.StatusBadRequest,
		},
		{ //06// success for month
			db:        mock.NewDB().WithDays(mock.Diary210101(), mock.Diary210201(), mock.Diary220301()),
			params:    httprouter.Params{{Key: "year", Value: "2021"}, {Key: "month", Value: "1"}},
			setCookie: true,
			out:       fmt.Sprintf(`[%s]`, mock.Diary210101Json),
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

		api.Diary(env)(res, req, data.params)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, body, data.out)
		}
	}
}

func TestDiaryEntries(t *testing.T) {
	for idx, data := range []struct {
		db        *mock.DB
		params    httprouter.Params
		setCookie bool

		out    string
		status int
	}{
		{ //00// invalid year format
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "year", Value: "thisone"}},
			status: http.StatusBadRequest,
		},
		{ //01// invalid month format
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "year", Value: "2021"}, {Key: "month", Value: "thisone"}},
			status: http.StatusBadRequest,
		},
		{ //02// invalid day format
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "year", Value: "2021"}, {Key: "month", Value: "1"}, {Key: "day", Value: "today"}},
			status: http.StatusBadRequest,
		},
		{ //03// no session
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "year", Value: "2021"}, {Key: "month", Value: "1"}, {Key: "day", Value: "2"}},
			status: http.StatusUnauthorized,
		},
		{ //04// connection failure
			db:        mock.NewDB().WithError(mock.ErrDOS),
			params:    httprouter.Params{{Key: "year", Value: "2021"}, {Key: "month", Value: "1"}, {Key: "day", Value: "2"}},
			setCookie: true,
			status:    http.StatusInternalServerError,
		},
		{ //05// diary doesn't exist
			db:        mock.NewDB().WithError(app.ErrNotFound),
			params:    httprouter.Params{{Key: "year", Value: "2021"}, {Key: "month", Value: "1"}, {Key: "day", Value: "2"}},
			setCookie: true,
			out:       "[]",
			status:    http.StatusOK,
		},
		{ //06// success
			db:        mock.NewDB().WithEntries(mock.Entry1(), mock.Entry2(), mock.Entry3()),
			params:    httprouter.Params{{Key: "year", Value: "2022"}, {Key: "month", Value: "3"}, {Key: "day", Value: "12"}},
			setCookie: true,
			out:       fmt.Sprintf(`[%s,%s]`, mock.Entry1Json, mock.Entry2Json),
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

		api.DiaryEntries(env)(res, req, data.params)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, body, data.out)
		}
	}
}
