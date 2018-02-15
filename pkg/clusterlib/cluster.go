package clusterlib

import (
	api "github.com/alejandroEsc/kubicorn-example-server/api"
)

type ClusterOptions struct {
	Name              string
	CloudProviderName string
	CloudID           string
}

type ClusterStatus struct {
	Code int64
	Msg  string
}

var (
	// find another name for initial uncreated state. (zero)
	Planned = ClusterStatus{0, "planned"}
	Created = ClusterStatus{1, "created"}
	Applied = ClusterStatus{2, "applied"}
	Deleted = ClusterStatus{3, "deleted"}
)

func (c *ClusterStatus) CreateClusterStatusMsg() *api.ClusterStatusMsg {
	return &api.ClusterStatusMsg{Status: c.Msg, Code: c.Code}
}

func ParseClusterDefinition(a *api.ClusterDefinition) (ClusterOptions, ProviderOptions) {
	clusterOpts := ClusterOptions{
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
