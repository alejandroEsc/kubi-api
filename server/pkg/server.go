package server

import (
	"context"
	"log"
	"time"

	"os"

	api "github.com/alejandroEsc/cluster-apis/api"
	"google.golang.org/grpc"

	"net"
)

type clusterServer struct {
}

func (s *clusterServer) Create(c context.Context, cd *api.ClusterDefinition) (*api.ClusterStatusMsg, error) {
	err, clusterOpts, providerOpts := parseClusterDefinition(cd)
	if err != nil {
		log.Printf("error parsing cluster definition: %s", err)
		return nil, err
	}

	err, provider := GetProvider(providerOpts, clusterOpts)
	if err != nil {
		log.Printf("error parsing cluster definition: %s", err)
		return nil, err
	}

	return provider.create()

}

func (s *clusterServer) Apply(c context.Context, cd *api.ClusterDefinition) (*api.ClusterStatusMsg, error) {
	err, clusterOpts, providerOpts := parseClusterDefinition(cd)
	if err != nil {
		log.Printf("error parsing cluster definition: %s", err)
		return nil, err
	}

	err, provider := GetProvider(providerOpts, clusterOpts)
	if err != nil {
		log.Printf("error parsing cluster definition: %s", err)
		return nil, err
	}

	return provider.apply()

}

func (s *clusterServer) Delete(c context.Context, cd *api.ClusterDefinition) (*api.ClusterStatusMsg, error) {
	err, clusterOpts, providerOpts := parseClusterDefinition(cd)
	if err != nil {
		log.Printf("error parsing cluster definition: %s", err)
		return nil, err
	}

	err, provider := GetProvider(providerOpts, clusterOpts)
	if err != nil {
		log.Printf("error parsing cluster definition: %s", err)
		return nil, err
	}

	return provider.delete()
}

func Start(addr string, gracefulStop chan os.Signal) error {
	var err error
	log.Print("starting server")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
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
