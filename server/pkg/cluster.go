package server

import (
	api "github.com/alejandroEsc/cluster-apis/api"
)

type ClusterOptions struct {
	Name              string
	CloudProviderName string
    CloudId           string
}

type ClusterStatus struct {
	Code int64
	Msg  string
}

var (
	// find another name for initial uncreated state. (zero)
	UnCreated = ClusterStatus{0, "uncreated"}
	Created = ClusterStatus{1, "created"}
	Applied = ClusterStatus{2, "applied"}
	Deleted = ClusterStatus{3, "deleted"}
)


func (c * ClusterStatus) createClusterStatusMsg() *api.ClusterStatusMsg {
	return &api.ClusterStatusMsg{Status: c.Msg, Code: c.Code}
}

func parseClusterDefinition(a *api.ClusterDefinition) (error, ClusterOptions, ProviderOptions) {
	var providerOpts ProviderOptions
	var error error
	var clusterOpts ClusterOptions

	clusterOpts = ClusterOptions {
		Name:              a.ClusterConfigs.Name,
		CloudProviderName: a.ClusterConfigs.CloudProviderName,
        CloudId:           a.CloudId,
	}

	providerOpts = ProviderOptions {
		Name:              a.ClusterProvider,
		AutoFetchProvider: a.AutoFetchClusterProvider,
        StorePath:         a.ProviderStorePath,
	}

	return error, clusterOpts, providerOpts
}
