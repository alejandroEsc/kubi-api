package server

import (
	api "github.com/alejandroEsc/cluster-apis/api"
)

type clusterOptions struct {
	Name              string
	CloudProviderName string
	CloudID           string
}

type clusterStatus struct {
	Code int64
	Msg  string
}

var (
	// find another name for initial uncreated state. (zero)
	unCreated = clusterStatus{0, "uncreated"}
	created   = clusterStatus{1, "created"}
	applied   = clusterStatus{2, "applied"}
	deleted   = clusterStatus{3, "deleted"}
)

func (c *clusterStatus) createClusterStatusMsg() *api.ClusterStatusMsg {
	return &api.ClusterStatusMsg{Status: c.Msg, Code: c.Code}
}

func parseClusterDefinition(a *api.ClusterDefinition) (clusterOptions, ProviderOptions) {
	clusterOpts := clusterOptions{
		Name:              a.ClusterConfigs.Name,
		CloudProviderName: a.ClusterConfigs.CloudProviderName,
		CloudID:           a.CloudID,
	}

	providerOpts := ProviderOptions{
		Name:              a.ClusterProvider,
		AutoFetchProvider: a.AutoFetchClusterProvider,
		StorePath:         a.ProviderStorePath,
	}

	return clusterOpts, providerOpts
}
