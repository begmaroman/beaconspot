module github.com/begmaroman/beaconspot

go 1.15

replace github.com/prysmaticlabs/prysm => github.com/prysmaticlabs/prysm v0.0.0-20210830193317-49dce52ae963

require (
	github.com/bloxapp/eth2-key-manager v1.0.4
	github.com/ethereum/go-ethereum v1.9.25
	github.com/go-log/log v0.2.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prysmaticlabs/eth2-types v0.0.0-20210303084904-c9735a06829d
	github.com/prysmaticlabs/ethereumapis v0.0.0-20210218172602-3f05f78bea9d
	github.com/prysmaticlabs/prysm v1.4.4
	github.com/stretchr/testify v1.7.0
	go.opencensus.io v0.23.0
	go.uber.org/zap v1.18.1
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.27.1
)
