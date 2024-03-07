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
	validator := NewValidator()
	if !validator.MatchPass(c.Pass) {
		return ErrWeakPass
	}

	if !validator.MatchEmail(c.Email) {
		return ErrNoEmail
	}

	// Doing this early comes with a performance penalty
	// that is actually desirable here since it helps
	// making user enumeration attacks a little harder.
	hash := NewCrypter().Encrypt(c.Pass)
	token := NewTokenizer().Create()

	_, err := db.UserByName(c.Email)
	if err == nil {
		return ErrExists
	}
	if err != ErrNotFound {
		return err
	}

	id, err := db.NewUser(c.Email, hash, token)
	if err != nil {
		return err
	}

	c.ID = id
	c.Pass = ""
	c.Token = token

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
	user, err := db.UserByName(q.Email)
	if err != nil {
		return err
	}

	if !NewCrypter().Match(user.Pass, q.Pass) {
		return ErrCredentials
	}

	if user.Perm&PermLogin != PermLogin {
		return ErrCredentials
	}

	q.ID = user.ID
	q.Perm = user.Perm
	q.Lang = user.Lang
	q.Pass = ""

	return nil
}

// Authorize is a query that issues a challenge to pass
// in order to make sure a user identified by ID has the
// right to perform a given action. It is intentionally
// not tied to a specific action or resource. The type
// of challenge is expressed through the fields of the
// query.
//
// If successful, the Ok field will be set to true. There
// is no error if the challenge itself fails, so you
// still have to check Ok to ultimately know the final
// result.
//
// Supported types of challenges are:
//   - Password: Make sure the user is still the same
//     user by asking for the password again. This type
//     of challenge is often issued before dangerous
//     actions or those with far-reaching consequences.
type Authorize struct {
	ID   int
	Pass string
	Ok   bool
}

func (q *Authorize) Fetch(db DB) error {
	if q.ID == 0 {
		return ErrNotFound
	}

	user, err := db.UserByID(q.ID)
	if err != nil {

		return err
	}

	q.Ok = NewCrypter().Match(user.Pass, q.Pass)
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
		tok, err := db.Token(c.Token)
		if err != nil {
			return err
		}

		data, ok := tok.Data.(string)
		if !ok {
			return ErrNotFound
		}

		if data != "reset" {
			return ErrNotFound
		}

		c.ID = tok.ID
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
