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

// Package handler defines handlers for static web content that conform
// to the http.Handler and the httprouter.Handle interfaces.
package handler

import (
	"github.com/and-rad/heyapple/internal/app"

	"github.com/and-rad/scs/v2"
)

// Environment is a collection of interfaces that http
// handlers can use to access various parts of the app.
type Environment struct {
	Log     app.Logger
	DB      app.DB
	Msg     app.Notifier
	L10n    app.Translator
	Session *scs.SessionManager
}
