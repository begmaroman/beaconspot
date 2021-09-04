package main

import (
	"github.com/begmaroman/beaconspot/grpcapi/microservice"
	"github.com/begmaroman/beaconspot/utils/logex"
	"github.com/herumi/bls-eth-go-binary/bls"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

// InitBLS initializes BLS
func init() {
	if err := bls.Init(bls.BLS12_381); err != nil {
		panic(err)
	}

	if err := bls.SetETHmode(bls.EthModeDraft07); err != nil {
		panic(err)
	}
}
