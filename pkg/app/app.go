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
	"time"
)

type (
	Notification     int
	NotificationData map[string]interface{}
)

const (
	RegisterNotification Notification = iota + 1
	RenameNotification
	ResetNotification
)

var (
	ErrCredentials = errors.New("nomatch")
	ErrExists      = errors.New("exists")
	ErrMissing     = errors.New("missing")
	ErrNotFound    = errors.New("notfound")
	ErrPermission  = errors.New("permission")
)

// DB provides access to persistent storage.
type DB interface {
	Execute(Command) error
	Fetch(Query) error

	NewUser(name, hash, token string) (int, error)
	SetUser(User) error
	UserByName(name string) (User, error)
	UserByID(id int) (User, error)
	NewToken(id int, hash string, data interface{}) error
	DeleteToken(string) error
	Token(string) (Token, error)

	NewFood() (int, error)
	SetFood(core.Food) error
	Food(id int) (core.Food, error)
	Foods(core.Filter) ([]core.Food, error)
	FoodExists(id int) (bool, error)

	NewRecipe(string) (int, error)
	SetRecipe(core.Recipe) error
	SetRecipeAccess(user, rec, perms int) error
	Recipe(id int) (core.Recipe, error)
	Recipes(uid int, f core.Filter) ([]core.Recipe, error)
	RecipeAccess(user, rec int) (int, error)

	NewDiaryEntries(id int, entries ...core.DiaryEntry) error
	SetDiaryEntries(id int, entries ...core.DiaryEntry) error
	DelDiaryEntries(id int, entries ...core.DiaryEntry) error
	DiaryEntries(id int, date time.Time) ([]core.DiaryEntry, error)
	DiaryDays(id, year, month, day int) ([]core.DiaryDay, error)

	SetShoppingListDone(id int, done map[int]bool) error
	ShoppingList(id int, date ...time.Time) ([]core.ShopItem, error)
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

// Notifier provides functions for sending messages to users
// of the application.
type Notifier interface {
	Send(to string, msg Notification, data NotificationData) error
}

// Translator defines functions for text localization.
type Translator interface {
	Translate(input interface{}, lang string) string
	Get(lang string) map[string]interface{}
	Default() string
}
