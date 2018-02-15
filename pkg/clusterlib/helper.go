package clusterlib

import (
    "log"
    "os/exec"
)

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
