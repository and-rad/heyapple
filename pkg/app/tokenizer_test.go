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

package app_test

import (
	"heyapple/pkg/app"
	"testing"
)

func TestTokenizerCreate(t *testing.T) {
	tok := app.NewTokenizer()

	t1 := tok.Create()
	if l := len(t1); l != 32 {
		t.Errorf("unexpected token length: %s, %d", t1, l)
	}

	t2 := tok.Create()
	if t1 == t2 {
		t.Errorf("identical tokens: %s %s", t1, t2)
	}

}
