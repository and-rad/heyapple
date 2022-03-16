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
	"time"
)

var (
	Date1 = time.Date(2022, 3, 12, 8, 43, 0, 0, time.UTC)
	Date2 = time.Date(2022, 3, 12, 16, 43, 0, 0, time.UTC)
	Date3 = time.Date(2022, 3, 13, 8, 33, 0, 0, time.UTC)
	Date4 = time.Date(2022, 3, 13, 16, 33, 0, 0, time.UTC)

	Date1p2 = Date1.Add(time.Minute * 2)
	Date2p2 = Date2.Add(time.Minute * 2)

	Date1ISO = Date1.Format(time.RFC3339)
	Date2ISO = Date2.Format(time.RFC3339)

	Date1Date = Date1.Format("2006-01-02")
	Date2Date = Date2.Format("2006-01-02")

	Date1Time = Date1.Format("15:04")
	Date2Time = Date2.Format("15:04")

	Day1 = Date1.Truncate(time.Hour * 24)
	Day2 = Date3.Truncate(time.Hour * 24)
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

// Recipes need to be functions to make sure the underlying
// data of the Items slice is unique for every call.
func Recipe0() core.Recipe {
	return core.Recipe{
		ID:    42,
		Size:  1,
		Items: []core.Ingredient{},
	}
}

func Recipe1() core.Recipe {
	return core.Recipe{
		ID:      1,
		Name:    "Banana Pie",
		Size:    1,
		KCal:    Food2.KCal * 1.5,
		Fat:     Food2.Fat * 1.5,
		Carbs:   Food2.Carbs * 1.5,
		Protein: Food2.Protein * 1.5,
		Items:   []core.Ingredient{{ID: Food2.ID, Amount: 150}},
	}
}

func Recipe2() core.Recipe {
	return core.Recipe{
		ID:    2,
		Name:  "Fruit Cake",
		Size:  3,
		KCal:  54*3 + 93*2.5,
		Items: []core.Ingredient{{ID: 1, Amount: 300}, {ID: 2, Amount: 250}},
	}
}

const (
	Recipe0Json = `{"name":"","items":[],"id":42,"size":1,"flags":0,"preptime":0,"cooktime":0,"misctime":0,"kcal":0,"fat":0,"carb":0,"prot":0}`
	Recipe1Json = `{"name":"Banana Pie","items":[{"id":2,"amount":150}],"id":1,"size":1,"flags":0,"preptime":0,"cooktime":0,"misctime":0,"kcal":139.5,"fat":0.3,"carb":30,"prot":1.5}`
	Recipe2Json = `{"name":"Fruit Cake","items":[{"id":1,"amount":300},{"id":2,"amount":250}],"id":2,"size":3,"flags":0,"preptime":0,"cooktime":0,"misctime":0,"kcal":394.5,"fat":0,"carb":0,"prot":0}`
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

func Entry1() core.DiaryEntry {
	return core.DiaryEntry{
		Food: core.Ingredient{ID: 2, Amount: 150},
		Date: Date1,
	}
}

func Entry1New() core.DiaryEntry {
	return core.DiaryEntry{
		Food: core.Ingredient{ID: 2, Amount: 75},
		Date: Date1p2,
	}
}

func Entry1Rec1() core.DiaryEntry {
	return core.DiaryEntry{
		Food:   core.Ingredient{ID: 2, Amount: 150},
		Date:   Date1,
		Recipe: "Rec1",
	}
}

func Entry2() core.DiaryEntry {
	return core.DiaryEntry{
		Food: core.Ingredient{ID: 1, Amount: 50},
		Date: Date2,
	}
}

func Entry2New() core.DiaryEntry {
	return core.DiaryEntry{
		Food: core.Ingredient{ID: 1, Amount: 90},
		Date: Date2p2,
	}
}

func Entry2Rec1() core.DiaryEntry {
	return core.DiaryEntry{
		Food:   core.Ingredient{ID: 1, Amount: 50},
		Date:   Date2,
		Recipe: "Rec1",
	}
}

func Entry3() core.DiaryEntry {
	return core.DiaryEntry{
		Food: core.Ingredient{ID: 1, Amount: 50},
		Date: Date3,
	}
}

func Entry3New() core.DiaryEntry {
	return core.DiaryEntry{
		Food: core.Ingredient{ID: 1, Amount: 520},
		Date: Date3,
	}
}

func Entry4() core.DiaryEntry {
	return core.DiaryEntry{
		Food: core.Ingredient{ID: 2, Amount: 250},
		Date: Date4,
	}
}

func Entry4New() core.DiaryEntry {
	return core.DiaryEntry{
		Food: core.Ingredient{ID: 2, Amount: 150},
		Date: Date4,
	}
}

const (
	Entry1Json = `{"date":"2022-03-12T08:43:00Z","recipe":"","food":{"id":2,"amount":150}}`
	Entry2Json = `{"date":"2022-03-12T16:43:00Z","recipe":"","food":{"id":1,"amount":50}}`
)

func Diary210101() core.DiaryDay {
	return core.DiaryDay{
		Date: "2021-01-01",
	}
}

func Diary210102() core.DiaryDay {
	return core.DiaryDay{
		Date: "2021-01-02",
	}
}

func Diary210201() core.DiaryDay {
	return core.DiaryDay{
		Date: "2021-02-01",
	}
}

func Diary210202() core.DiaryDay {
	return core.DiaryDay{
		Date: "2021-02-02",
	}
}

func Diary220301() core.DiaryDay {
	return core.DiaryDay{
		Date: "2022-03-01",
	}
}

func Diary220302() core.DiaryDay {
	return core.DiaryDay{
		Date: "2022-03-02",
	}
}
