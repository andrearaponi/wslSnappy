package backup

import (
	"testing"
)

func TestRunBackupSuccess(t *testing.T) {
	// Replace with a valid distribution name and backup directory
	distributionName := "Debian"
	backupDirectory := "/path/to/backup"

	result := RunBackup(distributionName, backupDirectory)

	if !result.Success {
		t.Errorf("Expected success, but got failure. Error message: %s", result.ErrorMessage)
	}
	if result.DistributionName != distributionName {
		t.Errorf("Expected DistributionName to be %s, but got %s", distributionName, result.DistributionName)
	}
	if result.ErrorMessage != "" {
		t.Errorf("Expected ErrorMessage to be empty, but got %s", result.ErrorMessage)
	}
}

func TestRunBackupFailure(t *testing.T) {
	// Replace with an invalid distribution name and backup directory
	distributionName := "InvalidDistribution"
	backupDirectory := "/path/to/backup"

	result := RunBackup(distributionName, backupDirectory)

	if result.Success {
		t.Errorf("Expected failure, but got success")
	}
	if result.DistributionName != distributionName {
		t.Errorf("Expected DistributionName to be %s, but got %s", distributionName, result.DistributionName)
	}
	if result.ErrorMessage == "" {
		t.Errorf("Expected ErrorMessage to be non-empty, but got empty")
	}
}

func TestRunBackupWithEmptyDistributionName(t *testing.T) {
	// Replace with an empty distribution name and valid backup directory
	distributionName := ""
	backupDirectory := "/path/to/backup"

	result := RunBackup(distributionName, backupDirectory)

	if result.Success {
		t.Errorf("Expected failure, but got success")
	}
	if result.DistributionName != distributionName {
		t.Errorf("Expected DistributionName to be empty, but got %s", result.DistributionName)
	}
	if result.ErrorMessage == "" {
		t.Errorf("Expected ErrorMessage to be non-empty, but got empty")
	}
}
