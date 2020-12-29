package docker

import (
	"fmt"
	"os/exec"
)

// Free your Docker instance of all unused containers, networks, images (dangling and unreferenced) and optionally, volumes
func docker() {
	// TODO: Allow the user to pass in options such as "-a" or "--volumes"
	err := exec.Command("/bin/sh", "-c", "docker system prune -f").Run()
	if err != nil {
		fmt.Printf("%s\n", "Failed to Prune Docker.")
		return
	}
	fmt.Printf("%s\n", "Docker instance pruned!")
}
