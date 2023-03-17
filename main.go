package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	var lastWindowTitle string
	for {
		activeWindowTitle, err := getActiveWindowTitle()
		if err != nil {
			log.Fatal(err)
		}

		if activeWindowTitle != lastWindowTitle {
			lastWindowTitle = activeWindowTitle
			fmt.Printf("Active window: %s\n", activeWindowTitle)
		}

		time.Sleep(1 * time.Second)
	}
}

func getActiveWindowTitle() (string, error) {
	cmd := exec.Command("osascript", "-e", `tell application "System Events" to tell process (name of first application process whose frontmost is true) to get name of window 1`)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	title := strings.TrimSpace(string(output))
	return title, nil
}
