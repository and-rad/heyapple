package app

import (
	"os"
	"strconv"
)

const (
	envEncryptCost = "HEYAPPLE_ENCRYPT_COST"
)

type config struct {
	encryptCost int
}

func getConfig() config {
	cfg := config{
		encryptCost: 10,
	}

	if cost, err := strconv.Atoi((os.Getenv(envEncryptCost))); err == nil {
		cfg.encryptCost = cost
	}

	return cfg
}
