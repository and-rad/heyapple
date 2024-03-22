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

package app_test

import (
	"testing"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/mock"
)

var baseRDI = app.RDIPrefs{
	KCal:       2000,
	Fat:        60,
	Carbs:      270,
	Protein:    80,
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

func TestPreferences_Fetch(t *testing.T) {
	for idx, data := range []struct {
		db   *mock.DB
		user int

		prefs app.Prefs
		err   error
	}{
		{ //00// id missing or invalid
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// user does not exist
			db:   mock.NewDB().WithUser(mock.User1),
			user: 23,
			err:  app.ErrNotFound,
		},
		{ //02// database error
			db:   mock.NewDB().WithUser(mock.User1).WithError(mock.ErrDOS),
			user: mock.User1.ID,
			err:  mock.ErrDOS,
		},
		{ //03// success
			db:   mock.NewDB().WithUser(mock.User1),
			user: mock.User1.ID,
			prefs: app.Prefs{
				Account: app.AccountPrefs{Email: mock.User1.Email, Name: mock.User1.Name},
				RDI:     baseRDI,
			},
		},
	} {
		qry := &app.Preferences{ID: data.user}
		err := qry.Fetch(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if qry.Prefs != data.prefs {
			t.Errorf("test case %d: preferences mismatch \nhave: %v \nwant: %v", idx, qry.Prefs, data.prefs)
		}
	}
}
