package providers

import (
	"log"
	"fmt"

	"github.com/alejandroEsc/kubicorn-example-server/api"
	cl "github.com/alejandroEsc/kubicorn-example-server/pkg/clusterlib"
)

type kubicornCLI struct {
	providerOpts cl.ProviderOptions
	clusterOpts  cl.ClusterOptions
	status       cl.ClusterStatus
}

var (
	errorRevertState = "cluster is in state: %s, we cannot revert to state: %s"
	errorReplayState = "cluster is already in state: %s, cannot replay state"
	awsCloudProvider = "aws"
)

func NewKubicornProviderCLI(p cl.ProviderOptions, c cl.ClusterOptions) *kubicornCLI {
	return &kubicornCLI{providerOpts: p,
		clusterOpts: c,
		status:      cl.Planned}
}

func (k *kubicornCLI) Apply() (*clusteror.ClusterStatusMsg, error) {
	if k.status.Code > cl.Applied.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, cl.Applied.Msg)
	}

	if k.status.Code == cl.Applied.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	commandString := fmt.Sprintf("kubicorn apply %s", k.clusterOpts.Name)
	err := cl.RunCommandPrintOutput(commandString)
	if err != nil {
		return nil, err
	}

	k.status = cl.Applied

	return k.status.CreateClusterStatusMsg(), nil
}

func (k *kubicornCLI) Create() (*clusteror.ClusterStatusMsg, error) {
	if k.status.Code > cl.Created.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, cl.Created.Msg)
	}

	if k.status.Code == cl.Created.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	commandString := fmt.Sprintf("kubicorn create %s --profile %s", k.clusterOpts.Name, k.clusterOpts.CloudProviderName)
	err := cl.RunCommandPrintOutput(commandString)
	if err != nil {
		return nil, err
	}

	k.status = cl.Created

	return k.status.CreateClusterStatusMsg(), nil
}

func (k *kubicornCLI) Delete() (*clusteror.ClusterStatusMsg, error) {
	if k.status.Code > cl.Deleted.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, cl.Deleted.Msg)
	}

	if k.status.Code == cl.Deleted.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	commandString := fmt.Sprintf("kubicorn delete %s", k.clusterOpts.Name)
	err := cl.RunCommandPrintOutput(commandString)
	if err != nil {
		return nil, err
	}

	k.status = cl.Deleted
	log.Print("done")

	return k.status.CreateClusterStatusMsg(), nil
}
