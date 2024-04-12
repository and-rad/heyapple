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

package mem

import (
	"reflect"
	"sort"
	"testing"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/core"
	"github.com/and-rad/heyapple/internal/mock"
)

func TestDB_UserByEmail(t *testing.T) {
	for idx, data := range []struct {
		db    *DB
		email string

		user app.User
		err  error
	}{
		{ //00// empty database
			db:  NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// user doesn't exist
			db: &DB{
				users:  map[int]app.User{1: {ID: 1, Email: "a@a.a"}},
				emails: map[string]int{"a@a.a": 1},
			},
			email: "b@b.b",
			err:   app.ErrNotFound,
		},
		{ //02// success
			db: &DB{
				users: map[int]app.User{
					1: {ID: 1, Email: "a@a.a"},
					2: {ID: 2, Email: "b@b.b"},
				},
				emails: map[string]int{
					"a@a.a": 1,
					"b@b.b": 2,
				},
			},
			email: "b@b.b",
			user:  app.User{ID: 2, Email: "b@b.b"},
		},
	} {
		user, err := data.db.UserByEmail(data.email)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if user != data.user {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, user, data.user)
		}
	}
}

func TestDB_UserByName(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		name string

		user app.User
		err  error
	}{
		{ //00// empty database
			db:  NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// username doesn't exist
			db: &DB{
				users:  map[int]app.User{1: {ID: 1, Email: "a@a.a"}, 2: {ID: 2, Email: "b@b.b", Name: "BadApple"}},
				emails: map[string]int{"a@a.a": 1, "b@b.b": 2},
			},
			name: "AnnoyingOrange",
			err:  app.ErrNotFound,
		},
		{ //02// success
			db: &DB{
				users: map[int]app.User{
					1: {ID: 1, Email: "a@a.a", Name: "GoodOrange"},
					2: {ID: 2, Email: "b@b.b", Name: "BadApple"},
				},
				emails: map[string]int{
					"a@a.a": 1,
					"b@b.b": 2,
				},
			},
			name: "BadApple",
			user: app.User{ID: 2, Email: "b@b.b", Name: "BadApple"},
		},
	} {
		user, err := data.db.UserByName(data.name)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if user != data.user {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, user, data.user)
		}
	}
}

func TestDB_UserNames(t *testing.T) {
	for idx, data := range []struct {
		db     *DB
		prefix string

		names []string
	}{
		{ //00// empty database
			db:     NewDB(),
			prefix: "BadApple",
			names:  []string{},
		},
		{ //01// non-empty database
			db: &DB{
				users: map[int]app.User{
					1: {ID: 1, Email: "a@a.a", Name: "GoodOrange"},
					2: {ID: 2, Email: "b@b.b", Name: "BadApple"},
					3: {ID: 3, Email: "c@c.c", Name: "BadApple4"},
				},
			},
			prefix: "BadApple",
			names:  []string{"BadApple", "BadApple4"},
		},
		{ //02// prefix not among existing names
			db: &DB{
				users: map[int]app.User{
					1: {ID: 1, Email: "a@a.a", Name: "GoodOrange"},
					2: {ID: 2, Email: "b@b.b", Name: "BadApple"},
					3: {ID: 3, Email: "c@c.c", Name: "BadApple4"},
				},
			},
			prefix: "GoodApple",
			names:  []string{},
		},
		{ //03// empty prefix returns all names
			db: &DB{
				users: map[int]app.User{
					1: {ID: 1, Email: "a@a.a", Name: "GoodOrange"},
					2: {ID: 2, Email: "b@b.b", Name: "BadApple"},
					3: {ID: 3, Email: "c@c.c", Name: "BadApple4"},
				},
			},
			prefix: "",
			names:  []string{"BadApple", "BadApple4", "GoodOrange"},
		},
	} {
		names, err := data.db.UserNames(data.prefix)

		// sort names for consistent test results
		sort.Strings(names)

		if err != nil {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, nil)
		}

		if !reflect.DeepEqual(names, data.names) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, names, data.names)
		}
	}
}

