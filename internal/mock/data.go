////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2021-2024 The HeyApple Authors.
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
	"bytes"
	"embed"
	"time"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/core"
)

var (
	//go:embed *.json
	convFS embed.FS
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
	Date3Date = Date3.Format("2006-01-02")

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
		ID:       1,
		Category: 9,
		Flags:    core.FlagVegan | core.FlagVegetarian,
		KCal:     54,
		Fat:      0.1,
		Carbs:    14.4,
		Sugar:    10.3,
		Fructose: 5,
		Glucose:  3,
		Sucrose:  2.3,
		Protein:  0.3,
		Fiber:    2,

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
		ID:       2,
		Category: 9,
		Flags:    core.FlagVegan | core.FlagVegetarian,
		KCal:     93,
		Fat:      0.2,
		Carbs:    20,
		Sugar:    17,
		Fructose: 10,
		Glucose:  4,
		Sucrose:  3,
		Protein:  1,
		Fiber:    2,

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

	Food0Json = `{"id":42,"brand":0,"cat":0,"flags":0,"kcal":0,"fat":0,"fatsat":0,"fato3":0,"fato6":0,"carb":0,"sug":0,"fruc":0,"gluc":0,"suc":0,"fib":0,"prot":0,"salt":0,"pot":0,"chl":0,"sod":0,"calc":0,"phos":0,"mag":0,"iron":0,"zinc":0,"mang":0,"cop":0,"iod":0,"chr":0,"mol":0,"sel":0,"vita":0,"vitb1":0,"vitb2":0,"vitb3":0,"vitb5":0,"vitb6":0,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":0,"vitd":0,"vite":0,"vitk":0}`
	Food1Json = `{"id":1,"brand":0,"cat":9,"flags":3,"kcal":54,"fat":0.1,"fatsat":0,"fato3":0,"fato6":0,"carb":14.4,"sug":10.3,"fruc":5,"gluc":3,"suc":2.3,"fib":2,"prot":0.3,"salt":0,"pot":119,"chl":2.2,"sod":1,"calc":5.3,"phos":11,"mag":5,"iron":0.2,"zinc":0.1,"mang":0,"cop":0.1,"iod":0.008,"chr":0,"mol":0,"sel":0,"vita":0.01,"vitb1":0.04,"vitb2":0.03,"vitb3":0,"vitb5":0,"vitb6":0.1,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":12,"vitd":0,"vite":0.49,"vitk":0}`
	Food2Json = `{"id":2,"brand":0,"cat":9,"flags":3,"kcal":93,"fat":0.2,"fatsat":0,"fato3":0,"fato6":0,"carb":20,"sug":17,"fruc":10,"gluc":4,"suc":3,"fib":2,"prot":1,"salt":0,"pot":367,"chl":109,"sod":1,"calc":7,"phos":22,"mag":30,"iron":0.4,"zinc":0.2,"mang":0.3,"cop":0.1,"iod":0.003,"chr":0,"mol":0,"sel":0,"vita":0.01,"vitb1":0.04,"vitb2":0.06,"vitb3":0,"vitb5":0,"vitb6":0.36,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":11,"vitd":0,"vite":0.27,"vitk":0}`
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
		ID:         1,
		Name:       "Banana Pie",
		Size:       1,
		Flags:      Food2.Flags,
		KCal:       Food2.KCal * 1.5,
		Fat:        Food2.Fat * 1.5,
		Carbs:      Food2.Carbs * 1.5,
		Sugar:      Food2.Sugar * 1.5,
		Fructose:   Food2.Fructose * 1.5,
		Glucose:    Food2.Glucose * 1.5,
		Sucrose:    Food2.Sucrose * 1.5,
		Protein:    Food2.Protein * 1.5,
		Fiber:      Food2.Fiber * 1.5,
		Iron:       Food2.Iron * 1.5,
		Zinc:       Food2.Zinc * 1.5,
		Magnesium:  Food2.Magnesium * 1.5,
		Chlorine:   Food2.Chlorine * 1.5,
		Sodium:     Food2.Sodium * 1.5,
		Calcium:    Food2.Calcium * 1.5,
		Potassium:  Food2.Potassium * 1.5,
		Phosphorus: Food2.Phosphorus * 1.5,
		Copper:     Food2.Copper * 1.5,
		Iodine:     Food2.Iodine * 1.5,
		Manganse:   Food2.Manganse * 1.5,
		VitA:       Food2.VitA * 1.5,
		VitB1:      Food2.VitB1 * 1.5,
		VitB2:      Food2.VitB2 * 1.5,
		VitB6:      Food2.VitB6 * 1.5,
		VitC:       Food2.VitC * 1.5,
		VitE:       Food2.VitE * 1.5,
		Items:      []core.Ingredient{{ID: Food2.ID, Amount: 150}},
	}
}

