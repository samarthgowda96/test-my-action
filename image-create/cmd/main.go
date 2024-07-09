package main

import (
	"log"
	"os/exec"
)

func buildDockerImage() {
	cmd := exec.Command("docker", "build", "-t", "ytest", "-f", "tes", ".")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Docker build failed with error: %v", err)
	}
}
