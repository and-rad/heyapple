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

package mock

import "heyapple/pkg/core"

var (
	// apple
	Food1 = core.Food{
		ID:      1,
		KCal:    54,
		Fat:     0.1,
		Carbs:   14.4,
		Sugar:   10.3,
		Protein: 0.3,
		Fiber:   2,

		Iron:       0.0002,
		Zinc:       0.0001,
		Magnesium:  0.005,
		Chlorine:   0.0022,
		Calcium:    0.0053,
		Potassium:  0.119,
		Phosphorus: 0.011,
		Copper:     0.0001,
		Iodine:     0.000008,

		VitA:  0.00001,
		VitB1: 0.00004,
		VitB2: 0.00003,
		VitB6: 0.0001,
		VitC:  0.012,
		VitE:  0.00049,
	}

	// banana
	Food2 = core.Food{
		ID:      2,
		KCal:    93,
		Fat:     0.2,
		Carbs:   20,
		Sugar:   17,
		Protein: 1,
		Fiber:   2,

		Iron:       0.0004,
		Zinc:       0.0002,
		Magnesium:  0.03,
		Chlorine:   0.109,
		Calcium:    0.007,
		Potassium:  0.367,
		Phosphorus: 0.022,
		Copper:     0.0001,
		Iodine:     0.000003,
		Manganse:   0.0003,

		VitA:  0.00001,
		VitB1: 0.00004,
		VitB2: 0.00006,
		VitB6: 0.00036,
		VitC:  0.011,
		VitE:  0.00027,
	}

	Food1Json = `{"id":1,"brand":0,"kcal":54,"fat":0.1,"fatsat":0,"fato3":0,"fato6":0,"carb":14.4,"sug":10.3,"prot":0.3,"fib":2,"pot":0.119,"chl":0.0022,"sod":0,"calc":0.0053,"phos":0.011,"mag":0.005,"iron":0.0002,"zinc":0.0001,"mang":0,"cop":0.0001,"iod":0.000008,"chr":0,"mol":0,"sel":0,"vita":0.00001,"vitb1":0.00004,"vitb2":0.00003,"vitb3":0,"vitb5":0,"vitb6":0.0001,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":0.012,"vitd":0,"vite":0.00049,"vitk":0}`
	Food2Json = `{"id":2,"brand":0,"kcal":93,"fat":0.2,"fatsat":0,"fato3":0,"fato6":0,"carb":20,"sug":17,"prot":1,"fib":2,"pot":0.367,"chl":0.109,"sod":0,"calc":0.007,"phos":0.022,"mag":0.03,"iron":0.0004,"zinc":0.0002,"mang":0.0003,"cop":0.0001,"iod":0.000003,"chr":0,"mol":0,"sel":0,"vita":0.00001,"vitb1":0.00004,"vitb2":0.00006,"vitb3":0,"vitb5":0,"vitb6":0.00036,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":0.011,"vitd":0,"vite":0.00027,"vitk":0}`
)