func TestDB_NewUser(t *testing.T) {
	for idx, data := range []struct {
		db    *DB
		email string
		hash  string
		token string

		user app.User
		err  error
	}{
		{ //00// empty database
			db:    NewDB(),
			email: "a@a.a",
			hash:  "djwrifkgh",
			user:  app.User{ID: 1, Email: "a@a.a", Pass: "djwrifkgh", Lang: getConfig().defaultLang},
		},
		{ //01// username already exists
			db: &DB{
				emails: map[string]int{"a@a.a": 1},
				users:  map[int]app.User{1: {ID: 1, Email: "a@a.a", Pass: "qpwoeirutz"}},
				userID: 1,
			},
			email: "a@a.a",
			hash:  "djwrifkgh",
			user:  app.User{ID: 1, Email: "a@a.a", Pass: "qpwoeirutz"},
			err:   app.ErrExists,
		},
		{ //02// success
			db: &DB{
				emails: map[string]int{"a@a.a": 1},
				users:  map[int]app.User{1: {ID: 1, Email: "a@a.a", Pass: "qpwoeirutz"}},
				tokens: map[string]app.Token{},
				userID: 1,
			},
			email: "b@b.b",
			hash:  "djwrifkgh",
			token: "aabbccdd",
			user:  app.User{ID: 2, Email: "b@b.b", Pass: "djwrifkgh", Lang: getConfig().defaultLang},
		},
	} {
		id, err := data.db.NewUser(data.email, data.hash, data.token)

		if err == nil && id != data.db.userID {
			t.Errorf("test case %d: id mismatch \nhave: %v\nwant: %v", idx, id, data.db.userID)
		}

		if err == nil && id != data.db.tokens[data.token].ID {
			t.Errorf("test case %d: token mismatch \nhave: %v\nwant: %v", idx, data.db.tokens[data.token].ID, id)
		}

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if u, _ := data.db.UserByEmail(data.email); u != data.user {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, u, data.user)
		}
	}
}

func TestDB_SetUser(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		user app.User

		err error
	}{
		{ //00// empty database
			db:  NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// user doesn't exist
			db:   &DB{users: map[int]app.User{2: {ID: 2}}},
			user: mock.User1,
			err:  app.ErrNotFound,
		},
		{ //02// success
			db: &DB{
				users:  map[int]app.User{1: mock.User1},
				emails: map[string]int{mock.User1.Email: 1},
			},
			user: app.User{ID: 1, Email: "b@b.b", Pass: "kdjfhghr", Perm: app.PermLogin},
		},
	} {
		err := data.db.SetUser(data.user)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if u, err := data.db.UserByID(data.user.ID); err == nil && u != data.user {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, u, data.user)
		}

		if err == nil && data.db.emails[data.user.Email] != data.user.ID {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.emails, data.user.Email)
		}
	}
}

func TestDB_DelUser(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int

		out *DB
		err error
	}{
		{ //00// empty database, no error
			db:  NewDB(),
			out: NewDB(),
			err: nil,
		},
		{ //01// user doesn't exist, no error
			db:  &DB{users: map[int]app.User{2: {ID: 2}}},
			id:  1,
			out: &DB{users: map[int]app.User{2: {ID: 2}}},
			err: nil,
		},
		{ //02// success, simple
			db: &DB{
				users:  map[int]app.User{1: mock.User1},
				emails: map[string]int{mock.User1.Email: 1},
			},
			id: 1,
			out: &DB{
				users:  map[int]app.User{},
				emails: map[string]int{},
			},
		},
		{ //03// success, complex
			db: &DB{
				users:   map[int]app.User{1: mock.User1, 2: mock.User2},
				prefs:   map[int]app.StoredPrefs{1: mock.StoredPrefs1, 2: mock.StoredPrefs1},
				emails:  map[string]int{mock.User1.Email: 1, mock.User2.Email: 2},
				recipes: map[int]core.Recipe{mock.Recipe1().ID: mock.Recipe1(), mock.Recipe2().ID: mock.Recipe2()},
				userRec: map[int]map[int]int{
					mock.User1.ID: {mock.Recipe1().ID: app.PermOwner, mock.Recipe2().ID: app.PermRead},
					mock.User2.ID: {mock.Recipe1().ID: app.PermRead, mock.Recipe2().ID: app.PermOwner},
				},
				recUser: map[int]map[int]int{
					mock.Recipe1().ID: {mock.User1.ID: app.PermOwner, mock.User2.ID: app.PermRead},
					mock.Recipe2().ID: {mock.User1.ID: app.PermRead, mock.User2.ID: app.PermOwner},
				},
				entries: entryMap{
					mock.User1.ID: {mock.Day1: {mock.Entry1(), mock.Entry2()}},
					mock.User2.ID: {mock.Day2: {mock.Entry3(), mock.Entry4()}},
				},
				days: dayMap{
					mock.User1.ID: {2022: {3: {mock.Diary220312()}}},
				},
				aisles: aisleMap{
					0:             {1: 0, 4: 0, 26: 0},
					mock.User1.ID: {1: 2, 4: 4, 26: 1},
					mock.User2.ID: {3: 9, 4: 12, 26: 2},
				},
				prices: priceMap{
					mock.User1.ID: {2: [2]float32{1.2, 3.5}},
					mock.User2.ID: {1: [2]float32{32, 45}},
				},
				done: doneMap{
					mock.User1.ID: {2: true, 3: true},
					mock.User2.ID: {12: false, 45: true},
				},
			},
			id: 1,
			out: &DB{
				users:   map[int]app.User{2: mock.User2},
				prefs:   map[int]app.StoredPrefs{2: mock.StoredPrefs1},
				emails:  map[string]int{mock.User2.Email: 2},
				recipes: map[int]core.Recipe{mock.Recipe1().ID: mock.Recipe1(), mock.Recipe2().ID: mock.Recipe2()},
				userRec: map[int]map[int]int{
					mock.User2.ID: {mock.Recipe1().ID: app.PermRead, mock.Recipe2().ID: app.PermOwner},
				},
				recUser: map[int]map[int]int{
					mock.Recipe1().ID: {mock.User2.ID: app.PermRead},
					mock.Recipe2().ID: {mock.User2.ID: app.PermOwner},
				},
				entries: entryMap{
					mock.User2.ID: {mock.Day2: {mock.Entry3(), mock.Entry4()}},
				},
				days: dayMap{},
				aisles: aisleMap{
					0:             {1: 0, 4: 0, 26: 0},
					mock.User2.ID: {3: 9, 4: 12, 26: 2},
				},
				prices: priceMap{
					mock.User2.ID: {1: [2]float32{32, 45}},
				},
				done: doneMap{
					mock.User2.ID: {12: false, 45: true},
				},
			},
		},
	} {
		err := data.db.DelUser(data.id)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.db, data.out) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db, data.out)
		}
	}
}

