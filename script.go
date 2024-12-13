package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Define the Ansible playbook command
	cmd := exec.Command("ansible-playbook", "deploy.yml")

	// Run the Ansible playbook
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ansible playbook executed successfully!")
}
