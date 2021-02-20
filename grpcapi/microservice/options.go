package microservice

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

// Options contains the configuration parameters of the service.
type Options struct {
	IsTest          bool
	PrysmAddrs      []string
	LighthouseAddrs []string
}

// Load loads missing options
func (opts *Options) Load() error {
	opts.PrysmAddrs = strings.Split(os.Getenv("BC_PRYSM_ADDRS"), ",")
	opts.LighthouseAddrs = strings.Split(os.Getenv("BC_LIGHTHOUSE_ADDRS"), ",")
	opts.IsTest = os.Getenv("DOCKER_COMPOSE") == "true"
	return nil
}

// Validate applies the validation logic to the options.
func (opts *Options) Validate() error {
	return nil
}

// ClientOptions represent external dependencies.
type ClientOptions struct {
	Name    string
	Version string
	Log     *zap.Logger
}
