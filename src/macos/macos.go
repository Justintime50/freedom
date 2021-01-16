package macos

import (
	"bytes"
	"fmt"
	"os/exec"
)

type execContext = func(name string, arg ...string) *exec.Cmd

// FreeFinderWindows frees your Mac of all Finder windows
func FreeFinderWindows(cmdContext execContext) (*bytes.Buffer, error) {
	cmd := cmdContext("osascript", "-e", `tell application "Finer" to close windows`)
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to close all finder windows: %s", err))
		return nil, err
	}

	fmt.Println("macOS Finder windows closed!")
	return &outb, nil
}
