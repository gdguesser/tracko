package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// Map to store the time spent on each window
	windowTimes := make(map[string]time.Duration)

	var lastWindowName string
	var lastSwitchTime time.Time

	// Create a CSV file to store the window activity data
	file, err := os.Create("window_activity.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header row to the CSV file
	err = writer.Write([]string{"Application Name", "Time Spent"})
	if err != nil {
		panic(err)
	}

	// Continuously update the CSV file with the latest data
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
				fmt.Printf("Time spent on %s: %s\n", lastWindowName, windowTimes[lastWindowName])
			}
			lastWindowName = activeWindowName
			lastSwitchTime = time.Now()
			fmt.Printf("Active window: %s\n", activeWindowName)
		}

		// Write the window activity data to the CSV file every 5 minutes
		if time.Now().Sub(lastSwitchTime) >= 5*time.Minute {
			// Clear the CSV file before writing the latest data to avoid duplicates
			file.Truncate(0)
			file.Seek(0, 0)

			// Write the header row to the CSV file
			err = writer.Write([]string{"Application Name", "Time Spent"})
			if err != nil {
				panic(err)
			}

			for name, duration := range windowTimes {
				record := []string{name, formatDuration(duration)}
				err = writer.Write(record)
				if err != nil {
					panic(err)
				}
			}

			writer.Flush()

			fmt.Println("Window activity data written to CSV file")
			lastSwitchTime = time.Now()
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
