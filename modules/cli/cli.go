package cli

import (
	"log"
	"os/exec"
)

type CLI struct {
	OutFile string
}

// Exec executes shell command given as parameter
func Exec(cli* CLI, cmdStr string, arg string, wait bool) {
	cmd := exec.Command(cmdStr, arg)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	if wait == true {
		err = cmd.Wait()
	}
	log.Printf("Command finished with error: %v", err)
}
