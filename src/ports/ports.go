package ports

import (
	"bytes"
	"fmt"
	"github.com/justintime50/mockcmd/mockcmd"
	"strconv"
)

// Kill frees a port of its process
func Kill(cmdContext mockcmd.ExecContext, portNumber int) (*bytes.Buffer, error) {
	port := strconv.Itoa(portNumber)
	pid := fmt.Sprintf("$(lsof -t -i:%s)", port)

	cmd := cmdContext("/bin/sh", "-c", fmt.Sprintf("kill %s", pid))
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Could not free port %s.\n", port)
		return nil, err
	}

	fmt.Printf("Port %s freed!\n", port)
	return &outb, nil
}
