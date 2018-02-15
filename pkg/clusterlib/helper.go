package clusterlib

import (
	"log"
	"os/exec"
)

// RunCommandPrintOutput is a tool to consistently run command strings.
func RunCommandPrintOutput(cmdS string) error {
	log.Printf("attempting to run command: %s ...", cmdS)

	cmd := exec.Command("sh", "-c", cmdS)
	cout, err := cmd.CombinedOutput()

	log.Print(string(cout))

	if err != nil {
		log.Printf("...found error attempting command: %s", err)
	}

	return err
}
