package proto

//go:generate protoc --proto_path=$GOPATH/src:. --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/beaconspot/proto/status/status.proto
//go:generate protoc --proto_path=$GOPATH/src:. --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/beaconspot/proto/health/health.proto
//go:generate protoc --proto_path=$GOPATH/src:. --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/beaconspot/proto/common/error_response.proto
//go:generate protoc --proto_path=$GOPATH/src:. --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/beaconspot/proto/common/types.proto

//go:generate protoc --proto_path=$GOPATH/src:. --go_out=plugins=grpc:$GOPATH/src $GOPATH/src/github.com/begmaroman/beaconspot/proto/beaconspot/beaconspot.proto
