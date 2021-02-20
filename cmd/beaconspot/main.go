package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	// We need these imports to register NATS broker, registry, and transport.
	// This type is defined through evars in docker-compose.yaml file.
	_ "github.com/micro/go-plugins/transport/grpc/v2"

	"github.com/begmaroman/beaconspot/grpcapi/microservice"
	"github.com/begmaroman/beaconspot/utils/logex"
)

const (
	serviceName = "BeaconSpot"
)

// Version may be changed during build via --ldflags parameter
var Version = "latest"

func main() {
	logger := logex.Build(serviceName, zapcore.InfoLevel)

	// Initialize service.
	microService, err := microservice.Init(&microservice.ClientOptions{
		Name:    serviceName,
		Version: Version,
		Log:     logger,
	})
	if err != nil {
		logger.Fatal("failed to initialize micro-service", zap.Error(err))
	}

	// Run service.
	if err := microService.Run(); err != nil {
		logger.Fatal("failed to run micro-service", zap.Error(err))
	}
}