func TestDB_UserByID(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int

		user app.User
		err  error
	}{
		{ //00// empty database
			db:  NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// user doesn't exist
			db: &DB{
				users: map[int]app.User{1: {ID: 1, Email: "a@a.a"}},
			},
			id:  2,
			err: app.ErrNotFound,
		},
		{ //02// success
			db: &DB{
				users: map[int]app.User{
					1: {ID: 1, Email: "a@a.a"},
					2: {ID: 2, Email: "b@b.b"},
				},
			},
			id:   2,
			user: app.User{ID: 2, Email: "b@b.b"},
		},
	} {
		user, err := data.db.UserByID(data.id)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if user != data.user {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, user, data.user)
		}
	}
}

func TestDB_UserPrefs(t *testing.T) {
	for idx, data := range []struct {
		db *DB
		id int

		prefs app.StoredPrefs
		err   error
	}{
		{ //00// empty database
			db:  NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// user doesn't exist
			db: &DB{
				users: map[int]app.User{1: {ID: 1, Email: "a@a.a"}},
			},
			id:  2,
			err: app.ErrNotFound,
		},
		{ //02// success
			db: &DB{
				users: map[int]app.User{
					1: mock.User1,
					2: mock.User2,
				},
				prefs: map[int]app.StoredPrefs{
					1: mock.StoredPrefs1,
				},
			},
			id:    1,
			prefs: mock.StoredPrefs1,
		},
		{ //03// success, but no user-specific prefs yet
			db: &DB{
				users: map[int]app.User{
					1: mock.User1,
					2: mock.User2,
				},
				prefs: map[int]app.StoredPrefs{
					2: mock.StoredPrefs1,
				},
			},
			id:    1,
			prefs: app.StoredPrefs{},
		},
	} {
		prefs, err := data.db.UserPrefs(data.id)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if prefs != data.prefs {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, prefs, data.prefs)
		}
	}
}

func TestDB_Token(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		hash string

		token app.Token
		err   error
	}{
		{ //00// empty database
			db:  NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// token not found
			db:   &DB{tokens: map[string]app.Token{"abcd": {ID: 1}}},
			hash: "bbbb",
			err:  app.ErrNotFound,
		},
		{ //02// success
			db:    &DB{tokens: map[string]app.Token{"abcd": {ID: 1, Data: "hi"}}},
			hash:  "abcd",
			token: app.Token{ID: 1, Data: "hi"},
		},
	} {
		token, err := data.db.Token(data.hash)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if token != data.token {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, token, data.token)
		}
	}
}

func TestDB_DelToken(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		hash string

		tokens map[string]app.Token
		err    error
	}{
		{ //00// empty database, no error
			db:     NewDB(),
			tokens: make(map[string]app.Token),
		},
		{ //01// token deleted
			db: &DB{tokens: map[string]app.Token{
				"abcd": {ID: 1, Data: "hi"},
				"efef": {ID: 2, Data: 9000},
			}},
			hash: "abcd",
			tokens: map[string]app.Token{
				"efef": {ID: 2, Data: 9000},
			},
		},
	} {
		err := data.db.DelToken(data.hash)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.db.tokens, data.tokens) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.tokens, data.tokens)
		}
	}
}

func TestDB_NewToken(t *testing.T) {
	for idx, data := range []struct {
		db   *DB
		id   int
		hash string
		data interface{}

		token app.Token
		err   error
	}{
		{ //00//
			db:    NewDB(),
			token: app.Token{},
		},
		{ //00//
			db:    NewDB(),
			id:    12,
			hash:  "abcd",
			data:  true,
			token: app.Token{ID: 12, Data: true},
		},
	} {
		err := data.db.NewToken(data.id, data.hash, data.data)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if tok, _ := data.db.Token(data.hash); tok != data.token {
			t.Errorf("test case %d: token mismatch \nhave: %v\nwant: %v", idx, tok, data.token)
		}
	}
}
