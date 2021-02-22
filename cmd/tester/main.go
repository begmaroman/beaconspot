package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	beaconspotproto "github.com/begmaroman/beaconspot/proto/beaconspot"
)

func main() {
	conn, err := grpc.Dial("5.9.6.37:5000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	beaconSpotService := beaconspotproto.NewBeaconSpotServiceClient(conn)

	resp, err := beaconSpotService.Ping(context.Background(), &empty.Empty{})
	fmt.Println("resp", resp)
	fmt.Println("err", err)
}
