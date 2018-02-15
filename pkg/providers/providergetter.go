package providers

import (
    "fmt"
    "log"
    pr "github.com/alejandroEsc/kubicorn-example-server/pkg/clusterlib"
)

func GetProvider(p pr.ProviderOptions, c pr.ClusterOptions) (pr.Provider, error) {
    switch p.Name {
    case "kubicorn":
        log.Println("provider: kubicorn")
        return NewKubicornProvider(p, c), nil
    case "kubicorn_cli":
        log.Println("provider: kubicorn_cli")
        return NewKubicornProviderCLI(p, c), nil
    default:
        return nil, fmt.Errorf("could not find provider for %s", p.Name)
    }
}