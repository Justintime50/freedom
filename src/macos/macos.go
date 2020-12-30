package macos

import (
	"fmt"
	"os"
	"os/exec"
)

// FreeFinderWindows frees your Mac of all Finder windows
func FreeFinderWindows() {
	err := exec.Command("osascript", "-e", `tell application "Finder" to close windows`).Run()
	if err != nil {
		fmt.Printf("%s\n", "Failed to close all finder windows.")
		os.Exit(1)
	}
	fmt.Printf("%s\n", "macOS Finder windows closed!")
}
