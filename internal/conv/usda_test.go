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
					Aisle: core.AisleProduce,
					Name:  "Apples, raw, with skin (Includes foods for USDA's Food Distribution Program)",
					Food: core.Food{
						ID:  1,
						Cat: core.CatFruit,

						KCal:    52,
						Fat:     0.17,
						FatSat:  0.028,
						FatO3:   0,
						FatO6:   0,
						Carbs:   13.8,
						Protein: 0.26,
						Fiber:   2.4,

						Calcium:    6,
						Chlorine:   0,
						Chromium:   0,
						Copper:     0.027,
						Iodine:     0,
						Iron:       0.12,
						Magnesium:  5,
						Manganse:   0.035,
						Molybdenum: 0,
						Phosphorus: 11,
						Potassium:  107,
						Selenium:   0,
						Sodium:     1,
						Zinc:       0.04,

						VitA:   0.003,
						VitB1:  0.017,
						VitB2:  0.026,
						VitB3:  0.091,
						VitB5:  0.061,
						VitB6:  0.041,
						VitB7:  0,
						VitB9:  0.003,
						VitB12: 0,
						VitC:   4.6,
						VitD:   0,
						VitE:   0.18,
						VitK:   0.0023,
					},
				},
				{
					Aisle: core.AisleSnacks,
					Name:  "Snacks, potato chips, made from dried potatoes, reduced fat",
					Food: core.Food{
						ID:  2,
						Cat: core.CatSnacks,

						KCal:    502,
						Fat:     26.1,
						FatSat:  6.76,
						FatO3:   0.066,
						FatO6:   12.003,
						Carbs:   64.8,
						Protein: 4.56,
						Fiber:   3.2,

						Calcium:    29,
						Chlorine:   0,
						Chromium:   0,
						Copper:     0.153,
						Iodine:     0,
						Iron:       1.11,
						Magnesium:  45,
						Manganse:   0.361,
						Molybdenum: 0,
						Phosphorus: 129,
						Potassium:  760,
						Selenium:   0.003,
						Sodium:     450,
						Zinc:       0.7,

						VitA:   0,
						VitB1:  0.212,
						VitB2:  0.015,
						VitB3:  3.6,
						VitB5:  0.83,
						VitB6:  0.412,
						VitB7:  0,
						VitB9:  0.027,
						VitB12: 0,
						VitC:   12,
						VitD:   0,
						VitE:   2.18,
						VitK:   0.008,
					},
				},
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
