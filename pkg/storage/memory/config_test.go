package memory

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
				backupDir:       "/tmp/heyapple/backup",
				storageDir:      "/tmp/heyapple/store",
				storageInterval: time.Minute * 5,
			},
		},
		{ //02// load all vars
			env: map[string]string{
				envStorageDir:      "/path/to/store",
				envBackupDir:       "/backup/is/here",
				envStorageInterval: "1h25m",
			},
			cfg: config{
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
