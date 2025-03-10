package utils

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// GetUptime returns the system uptime.
func GetUptime() (string, error) {
	output, err := exec.Command("uptime", "-p").Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// GetLoadAverage returns the system load average.
func GetLoadAverage() (string, error) {
	output, err := exec.Command("cat", "/proc/loadavg").Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// CheckDockerContainers returns the status of Docker containers.
func CheckDockerContainers() (string, error) {
	output, err := exec.Command("docker", "ps", "--format", "{{.Names}}: {{.Status}}").Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// AnalyzeLogs checks for errors in the specified log file.
func AnalyzeLogs(logFile string) ([]string, error) {
	file, err := os.Open(logFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var errors []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isErrorLine(line) {
			errors = append(errors, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return errors, nil
}

// isErrorLine checks if a log line indicates an error.
func isErrorLine(line string) bool {
	return strings.Contains(line, "ERROR") || strings.Contains(line, "FATAL")
}

// MonitorSystem performs a comprehensive system performance check.
func MonitorSystem() {
	uptime, err := GetUptime()
	if err != nil {
		fmt.Println("Error getting uptime:", err)
		return
	}
	fmt.Println("System Uptime:", uptime)

	loadAvg, err := GetLoadAverage()
	if err != nil {
		fmt.Println("Error getting load average:", err)
		return
	}
	fmt.Println("Load Average:", loadAvg)

	dockerStatus, err := CheckDockerContainers()
	if err != nil {
		fmt.Println("Error checking Docker containers:", err)
		return
	}
	fmt.Println("Docker Container Status:\n", dockerStatus)

	logErrors, err := AnalyzeLogs("/var/log/syslog")
	if err != nil {
		fmt.Println("Error analyzing logs:", err)
		return
	}
	if len(logErrors) > 0 {
		fmt.Println("Errors found in logs:")
		for _, errLine := range logErrors {
			fmt.Println(errLine)
		}
	} else {
		fmt.Println("No errors found in logs.")
	}
}