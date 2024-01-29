package backup

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type BackupResult struct {
	DistributionName string
	Success          bool
	ErrorMessage     string
}

func RunBackup(distributionName, backupDirectory string) BackupResult {

	currentDate := time.Now()
	backupFileName := fmt.Sprintf("%s_%s.tar", distributionName, currentDate.Format("20060102"))
	backupFilePath := backupDirectory + "\\" + backupFileName

	cmd := exec.Command("wsl", "--export", distributionName, backupFilePath)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	var result BackupResult
	result.DistributionName = distributionName

	if err != nil {
		result.Success = false
		result.ErrorMessage = fmt.Sprintf("%v", err)
	} else {
		result.Success = true
	}

	return result
}
