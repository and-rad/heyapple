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
						ID: 1,

						KCal:    60,
						Fat:     3.2,
						FatSat:  1.86,
						Carbs:   4.67,
						Protein: 3.28,
						Fiber:   0,

						Calcium:    123,
						Copper:     0.001,
						Iron:       0,
						Magnesium:  12,
						Phosphorus: 101,
						Potassium:  150,
						Selenium:   0.0019,
						Sodium:     38,
						Zinc:       0.41,

						VitA:   0.032,
						VitB1:  0.056,
						VitB2:  0.138,
						VitB3:  0.105,
						VitB6:  0.061,
						VitB9:  0,
						VitB12: 0.00056,
						VitC:   0,
						VitD:   0.00115,
						VitE:   0.05,
						VitK:   0.0003,
					},
				},
				{
					Aisle: 0,
					Name:  "Whopper with cheese (Burger King)",
					Food: core.Food{
						ID: 2,

						KCal:    268,
						Fat:     15.8,
						FatSat:  5.75,
						Carbs:   17.7,
						Protein: 13.7,
						Fiber:   1.1,

						Calcium:    73,
						Copper:     0.081,
						Iron:       2.3,
						Magnesium:  18,
						Phosphorus: 125,
						Potassium:  208,
						Selenium:   0.0165,
						Sodium:     473,
						Zinc:       2.26,

						VitA:   0.050,
						VitB1:  0.177,
						VitB2:  0.257,
						VitB3:  3.14,
						VitB6:  0.186,
						VitB9:  0.052,
						VitB12: 0.00095,
						VitC:   1.1,
						VitD:   0.00015,
						VitE:   0.57,
						VitK:   0.0142,
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
