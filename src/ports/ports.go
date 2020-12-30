package ports

import (
	"fmt"
	"os"
	"os/exec"
)

// FreePort frees a port of its process
func FreePort() {
	port := os.Args[2]
	pid := fmt.Sprintf("$(lsof -t -i:%s)", port)
	err := exec.Command("/bin/sh", "-c", fmt.Sprintf("kill %s", pid)).Run()
	if err != nil {
		fmt.Printf("%s\n", fmt.Sprintf("Cannot free port %s.", port))
		return
	}
	fmt.Printf("%s\n", fmt.Sprintf("Port %s freed!", port))
}
