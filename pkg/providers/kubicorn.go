package providers

import (
	"log"
	"strings"

	"fmt"

	"github.com/alejandroEsc/kubicorn-example-server/api"
	kubi "github.com/alejandroEsc/kubicorn-example-server/internal/app/clusterserver/kubicornlib"
	cl "github.com/alejandroEsc/kubicorn-example-server/pkg/clusterlib"
	"github.com/kris-nova/kubicorn/apis/cluster"
	"github.com/kris-nova/kubicorn/pkg"
	"github.com/kris-nova/kubicorn/pkg/agent"
	"github.com/kris-nova/kubicorn/pkg/initapi"
	"github.com/kris-nova/kubicorn/pkg/kubeconfig"
	"github.com/kris-nova/kubicorn/pkg/local"
	"github.com/kris-nova/kubicorn/pkg/task"
)

// Kubicorn represents a kubicorn provider via library calls
type Kubicorn struct {
	providerOpts cl.ProviderOptions
	clusterOpts  cl.ClusterOptions
	status       cl.ClusterStatus
}

// NewKubicornProvider returns provider for creating kubicorn cluster via library calls.
func NewKubicornProvider(p cl.ProviderOptions, c cl.ClusterOptions) *Kubicorn {
	return &Kubicorn{providerOpts: p,
		clusterOpts: c,
		status:      cl.Planned}
}

// Apply commits state changes.
func (k *Kubicorn) Apply() (*clusteror.ClusterStatusMsg, error) {
	log.Printf("[applying] cluster %s ...", k.clusterOpts.Name)
	defer log.Print("...done")

	if k.status.Code > cl.Applied.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, cl.Applied.Msg)
	}

	if k.status.Code == cl.Applied.Code {
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

	if k.clusterOpts.CloudProviderName == awsCloudProvider {
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

	k.status = cl.Applied

	return k.status.CreateClusterStatusMsg(), nil
}

// Create new cluster if it does not exists, meaning creates supporting state files
func (k *Kubicorn) Create() (*clusteror.ClusterStatusMsg, error) {
	log.Printf("[creating] cluster %s ...", k.clusterOpts.Name)
	defer log.Print("...done")

	if k.status.Code > cl.Created.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, cl.Created.Msg)
	}

	if k.status.Code == cl.Created.Code {
		return nil, fmt.Errorf(errorReplayState, k.status.Msg)
	}

	var newCluster *cluster.Cluster
	options := kubi.Options{StateStore: "fs"}

	if _, ok := kubi.ProfileMapIndexed[k.clusterOpts.CloudProviderName]; ok {
		newCluster = kubi.ProfileMapIndexed[k.clusterOpts.CloudProviderName].ProfileFunc(k.clusterOpts.Name)
	} else {
		return nil, fmt.Errorf("Invalid cloud provider [%s]", k.clusterOpts.CloudProviderName)
	}

	if newCluster.Cloud == cluster.CloudGoogle && k.clusterOpts.CloudID == "" {
		return nil, fmt.Errorf("CloudID is required for google cloud. Please set it to your project ID")
	}
	newCluster.CloudId = k.clusterOpts.CloudID

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

	k.status = cl.Created
	return k.status.CreateClusterStatusMsg(), nil
}

// Delete and destroy cluster and its resources
func (k *Kubicorn) Delete() (*clusteror.ClusterStatusMsg, error) {
	log.Printf("[deleting] cluster %s ...", k.clusterOpts.Name)
	defer log.Print("...done")

	var err error
	var acluster *cluster.Cluster
	var deleteCluster *cluster.Cluster

	if k.status.Code > cl.Deleted.Code {
		return nil, fmt.Errorf(errorRevertState, k.status.Msg, cl.Deleted.Msg)
	}

	if k.status.Code == cl.Deleted.Code {
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
	case awsCloudProvider:
		rp = pkg.RuntimeParameters{AwsProfile: "default"}
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

	k.status = cl.Deleted
	return k.status.CreateClusterStatusMsg(), nil
}
