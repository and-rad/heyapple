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

package app

// User represents a human user of the system, identified
// by their e-mail address or id. The password is always
// stored as an encrypted hash.
type User struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
	ID    int    `json:"id"`
}

// Authenticate is a query that authenticates a user by
// checking against e-mail and password. If successful,
// the user's id is stored in the command.
type Authenticate struct {
	Email string
	Pass  string
	ID    int
}

func (c *Authenticate) Fetch(db DB) error {
	if user, err := db.UserByName(c.Email); err != nil {
		return err
	} else if !NewCrypter().Match(user.Pass, c.Pass) {
		return ErrCredentials
	} else {
		c.ID = user.ID
	}

	return nil
}