func Recipe2() core.Recipe {
	return core.Recipe{
		ID:         2,
		Name:       "Fruit Cake",
		Size:       3,
		Flags:      Food1.Flags | Food2.Flags,
		KCal:       Food1.KCal*3 + Food2.KCal*2.5,
		Fat:        Food1.Fat*3 + Food2.Fat*2.5,
		Carbs:      Food1.Carbs*3 + Food2.Carbs*2.5,
		Sugar:      Food1.Sugar*3 + Food2.Sugar*2.5,
		Fructose:   Food1.Fructose*3 + Food2.Fructose*2.5,
		Glucose:    Food1.Glucose*3 + Food2.Glucose*2.5,
		Sucrose:    Food1.Sucrose*3 + Food2.Sucrose*2.5,
		Protein:    Food1.Protein*3 + Food2.Protein*2.5,
		Fiber:      Food1.Fiber*3 + Food2.Fiber*2.5,
		Iron:       Food1.Iron*3 + Food2.Iron*2.5,
		Zinc:       Food1.Zinc*3 + Food2.Zinc*2.5,
		Magnesium:  Food1.Magnesium*3 + Food2.Magnesium*2.5,
		Chlorine:   Food1.Chlorine*3 + Food2.Chlorine*2.5,
		Sodium:     Food1.Sodium*3 + Food2.Sodium*2.5,
		Calcium:    Food1.Calcium*3 + Food2.Calcium*2.5,
		Potassium:  Food1.Potassium*3 + Food2.Potassium*2.5,
		Phosphorus: Food1.Phosphorus*3 + Food2.Phosphorus*2.5,
		Copper:     Food1.Copper*3 + Food2.Copper*2.5,
		Iodine:     Food1.Iodine*3 + Food2.Iodine*2.5,
		Manganse:   Food1.Manganse*3 + Food2.Manganse*2.5,
		VitA:       Food1.VitA*3 + Food2.VitA*2.5,
		VitB1:      Food1.VitB1*3 + Food2.VitB1*2.5,
		VitB2:      Food1.VitB2*3 + Food2.VitB2*2.5,
		VitB6:      Food1.VitB6*3 + Food2.VitB6*2.5,
		VitC:       Food1.VitC*3 + Food2.VitC*2.5,
		VitE:       Food1.VitE*3 + Food2.VitE*2.5,
		Items:      []core.Ingredient{{ID: 1, Amount: 300}, {ID: 2, Amount: 250}},
	}
}

