package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	api "github.com/alejandroEsc/cluster-apis/api"
	configs "github.com/alejandroEsc/cluster-apis/server/pkg"
	"google.golang.org/grpc"
)

func createClusterDefinition() *api.ClusterDefinition {
	name := "ae_kluster"
	storePath := fmt.Sprintf("_state/%s", name)
	cloudProvider := "aws"
	clusterProvider := "kubicorn"

	cs := &api.ClusterConfigs{
		Name:              name,
		CloudProviderName: cloudProvider}

	cd := api.ClusterDefinition{
		ClusterProvider:          clusterProvider,
		AutoFetchClusterProvider: true,
		ClusterConfigs:           cs,
		ProviderStorePath:        storePath,
		CloudID:                  ""}

	return &cd
}

func runDoCreate(client api.ClusterCreatorClient) error {
	r, err := client.Create(context.Background(), createClusterDefinition())
	if err != nil {
		return err
	}

	log.Printf("reply message: %v", r)
	return err
}

func runDoApply(client api.ClusterCreatorClient) error {
	r, err := client.Apply(context.Background(), createClusterDefinition())
	if err != nil {
		return err
	}

	log.Printf("reply message: %v", r)
	return err
}

func runDoDelete(client api.ClusterCreatorClient) error {
	r, err := client.Delete(context.Background(), createClusterDefinition())
	if err != nil {
		return err
	}

	log.Printf("reply message: %v", r)
	return err
}

func runCommandPrintOutput(cmd string) error {
	log.Printf("attempting to run command: %s...", cmd)
	out, err := exec.Command("sh", "-c", cmd).Output()
	log.Print(string(out))

	if err != nil {
		log.Printf("found error attempting command: %s", err)
	}

	log.Printf(".. done")
	return err
}

func main() {
	if err := configs.InitEnvVars(); err != nil {
		log.Fatalf("failed to init config vars: %s", err)
	}

	port, address := configs.ParseServerEnvVars()

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	step, _ := configs.ParseClientEnvVars()

	conn, err := grpc.Dial(configs.FmtAddress(address, port), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewClusterCreatorClient(conn)

	switch strings.ToLower(step) {
	case "up":
		log.Println("Bringing a cluster up")
		err = runDoCreate(client)
		if err != nil {
			log.Printf("got an error message: %s", err)
		}

		err = runDoApply(client)
		if err != nil {
			log.Printf("got an error message: %s", err)
		}

	case "down":
		log.Println("Deleting a cluster")
		err = runDoDelete(client)
		if err != nil {
			log.Printf("got an error message: %s", err)
		}
	default:
		log.Printf("the command %s is not a valid task.", step)
	}

}
