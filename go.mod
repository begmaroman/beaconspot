module github.com/begmaroman/beaconspot

go 1.15

require (
	github.com/bloxapp/eth2-key-manager v1.0.4
	github.com/ethereum/go-ethereum v1.9.25
	github.com/go-log/log v0.2.0
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/transport/grpc/v2 v2.9.1
	github.com/pkg/errors v0.9.1
	github.com/prysmaticlabs/eth2-types v0.0.0-20210210115503-cf4ec6600a2d
	github.com/prysmaticlabs/ethereumapis v0.0.0-20210211220440-bfff608b8ba9
	github.com/prysmaticlabs/prysm v1.2.2
	github.com/stretchr/testify v1.6.1
	go.opencensus.io v0.22.5
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace google.golang.org/api => google.golang.org/api v0.15.0

replace github.com/ethereum/go-ethereum => github.com/prysmaticlabs/bazel-go-ethereum v0.0.0-20201126065335-1fb46e307951
