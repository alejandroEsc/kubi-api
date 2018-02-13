package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	s "github.com/alejandroEsc/cluster-apis/server/pkg"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	err := s.InitEnvVars()
	if err != nil {
		log.Fatalf("failed to init config vars: %s", err)
	}

	port, address := s.ParseServerEnvVars()

	//  Get notified that server is being asked to stop
	// Handle SIGINT and SIGTERM.
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	// Server Code
	err = s.Start(s.FmtAddress(address, port), gracefulStop)
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	} else {
		log.Printf("server started at: %s", s.FmtAddress(address, port))
	}
}
