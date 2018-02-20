package main

import (
	"context"
	"os/exec"
	"strings"
	api "github.com/alejandroEsc/kubicorn-example-server/api"
	pkg "github.com/alejandroEsc/kubicorn-example-server/internal/pkg"
	"github.com/juju/loggo"
	cl "github.com/alejandroEsc/kubicorn-example-server/pkg/clusterlib"
	"google.golang.org/grpc"
	"os"
)

var logger loggo.Logger

func createClusterDefinition() *api.ClusterDefinition {
	cs := &api.ClusterConfigs{
		Name:              "ae_kluster",
		CloudProviderName: "aws"}

	cd := api.ClusterDefinition{
		ClusterProvider:          "kubicorn_cli",
		AutoFetchClusterProvider: true,
		ClusterConfigs:           cs}

	return &cd
}

func runDoCreate(client api.ClusterCreatorClient) error {
	r, err := client.Create(context.Background(), createClusterDefinition())
	if err != nil {
		return err
	}

	logger.Infof("reply message: %+v", r)
	return err
}

func runDoApply(client api.ClusterCreatorClient) error {
	r, err := client.Apply(context.Background(), createClusterDefinition())
	if err != nil {
		return err
	}

	logger.Infof("reply message: %+v", r)
	return err
}

func runDoDelete(client api.ClusterCreatorClient) error {
	r, err := client.Delete(context.Background(), createClusterDefinition())
	if err != nil {
		return err
	}

	logger.Infof("reply message: %+v", r)
	return err
}

func runCommandPrintOutput(cmd string) error {
	logger.Infof("attempting to run command: %s...", cmd)
	out, err := exec.Command("sh", "-c", cmd).Output()
	logger.Infof(string(out))

	if err != nil {
		logger.Errorf("found error attempting command: %s", err)
	}

	logger.Infof(".. done")
	return err
}

func main() {
	if err := pkg.InitEnvVars(); err != nil {
		logger.Criticalf("failed to init config vars: %s", err)
		os.Exit(1)
	}

	port, address := pkg.ParseServerEnvVars()

	logLevel := pkg.ParseLogLevel()
	logger = cl.GetModuleLogger("cmd.clientkubicorn", logLevel)

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	step, destroyAll := pkg.ParseClientEnvVars()

	conn, err := grpc.Dial(pkg.FmtAddress(address, port), opts...)
	if err != nil {
		logger.Criticalf("fail to dial: %v", err)
		os.Exit(1)

	}
	defer conn.Close()
	client := api.NewClusterCreatorClient(conn)

	switch strings.ToLower(step) {
	case "up":
		logger.Infof("Bringing a cluster up")
		err = runDoCreate(client)
		if err != nil {
			logger.Infof("got an error message: %s", err)
		}

		err = runDoApply(client)
		if err != nil {
			logger.Infof("got an error message: %s", err)
		}

	case "down":
		logger.Infof("Deleting a cluster")
		err = runDoDelete(client)
		if err != nil {
			logger.Infof("got an error message: %s", err)
		}

		if destroyAll {
			err = runCommandPrintOutput("rm -rf _state")
			if err != nil {
				logger.Infof("got an error message: %s", err)
			}
		}
	default:
		logger.Infof("the command %s is not a valid task.", step)
	}

}
