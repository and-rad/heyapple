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

// Package app contains the core functionality of the aplication.
//
// This package makes heavy use of the Command pattern: Functions are
// encapsulated in individual small structs that do exactly one thing and
// nothing else. The Command and Query types are both commands in this
// sense, their names reflecting whether they perform read (query) or
// write (command) actions.
package app

import (
	"errors"
	"heyapple/pkg/core"
)

var (
	ErrCredentials = errors.New("nomatch")
	ErrExists      = errors.New("exists")
	ErrNotFound    = errors.New("notfound")
)

// DB provides access to persistent storage.
type DB interface {
	Execute(Command) error
	Fetch(Query) error

	NewUser(name string, hash string) (int, error)
	UserByName(name string) (User, error)

	Food(id int) (core.Food, error)
	Foods() ([]core.Food, error)
	NewFood() (int, error)
	SetFood(core.Food) error
}

// A Command encapsulates a single action that changes the
// underlying data. It can carry input and output parameters.
type Command interface {
	Execute(db DB) error
}

// A Query encapsulates a single read action on the underlying
// data. It should not make any changes to the data. It can
// carry input and output parameters.
type Query interface {
	Fetch(db DB) error
}
