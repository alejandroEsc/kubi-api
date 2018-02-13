package server

import (
	"fmt"
    "log"
    "os/exec"
	api "github.com/alejandroEsc/cluster-apis/api"
)

type ProviderOptions struct {
	Name              string
	AutoFetchProvider bool
	StorePath         string
}

func GetProvider(p ProviderOptions, c ClusterOptions) (error, Provider) {
	switch p.Name {
	case "kubicorn":
		log.Println("provider: kubicorn")
		return nil, NewKubicornProvider(p, c)
    case "kubicorn_cli":
		log.Println("provider: kubicorn_cli")
        return nil, NewKubicornProviderCLI(p, c)
	default:
		return fmt.Errorf("could not find provider for %s.", p.Name), nil
	}
}

type Provider interface {
	apply() (*api.ClusterStatusMsg, error)
	create() (*api.ClusterStatusMsg, error)
	delete() (*api.ClusterStatusMsg, error)
}

func runCommandPrintOutput(cmdS string) error {
    log.Printf("attempting to run command: %s ...", cmdS)

    cmd := exec.Command("sh","-c", cmdS)
    cout, err := cmd.CombinedOutput()

    log.Print(string(cout))

    if err != nil {
        log.Printf("...found error attempting command: %s", err)
    }

    return err
}