package main

import (
	"flag"
	"fmt"
	"github.com/Justintime50/freedom/src/docker"
	"github.com/Justintime50/freedom/src/macos"
	"github.com/Justintime50/freedom/src/ports"
	"os"
	"os/exec"
)

func main() {
	// main function that accepts CLI args to invoke different functionality
	closeMacosFinderWindows := flag.Bool("finder", false, "Close all your macOS Finder Windows.")
	pruneDocker := flag.Bool("docker", false, "Prune your Docker instance.")
	killPort := flag.Int("port", 0, "Free a port of its process (pass a port number as an argument).")
	flag.Parse()

	if *closeMacosFinderWindows {
		macos.FreeFinderWindows(exec.Command)
		return
	} else if *pruneDocker {
		docker.Prune(exec.Command)
		return
	} else if *killPort != 0 {
		ports.Kill(exec.Command, *killPort)
		return
	}

	fmt.Println("No action taken as no flag was passed.")
	os.Exit(1)
}
