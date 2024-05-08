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

var (
	prefBefore = app.Prefs{
		Macros: app.BaseMacroPrefs(),
		UI: app.UIPrefs{
			NeutralCharts:     true,
			TrackSaltAsSodium: true,
		},
	}

	prefAfter = app.Prefs{
		Account: app.AccountPrefs{
			Email: mock.User1.Email,
			Name:  mock.User1.Name,
		},
		Macros: app.BaseMacroPrefs(),
		RDI:    app.BaseRDIPrefs(),
		UI: app.UIPrefs{
			NeutralCharts:     true,
			TrackSaltAsSodium: true,
		},
	}
)

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
			db:   mock.NewDB().WithError(mock.ErrDOS),
			user: 1,
			err:  mock.ErrDOS,
		},
		{ //03// delayed database error
			db:   mock.NewDB().WithUser(mock.User1).WithError(nil, mock.ErrDOS),
			user: 1,
			err:  mock.ErrDOS,
		},
		{ //04// success
			db:    mock.NewDB().WithUser(mock.User1).WithPrefs(mock.StoredPrefs1),
			user:  mock.User1.ID,
			prefs: mock.Prefs1,
		},
		{ //05// success, user doesn't have stored prefs yet
			db:    mock.NewDB().WithUser(mock.User1),
			user:  mock.User1.ID,
			prefs: mock.Prefs0(mock.User1),
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

func TestSavePreferences_Execute(t *testing.T) {
	for idx, data := range []struct {
		db   *mock.DB
		user int
		in   app.Prefs

		out app.Prefs
		err error
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
			user: 1,
			err:  mock.ErrDOS,
		},
		{ //03// delayed error
			db:   mock.NewDB().WithUser(mock.User1).WithError(nil, mock.ErrDOS),
			user: 1,
			err:  mock.ErrDOS,
		},
		{ //04// success
			db:   mock.NewDB().WithUser(mock.User1).WithPrefs(mock.StoredPrefs1),
			user: 1,
			in:   prefBefore,
			out:  prefAfter,
		},
	} {
		cmd := &app.SavePreferences{ID: data.user, Prefs: data.in}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if err == nil && cmd.Prefs != data.out {
			t.Errorf("test case %d: preferences mismatch \nhave: %v \nwant: %v", idx, cmd.Prefs, data.out)
		}
	}
}
