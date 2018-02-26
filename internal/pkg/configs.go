package internalpkg

import (
	"github.com/juju/loggo"
	"github.com/spf13/viper"
)

var (
	logger = loggo.GetLogger("internal.pkg")

	envPrefix                = "cluster_creator"
	envServiceAddress        = "service_address"
	envServicePort           = "service_port"
	envGateWayServiceAddress = "gw_service_address"
	envGateWayPort           = "gw_port"
	envGateWaySwaggerDir     = "gw_swagger_dir"

	envLogLevel = "log_level"

	envClientStep             = "step"
	envClientDestroyArtifacts = "destroy_artifacts"

	defaultAddress      = "localhost"
	defaultPort         = 8501
	defaultGWAddress    = "localhost"
	defaultGWPort       = 8502
	defaultGWSwaggerDir = "swagger"
	defaultClientStep   = "up"

	defaultLogLevel = "info"
)

// InitEnvVars allows you to initiate gathering environment variables
func InitEnvVars() error {
	var err error
	viper.SetEnvPrefix(envPrefix)

	if err = viper.BindEnv(envServicePort); err != nil {
		return err
	}

	if err = viper.BindEnv(envServiceAddress); err != nil {
		return err
	}

	if err = viper.BindEnv(envGateWayServiceAddress); err != nil {
		return err
	}

	if err = viper.BindEnv(envGateWayPort); err != nil {
		return err
	}

	if err = viper.BindEnv(envGateWaySwaggerDir); err != nil {
		return err
	}

	if err = viper.BindEnv(envClientStep); err != nil {
		return err
	}

	if err = viper.BindEnv(envClientDestroyArtifacts); err != nil {
		return err
	}

	err = viper.BindEnv(envLogLevel)

	return err
}

// ParseGWSwaggerEnvVars parses environment variables consumed by swagger server
func ParseGWSwaggerEnvVars() string {
	gwSwaggerDir := viper.GetString(envGateWaySwaggerDir)
	if gwSwaggerDir == "" {
		gwSwaggerDir = defaultGWSwaggerDir
	}
	return gwSwaggerDir
}

// ParseServerEnvVars parses environment variables consumed by the gateway service
func ParseServerEnvVars() (int, string) {
	port := viper.GetInt(envServicePort)
	if port == 0 {
		port = defaultPort
	}

	serviceAddress := viper.GetString(envServiceAddress)
	if serviceAddress == "" {
		serviceAddress = defaultAddress
	}

	return port, serviceAddress
}

// ParseClientEnvVars parses environment variables consumed by clients
func ParseClientEnvVars() (string, bool) {
	clientStep := viper.GetString(envClientStep)
	if clientStep == "" {
		clientStep = defaultClientStep
	}

	destroyAll := viper.GetBool(envClientDestroyArtifacts)

	return clientStep, destroyAll
}

// ParseGateWayEnvVars parses environment variables consumed by the gateway service
func ParseGateWayEnvVars() (int, string) {
	gwPort := viper.GetInt(envGateWayPort)
	if gwPort == 0 {
		gwPort = defaultGWPort
	}

	gwServiceAddress := viper.GetString(envGateWayServiceAddress)
	if gwServiceAddress == "" {
		gwServiceAddress = defaultGWAddress
	}

	return gwPort, gwServiceAddress
}

// ParseLogLevel parses environment variables for log levels
func ParseLogLevel() loggo.Level {
	logString := viper.GetString(envLogLevel)
	logL, ok := loggo.ParseLevel(logString)

	if !ok {
		logL = loggo.INFO
	}

	return logL
}
