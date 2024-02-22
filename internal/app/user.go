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

const (
	PermNone  = 0x00000000
	PermLogin = 0x00000001

	PermCreate = 0x00000100
	PermRead   = 0x00000200
	PermEdit   = 0x00000400
	PermDelete = 0x00000800
	PermOwner  = PermCreate | PermRead | PermEdit | PermDelete

	PermCreateFood = 0x00010000
	PermEditFood   = 0x00020000
	PermAdmin      = PermCreateFood | PermEditFood
)

// User represents a human user of the system, identified
// by their e-mail address or id. The password is always
// stored as an encrypted hash.
type User struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
	Lang  string `json:"lang"`
	Perm  int    `json:"perm"`
	ID    int    `json:"id"`
}

// CreateUser is a command to create a new user with a
// unique id and e-mail address. If successful, the new
// user's id is stored in the command.
type CreateUser struct {
	Email string
	Pass  string
	Token string
	ID    int
}

func (c *CreateUser) Execute(db DB) error {
	if !NewValidator().MatchPass(c.Pass) {
		return ErrWeakPass
	}

	// Doing this early comes with a performance penalty
	// that is actually desirable here since it helps
	// making user enumeration attacks a little harder.
	hash := NewCrypter().Encrypt(c.Pass)
	token := NewTokenizer().Create()

	if _, err := db.UserByName(c.Email); err == nil {
		return ErrExists
	} else if err != ErrNotFound {
		return err
	}

	if id, err := db.NewUser(c.Email, hash, token); err != nil {
		return err
	} else {
		c.ID = id
		c.Pass = ""
		c.Token = token
	}

	return nil
}

// Authenticate is a query that authenticates a user by
// checking against e-mail and password. If successful,
// the user's id, language, and permissions are stored
// in the command.
type Authenticate struct {
	Email string
	Pass  string
	Lang  string
	Perm  int
	ID    int
}

func (q *Authenticate) Fetch(db DB) error {
	if user, err := db.UserByName(q.Email); err != nil {
		return err
	} else if !NewCrypter().Match(user.Pass, q.Pass) {
		return ErrCredentials
	} else if user.Perm&PermLogin != PermLogin {
		return ErrCredentials
	} else {
		q.ID = user.ID
		q.Perm = user.Perm
		q.Lang = user.Lang
		q.Pass = ""
	}

	return nil
}

// Activate is a command to unlock a user's ability
// to log into the application and to change their
// e-mail address after a change request was issued.
// If successful, the token is deleted and cannot be
// used again.
type Activate struct {
	Token string
}

func (c *Activate) Execute(db DB) error {
	tok, err := db.Token(c.Token)
	if err != nil {
		return err
	}

	user, err := db.UserByID(tok.ID)
	if err != nil {
		return err
	}

	if tok.Data == nil {
		user.Perm |= PermLogin
	} else if data, ok := tok.Data.(string); !ok {
		return ErrNotFound
	} else if !NewValidator().MatchEmail(data) {
		return ErrNotFound
	} else {
		user.Email = data
	}

	if err = db.SetUser(user); err == nil {
		err = db.DelToken(c.Token)
	}

	return err
}

// SwitchLanguage is a command to change a user's UI
// language preference.
type SwitchLanguage struct {
	Lang string
	ID   int
}

func (c *SwitchLanguage) Execute(db DB) error {
	if c.Lang == "" || c.ID == 0 {
		return ErrNotFound
	}

	u, err := db.UserByID(c.ID)
	if err != nil {
		return err
	}

	u.Lang = c.Lang
	return db.SetUser(u)
}

// ResetPassword is a command to create a password reset
// token that allows an anonymous user to change the
// password that's associated with a registered user.
type ResetPassword struct {
	Email string
	Token string
}

func (c *ResetPassword) Execute(db DB) error {
	if c.Email == "" {
		return ErrNotFound
	}

	user, err := db.UserByName(c.Email)
	if err != nil {
		return err
	}

	token := NewTokenizer().Create()
	if err := db.NewToken(user.ID, token, "reset"); err != nil {
		return err
	}
	c.Token = token

	return nil
}

// ChangePassword is a command to change a registered
// user's password. It expects either the ID or the
// Token field to be set, where ID points to a valid user
// id or the token can be used to retrieve such an id.
type ChangePassword struct {
	Token string
	Pass  string
	ID    int
}

func (c *ChangePassword) Execute(db DB) error {
	if !NewValidator().MatchPass(c.Pass) {
		return ErrWeakPass
	}

	if c.Token == "" && c.ID < 1 {
		return ErrNotFound
	}

	if c.Token != "" {
		if tok, err := db.Token(c.Token); err != nil {
			return err
		} else if data, ok := tok.Data.(string); !ok {
			return ErrNotFound
		} else if data != "reset" {
			return ErrNotFound
		} else {
			c.ID = tok.ID
		}
	}

	user, err := db.UserByID(c.ID)
	if err != nil {
		return err
	}

	user.Pass = NewCrypter().Encrypt(c.Pass)
	if err = db.SetUser(user); err == nil {
		if c.Token != "" {
			err = db.DelToken(c.Token)
		}
	}

	return err
}
