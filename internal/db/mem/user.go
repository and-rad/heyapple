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
	"strings"

	"github.com/and-rad/heyapple/internal/app"
)

func (db *DB) NewUser(email, hash, token string) (int, error) {
	if _, ok := db.emails[email]; ok {
		return 0, app.ErrExists
	}

	db.userID++
	db.emails[email] = db.userID
	db.tokens[token] = app.Token{ID: db.userID}
	db.users[db.userID] = app.User{
		ID:    db.userID,
		Email: email,
		Pass:  hash,
		Lang:  getConfig().defaultLang,
	}

	return db.userID, nil
}

func (db *DB) SetUser(user app.User) error {
	if u, ok := db.users[user.ID]; ok {
		delete(db.emails, u.Email)
		db.users[user.ID] = user
		db.emails[user.Email] = user.ID
		return nil
	}
	return app.ErrNotFound
}

func (db *DB) DelUser(id int) error {
	user, ok := db.users[id]
	if !ok {
		return nil
	}

	delete(db.emails, user.Email)
	delete(db.users, id)
	delete(db.entries, id)
	delete(db.days, id)
	delete(db.aisles, id)
	delete(db.prices, id)
	delete(db.done, id)

	recs, ok := db.userRec[id]
	if !ok {
		return nil
	}

	delete(db.userRec, id)
	for r := range recs {
		if rec, ok := db.recUser[r]; ok {
			delete(rec, id)
		}
	}

	return nil
}

func (db *DB) UserByEmail(email string) (app.User, error) {
	if id, ok := db.emails[email]; ok {
		if user, ok := db.users[id]; ok {
			return user, nil
		}
	}
	return app.User{}, app.ErrNotFound
}

func (db *DB) UserByName(name string) (app.User, error) {
	for _, user := range db.users {
		if user.Name == name {
			return user, nil
		}
	}
	return app.User{}, app.ErrNotFound
}

func (db *DB) UserByID(id int) (app.User, error) {
	if user, ok := db.users[id]; ok {
		return user, nil
	}
	return app.User{}, app.ErrNotFound
}

func (db *DB) UserNames(prefix string) ([]string, error) {
	names := []string{}
	for _, user := range db.users {
		if strings.HasPrefix(user.Name, prefix) {
			names = append(names, user.Name)
		}
	}
	return names, nil
}

func (db *DB) Token(hash string) (app.Token, error) {
	if token, ok := db.tokens[hash]; ok {
		return token, nil
	}
	return app.Token{}, app.ErrNotFound
}

func (db *DB) NewToken(id int, hash string, data interface{}) error {
	db.tokens[hash] = app.Token{ID: id, Data: data}
	return nil
}

func (db *DB) DelToken(hash string) error {
	delete(db.tokens, hash)
	return nil
}
