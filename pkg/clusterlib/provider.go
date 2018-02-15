package clusterlib

import (
	api "github.com/alejandroEsc/kubicorn-example-server/api"
)

// ProviderOptions to pass on to a cluster provider
type ProviderOptions struct {
	Name              string
	AutoFetchProvider bool
	StorePath         string
}

// Provider is cluster provider object
type Provider interface {
	Apply() (*api.ClusterStatusMsg, error)
	Create() (*api.ClusterStatusMsg, error)
	Delete() (*api.ClusterStatusMsg, error)
}
