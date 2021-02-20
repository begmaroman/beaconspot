package grpcex

import (
	"fmt"
	"strings"
	"time"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"go.opencensus.io/plugin/ocgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/begmaroman/beaconspot/utils/retry"
)

// Default values.
const (
	DefaultMaxCallRecvMsgSize      = 10 * 5 << 20 // Default 50Mb
	DefaultGRPCRetries        uint = 3
	defaultGRPCTimeout             = time.Second
)

// Default values.
var (
	defaultGRPCHeaders = strings.Split("", ",")
)

// Dial creates a new gRPC connection.
func Dial(logger *zap.Logger, endpoint string) (*grpc.ClientConn, error) {
	// Prepare options.
	opts, err := constructDialOptions(
		logger,
		DefaultMaxCallRecvMsgSize,
		defaultGRPCHeaders,
		DefaultGRPCRetries,
		defaultGRPCTimeout,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct dial options")
	}

	// Dial connection
	x, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial context")
	}

	return x, nil
}

// constructDialOptions constructs a list of grpc dial options
func constructDialOptions(
	logger *zap.Logger,
	maxCallRecvMsgSize int,
	grpcHeaders []string,
	grpcRetries uint,
	grpcTimeout time.Duration,
) ([]grpc.DialOption, error) {
	if maxCallRecvMsgSize == 0 {
		maxCallRecvMsgSize = 10 * 5 << 20 // Default 50Mb
	}

	md := make(metadata.MD)
	for _, hdr := range grpcHeaders {
		if hdr != "" {
			ss := strings.Split(hdr, "=")
			if len(ss) != 2 {
				return nil, fmt.Errorf("incorrect gRPC header flag format, skipping %v", hdr)
			}
			md.Set(ss[0], ss[1])
		}
	}

	// Use 15 sec timeout by default
	if grpcTimeout == 0 {
		grpcTimeout = time.Second * 15
	}

	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(maxCallRecvMsgSize),
			grpc_retry.WithMax(grpcRetries),
			grpc.Header(&md),
		),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
		grpc.WithUnaryInterceptor(middleware.ChainUnaryClient(
			grpc_opentracing.UnaryClientInterceptor(),
			grpc_prometheus.UnaryClientInterceptor,
			retry.UnaryInterceptor(logger, grpcRetries, time.Millisecond*100),
		)),
	}

	return dialOpts, nil
}
