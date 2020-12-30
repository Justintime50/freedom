package main

import (
	"flag"
	"fmt"
	"github.com/Justintime50/freedom/src/docker"
	"github.com/Justintime50/freedom/src/macos"
	"github.com/Justintime50/freedom/src/ports"
)

func main() {
	// main function that accepts CLI args to invoke different functionality
	freeFinder := flag.Bool("finder", false, "Close all your macOS Finder Windows.")
	freeDocker := flag.Bool("docker", false, "Prune your Docker instance.")
	freePort := flag.Int("port", 0, "Free a port of its process (pass a port number as an argument).")
	flag.Parse()

	if *freeFinder {
		macos.FreeFinderWindows()
		return
	} else if *freeDocker {
		docker.PruneDocker()
		return
	} else if *freePort != 0 {
		ports.FreePort(*freePort)
		return
	}

	fmt.Printf("%s\n", "No action taken as no flag was passed.")
}
