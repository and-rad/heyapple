package memory

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
				storageDir: "/tmp/heyapple/store",
				backupDir:  "/tmp/heyapple/backup",
			},
		},
		{ //01// ignore other vars
			env: map[string]string{
				"PATH": "/usr/bin",
			},
			cfg: config{
				storageDir: "/tmp/heyapple/store",
				backupDir:  "/tmp/heyapple/backup",
			},
		},
		{ //02// load all vars
			env: map[string]string{
				envStorageDir: "/path/to/store",
				envBackupDir:  "/backup/is/here",
			},
			cfg: config{
				storageDir: "/path/to/store",
				backupDir:  "/backup/is/here",
			},
		},
	} {
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
