package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/gen2brain/beeep"
)

type Config struct {
	BackupDirectory string
	Distributions   []string
}

func NewConfig() (*Config, error) {
	backupDir := os.Getenv("WSL_BACKUP_DIR")
	if backupDir == "" {
		beeep.Alert("WSL Snappy Alert", "Environment variable WSL_BACKUP_DIR is not set.", "")
		return nil, fmt.Errorf("environment variable WSL_BACKUP_DIR is not set")
	}

	distributions := os.Getenv("WSL_DISTRIBUTIONS")
	if distributions == "" {
		beeep.Alert("WSL Snappy Alert", "Environment variable WSL_DISTRIBUTIONS is not set.", "")
		return nil, fmt.Errorf("environment variable WSL_DISTRIBUTIONS is not set")
	}

	return &Config{
		BackupDirectory: backupDir,
		Distributions:   strings.Split(distributions, ","),
	}, nil
}
