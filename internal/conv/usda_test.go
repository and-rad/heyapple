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

package conv_test

import (
	"reflect"
	"testing"

	"github.com/and-rad/heyapple/internal/conv"
	"github.com/and-rad/heyapple/internal/core"
	"github.com/and-rad/heyapple/internal/mock"
)

func TestFromUSDA(t *testing.T) {
	for idx, data := range []struct {
		in []byte

		out []conv.Food
		err error
	}{
		{ //00//
			in:  []byte(`{`),
			err: conv.ErrFromJSON,
		},
		{ //00//
			in:  []byte(`{}`),
			out: []conv.Food{},
		},
		{ //00//
			in: mock.USDA1(),
			out: []conv.Food{
				{
					Aisle: 0,
					Name:  "Milk, whole",
					Food: core.Food{
						ID:      1,
						Protein: 3.28,
					},
				},
				{
					Aisle: 0,
					Name:  "Whopper with cheese (Burger King)",
					Food: core.Food{
						ID:      2,
						Protein: 13.7,
					}},
			},
		},
	} {
		out, err := conv.FromUSDA(data.in)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(out, data.out) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, out, data.out)
		}
	}
}
