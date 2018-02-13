package server

import (
	"log"

	"fmt"

	"github.com/alejandroEsc/cluster-apis/api"
)

type kubicornCLI struct {
	providerOpts ProviderOptions
	clusterOpts  ClusterOptions
	status       ClusterStatus
}

var (
	errorRevertState = "cluster is in state: %s, we cannot revert to state: %s"
	errorReplayState = "cluster is already in state: %s, cannot replay state"
)

func NewKubicornProviderCLI(p ProviderOptions, c ClusterOptions) *kubicornCLI {
	return &kubicornCLI{providerOpts: p,
		clusterOpts: c,
		status:      UnCreated}
}

func (k *kubicornCLI) apply() (*clusteror.ClusterStatusMsg, error) {
	if k.status.Code > Applied.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, Applied.Msg)
	}

	if k.status.Code == Applied.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	commandString := fmt.Sprintf("kubicorn apply %s", k.clusterOpts.Name)
	err := runCommandPrintOutput(commandString)
	if err != nil {
		return nil, err
	}

	k.status = Applied

	return k.status.createClusterStatusMsg(), nil
}

func (k *kubicornCLI) create() (*clusteror.ClusterStatusMsg, error) {
	if k.status.Code > Created.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, Created.Msg)
	}

	if k.status.Code == Created.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	commandString := fmt.Sprintf("kubicorn create %s --profile %s", k.clusterOpts.Name, k.clusterOpts.CloudProviderName)
	err := runCommandPrintOutput(commandString)
	if err != nil {
		return nil, err
	}

	k.status = Created

	return k.status.createClusterStatusMsg(), nil
}

func (k *kubicornCLI) delete() (*clusteror.ClusterStatusMsg, error) {
	if k.status.Code > Deleted.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, Deleted.Msg)
	}

	if k.status.Code == Deleted.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	commandString := fmt.Sprintf("kubicorn delete %s", k.clusterOpts.Name)
	err := runCommandPrintOutput(commandString)
	if err != nil {
		return nil, err
	}

	k.status = Deleted
	log.Print("done")

	return k.status.createClusterStatusMsg(), nil
}
