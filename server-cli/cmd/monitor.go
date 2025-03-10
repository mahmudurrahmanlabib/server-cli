package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/flywp/server-cli/internal/utils"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor system performance and service health",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			systemStats, err := utils.GetSystemStats()
			if err != nil {
				color.Red("Error retrieving system stats: %v\n", err)
				os.Exit(1)
			}

			containerStats, err := utils.GetDockerContainerStats()
			if err != nil {
				color.Red("Error retrieving Docker container stats: %v\n", err)
				os.Exit(1)
			}

			serviceHealth, err := utils.CheckServiceHealth()
			if err != nil {
				color.Red("Error checking service health: %v\n", err)
				os.Exit(1)
			}

			uptime, loadAvg, err := utils.GetUptimeAndLoadAverage()
			if err != nil {
				color.Red("Error retrieving uptime and load average: %v\n", err)
				os.Exit(1)
			}

			logErrors, err := utils.AnalyzeLogsForErrors()
			if err != nil {
				color.Red("Error analyzing logs: %v\n", err)
				os.Exit(1)
			}

			// Display the collected statistics
			fmt.Printf("System Stats: %+v\n", systemStats)
			fmt.Printf("Docker Container Stats: %+v\n", containerStats)
			fmt.Printf("Service Health: %+v\n", serviceHealth)
			fmt.Printf("Uptime: %s, Load Average: %v\n", uptime, loadAvg)
			fmt.Printf("Log Errors: %+v\n", logErrors)

			// Wait for a specified interval before the next check
			time.Sleep(10 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
}