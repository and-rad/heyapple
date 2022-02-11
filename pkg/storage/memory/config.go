package memory

import (
	"os"
	"time"
)

const (
	envBackupDir       = "HEYAPPLE_BACKUP_DIR"
	envStorageDir      = "HEYAPPLE_STORAGE_DIR"
	envStorageInterval = "HEYAPPLE_STORAGE_INTERVAL"
)

type config struct {
	backupDir       string
	storageDir      string
	storageInterval time.Duration
}

func getConfig() config {
	cfg := config{
		backupDir:       "/tmp/heyapple/backup",
		storageDir:      "/tmp/heyapple/store",
		storageInterval: time.Minute * 15,
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
