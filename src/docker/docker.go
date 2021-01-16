package docker

import (
	"bytes"
	"fmt"
	"os/exec"
)

type execContext = func(name string, arg ...string) *exec.Cmd

// Prune frees your Docker instance of all unused containers, networks, images (dangling and unreferenced) and optionally, volumes
func Prune(cmdContext execContext) (*bytes.Buffer, error) {
	// TODO: Allow the user to pass in options such as "-a" or "--volumes"
	cmd := cmdContext("/bin/sh", "-c", "docker system prune -f")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to Prune Docker: %s", err))
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Docker instance pruned!\n%s", outb.String()))
	return &outb, nil
}
