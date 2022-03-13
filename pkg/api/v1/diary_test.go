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
			in:        url.Values{"id": {"hi"}, "amount": {"120"}, "date": {mock.Date1ISO}},
			setCookie: true,
			entries:   []core.DiaryEntry{},
			status:    http.StatusNoContent,
		},
		{ //06// invalid amount param
			db:        mock.NewDB(),
			method:    http.MethodPost,
			params:    httprouter.Params{{Key: "date", Value: "2022-03-12"}},
			in:        url.Values{"id": {"1"}, "amount": {"lots"}, "date": {mock.Date1ISO}},
			setCookie: true,
			entries:   []core.DiaryEntry{},
			status:    http.StatusNoContent,
		},
		{ //07// invalid date param
			db:        mock.NewDB(),
			method:    http.MethodPost,
			params:    httprouter.Params{{Key: "date", Value: "2022-03-12"}},
			in:        url.Values{"id": {"1"}, "amount": {"120"}, "date": {"today"}},
			setCookie: true,
			entries:   []core.DiaryEntry{},
			status:    http.StatusNoContent,
		},
		{ //08// connection failure
			db:        mock.NewDB().WithError(mock.ErrDOS),
			method:    http.MethodPost,
			params:    httprouter.Params{{Key: "date", Value: mock.Date1yyyymmdd}},
			in:        url.Values{"id": {"1"}, "amount": {"120"}, "date": {mock.Date1ISO}},
			setCookie: true,
			entries:   []core.DiaryEntry{},
			status:    http.StatusInternalServerError,
		},
		{ //09// diary id does not exist
			db:        mock.NewDB().WithError(app.ErrNotFound),
			method:    http.MethodPut,
			params:    httprouter.Params{{Key: "date", Value: mock.Date1yyyymmdd}},
			in:        url.Values{"id": {"1"}, "amount": {"120"}, "date": {mock.Date1ISO}},
			setCookie: true,
			entries:   []core.DiaryEntry{},
			status:    http.StatusNotFound,
		},
		{ //10// add entries
			db:     mock.NewDB(),
			method: http.MethodPost,
			params: httprouter.Params{{Key: "date", Value: mock.Date1yyyymmdd}},
			in: url.Values{
				"id":     {"1", "2"},
				"amount": {"120", "210"},
				"date":   {mock.Date1ISO, mock.Date1ISO}},
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
			params: httprouter.Params{{Key: "date", Value: mock.Date1yyyymmdd}},
			in: url.Values{
				"id":     {"1", "2"},
				"amount": {"120", "210"},
				"date":   {mock.Date1ISO, mock.Date1ISO}},
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
