package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "server-cli",
	Short: "A CLI tool for managing server operations",
}

func init() {
	cobra.OnInitialize()
	rootCmd.AddCommand(checkServersCmd)
	rootCmd.AddCommand(monitorCmd)
}