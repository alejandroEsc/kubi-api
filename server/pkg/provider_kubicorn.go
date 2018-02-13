package server

import (
	"log"
	"strings"

	"fmt"
	"github.com/alejandroEsc/cluster-apis/api"
	"github.com/kris-nova/kubicorn/pkg"
	"github.com/kris-nova/kubicorn/pkg/agent"
	"github.com/kris-nova/kubicorn/pkg/local"
	"github.com/kris-nova/kubicorn/pkg/task"
	kubi "github.com/alejandroEsc/cluster-apis/server/pkg/kubicorn_lib"
	"github.com/kris-nova/kubicorn/apis/cluster"
	"github.com/kris-nova/kubicorn/pkg/kubeconfig"
	"github.com/kris-nova/kubicorn/pkg/initapi"
)

type kubicorn struct {
	providerOpts ProviderOptions
	clusterOpts  ClusterOptions
	status       ClusterStatus
}


func NewKubicornProvider(p ProviderOptions, c ClusterOptions) *kubicorn {
	return &kubicorn{providerOpts: p,
		clusterOpts: c,
		status: UnCreated}
}

func (k *kubicorn) apply() (*clusteror.ClusterStatusMsg, error) {
	log.Printf("[applying] cluster %s ...", k.clusterOpts.Name)
	defer log.Print("...done")

	if k.status.Code > Applied.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, Applied.Msg)
	}

	if k.status.Code == Applied.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}
	var cluster *cluster.Cluster

	options := kubi.Options{StateStore: "fs"}

	// Expand state store path
	options.StateStorePath = kubi.ExpandPath(k.providerOpts.StorePath)
	stateStoreObj, err := options.NewStateStore()
	if err != nil {
		log.Printf("error found to be: %s", err)
		return nil, err
	}

	cluster, err = stateStoreObj.GetCluster()
	if err != nil {
		return nil, fmt.Errorf("Unable to get cluster [%s]: %v", k.clusterOpts.Name, err)
	}

	cluster, err = initapi.InitCluster(cluster)
	if err != nil {
		return nil, err
	}

	runtimeParams := &pkg.RuntimeParameters{}

	if k.clusterOpts.CloudProviderName == "aws" {
		runtimeParams.AwsProfile = "default"
	}

	reconciler, err := pkg.GetReconciler(cluster, runtimeParams)
	if err != nil {
		log.Printf("error found to be: %s", err)
		return nil, err
	}

	log.Printf("...getting expected")
	expected, err := reconciler.Expected(cluster)
	if err != nil {
		log.Printf("error found to be: %s", err)
		return nil, err
	}

	log.Printf("...getting actual")
	actual, err := reconciler.Actual(cluster)
	if err != nil {
		log.Printf("error found to be: %s", err)
		return nil, err
	}

	log.Printf("...reconciling actual with expected")
	reconciled, err := reconciler.Reconcile(actual, expected)
	if err != nil {
		log.Printf("error found to be: %s", err)
		return nil, err
	}

	log.Print("...committing reconciliation")
	err = stateStoreObj.Commit(reconciled)
	if err != nil {
		return nil, fmt.Errorf("Unable to commit state store: %v", err)
	}

	// Ensure we have SSH agent
	agent := agent.NewAgent()

	err = kubeconfig.RetryGetConfig(reconciled, agent)
	if err != nil {
		return nil, fmt.Errorf("Unable to write kubeconfig: %v", err)
	}

	log.Printf("The [%s] cluster has applied successfully!", reconciled.Name)
	if path, ok := reconciled.Annotations[kubeconfig.ClusterAnnotationKubeconfigLocalFile]; ok {
		path = local.Expand(path)
		log.Println("To start using your cluster, you need to run")
		log.Printf("  export KUBECONFIG=\"${KUBECONFIG}:%s\"", path)
	}
	log.Println("You can now `kubectl get nodes`")
	privKeyPath := strings.Replace(cluster.SSH.PublicKeyPath, ".pub", "", 1)
	log.Printf("You can SSH into your cluster ssh -i %s %s@%s", privKeyPath, reconciled.SSH.User, reconciled.KubernetesAPI.Endpoint)


	k.status = Applied

	return k.status.createClusterStatusMsg(), nil
}

