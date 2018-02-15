package server

import (
	"context"
	"log"
	"time"

	"os"

	api "github.com/alejandroEsc/kubicorn-example-server/api"
	ipkg "github.com/alejandroEsc/kubicorn-example-server/internal/pkg"
	cl "github.com/alejandroEsc/kubicorn-example-server/pkg/clusterlib"
	prs "github.com/alejandroEsc/kubicorn-example-server/pkg/providers"
	"google.golang.org/grpc"

	"fmt"
	"net"
)

type clusterServer struct {
	provider cl.Provider
}

func (s *clusterServer) Create(c context.Context, cd *api.ClusterDefinition) (*api.ClusterStatusMsg, error) {
	var err error
	s.provider, err = s.getProviderParseOptions(cd)
	if err != nil {
		log.Printf("error parsing cluster definition: %s", err)
		return nil, err
	}

	return s.provider.Create()

}

func (s *clusterServer) Apply(c context.Context, cd *api.ClusterDefinition) (*api.ClusterStatusMsg, error) {
	var err error
	s.provider, err = s.getProviderParseOptions(cd)
	if err != nil {
		log.Printf("error parsing cluster definition: %s", err)
		return nil, err
	}

	return s.provider.Apply()

}

func (s *clusterServer) Delete(c context.Context, cd *api.ClusterDefinition) (*api.ClusterStatusMsg, error) {
	var err error
	s.provider, err = s.getProviderParseOptions(cd)
	if err != nil {
		log.Printf("error parsing cluster definition: %s", err)
		return nil, err
	}

	return s.provider.Delete()
}

func (s *clusterServer) getProviderParseOptions(cd *api.ClusterDefinition) (cl.Provider, error) {
	clusterOpts, providerOpts := cl.ParseClusterDefinition(cd)

	provider, err := prs.GetProvider(providerOpts, clusterOpts)
	if err != nil {
		log.Printf("error parsing cluster definition: %s", err)
		return nil, err
	}

	return provider, nil
}

// Start the server here.
func Start(gracefulStop chan os.Signal) error {
	err := ipkg.InitEnvVars()
	if err != nil {
		return fmt.Errorf("failed to init config vars: %s", err)
	}

	port, address := ipkg.ParseServerEnvVars()
	addr := ipkg.FmtAddress(address, port)

	log.Print("starting server")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	api.RegisterClusterCreatorServer(grpcServer, &clusterServer{})

	// Chance here to gracefully handle being stopped.
	go func() {
		sig := <-gracefulStop
		log.Printf("caught sig: %+v", sig)
		log.Println("waiting for 2 second to finish processing")
		time.Sleep(2 * time.Second)
		grpcServer.Stop()
		log.Print("server terminated")
		os.Exit(0)
	}()

	log.Printf("attempting to start server in address: %s", addr)

	return grpcServer.Serve(listener)
}
