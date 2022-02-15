package app

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
				encryptCost: 10,
			},
		},
		{ //01// ignore other vars
			env: map[string]string{
				"PATH": "/usr/bin",
			},
			cfg: config{
				encryptCost: 10,
			},
		},
		{ //02// wrong data type
			env: map[string]string{
				envEncryptCost: "twelve",
			},
			cfg: config{
				encryptCost: 10,
			},
		},
		{ //03// load all fields
			env: map[string]string{
				envEncryptCost: "8",
			},
			cfg: config{
				encryptCost: 8,
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
