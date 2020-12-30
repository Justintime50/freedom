package ports

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// FreePort frees a port of its process
func FreePort(portNumber int) {
	port := strconv.Itoa(portNumber)
	pid := fmt.Sprintf("$(lsof -t -i:%s)", port)
	err := exec.Command("/bin/sh", "-c", fmt.Sprintf("kill %s", pid)).Run()
	if err != nil {
		fmt.Printf("%s\n", fmt.Sprintf("Cannot free port %s.", port))
		os.Exit(1)
	}
	fmt.Printf("%s\n", fmt.Sprintf("Port %s freed!", port))
}