const (
	Recipe0Json = `{"name":"","items":[],"id":42,"size":1,"flags":0,"preptime":0,"cooktime":0,"misctime":0,"kcal":0,"fat":0,"fatsat":0,"fato3":0,"fato6":0,"carb":0,"sug":0,"fruc":0,"gluc":0,"suc":0,"fib":0,"prot":0,"salt":0,"pot":0,"chl":0,"sod":0,"calc":0,"phos":0,"mag":0,"iron":0,"zinc":0,"mang":0,"cop":0,"iod":0,"chr":0,"mol":0,"sel":0,"vita":0,"vitb1":0,"vitb2":0,"vitb3":0,"vitb5":0,"vitb6":0,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":0,"vitd":0,"vite":0,"vitk":0}`
	Recipe1Json = `{"name":"Banana Pie","items":[{"id":2,"amount":150}],"id":1,"size":1,"flags":3,"preptime":0,"cooktime":0,"misctime":0,"kcal":139.5,"fat":0.3,"fatsat":0,"fato3":0,"fato6":0,"carb":30,"sug":25.5,"fruc":15,"gluc":6,"suc":4.5,"fib":3,"prot":1.5,"salt":0,"pot":550.5,"chl":163.5,"sod":1.5,"calc":10.5,"phos":33,"mag":45,"iron":0.6,"zinc":0.3,"mang":0.45000002,"cop":0.15,"iod":0.0045,"chr":0,"mol":0,"sel":0,"vita":0.015,"vitb1":0.06,"vitb2":0.089999996,"vitb3":0,"vitb5":0,"vitb6":0.54,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":16.5,"vitd":0,"vite":0.40500003,"vitk":0}`
	Recipe2Json = `{"name":"Fruit Cake","items":[{"id":1,"amount":300},{"id":2,"amount":250}],"id":2,"size":3,"flags":3,"preptime":0,"cooktime":0,"misctime":0,"kcal":394.5,"fat":0.8,"fatsat":0,"fato3":0,"fato6":0,"carb":93.2,"sug":73.4,"fruc":40,"gluc":19,"suc":14.4,"fib":11,"prot":3.4,"salt":0,"pot":1274.5,"chl":279.1,"sod":5.5,"calc":33.4,"phos":88,"mag":90,"iron":1.6,"zinc":0.8,"mang":0.75,"cop":0.55,"iod":0.0315,"chr":0,"mol":0,"sel":0,"vita":0.055,"vitb1":0.22,"vitb2":0.23999998,"vitb3":0,"vitb5":0,"vitb6":1.2,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":63.5,"vitd":0,"vite":2.145,"vitk":0}`
)

const (
	Instructions1 = "Cook it well."
)

var (
	User1 = app.User{
		ID:    1,
		Email: "a@a.a",
		Name:  "AnnoyingOrange",
		Pass:  "$2a$10$CpVy94BcePvhBH3QS/mMnOtFVrfN0DvwdooEUc0T8tWdKNi3ayFXC", // plaintext: Tr0ub4dor&3
		Lang:  "en",
		Perm:  app.PermLogin,
	}

	User1Pass = "Tr0ub4dor&3"
	User1Json = `{"email":"a@a.a","name":"AnnoyingOrange","pass":"$2a$10$CpVy94BcePvhBH3QS/mMnOtFVrfN0DvwdooEUc0T8tWdKNi3ayFXC","lang":"en","perm":1,"id":1}`

	User2 = app.User{
		ID:    2,
		Email: "b@b.b",
		Name:  "BadApple",
		Pass:  "$2a$10$CpVy94BcePvhBH3QS/mMnOtFVrfN0DvwdooEUc0T8tWdKNi3ayFXC", // plaintext: Tr0ub4dor&3
		Lang:  "de",
		Perm:  app.PermAdmin,
	}
)

