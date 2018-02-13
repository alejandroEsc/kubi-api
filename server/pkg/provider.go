package server

import (
	"fmt"
	"log"
	"os/exec"

	api "github.com/alejandroEsc/cluster-apis/api"
)

// ProviderOptions to pass on to a cluster provider
type ProviderOptions struct {
	Name              string
	AutoFetchProvider bool
	StorePath         string
}

func getProvider(p ProviderOptions, c clusterOptions) (Provider, error) {
	switch p.Name {
	case "kubicorn":
		log.Println("provider: kubicorn")
		return newKubicornProvider(p, c), nil
	case "kubicorn_cli":
		log.Println("provider: kubicorn_cli")
		return newKubicornProviderCLI(p, c), nil
	default:
		return nil, fmt.Errorf("could not find provider for %s", p.Name)
	}
}

// Provider is cluster provider object
type Provider interface {
	apply() (*api.ClusterStatusMsg, error)
	create() (*api.ClusterStatusMsg, error)
	delete() (*api.ClusterStatusMsg, error)
}

func runCommandPrintOutput(cmdS string) error {
	log.Printf("attempting to run command: %s ...", cmdS)

	cmd := exec.Command("sh", "-c", cmdS)
	cout, err := cmd.CombinedOutput()

	log.Print(string(cout))

	if err != nil {
		log.Printf("...found error attempting command: %s", err)
	}

	return err
}
