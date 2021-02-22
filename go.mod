module github.com/begmaroman/beaconspot

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.27.0

replace google.golang.org/api => google.golang.org/api v0.15.0

require (
	github.com/bloxapp/eth2-key-manager v1.0.4
	github.com/ethereum/go-ethereum v1.9.25
	github.com/go-log/log v0.2.0
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pkg/errors v0.9.1
	github.com/prysmaticlabs/eth2-types v0.0.0-20210210115503-cf4ec6600a2d
	github.com/prysmaticlabs/ethereumapis v0.0.0-20210211220440-bfff608b8ba9
	github.com/prysmaticlabs/prysm v1.1.0
	github.com/stretchr/testify v1.6.1
	go.opencensus.io v0.22.5
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)

replace github.com/ethereum/go-ethereum => github.com/prysmaticlabs/bazel-go-ethereum v0.0.0-20201126065335-1fb46e307951
