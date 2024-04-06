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

package app

// AccountPrefs holds all user settings directly related to
// the user's account.
type AccountPrefs struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// UIPrefs holds all user settings related to the presentation
// of data in a client application. It contains options for
// filtering, sorting, and display preferences.
type UIPrefs struct {
	NeutralCharts     bool `json:"neutralCharts"`
	TrackSaltAsSodium bool `json:"trackSaltAsSodium"`
}

// MacroPrefs holds all user settings related to personal
// macronutrient targets.
type MacroPrefs struct {
	KCal    float32 `json:"kcal"`
	Fat     float32 `json:"fat"`
	Carbs   float32 `json:"carb"`
	Protein float32 `json:"prot"`
}

// RDIPrefs holds all user settings related to the
// recommended daily intake of various nutrients.
type RDIPrefs struct {
	FatSat float32 `json:"fatsat"`
	FatO3  float32 `json:"fato3"`
	FatO6  float32 `json:"fato6"`
	Fiber  float32 `json:"fib"`
	Salt   float32 `json:"salt"`

	Potassium  float32 `json:"pot"`
	Chlorine   float32 `json:"chl"`
	Sodium     float32 `json:"sod"`
	Calcium    float32 `json:"calc"`
	Phosphorus float32 `json:"phos"`
	Magnesium  float32 `json:"mag"`
	Iron       float32 `json:"iron"`
	Zinc       float32 `json:"zinc"`
	Manganse   float32 `json:"mang"`
	Copper     float32 `json:"cop"`
	Iodine     float32 `json:"iod"`
	Chromium   float32 `json:"chr"`
	Molybdenum float32 `json:"mol"`
	Selenium   float32 `json:"sel"`

	VitA   float32 `json:"vita"`
	VitB1  float32 `json:"vitb1"`
	VitB2  float32 `json:"vitb2"`
	VitB3  float32 `json:"vitb3"`
	VitB5  float32 `json:"vitb5"`
	VitB6  float32 `json:"vitb6"`
	VitB7  float32 `json:"vitb7"`
	VitB9  float32 `json:"vitb9"`
	VitB12 float32 `json:"vitb12"`
	VitC   float32 `json:"vitc"`
	VitD   float32 `json:"vitd"`
	VitE   float32 `json:"vite"`
	VitK   float32 `json:"vitk"`
}

// Prefs holds all individual preferences types and is the
// object that gets passed around the application.
type Prefs struct {
	Account AccountPrefs  `json:"account"`
	Macros  [7]MacroPrefs `json:"macros"`
	RDI     RDIPrefs      `json:"rdi"`
	UI      UIPrefs       `json:"ui"`
}

// This contains the average recommended macro targets for
// an average human being.
var baseMacroPrefs = MacroPrefs{
	KCal:    2000,
	Fat:     60,
	Carbs:   270,
	Protein: 80,
}

// This contains an average person's recommended daily intake
// for all the nutrients tracked by the system. It should be
// multiplied by values computed from a person's specific
// body composition and living circumstances in order to
// represent a correct estimation of their actual needs.
var baseRDIPrefs = RDIPrefs{
	Fiber:      32,
	Salt:       5.8,
	FatSat:     22,
	FatO3:      1.6,
	FatO6:      3.2,
	VitA:       0.9,
	VitB1:      1.2,
	VitB2:      1.3,
	VitB3:      16,
	VitB5:      5,
	VitB6:      1.7,
	VitB7:      0.03,
	VitB9:      0.4,
	VitB12:     0.003,
	VitC:       90,
	VitD:       0.015,
	VitE:       15,
	VitK:       0.12,
	Potassium:  3400,
	Chlorine:   2300,
	Sodium:     2300,
	Calcium:    1000,
	Phosphorus: 700,
	Magnesium:  400,
	Iron:       8,
	Zinc:       11,
	Manganse:   2.3,
	Copper:     0.9,
	Iodine:     0.15,
	Chromium:   0.035,
	Molybdenum: 0.045,
	Selenium:   0.055,
}

// Preferences is a query to collect all user preferences
// from the database. If successful, the Prefs field will
// contains all preferences that were found.
type Preferences struct {
	ID    int
	Prefs Prefs
}

func (q *Preferences) Fetch(db DB) error {
	if q.ID == 0 {
		return ErrNotFound
	}

	user, err := db.UserByID(q.ID)
	if err != nil {
		return err
	}

	q.Prefs.Account = AccountPrefs{
		Email: user.Email,
		Name:  user.Name,
	}
	q.Prefs.RDI = baseRDIPrefs

	for i := range q.Prefs.Macros {
		q.Prefs.Macros[i] = baseMacroPrefs
	}

	return nil
}
