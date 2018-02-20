package providers

import (
	"fmt"

	pr "github.com/alejandroEsc/kubicorn-example-server/pkg/clusterlib"
	"github.com/juju/loggo"
)

var (
	logger = pr.GetModuleLogger("pkg.providers", loggo.INFO)
)

// GetProvider returns provider object based on the provider option's name
func GetProvider(p pr.ProviderOptions, c pr.ClusterOptions) (pr.Provider, error) {
	switch p.Name {
	case pr.Kubicorn:
		logger.Infof("provider: %s", pr.Kubicorn)
		return NewKubicornProvider(p, c), nil
	case pr.KubicornCLI:
		logger.Infof("provider: %s", pr.KubicornCLI)
		return NewKubicornProviderCLI(p, c), nil
	default:
		return nil, fmt.Errorf("could not find provider for %s", p.Name)
	}
}
