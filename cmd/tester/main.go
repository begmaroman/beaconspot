package main

import (
	"context"
	"fmt"

	beaconspotproto "github.com/begmaroman/beaconspot/proto/beaconspot"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	beaconSpotService := beaconspotproto.NewBeaconSpotServiceClient(conn)

	resp, err := beaconSpotService.Ping(context.Background(), &empty.Empty{})
	fmt.Println("resp", resp)
	fmt.Println("err", err)
}
