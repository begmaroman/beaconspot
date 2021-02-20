package microservice

import (
	"github.com/micro/cli/v2"
)

var opts Options

// flags contains the list of configuration parameters
var flags = []cli.Flag{
	&cli.BoolFlag{
		Name:        "docker-compose",
		EnvVars:     []string{"DOCKER_COMPOSE"},
		Usage:       "Set to true if we are running in docker-compose",
		Destination: &opts.IsTest,
	},

	// Beacon chain addresses
	&cli.StringSliceFlag{
		Name:    "prysm-addrs",
		EnvVars: []string{"BC_PRYSM_ADDRS"},
		Usage:   "Comma separated Prysm node addresses",
	},
	&cli.StringSliceFlag{
		Name:    "lighthouse-addrs",
		EnvVars: []string{"BC_LIGHTHOUSE_ADDRS"},
		Usage:   "Comma separated LightHouse node addresses",
	},
}

// Load loads missing options
func (opts *Options) Load(ctx *cli.Context) error {
	opts.PrysmAddrs = ctx.StringSlice("prysm-addrs")
	opts.LighthouseAddrs = ctx.StringSlice("lighthouse-addrs")
	return nil
}
