package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
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
			lastWindowName = activeWindowName
			lastSwitchTime = time.Now()
			fmt.Printf("Active window: %s\n", activeWindowName)
		} else {
			// If it's the same window, calculate the time spent on it
			duration := time.Since(lastSwitchTime)
			fmt.Printf("Time spent on %s: %s\n", activeWindowName, duration)
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
