# BeaconSpot

Ethereum 2.0 multi-Beacon-Chain-node validator client written in Go.

This is the service to communicate with multiple different Beacon Chain nodes such as [Prysm](https://github.com/prysmaticlabs/prysm), [Lighthouse](https://github.com/sigp/lighthouse), etc. 

There are two transports of the BeaconSpot's API: [RPC](./proto/beaconspot/beaconspot.proto) and HTTP (*in development*).

The purposes of creation of the service are:
- Aggregate API of all Beacon Chain implementations to provide one common API so there is no needed to develop API clients for each of the Beacon Chain implementations.
- Improve response time during ETH2 validation process. Meaning a response from faster Beacon Node will be returned.
For instance, if Prysm node takes 1 second to get attestation data, and Lighthouse one takes 2 seconds, the response from Prysm node will be returned.
Multiple Beacon Chain nodes of one type could be configured as well.

## How does it work

Service supports multiple clients for multiple Beacon Node implementations. 

For instance, BeaconSpot can be configured to support 2 Prysm Beacon Nodes, 3 LightHouse Beacon Nodes, 2 Tekus, and 1 Nimbus.

All incoming requests to the BeaconSpot will be translated and sent to each of the configured nodes. 
The first response from one of the nodes will be returned, so if one of the nodes responds too slow, other node will respond faster so this response will be returned.
Service waits for the response from ALL nodes only for endpoints to submits something to the blockchain such as submitting attestation.

All needed configuration should be provided via environmental variables.

## Install

First, BeaconSpot server should be launched. This step requires running Beacon Chain nodes such as Prysm, Lighthouse, etc.

#### Docker

This is the example of the command how to run BeaconSpot in Docker container:

```bash
docker run \
-p 5000:5000 \
-e BC_PRYSM_ADDRS=http://lighthouse-address-1:5052,http://lighthouse-address-2:5052 \
-e BC_PRYSM_ADDRS=prysm-address-1:4000 \
begmaroman/beaconspot:latest
```

#### Source

Once there is running BeaconSpot service, the code should be prepared to use a client for the service.
BeaconSpot is a standard Go module which can be installed with:

```bash
go get github.com/begmaroman/beaconspot
```

Example how to use BeaconSpot gRPC client

```go
package main

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	beaconspotproto "github.com/begmaroman/beaconspot/proto/beaconspot"
)

func main() {
	// Open gRPC connection with BeaconSpot instance
	conn, err := grpc.Dial("beacon-spot:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	// Create BeaconSpot client for gRPC transport
	beaconSpotClient := beaconspotproto.NewBeaconSpotServiceClient(conn)

	// Send health request to make sure it works well
	_, err = beaconSpotClient.Health(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatal(err)
	}
}
```

## Inspired by

- [Ethereum foundation](https://ethereum.org/en/eth2/)
- [Prysm](https://github.com/prysmaticlabs/prysm)
- [Lighthouse](https://github.com/sigp/lighthouse)

## Tech stack

- [gRPC](https://grpc.io/) and [protocol buffer](https://developers.google.com/protocol-buffers)
- [Swagger (OpenAPI)](https://swagger.io/)
- [GoLang](https://golang.org/)
- [Docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/)

## Contribute

Contributions welcome. Please check out the [issues](https://github.com/begmaroman/beaconspot/issues).
You can also create your own issues with bugs or features.

## TODO

- Add more tests
- ~~Support Pryms client~~
- ~~Support Lighthouse client~~
- Support Teku client
- Support Nimbus client

## License

[Apache-2.0](./LICENSE) © 2020
