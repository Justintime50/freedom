package macos

import (
	"bytes"
	"fmt"
	"github.com/justintime50/mockcmd/mockcmd"
)

// FreeFinderWindows frees your Mac of all Finder windows
func FreeFinderWindows(cmdContext mockcmd.ExecContext) (*bytes.Buffer, error) {
	cmd := cmdContext("osascript", "-e", `tell application "Finer" to close windows`)
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to close all finder windows: %s\n", err)
		return nil, err
	}

	fmt.Println("macOS Finder windows closed!")
	return &outb, nil
}
