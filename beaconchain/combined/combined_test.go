package combined

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"

	"github.com/begmaroman/beaconspot/beaconchain"
	"github.com/begmaroman/beaconspot/beaconchain/mock"
)

func Test_combined_GetAttestationData(t *testing.T) {
	type fields struct {
		beaconChains []beaconchain.BeaconChain
	}
	type args struct {
		ctx   context.Context
		slot  uint64
		index uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ethpb.AttestationData
		wantErr bool
	}{
		{
			name: "return faster response",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						GetAttestationDataFn: func(ctx context.Context, slot, index uint64) (*ethpb.AttestationData, error) {
							time.Sleep(time.Second / 2)
							return &ethpb.AttestationData{
								Slot: 1,
							}, nil
						},
					},
					&mock.BeaconChain{
						GetAttestationDataFn: func(ctx context.Context, slot, index uint64) (*ethpb.AttestationData, error) {
							time.Sleep(time.Second / 4)
							return &ethpb.AttestationData{
								Slot: 2,
							}, nil
						},
					},
					&mock.BeaconChain{
						GetAttestationDataFn: func(ctx context.Context, slot, index uint64) (*ethpb.AttestationData, error) {
							return &ethpb.AttestationData{
								Slot: 3,
							}, nil
						},
					},
				},
			},
			want: &ethpb.AttestationData{
				Slot: 3,
			},
		},
		{
			name: "return successful response",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						GetAttestationDataFn: func(ctx context.Context, slot, index uint64) (*ethpb.AttestationData, error) {
							return nil, errors.New("test error")
						},
					},
					&mock.BeaconChain{
						GetAttestationDataFn: func(ctx context.Context, slot, index uint64) (*ethpb.AttestationData, error) {
							return &ethpb.AttestationData{
								Slot: 2,
							}, nil
						},
					},
				},
			},
			want: &ethpb.AttestationData{
				Slot: 2,
			},
		},
		{
			name: "return error from all nodes",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						GetAttestationDataFn: func(ctx context.Context, slot, index uint64) (*ethpb.AttestationData, error) {
							return nil, errors.New("test error 1")
						},
					},
					&mock.BeaconChain{
						GetAttestationDataFn: func(ctx context.Context, slot, index uint64) (*ethpb.AttestationData, error) {
							return nil, errors.New("test error 2")
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &combined{
				beaconChains: tt.fields.beaconChains,
			}
			got, err := c.GetAttestationData(tt.args.ctx, tt.args.slot, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAttestationData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAttestationData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combined_ProposeAttestation(t *testing.T) {
	type fields struct {
		beaconChains []beaconchain.BeaconChain
	}
	type args struct {
		ctx             context.Context
		data            *ethpb.AttestationData
		aggregationBits []byte
		signature       []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "return successful response",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						ProposeAttestationFn: func(ctx context.Context, data *ethpb.AttestationData, aggregationBits, signature []byte) error {
							return nil
						},
					},
					&mock.BeaconChain{
						ProposeAttestationFn: func(ctx context.Context, data *ethpb.AttestationData, aggregationBits, signature []byte) error {
							return nil
						},
					},
				},
			},
		},
		{
			name: "return error from all nodes",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						ProposeAttestationFn: func(ctx context.Context, data *ethpb.AttestationData, aggregationBits, signature []byte) error {
							return errors.New("test error")
						},
					},
					&mock.BeaconChain{
						ProposeAttestationFn: func(ctx context.Context, data *ethpb.AttestationData, aggregationBits, signature []byte) error {
							return nil
						},
					},
				},
			},
		},
		{
			name: "return error from all nodes",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						ProposeAttestationFn: func(ctx context.Context, data *ethpb.AttestationData, aggregationBits, signature []byte) error {
							return errors.New("test error")
						},
					},
					&mock.BeaconChain{
						ProposeAttestationFn: func(ctx context.Context, data *ethpb.AttestationData, aggregationBits, signature []byte) error {
							return errors.New("test error")
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &combined{
				beaconChains: tt.fields.beaconChains,
			}
			if err := c.ProposeAttestation(tt.args.ctx, tt.args.data, tt.args.aggregationBits, tt.args.signature); (err != nil) != tt.wantErr {
				t.Errorf("ProposeAttestation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
