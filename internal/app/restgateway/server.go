package restgateway

import (
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	api "github.com/alejandroEsc/kubicorn-example-server/api"
	ipkg "github.com/alejandroEsc/kubicorn-example-server/internal/pkg"
	cl "github.com/alejandroEsc/kubicorn-example-server/pkg/clusterlib"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/juju/loggo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var logger loggo.Logger

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	logger.Infof("call to find swagger resource.... %s", r.URL.Path)
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		logger.Errorf("Not a swagger file %s, missing suffix .swagger.json", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	gwSwaggerDir := ipkg.ParseGWSwaggerEnvVars()

	logger.Infof("Serving %s", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(gwSwaggerDir, p)
	http.ServeFile(w, r, p)
}

func Start(gracefulStop chan os.Signal) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	port, address := ipkg.ParseServerEnvVars()
	gwPort, gwAddress := ipkg.ParseGateWayEnvVars()

	grpcEndpoint := ipkg.FmtAddress(address, port)

	logLevel := ipkg.ParseLogLevel()
	logger = cl.GetModuleLogger("internal.app.restgateway", logLevel)

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", serveSwagger)

	muxGateway := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := api.RegisterClusterCreatorHandlerFromEndpoint(ctx, muxGateway, grpcEndpoint, opts)
	if err != nil {
		logger.Errorf("could not register from endpoints: %s", err)
		return err
	}

	mux.Handle("/", muxGateway)

	// Chance here to gracefully handle being stopped.
	go func() {
		sig := <-gracefulStop
		logger.Infof("caught sig: %+v", sig)
		logger.Infof("Wait for 2 second to finish processing")
		time.Sleep(2 * time.Second)
		cancel()
		logger.Infof("service terminated")
		os.Exit(0)
	}()

	return http.ListenAndServe(ipkg.FmtAddress(gwAddress, gwPort), mux)
}
