package defaults_test

import (
	"embed"
	"os"
	"testing"

	"github.com/and-rad/heyapple/internal/defaults"
	"github.com/and-rad/heyapple/internal/defaults/dev"
	"github.com/and-rad/heyapple/internal/defaults/prod"
)

func TestGet(t *testing.T) {
	for idx, data := range []struct {
		env string
		fs  embed.FS
	}{
		{ //00//
			env: "",
			fs:  prod.FS,
		},
		{ //01//
			env: "production",
			fs:  prod.FS,
		},
		{ //02//
			env: "develop",
			fs:  dev.FS,
		},
		{ //03//
			env: "none",
			fs:  embed.FS{},
		},
	} {
		os.Setenv("HEYAPPLE_STORAGE_PRESET", data.env)
		defer os.Unsetenv("HEYAPPLE_STORAGE_PRESET")

		if fs := defaults.Get(); fs != data.fs {
			t.Errorf("test case %d: files mismatch \nhave: %v \nwant: %v", idx, fs, data.fs)
		}
	}
}
