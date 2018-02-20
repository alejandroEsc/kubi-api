package server

import (
	"context"
	"time"

	"os"

	api "github.com/alejandroEsc/kubicorn-example-server/api"
	ipkg "github.com/alejandroEsc/kubicorn-example-server/internal/pkg"
	cl "github.com/alejandroEsc/kubicorn-example-server/pkg/clusterlib"
	prs "github.com/alejandroEsc/kubicorn-example-server/pkg/providers"
	"google.golang.org/grpc"

	"fmt"
	"net"

	"github.com/juju/loggo"
)

type clusterServer struct {
	provider cl.Provider
}

var logger loggo.Logger

func (s *clusterServer) Create(c context.Context, cd *api.ClusterDefinition) (*api.ClusterStatusMsg, error) {
	var err error
	s.provider, err = s.getProviderParseOptions(cd)
	if err != nil {
		logger.Errorf("error parsing cluster definition: %s", err)
		return nil, err
	}

	return s.provider.Create()

}

func (s *clusterServer) Apply(c context.Context, cd *api.ClusterDefinition) (*api.ClusterStatusMsg, error) {
	var err error
	s.provider, err = s.getProviderParseOptions(cd)
	if err != nil {
		logger.Errorf("error parsing cluster definition: %s", err)
		return nil, err
	}

	return s.provider.Apply()

}

func (s *clusterServer) Delete(c context.Context, cd *api.ClusterDefinition) (*api.ClusterStatusMsg, error) {
	var err error
	s.provider, err = s.getProviderParseOptions(cd)
	if err != nil {
		logger.Errorf("error parsing cluster definition: %s", err)
		return nil, err
	}

	return s.provider.Delete()
}

func (s *clusterServer) getProviderParseOptions(cd *api.ClusterDefinition) (cl.Provider, error) {
	clusterOpts, providerOpts := cl.ParseClusterDefinition(cd)

	provider, err := prs.GetProvider(providerOpts, clusterOpts)
	if err != nil {
		logger.Errorf("error parsing cluster definition: %s", err)
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

	logLevel := ipkg.ParseLogLevel()
	logger = cl.GetModuleLogger("internal.app.clusterserver", logLevel)

	logger.Infof("starting server")

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
		logger.Infof("caught sig: %+v", sig)
		logger.Infof("waiting for 2 second to finish processing")
		time.Sleep(2 * time.Second)
		grpcServer.Stop()
		logger.Infof("server terminated")
		os.Exit(0)
	}()

	logger.Infof("attempting to start server in address: %s", addr)

	return grpcServer.Serve(listener)
}
