package docker

import (
	"bytes"
	"fmt"
	"github.com/justintime50/mockcmd/mockcmd"
)

// Prune frees your Docker instance of all unused containers, networks, images (dangling and unreferenced) and optionally, volumes
func Prune(cmdContext mockcmd.ExecContext) (*bytes.Buffer, error) {
	// TODO: Allow the user to pass in options such as "-a" or "--volumes"
	cmd := cmdContext("/bin/sh", "-c", "docker system prune -f")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to Prune Docker: %s\n", err)
		return nil, err
	}

	fmt.Printf("Docker instance pruned!\n%s\n", outb.String())
	return &outb, nil
}
