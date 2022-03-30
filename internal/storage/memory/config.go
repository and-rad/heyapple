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

package memory

import (
	"os"
	"time"
)

const (
	envBackupDir       = "HEYAPPLE_BACKUP_DIR"
	envDefaultLang     = "HEYAPPLE_APP_LANG"
	envStorageDir      = "HEYAPPLE_STORAGE_DIR"
	envStorageInterval = "HEYAPPLE_STORAGE_INTERVAL"
)

type config struct {
	defaultLang     string
	backupDir       string
	storageDir      string
	storageInterval time.Duration
}

func getConfig() config {
	cfg := config{
		defaultLang:     "en",
		backupDir:       "/tmp/heyapple/backup",
		storageDir:      "/tmp/heyapple/store",
		storageInterval: time.Minute * 15,
	}

	if lang := os.Getenv(envDefaultLang); lang != "" {
		cfg.defaultLang = lang
	}
	if dir := os.Getenv(envBackupDir); dir != "" {
		cfg.backupDir = dir
	}
	if dir := os.Getenv(envStorageDir); dir != "" {
		cfg.storageDir = dir
	}

	if val, err := time.ParseDuration(os.Getenv(envStorageInterval)); err == nil {
		cfg.storageInterval = val
	}

	return cfg
}
