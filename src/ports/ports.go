package ports

import (
	"fmt"
	"os"
	"os/exec"
)

// Free a port of its process
func ports() {
	port := os.Args[1]
	pid := fmt.Sprintf("$(lsof -t -i:%s)", port)
	err := exec.Command("/bin/sh", "-c", fmt.Sprintf("kill %s", pid)).Run()
	if err != nil {
		fmt.Printf("%s\n", fmt.Sprintf("Cannot free port %s.", os.Args[1]))
		return
	}
	fmt.Printf("%s\n", fmt.Sprintf("Port %s freed!", port))
}
