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

		id   int
		lang string
		err  error
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
		{ //04// missing login permissions
			email: "a@a.a",
			pass:  "password123",
			db: mock.NewDB().WithUser(app.User{
				ID:    42,
				Email: "a@a.a",
				Pass:  "$2a$10$ADm2JBRbt8UvB0uI7NNFBupOdTq7XKae6Dvc7NfVCnw89rPZr3.zK",
			}),
			err: app.ErrCredentials,
		},
		{ //05// success
			email: "a@a.a",
			pass:  "password123",
			db: mock.NewDB().WithUser(app.User{
				ID:    42,
				Email: "a@a.a",
				Pass:  "$2a$10$ADm2JBRbt8UvB0uI7NNFBupOdTq7XKae6Dvc7NfVCnw89rPZr3.zK",
				Lang:  "en",
				Perm:  app.PermLogin,
			}),
			id:   42,
			lang: "en",
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

		if qry.Lang != data.lang {
			t.Errorf("test case %d: language mismatch \nhave: %v \nwant: %v", idx, qry.Lang, data.lang)
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

		if id := data.db.Tok.ID; id != cmd.ID {
			t.Errorf("test case %d: token mismatch \nhave: %v\nwant: %v", idx, id, cmd.ID)
		}

		if err == nil && !app.NewCrypter().Match(data.db.User.Pass, data.pass) {
			t.Errorf("test case %d: password mismatch", idx)
		}

		if err == nil && data.db.User.Perm != app.PermNone {
			t.Errorf("test case %d: permission mismatch \nhave: %v\nwant: %v", idx, data.db.User.Perm, app.PermNone)
		}
	}
}

func TestActivate_Execute(t *testing.T) {
	for idx, data := range []struct {
		db    *mock.DB
		token string

		perm int
		err  error
	}{
		{ //00// database failure
			db:  mock.NewDB().WithError(mock.ErrDOS),
			err: mock.ErrDOS,
		},
		{ //01// no token
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //02// user doesn't exist
			db:  mock.NewDB().WithToken(app.Token{ID: 1}),
			err: app.ErrNotFound,
		},
		{ //03// apply activation token
			db:   mock.NewDB().WithToken(app.Token{ID: 1}).WithUser(app.User{ID: 1}),
			perm: app.PermLogin,
		},
		{ //04// unexpected data type
			db:  mock.NewDB().WithToken(app.Token{ID: 1, Data: 12}).WithUser(app.User{ID: 1}),
			err: app.ErrNotFound,
		},
		{ //05// not a valid email address
			db:  mock.NewDB().WithToken(app.Token{ID: 1, Data: "noemail"}).WithUser(app.User{ID: 1}),
			err: app.ErrNotFound,
		},
		{ //06// apply email token
			db: mock.NewDB().WithToken(app.Token{ID: 1, Data: "a@a.a"}).WithUser(app.User{ID: 1}),
		},
	} {
		cmd := &app.Activate{Token: data.token}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if data.db.User.Perm != data.perm {
			t.Errorf("test case %d: permission mismatch \nhave: %v\nwant: %v", idx, data.db.User.Perm, data.perm)
		}

		if err == nil && data.db.Tok != (app.Token{}) {
			t.Errorf("test case %d: token mismatch \nhave: %v\nwant: %v", idx, data.db.Tok, app.Token{})
		}
	}
}

func TestSwitchLanguage_Execute(t *testing.T) {
	for idx, data := range []struct {
		db   *mock.DB
		id   int
		lang string

		err error
	}{
		{ //00// empty data
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// database failure
			db:   mock.NewDB().WithError(mock.ErrDOS),
			id:   1,
			lang: "fr",
			err:  mock.ErrDOS,
		},
		{ //02// user doesn't exist
			db:   mock.NewDB(),
			id:   1,
			lang: "fr",
			err:  app.ErrNotFound,
		},
		{ //03// success
			db:   mock.NewDB().WithUser(app.User{ID: 1}),
			id:   1,
			lang: "fr",
		},
	} {
		err := (&app.SwitchLanguage{ID: data.id, Lang: data.lang}).Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if err == nil && data.db.User.Lang != data.lang {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.User.Lang, data.lang)
		}
	}
}

func TestResetPassword_Execute(t *testing.T) {
	for idx, data := range []struct {
		db    *mock.DB
		email string

		token app.Token
		err   error
	}{
		{ //00// empty data
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// database failure
			db:    mock.NewDB().WithError(mock.ErrDOS),
			email: "a@a.a",
			err:   mock.ErrDOS,
		},
		{ //02// user doesn't exist
			db:    mock.NewDB(),
			email: "a@a.a",
			err:   app.ErrNotFound,
		},
		{ //03// database failure
			db:    mock.NewDB().WithUser(mock.User1).WithError(nil, mock.ErrDOS),
			email: "a@a.a",
			err:   mock.ErrDOS,
		},
		{ //04// success
			db:    mock.NewDB().WithUser(mock.User1),
			email: "a@a.a",
			token: app.Token{ID: mock.User1.ID, Data: "reset"},
		},
	} {
		cmd := &app.ResetPassword{Email: data.email}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if token := data.db.Tok; token != data.token {
			t.Errorf("test case %d: token mismatch \nhave: %v\nwant: %v", idx, token, data.token)
		}
	}
}

func TestChangePassword_Execute(t *testing.T) {
	for idx, data := range []struct {
		db    *mock.DB
		pass  string
		token string
		id    int

		err error
	}{
		{ //00// empty data
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// token doesn't exist
			db:    mock.NewDB(),
			token: "abcd",
			err:   app.ErrNotFound,
		},
		{ //02// missing token data
			db:    mock.NewDB().WithToken(app.Token{ID: 1}),
			token: "abcd",
			err:   app.ErrNotFound,
		},
		{ //03// wrong token data
			db:    mock.NewDB().WithToken(app.Token{ID: 1, Data: "wrong"}),
			token: "abcd",
			err:   app.ErrNotFound,
		},
		{ //04// user doesn't exist
			db:    mock.NewDB().WithToken(app.Token{ID: 1, Data: "reset"}),
			token: "abcd",
			err:   app.ErrNotFound,
		},
		{ //05// user doesn't exist
			db:  mock.NewDB(),
			id:  1,
			err: app.ErrNotFound,
		},
		{ //06// deferred failure
			db:  mock.NewDB().WithUser(mock.User1).WithError(nil, mock.ErrDOS),
			id:  1,
			err: mock.ErrDOS,
		},
		{ //07// success for id path
			db:   mock.NewDB().WithUser(mock.User1),
			pass: "supersecret",
			id:   1,
		},
		{ //08// success for token path
			db:    mock.NewDB().WithUser(mock.User1).WithToken(app.Token{ID: mock.User1.ID, Data: "reset"}),
			pass:  "supersecret",
			token: "abcd",
		},
	} {
		cmd := &app.ChangePassword{Token: data.token, ID: data.id, Pass: data.pass}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if err == nil {
			if data.token != "" && data.db.Tok != (app.Token{}) {
				t.Errorf("test case %d: token mismatch \nhave: %v\nwant: %v", idx, data.db.Tok, app.Token{})
			}

			if !app.NewCrypter().Match(data.db.User.Pass, data.pass) {
				t.Errorf("test case %d: password mismatch", idx)
			}
		}

	}
}
