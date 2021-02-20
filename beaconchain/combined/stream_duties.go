package combined

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"google.golang.org/grpc/metadata"
)

// streamDuties implements ethpb.BeaconNodeValidator_StreamDutiesClient interface
type streamDuties struct {
	clients []ethpb.BeaconNodeValidator_StreamDutiesClient
}

// newStreamDuties is the constructor of streamDuties
func newStreamDuties(clients ...ethpb.BeaconNodeValidator_StreamDutiesClient) ethpb.BeaconNodeValidator_StreamDutiesClient {
	return &streamDuties{
		clients: clients,
	}
}

func (c *streamDuties) Recv() (*ethpb.DutiesResponse, error) {
	type result struct {
		res *ethpb.DutiesResponse
		err error
	}

	feed := &event.Feed{}

	dataChan := make(chan *result)
	dataSub := feed.Subscribe(dataChan)
	defer dataSub.Unsubscribe()

	for _, client := range c.clients {
		go func(client ethpb.BeaconNodeValidator_StreamDutiesClient) {
			if client == nil {
				feed.Send(&result{
					err: errors.New("nil client"),
				})
				return
			}

			res, err := client.Recv()
			feed.Send(&result{
				res: res,
				err: err,
			})
		}(client)
	}

	var errs []string
	for i := 0; i < len(c.clients); i++ {
		res := <-dataChan
		if res.err != nil {
			errs = append(errs, res.err.Error())
		} else {
			return res.res, nil
		}
	}

	return nil, errors.Errorf("failed to receive duties from clients: %s", strings.Join(errs, ", "))
}

func (c *streamDuties) Header() (metadata.MD, error) {
	panic("not implemented")
}

func (c *streamDuties) Trailer() metadata.MD {
	panic("not implemented")
}

func (c *streamDuties) CloseSend() error {
	type result struct {
		err error
	}

	feed := &event.Feed{}

	dataChan := make(chan *result)
	dataSub := feed.Subscribe(dataChan)
	defer dataSub.Unsubscribe()

	for _, client := range c.clients {
		go func(client ethpb.BeaconNodeValidator_StreamDutiesClient) {
			feed.Send(&result{
				err: client.CloseSend(),
			})
		}(client)
	}

	var errs []string
	for i := 0; i < len(c.clients); i++ {
		res := <-dataChan
		if res.err != nil {
			errs = append(errs, res.err.Error())
		} else {
			return nil
		}
	}

	return errors.Errorf("failed to close: %s", strings.Join(errs, ", "))
}

func (c *streamDuties) Context() context.Context {
	type result struct {
		res context.Context
	}

	feed := &event.Feed{}

	dataChan := make(chan *result)
	dataSub := feed.Subscribe(dataChan)
	defer dataSub.Unsubscribe()

	for _, client := range c.clients {
		go func(client ethpb.BeaconNodeValidator_StreamDutiesClient) {
			feed.Send(&result{
				res: client.Context(),
			})
		}(client)
	}

	res := <-dataChan

	return res.res
}

func (c *streamDuties) SendMsg(m interface{}) error {
	panic("not implemented")
}

func (c *streamDuties) RecvMsg(m interface{}) error {
	panic("not implemented")
}
