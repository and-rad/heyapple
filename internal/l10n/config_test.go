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
	"testing"
)

func Test_getConfig(t *testing.T) {
	for idx, data := range []struct {
		env map[string]string
		cfg config
	}{
		{ //00// empty environment, default values
			cfg: config{
				debugMode:   false,
				defaultLang: "en",
				dataDir:     "/tmp/heyapple/data",
			},
		},
		{ //01// ignore other vars
			env: map[string]string{
				"PATH": "/usr/bin",
			},
			cfg: config{
				debugMode:   false,
				defaultLang: "en",
				dataDir:     "/tmp/heyapple/data",
			},
		},
		{ //02// invalid type
			env: map[string]string{
				envDebug: "Maybe",
			},
			cfg: config{
				debugMode:   false,
				defaultLang: "en",
				dataDir:     "/tmp/heyapple/data",
			},
		},
		{ //03// load all vars
			env: map[string]string{
				envDebug:   "true",
				envDefault: "fr",
				envDir:     "/tmp/overrides",
			},
			cfg: config{
				debugMode:   true,
				defaultLang: "fr",
				dataDir:     "/tmp/overrides",
			},
		},
	} {
		os.Clearenv()
		for k, v := range data.env {
			os.Setenv(k, v)
			defer os.Unsetenv(k)
		}

		cfg := getConfig()

		if cfg != data.cfg {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, cfg, data.cfg)
		}
	}
}
