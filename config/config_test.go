package config

import (
	"os"
	"testing"
)

func TestNewConfigWithValidEnvVars(t *testing.T) {
	// Set valid environment variables for testing
	os.Setenv("WSL_BACKUP_DIR", "/path/to/backup")
	os.Setenv("WSL_DISTRIBUTIONS", "Debian,Kali,Ubuntu")

	config, err := NewConfig()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedBackupDir := "/path/to/backup"
	if config.BackupDirectory != expectedBackupDir {
		t.Errorf("Expected BackupDirectory to be %s, but got %s", expectedBackupDir, config.BackupDirectory)
	}

	expectedDistributions := []string{"Debian", "Kali", "Ubuntu"}
	for i, dist := range expectedDistributions {
		if config.Distributions[i] != dist {
			t.Errorf("Expected Distributions[%d] to be %s, but got %s", i, dist, config.Distributions[i])
		}
	}
}

func TestNewConfigWithMissingBackupDirEnvVar(t *testing.T) {
	// Unset the environment variable
	os.Unsetenv("WSL_BACKUP_DIR")

	_, err := NewConfig()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	expectedErrorMessage := "environment variable WSL_BACKUP_DIR is not set"
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message to be %s, but got %s", expectedErrorMessage, err.Error())
	}
}

func TestNewConfigWithMissingDistributionsEnvVar(t *testing.T) {
	// Set valid backup directory but unset distributions
	os.Setenv("WSL_BACKUP_DIR", "/path/to/backup")
	os.Unsetenv("WSL_DISTRIBUTIONS")

	_, err := NewConfig()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	expectedErrorMessage := "environment variable WSL_DISTRIBUTIONS is not set"
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message to be %s, but got %s", expectedErrorMessage, err.Error())
	}
}
