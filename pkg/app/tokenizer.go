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

import (
	"crypto/rand"
	"fmt"
)

// Token is a piece of data used for anonymous authorization.
type Token struct {
	Data interface{} `json:"data,omitempty"`
	ID   int         `json:"id"`
}

// Tokenizer creates cryptographically secure tokens.
type Tokenizer interface {
	Create() string
}

type tokenizer struct{}

// NewTokenizer returns a default implementation of the Tokenizer interface.
func NewTokenizer() Tokenizer {
	return tokenizer{}
}

// Create creates a new 16-bit token hash.
func (t tokenizer) Create() string {
	token := make([]byte, 16)
	rand.Read(token)

	return fmt.Sprintf("%x", token)
}
