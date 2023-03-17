package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// Map to store the time spent on each window
	windowTimes := make(map[string]time.Duration)

	var lastWindowName string
	var lastSwitchTime time.Time

	for {
		// Get the name of the currently active window
		activeWindowName, err := getActiveApplicationName()
		if err != nil {
			panic(err)
		}

		// Check if the active window name is the same as the last one
		if activeWindowName != lastWindowName {
			// If it's a new window, record the current time and window name
			if lastWindowName != "" {
				duration := time.Since(lastSwitchTime)
				windowTimes[lastWindowName] += duration
				fmt.Printf("Time spent on %s: %s\n", lastWindowName, formatDuration(windowTimes[lastWindowName]))
			}
			lastWindowName = activeWindowName
			lastSwitchTime = time.Now()
			fmt.Printf("Active window: %s\n", activeWindowName)
		}

		time.Sleep(1 * time.Second)
	}
}

func getActiveApplicationName() (string, error) {
	script := `tell application "System Events" to get name of first application process whose frontmost is true`
	cmd := exec.Command("osascript", "-e", script)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func formatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
