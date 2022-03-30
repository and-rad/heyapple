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

package mem

import (
	"os"
	"testing"
	"time"
)

func Test_getConfig(t *testing.T) {
	for idx, data := range []struct {
		env map[string]string
		cfg config
	}{
		{ //00// empty environment, default values
			cfg: config{
				defaultLang:     "en",
				backupDir:       "/tmp/heyapple/backup",
				storageDir:      "/tmp/heyapple/store",
				storageInterval: time.Minute * 15,
			},
		},
		{ //01// ignore other vars
			env: map[string]string{
				"PATH":             "/usr/bin",
				envStorageInterval: "5m",
			},
			cfg: config{
				defaultLang:     "en",
				backupDir:       "/tmp/heyapple/backup",
				storageDir:      "/tmp/heyapple/store",
				storageInterval: time.Minute * 5,
			},
		},
		{ //02// load all vars
			env: map[string]string{
				envDefaultLang:     "fr",
				envStorageDir:      "/path/to/store",
				envBackupDir:       "/backup/is/here",
				envStorageInterval: "1h25m",
			},
			cfg: config{
				defaultLang:     "fr",
				storageDir:      "/path/to/store",
				backupDir:       "/backup/is/here",
				storageInterval: time.Hour + time.Minute*25,
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
