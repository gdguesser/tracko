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
		activeWindowTitle, err := getActiveApplicationName()
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

func getActiveApplicationName() (string, error) {
	script := `tell application "System Events" to get name of first application process whose frontmost is true`
	cmd := exec.Command("osascript", "-e", script)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
