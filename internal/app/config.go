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

import (
	"os"
	"strconv"
)

const (
	envEncryptCost = "HEYAPPLE_ENCRYPT_COST"
)

type config struct {
	encryptCost int
}

func getConfig() config {
	cfg := config{
		encryptCost: 10,
	}

	if cost, err := strconv.Atoi((os.Getenv(envEncryptCost))); err == nil {
		cfg.encryptCost = cost
	}

	return cfg
}
