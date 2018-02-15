package main

import (
	"context"
	"log"
	"os/exec"
	"strings"

	pkg "github.com/alejandroEsc/kubicorn-example-server/internal/pkg"
	api "github.com/alejandroEsc/kubicorn-example-server/api"
	"google.golang.org/grpc"
)

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

	log.Printf("reply message: %+v", r)
	return err
}

func runDoApply(client api.ClusterCreatorClient) error {
	r, err := client.Apply(context.Background(), createClusterDefinition())
	if err != nil {
		return err
	}

	log.Printf("reply message: %+v", r)
	return err
}

func runDoDelete(client api.ClusterCreatorClient) error {
	r, err := client.Delete(context.Background(), createClusterDefinition())
	if err != nil {
		return err
	}

	log.Printf("reply message: %+v", r)
	return err
}

func runCommandPrintOutput(cmd string) error {
	log.Printf("attempting to run command: %s...", cmd)
	out, err := exec.Command("sh", "-c", cmd).Output()
	log.Print(string(out))

	if err != nil {
		log.Printf("found error attempting command: %s", err)
	}

	log.Print(".. done")
	return err
}

func main() {
	if err := pkg.InitEnvVars(); err != nil {
		log.Fatalf("failed to init config vars: %s", err)
	}

	port, address := pkg.ParseServerEnvVars()

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	step, destroyAll := pkg.ParseClientEnvVars()

	conn, err := grpc.Dial(pkg.FmtAddress(address, port), opts...)
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

		if destroyAll {
			err = runCommandPrintOutput("rm -rf _state")
			if err != nil {
				log.Printf("got an error message: %s", err)
			}
		}
	default:
		log.Printf("the command %s is not a valid task.", step)
	}

}
