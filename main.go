package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	// Ottieni la directory di backup dalla variabile d'ambiente
	backupDirectory := os.Getenv("WSL_BACKUP_DIR")
	if backupDirectory == "" {
		fmt.Println("Environment variable WSL_BACKUP_DIR is not set.")
		os.Exit(1)
	}

	// Ottieni la lista delle distribuzioni da una variabile d'ambiente
	distributions := os.Getenv("WSL_DISTRIBUTIONS")
	if distributions == "" {
		fmt.Println("Environment variable WSL_DISTRIBUTIONS is not set.")
		os.Exit(1)
	}
	distributionList := strings.Split(distributions, ",")

	var wg sync.WaitGroup

	for _, distributionName := range distributionList {
		wg.Add(1)
		go func(distributionName string) {
			defer wg.Done()

			currentDate := time.Now()
			backupFileName := fmt.Sprintf("%s_%s.tar", distributionName, currentDate.Format("20060102"))
			backupFilePath := backupDirectory + "\\" + backupFileName

			cmd := exec.Command("wsl", "--export", distributionName, backupFilePath)
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				beeep.Alert("Backup Failed", fmt.Sprintf("Error during the export of the distribution %s: %v\n", distributionName, err), "")
				return
			}

			fmt.Printf("Backup of the distribution %s created in %s.\n", distributionName, backupFilePath)
			beeep.Notify("Backup Completed", fmt.Sprintf("The backup of %s has been created in %s.", distributionName, backupFilePath), "")
		}(distributionName)
	}

	wg.Wait()
}
