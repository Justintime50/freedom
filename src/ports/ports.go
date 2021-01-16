package ports

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

type execContext = func(name string, arg ...string) *exec.Cmd

// Kill frees a port of its process
func Kill(cmdContext execContext, portNumber int) (*bytes.Buffer, error) {
	port := strconv.Itoa(portNumber)
	pid := fmt.Sprintf("$(lsof -t -i:%s)", port)

	cmd := cmdContext("/bin/sh", "-c", fmt.Sprintf("kill %s", pid))
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("Could not free port %s.", port))
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Port %s freed!", port))
	return &outb, nil
}