var (
	Prefs1 = app.Prefs{
		Account: app.AccountPrefs{
			Email: User1.Email,
			Name:  User1.Name,
		},
		Macros: [7]app.MacroPrefs{
			{KCal: 2200, Fat: 60, Carbs: 250, Protein: 165},
			{KCal: 2200, Fat: 60, Carbs: 250, Protein: 165},
			{KCal: 2600, Fat: 80, Carbs: 285, Protein: 185},
			{KCal: 2200, Fat: 60, Carbs: 250, Protein: 165},
			{KCal: 2200, Fat: 60, Carbs: 250, Protein: 165},
			{KCal: 2200, Fat: 60, Carbs: 250, Protein: 165},
			{KCal: 2600, Fat: 80, Carbs: 285, Protein: 185},
		},
		RDI: app.BaseRDIPrefs(),
	}

	StoredPrefs1 = app.StoredPrefs{
		Macros: [7]app.MacroPrefs{
			{KCal: 2200, Fat: 60, Carbs: 250, Protein: 165},
			{KCal: 2200, Fat: 60, Carbs: 250, Protein: 165},
			{KCal: 2600, Fat: 80, Carbs: 285, Protein: 185},
			{KCal: 2200, Fat: 60, Carbs: 250, Protein: 165},
			{KCal: 2200, Fat: 60, Carbs: 250, Protein: 165},
			{KCal: 2200, Fat: 60, Carbs: 250, Protein: 165},
			{KCal: 2600, Fat: 80, Carbs: 285, Protein: 185},
		},
	}

	Prefs1Json       = `{"account":{"email":"a@a.a","name":"AnnoyingOrange"},"macros":[{"kcal":2200,"fat":60,"carb":250,"prot":165},{"kcal":2200,"fat":60,"carb":250,"prot":165},{"kcal":2600,"fat":80,"carb":285,"prot":185},{"kcal":2200,"fat":60,"carb":250,"prot":165},{"kcal":2200,"fat":60,"carb":250,"prot":165},{"kcal":2200,"fat":60,"carb":250,"prot":165},{"kcal":2600,"fat":80,"carb":285,"prot":185}],"rdi":{"fatsat":22,"fato3":1.6,"fato6":3.2,"fib":32,"salt":5.8,"pot":3400,"chl":2300,"sod":2300,"calc":1000,"phos":700,"mag":400,"iron":8,"zinc":11,"mang":2.3,"cop":0.9,"iod":0.15,"chr":0.035,"mol":0.045,"sel":0.055,"vita":0.9,"vitb1":1.2,"vitb2":1.3,"vitb3":16,"vitb5":5,"vitb6":1.7,"vitb7":0.03,"vitb9":0.4,"vitb12":0.003,"vitc":90,"vitd":0.015,"vite":15,"vitk":0.12},"ui":{"neutralCharts":false,"trackSaltAsSodium":false}}`
	StoredPrefs1Json = `{"macros":[{"kcal":2200,"fat":60,"carb":250,"prot":165},{"kcal":2200,"fat":60,"carb":250,"prot":165},{"kcal":2600,"fat":80,"carb":285,"prot":185},{"kcal":2200,"fat":60,"carb":250,"prot":165},{"kcal":2200,"fat":60,"carb":250,"prot":165},{"kcal":2200,"fat":60,"carb":250,"prot":165},{"kcal":2600,"fat":80,"carb":285,"prot":185}],"neutralCharts":false,"trackSaltAsSodium":false}`

	StoredPrefs2 = app.StoredPrefs{
		Macros: [7]app.MacroPrefs{
			{KCal: 2000, Fat: 40, Carbs: 220, Protein: 60},
			{KCal: 2100, Fat: 45, Carbs: 230, Protein: 65},
			{KCal: 2200, Fat: 50, Carbs: 240, Protein: 70},
			{KCal: 2300, Fat: 55, Carbs: 250, Protein: 75},
			{KCal: 2400, Fat: 60, Carbs: 260, Protein: 80},
			{KCal: 2500, Fat: 65, Carbs: 270, Protein: 85},
			{KCal: 2600, Fat: 70, Carbs: 280, Protein: 90},
		},
		UIPrefs: app.UIPrefs{
			NeutralCharts: true,
		},
	}
)

func Prefs0(user app.User) app.Prefs {
	return app.Prefs{
		Account: app.AccountPrefs{
			Email: user.Email,
			Name:  user.Name,
		},
		Macros: app.BaseMacroPrefs(),
		RDI:    app.BaseRDIPrefs(),
	}
}

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
	Entry3Json = `{"date":"2022-03-13T08:33:00Z","recipe":"","food":{"id":1,"amount":50}}`
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

