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
	"heyapple/pkg/core"
)

var (
	food1 = core.Food{ID: 1, KCal: 2, Fat: 3, Carbs: 4, Protein: 5}
	food2 = core.Food{ID: 2, KCal: 3, Fat: 4, Carbs: 5, Protein: 6}

	food1json = `{"id":1,"brand":0,"kcal":2,"fat":3,"fatsat":0,"fato3":0,"fato6":0,"carb":4,"sug":0,"prot":5,"fib":0,"pot":0,"chl":0,"sod":0,"calc":0,"phos":0,"mag":0,"iron":0,"zinc":0,"mang":0,"cop":0,"iod":0,"chr":0,"mol":0,"sel":0,"vita":0,"vitb1":0,"vitb2":0,"vitb3":0,"vitb5":0,"vitb6":0,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":0,"vitd":0,"vite":0,"vitk":0}`
	food2json = `{"id":2,"brand":0,"kcal":3,"fat":4,"fatsat":0,"fato3":0,"fato6":0,"carb":5,"sug":0,"prot":6,"fib":0,"pot":0,"chl":0,"sod":0,"calc":0,"phos":0,"mag":0,"iron":0,"zinc":0,"mang":0,"cop":0,"iod":0,"chr":0,"mol":0,"sel":0,"vita":0,"vitb1":0,"vitb2":0,"vitb3":0,"vitb5":0,"vitb6":0,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":0,"vitd":0,"vite":0,"vitk":0}`
)
