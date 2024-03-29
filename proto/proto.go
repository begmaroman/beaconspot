package proto

//go:generate protoc --proto_path=$GOPATH/src:. --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/beaconspot/proto/status/status.proto
//go:generate protoc --proto_path=$GOPATH/src:. --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/beaconspot/proto/health/health.proto
//go:generate protoc --proto_path=$GOPATH/src:. --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/beaconspot/proto/common/error_response.proto
//go:generate protoc --proto_path=$GOPATH/src:. --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/beaconspot/proto/common/types.proto

//go:generate protoc -I $GOPATH/src/github.com/prysmaticlabs/prysm -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --proto_path=$GOPATH/src:. --go_out=$GOPATH/src --go_opt=paths=source_relative --go-grpc_out=$GOPATH/src --go-grpc_opt=paths=source_relative $GOPATH/src/github.com/begmaroman/beaconspot/proto/beaconspot/beaconspot.proto
