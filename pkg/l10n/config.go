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

package l10n

import (
	"os"
	"strconv"
)

const (
	envDebug   = "HEYAPPLE_DEBUG_MODE"
	envDefault = "HEYAPPLE_LANG_DEFAULT"
)

type config struct {
	defaultLang string
	debugMode   bool
}

func getConfig() config {
	cfg := config{
		debugMode:   false,
		defaultLang: "en",
	}

	if lang := os.Getenv(envDefault); lang != "" {
		cfg.defaultLang = lang
	}
	if debug, err := strconv.ParseBool(os.Getenv(envDebug)); err == nil {
		cfg.debugMode = debug
	}

	return cfg
}
