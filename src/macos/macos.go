package macos

import (
	"fmt"
	"os/exec"
)

// Free your Mac of all Finder windows
func macos() {
	err := exec.Command("osascript", "-e", `tell application "Finder" to close windows`).Run()
	if err != nil {
		fmt.Printf("%s\n", "Failed to close all finder windows.")
		return
	}
	fmt.Printf("%s\n", "macOS Finder windows closed!")
}
