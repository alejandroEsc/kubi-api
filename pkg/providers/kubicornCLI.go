package providers

import (
	"fmt"

	"github.com/alejandroEsc/kubicorn-example-server/api"
	cl "github.com/alejandroEsc/kubicorn-example-server/pkg/clusterlib"
)

// KubicornCLI represents a kubicorn provider via library calls
type KubicornCLI struct {
	providerOpts cl.ProviderOptions
	clusterOpts  cl.ClusterOptions
	status       cl.ClusterStatus
}

var (
	errorRevertState = "cluster is in state: %s, we cannot revert to state: %s"
	errorReplayState = "cluster is already in state: %s, cannot replay state"
	awsCloudProvider = "aws"
)

// NewKubicornProviderCLI returns a new kubicornCLI providor
func NewKubicornProviderCLI(p cl.ProviderOptions, c cl.ClusterOptions) *KubicornCLI {
	return &KubicornCLI{providerOpts: p,
		clusterOpts: c,
		status:      cl.Planned}
}

// Apply commits state changes.
func (k *KubicornCLI) Apply() (*clusteror.ClusterStatusMsg, error) {
	if k.status.Code > cl.Applied.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, cl.Applied.Msg)
	}

	if k.status.Code == cl.Applied.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	commandString := fmt.Sprintf("apply %s", k.clusterOpts.Name)
	err := cl.RunCommandPrintOutput(commandString)
	if err != nil {
		return nil, err
	}

	k.status = cl.Applied

	return k.status.CreateClusterStatusMsg(), nil
}

// Create new cluster if it does not exists, meaning creates supporting state files
func (k *KubicornCLI) Create() (*clusteror.ClusterStatusMsg, error) {
	if k.status.Code > cl.Created.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, cl.Created.Msg)
	}

	if k.status.Code == cl.Created.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	commandString := fmt.Sprintf("create %s --profile %s", k.clusterOpts.Name, k.clusterOpts.CloudProviderName)
	err := cl.RunCommandPrintOutput(commandString)
	if err != nil {
		return nil, err
	}

	k.status = cl.Created

	return k.status.CreateClusterStatusMsg(), nil
}

// Delete and destroy cluster and its resources
func (k *KubicornCLI) Delete() (*clusteror.ClusterStatusMsg, error) {
	defer logger.Infof("done")

	if k.status.Code > cl.Deleted.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, cl.Deleted.Msg)
	}

	if k.status.Code == cl.Deleted.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	commandString := fmt.Sprintf("delete %s", k.clusterOpts.Name)
	err := cl.RunCommandPrintOutput(commandString)
	if err != nil {
		return nil, err
	}

	k.status = cl.Deleted

	return k.status.CreateClusterStatusMsg(), nil
}
