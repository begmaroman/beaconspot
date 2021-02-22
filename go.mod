module github.com/begmaroman/beaconspot

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace google.golang.org/api => google.golang.org/api v0.15.0

replace github.com/ethereum/go-ethereum => github.com/prysmaticlabs/bazel-go-ethereum v0.0.0-20201113091623-013fd65b3791

replace github.com/prysmaticlabs/prysm => github.com/prysmaticlabs/prysm v1.1.0

replace github.com/prysmaticlabs/ethereumapis => github.com/prysmaticlabs/ethereumapis v0.0.0-20210105190001-13193818c0df

require (
	github.com/bloxapp/eth2-key-manager v1.0.4
	github.com/ethereum/go-ethereum v1.9.24
	github.com/go-log/log v0.2.0
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prysmaticlabs/ethereumapis v0.0.0-20210218172602-3f05f78bea9d
	github.com/prysmaticlabs/prysm v1.1.0
	github.com/stretchr/testify v1.7.0
	go.opencensus.io v0.22.6
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
)
