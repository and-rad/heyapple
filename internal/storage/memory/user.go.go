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

package memory

import (
	"heyapple/internal/app"
)

func (db *DB) NewUser(name, hash, token string) (int, error) {
	if _, ok := db.emails[name]; ok {
		return 0, app.ErrExists
	}

	db.userID++
	db.emails[name] = db.userID
	db.tokens[token] = app.Token{ID: db.userID}
	db.users[db.userID] = app.User{
		ID:    db.userID,
		Email: name,
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

func (db *DB) UserByName(name string) (app.User, error) {
	if id, ok := db.emails[name]; ok {
		if user, ok := db.users[id]; ok {
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

func (db *DB) DeleteToken(hash string) error {
	delete(db.tokens, hash)
	return nil
}
