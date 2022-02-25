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

package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/and-rad/scs/v2"
)

func Test_sessionLang(t *testing.T) {
	for idx, data := range []struct {
		setCookie bool
		setHeader bool
		out       string
	}{
		{ //00// empty data
			out: "",
		},
		{ //01// from session
			setCookie: true,
			out:       "en",
		},
		{ //02// from header
			setHeader: true,
			out:       "en,fr",
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		sm := scs.New()

		if data.setCookie {
			if ctx, err := sm.Load(req.Context(), "abc"); err == nil {
				req = req.WithContext(ctx)
				sm.Put(req.Context(), "lang", "en")
			}
		}

		if data.setHeader {
			req.Header.Set("Accept-Language", "en,fr")
		}

		if out := sessionLang(sm, req); out != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, out, data.out)
		}
	}
}
