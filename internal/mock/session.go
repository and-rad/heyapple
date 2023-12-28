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

package mock

import (
	"context"
	"errors"

	"github.com/and-rad/scs/v2/memstore"
)

var (
	ErrFailSessionDestroy = errors.New("faildestroy")
)

type SessionStore struct {
	memstore.MemStore
	failDestroy bool
}

func NewSessionStore() *SessionStore {
	return &SessionStore{}
}

func (s *SessionStore) WithFailDestroy() *SessionStore {
	s.failDestroy = true
	return s
}

func (s *SessionStore) DeleteCtx(ctx context.Context, token string) error {
	if s.failDestroy {
		return ErrFailSessionDestroy
	}
	return s.Delete(token)
}
