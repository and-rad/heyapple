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

import (
	"heyapple/pkg/app"
	"heyapple/pkg/core"
)

var (
	Food0 = core.Food{
		ID: 42,
	}

	// apple
	Food1 = core.Food{
		ID:      1,
		KCal:    54,
		Fat:     0.1,
		Carbs:   14.4,
		Sugar:   10.3,
		Protein: 0.3,
		Fiber:   2,

		Iron:       0.2,
		Zinc:       0.1,
		Magnesium:  5,
		Chlorine:   2.2,
		Sodium:     1,
		Calcium:    5.3,
		Potassium:  119,
		Phosphorus: 11,
		Copper:     0.1,
		Iodine:     0.008,

		VitA:  0.01,
		VitB1: 0.04,
		VitB2: 0.03,
		VitB6: 0.1,
		VitC:  12,
		VitE:  0.49,
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

		Iron:       0.4,
		Zinc:       0.2,
		Magnesium:  30,
		Chlorine:   109,
		Sodium:     1,
		Calcium:    7,
		Potassium:  367,
		Phosphorus: 22,
		Copper:     0.1,
		Iodine:     0.003,
		Manganse:   0.3,

		VitA:  0.01,
		VitB1: 0.04,
		VitB2: 0.06,
		VitB6: 0.36,
		VitC:  11,
		VitE:  0.27,
	}

	Food0Json = `{"id":42,"brand":0,"kcal":0,"fat":0,"fatsat":0,"fato3":0,"fato6":0,"carb":0,"sug":0,"prot":0,"fib":0,"salt":0,"pot":0,"chl":0,"sod":0,"calc":0,"phos":0,"mag":0,"iron":0,"zinc":0,"mang":0,"cop":0,"iod":0,"chr":0,"mol":0,"sel":0,"vita":0,"vitb1":0,"vitb2":0,"vitb3":0,"vitb5":0,"vitb6":0,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":0,"vitd":0,"vite":0,"vitk":0}`
	Food1Json = `{"id":1,"brand":0,"kcal":54,"fat":0.1,"fatsat":0,"fato3":0,"fato6":0,"carb":14.4,"sug":10.3,"prot":0.3,"fib":2,"salt":0,"pot":119,"chl":2.2,"sod":1,"calc":5.3,"phos":11,"mag":5,"iron":0.2,"zinc":0.1,"mang":0,"cop":0.1,"iod":0.008,"chr":0,"mol":0,"sel":0,"vita":0.01,"vitb1":0.04,"vitb2":0.03,"vitb3":0,"vitb5":0,"vitb6":0.1,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":12,"vitd":0,"vite":0.49,"vitk":0}`
	Food2Json = `{"id":2,"brand":0,"kcal":93,"fat":0.2,"fatsat":0,"fato3":0,"fato6":0,"carb":20,"sug":17,"prot":1,"fib":2,"salt":0,"pot":367,"chl":109,"sod":1,"calc":7,"phos":22,"mag":30,"iron":0.4,"zinc":0.2,"mang":0.3,"cop":0.1,"iod":0.003,"chr":0,"mol":0,"sel":0,"vita":0.01,"vitb1":0.04,"vitb2":0.06,"vitb3":0,"vitb5":0,"vitb6":0.36,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":11,"vitd":0,"vite":0.27,"vitk":0}`
)

var (
	Recipe0 = core.Recipe{
		ID:    42,
		Size:  1,
		Items: []core.Ingredient{},
	}

	Recipe1 = core.Recipe{
		ID:    1,
		Size:  1,
		Items: []core.Ingredient{{ID: 2, Amount: 150}},
	}

	RecMeta1 = core.RecipeMeta{
		ID:   1,
		Name: "Apple Pie",
		KCal: 54 * 1.5,
	}

	Recipe2 = core.Recipe{
		ID:    2,
		Size:  3,
		Items: []core.Ingredient{{ID: 1, Amount: 300}, {ID: 2, Amount: 250}},
	}

	RecMeta2 = core.RecipeMeta{
		ID:   2,
		Name: "Fruit Cake",
		KCal: 54*3 + 93*2.5,
	}

	Recipe0Json  = `{"items":[],"size":1,"id":42}`
	Recipe1Json  = `{"items":[{"id":2,"amount":150}],"size":1,"id":1}`
	Recipe2Json  = `{"items":[{"id":1,"amount":300},{"id":2,"amount":250}],"size":3,"id":2}`
	RecMeta1Json = `{"name":"Apple Pie","instructions":"","id":1,"flags":0,"preptime":0,"cooktime":0,"misctime":0,"kcal":81,"fat":0,"carb":0,"prot":0}`
	RecMeta2Json = `{"name":"Fruit Cake","instructions":"","id":2,"flags":0,"preptime":0,"cooktime":0,"misctime":0,"kcal":394.5,"fat":0,"carb":0,"prot":0}`
)

var (
	User1 = app.User{
		ID:    1,
		Email: "a@a.a",
		Pass:  "$2a$10$CpVy94BcePvhBH3QS/mMnOtFVrfN0DvwdooEUc0T8tWdKNi3ayFXC",
		Lang:  "en",
		Perm:  app.PermLogin,
	}

	User1Json = `{"email":"a@a.a","pass":"$2a$10$CpVy94BcePvhBH3QS/mMnOtFVrfN0DvwdooEUc0T8tWdKNi3ayFXC","lang":"en","perm":1,"id":1}`
)
