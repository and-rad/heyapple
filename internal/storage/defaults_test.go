package storage_test

import (
	"embed"
	"os"
	"testing"

	"github.com/and-rad/heyapple/internal/data/dev"
	"github.com/and-rad/heyapple/internal/data/prod"
	"github.com/and-rad/heyapple/internal/storage"
)

func TestDefaults(t *testing.T) {
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

		if fs := storage.Defaults(); fs != data.fs {
			t.Errorf("test case %d: files mismatch \nhave: %v \nwant: %v", idx, fs, data.fs)
		}
	}
}
