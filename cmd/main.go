package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
	"wslSnappy/backup"
	"wslSnappy/config"

	"github.com/gen2brain/beeep"
)

func main() {
	// Parse command-line flags
	rocketMode := flag.Bool("rocket", false, "Run backups in parallel using goroutines")

	flag.Parse()

	start := time.Now()
	beeep.Alert("WSLSnappy", "Backup Started", "")

	// Load configuration settings
	cfg, err := config.NewConfig()
	if err != nil {
		os.Exit(1)
	}

	// Create a channel to receive backup results
	resultsChan := make(chan backup.BackupResult, len(cfg.Distributions))
	var wg sync.WaitGroup

	// Start backup routines for each distribution
	for _, distributionName := range cfg.Distributions {
		if *rocketMode {
			wg.Add(1)
			go func(name string) {
				defer wg.Done()
				result := backup.RunBackup(name, cfg.BackupDirectory)
				resultsChan <- result
			}(distributionName)
		} else {
			result := backup.RunBackup(distributionName, cfg.BackupDirectory)
			resultsChan <- result
		}
	}

	// If in rocket mode, wait for goroutines to complete
	if *rocketMode {
		wg.Wait()
	}
	close(resultsChan)
	elapsed := time.Since(start)

	// Process backup results
	var successMsg, errorMsg string
	for result := range resultsChan {
		if result.Success {
			successMsg += result.DistributionName + " backup succeeded\n"
		} else {
			errorMsg += result.DistributionName + ": backup failures" + result.ErrorMessage + "\n"
		}
	}

	// If there were successful backups, display the total time taken
	if successMsg != "" {
		timeTaken := fmt.Sprintf("Total time taken for backup: %s", elapsed)

		// Print the time and send a notification
		fmt.Println(timeTaken)
		beeep.Notify("WSLSnappy", successMsg+"\n"+timeTaken, "")
	}

	// If there were backup failures, display an alert with error messages
	if errorMsg != "" {
		beeep.Alert("WSLSnappy", errorMsg, "")
	}
}
