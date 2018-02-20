package clusterlib

import (
	"os/exec"
)

// RunCommandPrintOutput is a tool to consistently run command strings.
func RunCommandPrintOutput(cmdS string) error {
	logger.Infof("attempting to run command: %s ...", cmdS)

	cmd := exec.Command("sh", "-c", cmdS)
	cout, err := cmd.CombinedOutput()

	logger.Infof(string(cout))

	if err != nil {
		logger.Infof("...found error attempting command: %s", err)
	}

	return err
}
