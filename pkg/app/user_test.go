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

package app_test

import (
	"heyapple/internal/mock"
	"heyapple/pkg/app"
	"testing"
)

func TestAuthenticate_Fetch(t *testing.T) {
	for idx, data := range []struct {
		email string
		pass  string
		db    *mock.DB

		id  int
		err error
	}{
		{ //00// connection failed
			db:  mock.NewDB().WithError(mock.ErrDOS),
			err: mock.ErrDOS,
		},
		{ //01// empty DB
			email: "a@a.a",
			db:    mock.NewDB(),
			err:   app.ErrNotFound,
		},
		{ //02// user not found
			email: "a@a.a",
			db:    mock.NewDB().WithUser(app.User{Email: "b@b.b"}),
			err:   app.ErrNotFound,
		},
		{ //03// passwords don't match
			email: "a@a.a",
			pass:  "Tr0ub4dor&3",
			db: mock.NewDB().WithUser(app.User{
				Email: "a@a.a",
				Pass:  "$2a$10$ADm2JBRbt8UvB0uI7NNFBupOdTq7XKae6Dvc7NfVCnw89rPZr3.zK",
			}),
			err: app.ErrCredentials,
		},
		{ //04// success
			email: "a@a.a",
			pass:  "password123",
			db: mock.NewDB().WithUser(app.User{
				ID:    42,
				Email: "a@a.a",
				Pass:  "$2a$10$ADm2JBRbt8UvB0uI7NNFBupOdTq7XKae6Dvc7NfVCnw89rPZr3.zK",
			}),
			id: 42,
		},
	} {
		qry := &app.Authenticate{Email: data.email, Pass: data.pass}
		err := qry.Fetch(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v \nwant: %v", idx, err, data.err)
		}

		if err == nil && qry.Pass != "" {
			t.Errorf("test case %d: password not cleared", idx)
		}

		if qry.ID != data.id {
			t.Errorf("test case %d: id mismatch \nhave: %v \nwant: %v", idx, qry.ID, data.id)
		}
	}
}

func TestCreateUser_Execute(t *testing.T) {
	for idx, data := range []struct {
		db    *mock.DB
		email string
		pass  string

		err error
	}{
		{ //00// database failure
			db:  mock.NewDB().WithError(mock.ErrDOS),
			err: mock.ErrDOS,
		},
		{ //01// username already exists
			db:    mock.NewDB().WithUser(app.User{Email: "a@a.a"}),
			email: "a@a.a",
			err:   app.ErrExists,
		},
		{ //02// deferred database failure
			db:    mock.NewDB().WithError(nil, mock.ErrDOS),
			email: "b@b.b",
			pass:  "topsecret",
			err:   mock.ErrDOS,
		},
		{ //03// success
			db:    mock.NewDB().WithUser(app.User{Email: "a@a.a"}),
			email: "b@b.b",
			pass:  "topsecret",
		},
	} {
		cmd := &app.CreateUser{Email: data.email, Pass: data.pass}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if cmd.ID != data.db.ID {
			t.Errorf("test case %d: id mismatch \nhave: %v\nwant: %v", idx, cmd.ID, data.db.ID)
		}

		if id := data.db.Token.ID; id != cmd.ID {
			t.Errorf("test case %d: token mismatch \nhave: %v\nwant: %v", idx, id, cmd.ID)
		}

		if err == nil && !app.NewCrypter().Match(data.db.User.Pass, data.pass) {
			t.Errorf("test case %d: password mismatch", idx)
		}
	}
}
