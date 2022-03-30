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

// Package conv provides functionality to convert to or from various
// external formats. Its main use case lies in converting data from
// third-party food databases into the format used internally by HeyApple.
package conv

import (
	"errors"

	"github.com/and-rad/heyapple/internal/core"
)

var (
	ErrFromJSON = errors.New("fromjson")
)

type Food struct {
	Name  string     `json:"name"`
	Aisle core.Aisle `json:"aisle"`
	core.Food
}
