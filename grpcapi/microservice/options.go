package microservice

import (
	"go.uber.org/zap"
)

// Options contains the configuration parameters of the service.
type Options struct {
	IsTest          bool
	PrysmAddrs      []string
	LighthouseAddrs []string
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