func Diary220312() core.DiaryDay {
	return core.DiaryDay{
		Date:       "2022-03-12",
		KCal:       Food1.KCal*0.5 + Food2.KCal*1.5,
		Fat:        Food1.Fat*0.5 + Food2.Fat*1.5,
		Carbs:      Food1.Carbs*0.5 + Food2.Carbs*1.5,
		Sugar:      Food1.Sugar*0.5 + Food2.Sugar*1.5,
		Fructose:   Food1.Fructose*0.5 + Food2.Fructose*1.5,
		Glucose:    Food1.Glucose*0.5 + Food2.Glucose*1.5,
		Sucrose:    Food1.Sucrose*0.5 + Food2.Sucrose*1.5,
		Protein:    Food1.Protein*0.5 + Food2.Protein*1.5,
		Fiber:      Food1.Fiber*0.5 + Food2.Fiber*1.5,
		Iron:       Food1.Iron*0.5 + Food2.Iron*1.5,
		Zinc:       Food1.Zinc*0.5 + Food2.Zinc*1.5,
		Magnesium:  Food1.Magnesium*0.5 + Food2.Magnesium*1.5,
		Chlorine:   Food1.Chlorine*0.5 + Food2.Chlorine*1.5,
		Sodium:     Food1.Sodium*0.5 + Food2.Sodium*1.5,
		Calcium:    Food1.Calcium*0.5 + Food2.Calcium*1.5,
		Potassium:  Food1.Potassium*0.5 + Food2.Potassium*1.5,
		Phosphorus: Food1.Phosphorus*0.5 + Food2.Phosphorus*1.5,
		Copper:     Food1.Copper*0.5 + Food2.Copper*1.5,
		Iodine:     Food1.Iodine*0.5 + Food2.Iodine*1.5,
		Manganse:   Food1.Manganse*0.5 + Food2.Manganse*1.5,
		VitA:       Food1.VitA*0.5 + Food2.VitA*1.5,
		VitB1:      Food1.VitB1*0.5 + Food2.VitB1*1.5,
		VitB2:      Food1.VitB2*0.5 + Food2.VitB2*1.5,
		VitB6:      Food1.VitB6*0.5 + Food2.VitB6*1.5,
		VitC:       Food1.VitC*0.5 + Food2.VitC*1.5,
		VitE:       Food1.VitE*0.5 + Food2.VitE*1.5,
	}
}

func Diary220313() core.DiaryDay {
	return core.DiaryDay{
		Date:       "2022-03-13",
		KCal:       Food1.KCal * 0.5,
		Fat:        Food1.Fat * 0.5,
		Carbs:      Food1.Carbs * 0.5,
		Sugar:      Food1.Sugar * 0.5,
		Fructose:   Food1.Fructose * 0.5,
		Glucose:    Food1.Glucose * 0.5,
		Sucrose:    Food1.Sucrose * 0.5,
		Protein:    Food1.Protein * 0.5,
		Fiber:      Food1.Fiber * 0.5,
		Iron:       Food1.Iron * 0.5,
		Zinc:       Food1.Zinc * 0.5,
		Magnesium:  Food1.Magnesium * 0.5,
		Chlorine:   Food1.Chlorine * 0.5,
		Sodium:     Food1.Sodium * 0.5,
		Calcium:    Food1.Calcium * 0.5,
		Potassium:  Food1.Potassium * 0.5,
		Phosphorus: Food1.Phosphorus * 0.5,
		Copper:     Food1.Copper * 0.5,
		Iodine:     Food1.Iodine * 0.5,
		Manganse:   Food1.Manganse * 0.5,
		VitA:       Food1.VitA * 0.5,
		VitB1:      Food1.VitB1 * 0.5,
		VitB2:      Food1.VitB2 * 0.5,
		VitB6:      Food1.VitB6 * 0.5,
		VitC:       Food1.VitC * 0.5,
		VitE:       Food1.VitE * 0.5,
	}
}

