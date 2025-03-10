package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// FindComposeFile searches for the Docker Compose file in the specified domain.
func FindComposeFile(domain string) string {
	// Implementation to find the Docker Compose file
	// This is a placeholder; actual implementation will depend on the project structure
	return fmt.Sprintf("%s/docker-compose.yml", domain)
}

// ShowNoComposeError displays an error message when no Docker Compose file is found.
func ShowNoComposeError() {
	fmt.Println("No Docker Compose file found. Please ensure you are in the correct directory.")
}

// CheckServers checks the status of servers defined in the Docker Compose file.
func CheckServers(composePath string) error {
	cmd := exec.Command("docker-compose", "-f", composePath, "ps")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to check servers: %v", err)
	}

	if strings.Contains(string(output), "Exit") {
		return fmt.Errorf("some servers are not running")
	}

	return nil
}