func (k *kubicorn) create() (*clusteror.ClusterStatusMsg, error) {
	log.Printf("[creating] cluster %s ...", k.clusterOpts.Name)
	defer log.Print("...done")

	if k.status.Code > Created.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, Created.Msg)
	}

	if k.status.Code == Created.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	var newCluster *cluster.Cluster
	options := kubi.Options{StateStore: "fs"}

	if _, ok := kubi.ProfileMapIndexed[k.clusterOpts.CloudProviderName]; ok {
		newCluster = kubi.ProfileMapIndexed[k.clusterOpts.CloudProviderName].ProfileFunc(k.clusterOpts.Name)
	} else {
		return nil, fmt.Errorf("Invalid cloud provider [%s]", k.clusterOpts.CloudProviderName)
	}

	if newCluster.Cloud == cluster.CloudGoogle && k.clusterOpts.CloudId == "" {
		return nil, fmt.Errorf("CloudID is required for google cloud. Please set it to your project ID")
	}
	newCluster.CloudId = k.clusterOpts.CloudId

	// Expand state store path
	options.StateStorePath = kubi.ExpandPath(k.providerOpts.StorePath)

	// Register state store
	stateStoreObj, err := options.NewStateStore()
	if err != nil {
		return nil, err
	} else if stateStoreObj.Exists() {
		return nil, fmt.Errorf("State store [%s] exists, will not overwrite. Delete existing profile [%s] and retry", k.clusterOpts.Name, options.StateStorePath+"/"+k.clusterOpts.Name)
	}

	// Init new state store with the cluster resource
	err = stateStoreObj.Commit(newCluster)
	if err != nil {
		return nil, fmt.Errorf("Unable to init state store: %v", err)
	}

	k.status = Created
	return k.status.createClusterStatusMsg(), nil
}

func (k *kubicorn) delete() (*clusteror.ClusterStatusMsg, error) {
	log.Printf("[deleting] cluster %s ...", k.clusterOpts.Name)
	defer log.Print("...done")

	var err error
	var acluster *cluster.Cluster
	var deleteCluster *cluster.Cluster

	if k.status.Code > Deleted.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, Deleted.Msg)
	}

	if k.status.Code == Deleted.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	options := kubi.Options{StateStore: "fs"}

	// Expand state store path
	options.StateStorePath = kubi.ExpandPath(k.providerOpts.StorePath)

	// Register state store
	stateStoreObj, err := options.NewStateStore()
	if err != nil {
		return nil, err
	}

	acluster, err = stateStoreObj.GetCluster()
	if err != nil {
		return nil, fmt.Errorf("Unable to get cluster [%s]: %v", k.clusterOpts.Name, err)
	}

	var rp pkg.RuntimeParameters
	switch k.clusterOpts.CloudProviderName {
	case "aws":
		rp =  pkg.RuntimeParameters{AwsProfile: "default"}
	}

	reconciler, err := pkg.GetReconciler(acluster, &rp)
	if err != nil {
		return nil, err
	}

	var deleteClusterTask = func() error {
		deleteCluster, err = reconciler.Destroy()
		return err
	}

	err = task.RunAnnotated(deleteClusterTask, fmt.Sprintf("\nDestroying resources for cluster [%s]:\n", options.Name), "")
	if err != nil {
		return nil, fmt.Errorf("Unable to destroy resources for cluster [%s]: %v", options.Name, err)
	}

	err = stateStoreObj.Commit(deleteCluster)
	if err != nil {
		return nil, fmt.Errorf("Unable to save state store: %v", err)
	}

	err = stateStoreObj.Destroy()
	if err != nil {
		return nil, fmt.Errorf("Unable to remove state store for cluster [%s]: %v", options.Name, err)
	}

	k.status = Deleted
	return k.status.createClusterStatusMsg(), nil
}