const (
	Diary210101Json = `{"date":"2021-01-01","kcal":0,"fat":0,"fatsat":0,"fato3":0,"fato6":0,"carb":0,"sug":0,"fruc":0,"gluc":0,"suc":0,"fib":0,"prot":0,"salt":0,"pot":0,"chl":0,"sod":0,"calc":0,"phos":0,"mag":0,"iron":0,"zinc":0,"mang":0,"cop":0,"iod":0,"chr":0,"mol":0,"sel":0,"vita":0,"vitb1":0,"vitb2":0,"vitb3":0,"vitb5":0,"vitb6":0,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":0,"vitd":0,"vite":0,"vitk":0}`
	Diary210102Json = `{"date":"2021-01-02","kcal":0,"fat":0,"fatsat":0,"fato3":0,"fato6":0,"carb":0,"sug":0,"fruc":0,"gluc":0,"suc":0,"fib":0,"prot":0,"salt":0,"pot":0,"chl":0,"sod":0,"calc":0,"phos":0,"mag":0,"iron":0,"zinc":0,"mang":0,"cop":0,"iod":0,"chr":0,"mol":0,"sel":0,"vita":0,"vitb1":0,"vitb2":0,"vitb3":0,"vitb5":0,"vitb6":0,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":0,"vitd":0,"vite":0,"vitk":0}`
	Diary220301Json = `{"date":"2022-03-01","kcal":0,"fat":0,"fatsat":0,"fato3":0,"fato6":0,"carb":0,"sug":0,"fruc":0,"gluc":0,"suc":0,"fib":0,"prot":0,"salt":0,"pot":0,"chl":0,"sod":0,"calc":0,"phos":0,"mag":0,"iron":0,"zinc":0,"mang":0,"cop":0,"iod":0,"chr":0,"mol":0,"sel":0,"vita":0,"vitb1":0,"vitb2":0,"vitb3":0,"vitb5":0,"vitb6":0,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":0,"vitd":0,"vite":0,"vitk":0}`
	Diary220312Json = `{"date":"2022-03-12","kcal":166.5,"fat":0.35000002,"fatsat":0,"fato3":0,"fato6":0,"carb":37.2,"sug":30.65,"fruc":17.5,"gluc":7.5,"suc":5.65,"fib":4,"prot":1.65,"salt":0,"pot":610,"chl":164.6,"sod":2,"calc":13.15,"phos":38.5,"mag":47.5,"iron":0.70000005,"zinc":0.35000002,"mang":0.45000002,"cop":0.2,"iod":0.0085,"chr":0,"mol":0,"sel":0,"vita":0.02,"vitb1":0.08,"vitb2":0.105,"vitb3":0,"vitb5":0,"vitb6":0.59000003,"vitb7":0,"vitb9":0,"vitb12":0,"vitc":22.5,"vitd":0,"vite":0.65000004,"vitk":0}`
)

func List1() []core.ShopItem {
	return []core.ShopItem{
		{
			ID:     Entry2().Food.ID,
			Amount: Entry2().Food.Amount,
			Price:  [2]float32{},
			Aisle:  0,
			Done:   false,
		},
		{
			ID:     Entry1().Food.ID,
			Amount: Entry1().Food.Amount,
			Price:  [2]float32{},
			Aisle:  0,
			Done:   false,
		},
	}
}

func List12() []core.ShopItem {
	return []core.ShopItem{
		{
			ID:     Entry2().Food.ID,
			Amount: Entry2().Food.Amount + Entry3().Food.Amount,
			Price:  [2]float32{},
			Aisle:  0,
			Done:   false,
		},
		{
			ID:     Entry1().Food.ID,
			Amount: Entry1().Food.Amount + Entry4().Food.Amount,
			Price:  [2]float32{},
			Aisle:  0,
			Done:   false,
		},
	}
}

const (
	List1Json = `[{"price":[0,0],"done":false,"amount":50,"aisle":0,"id":1},{"price":[0,0],"done":false,"amount":150,"aisle":0,"id":2}]`
)

func USDA1() []byte {
	if file, err := convFS.Open("usda.json"); err == nil {
		defer file.Close()
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(file); err == nil {
			return buf.Bytes()
		}
	}
	panic("usda")
}
