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

package api

import (
	"heyapple/internal/core"
	"heyapple/internal/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_sendResponse(t *testing.T) {
	for idx, data := range []struct {
		data   interface{}
		out    string
		status int
	}{
		{ //00// nil data
			data:   nil,
			out:    "null",
			status: http.StatusOK,
		},
		{ //01// JSON marshal failure
			data:   mock.MarshalFailer{},
			status: http.StatusInternalServerError,
		},
		{ //02// strings
			data:   "action successful",
			out:    `"action successful"`,
			status: http.StatusOK,
		},
		{ //03// food
			data:   core.Food{ID: 1, Protein: 23, Magnesium: 144},
			out:    `{"id":1,"brand":0,"kcal":0,"fat":0,"fatsat":0,"fato3":0,"fato6":0,"carb":0,"sug":0,"prot":23,"fib":0,"salt":0,"pot":0,"chl":0,"sod":0,"calc":0,"phos":0,"mag":144,"iron":0,"zinc":0,"mang":0,"cop":0,"iod":0,"chr":0,"mol":0,"sel":0,"vita":0,"vitb1":0,"vitb2":0,"vitb3":0,"vitb5":0,"vitb6":0,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":0,"vitd":0,"vite":0,"vitk":0}`,
			status: http.StatusOK,
		},
	} {
		res := httptest.NewRecorder()
		sendResponse(data.data, res)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %#v \nwant: %#v", idx, body, data.out)
		}
	}
}
