package memory

import (
	"os"
)

const (
	envBackupDir  = "HEYAPPLE_BACKUP_DIR"
	envStorageDir = "HEYAPPLE_STORAGE_DIR"
)

type config struct {
	backupDir  string
	storageDir string
}

func getConfig() config {
	cfg := config{
		backupDir:  "/tmp/heyapple/backup",
		storageDir: "/tmp/heyapple/store",
	}

	if dir := os.Getenv(envBackupDir); dir != "" {
		cfg.backupDir = dir
	}
	if dir := os.Getenv(envStorageDir); dir != "" {
		cfg.storageDir = dir
	}

	return cfg
}
