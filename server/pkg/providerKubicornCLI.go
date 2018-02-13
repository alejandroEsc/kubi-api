package server

import (
	"log"

	"fmt"

	"github.com/alejandroEsc/cluster-apis/api"
)

type kubicornCLI struct {
	providerOpts ProviderOptions
	clusterOpts  clusterOptions
	status       clusterStatus
}

var (
	errorRevertState = "cluster is in state: %s, we cannot revert to state: %s"
	errorReplayState = "cluster is already in state: %s, cannot replay state"
	awsCloudProvider = "aws"
)

func newKubicornProviderCLI(p ProviderOptions, c clusterOptions) *kubicornCLI {
	return &kubicornCLI{providerOpts: p,
		clusterOpts: c,
		status:      unCreated}
}

func (k *kubicornCLI) apply() (*clusteror.ClusterStatusMsg, error) {
	if k.status.Code > applied.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, applied.Msg)
	}

	if k.status.Code == applied.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	commandString := fmt.Sprintf("kubicorn apply %s", k.clusterOpts.Name)
	err := runCommandPrintOutput(commandString)
	if err != nil {
		return nil, err
	}

	k.status = applied

	return k.status.createClusterStatusMsg(), nil
}

func (k *kubicornCLI) create() (*clusteror.ClusterStatusMsg, error) {
	if k.status.Code > created.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, created.Msg)
	}

	if k.status.Code == created.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	commandString := fmt.Sprintf("kubicorn create %s --profile %s", k.clusterOpts.Name, k.clusterOpts.CloudProviderName)
	err := runCommandPrintOutput(commandString)
	if err != nil {
		return nil, err
	}

	k.status = created

	return k.status.createClusterStatusMsg(), nil
}

func (k *kubicornCLI) delete() (*clusteror.ClusterStatusMsg, error) {
	if k.status.Code > deleted.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, deleted.Msg)
	}

	if k.status.Code == deleted.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	commandString := fmt.Sprintf("kubicorn delete %s", k.clusterOpts.Name)
	err := runCommandPrintOutput(commandString)
	if err != nil {
		return nil, err
	}

	k.status = deleted
	log.Print("done")

	return k.status.createClusterStatusMsg(), nil
}
