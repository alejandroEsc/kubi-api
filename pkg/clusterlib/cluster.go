package clusterlib

import (
	api "github.com/alejandroEsc/kubicorn-example-server/api"
)

// ClusterOptions are those that correspond only to the cluster to be managed.
type ClusterOptions struct {
	Name              string
	CloudProviderName string
	CloudID           string
}

// ClusterStatus corresponds to the state of the cluster.
type ClusterStatus struct {
	Code int64
	Msg  string
}

var (
	// Planned is the zero state, as in no cluster or config info exists.
	Planned = ClusterStatus{0, "planned"}
	// Created represents state in which actual config reprenseting cluster now exists.
	Created = ClusterStatus{1, "created"}
	// Applied represents that action to update state has been committed.
	Applied = ClusterStatus{2, "applied"}
	// Deleted represents a deleted/destroyed cluster state.
	Deleted = ClusterStatus{3, "deleted"}
)

// CreateClusterStatusMsg based on internal types convert to grpc ClusterStatusMsg type.
func (c *ClusterStatus) CreateClusterStatusMsg() *api.ClusterStatusMsg {
	return &api.ClusterStatusMsg{Status: c.Msg, Code: c.Code}
}

// ParseClusterDefinition parses server cluster definition into provider options and cluster configs
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
