package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	s "github.com/alejandroEsc/kubicorn-example-server/internal/app/clusterserver"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	//  Get notified that server is being asked to stop
	// Handle SIGINT and SIGTERM.
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	// Server Code
	err := s.Start(gracefulStop)
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
