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

// StoredPrefs contains preferences that cannot be inferred
// from other objects. These are the absolute minimum that
// need to be stored in persistent storage in order to
// reconstruct a complete set of user preferences.
//
// Most notably, we really don't want to store individual
// RDI values because doing so would bloat the stored data
// size with largely the same values for most users.
type StoredPrefs struct {
	Macros [7]MacroPrefs `json:"macros"`
	UIPrefs
}

// Prefs holds all individual preferences types and is the
// object that gets passed around the application.
type Prefs struct {
	Account AccountPrefs  `json:"account"`
	Macros  [7]MacroPrefs `json:"macros"`
	RDI     RDIPrefs      `json:"rdi"`
	UI      UIPrefs       `json:"ui"`
}

// This contains all zeroed-out macro targets and represents
// the weekly macro prefs' zero value. Implemented as a function
// to ensure constness.
func EmptyMacroPrefs() [7]MacroPrefs {
	return [7]MacroPrefs{}
}

// This contains the average recommended macro targets for
// an average human being. Implemented as a function to
// ensure constness.
func BaseMacroPrefs() [7]MacroPrefs {
	return [7]MacroPrefs{
		{KCal: 2000, Fat: 60, Carbs: 270, Protein: 95},
		{KCal: 2000, Fat: 60, Carbs: 270, Protein: 95},
		{KCal: 2000, Fat: 60, Carbs: 270, Protein: 95},
		{KCal: 2000, Fat: 60, Carbs: 270, Protein: 95},
		{KCal: 2000, Fat: 60, Carbs: 270, Protein: 95},
		{KCal: 2000, Fat: 60, Carbs: 270, Protein: 95},
		{KCal: 2000, Fat: 60, Carbs: 270, Protein: 95},
	}
}

// This contains an average person's recommended daily intake
// for all the nutrients tracked by the system. It should be
// multiplied by values computed from a person's specific
// body composition and living circumstances in order to
// represent a correct estimation of their actual needs.
// Implemented as a function to ensure constness.
func BaseRDIPrefs() RDIPrefs {
	return RDIPrefs{
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
}

// Preferences is a query to collect all user preferences
// from the database. If successful, the Prefs field will
// contain all preferences that were found.
//
// As long as the user exists, a valid set of preferences
// will be generated. For users who never changed their
// settings in any way, this is identical to the set of
// default preferences.
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

	prefs, err := db.UserPrefs(q.ID)
	if err != nil {
		return err
	}

	q.Prefs.Account.Email = user.Email
	q.Prefs.Account.Name = user.Name
	q.Prefs.UI = prefs.UIPrefs

	if prefs.Macros == EmptyMacroPrefs() {
		q.Prefs.Macros = BaseMacroPrefs()
	} else {
		q.Prefs.Macros = prefs.Macros
	}

	q.Prefs.RDI = BaseRDIPrefs()

	return nil
}

// SavePreferences is a command to store all user
// preferences in the database. If successful, the
// Prefs field will contain an updated set of preferences.
// When the command fails, the value of Prefs is undefined
// and should not be relied on!
type SavePreferences struct {
	ID    int
	Prefs Prefs
}

func (c *SavePreferences) Execute(db DB) error {
	if c.ID == 0 {
		return ErrNotFound
	}

	prefs := StoredPrefs{
		Macros:  c.Prefs.Macros,
		UIPrefs: c.Prefs.UI,
	}
	if err := db.SetUserPrefs(c.ID, prefs); err != nil {
		return err
	}

	query := Preferences{ID: c.ID}
	if err := query.Fetch(db); err != nil {
		return err
	}

	c.Prefs = query.Prefs
	return nil
}
