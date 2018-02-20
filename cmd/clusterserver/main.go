package main

import (
	"os"
	"os/signal"
	"syscall"

	s "github.com/alejandroEsc/kubicorn-example-server/internal/app/clusterserver"
	cl "github.com/alejandroEsc/kubicorn-example-server/pkg/clusterlib"
	"github.com/juju/loggo"
	"golang.org/x/net/context"
)

func main() {
	logger := cl.GetModuleLogger("cmd.restgateway", loggo.INFO)
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	//  Get notified that server is being asked to stop
	// Handle SIGINT and SIGTERM.
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	// Server Code
	logger.Infof("starting server...")
	err := s.Start(gracefulStop)
	if err != nil {
		logger.Criticalf("failed to start server: %s", err)
		os.Exit(1)
	}
}
