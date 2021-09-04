package combined

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	types "github.com/prysmaticlabs/eth2-types"
	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"

	"github.com/begmaroman/beaconspot/beaconchain"
	"github.com/begmaroman/beaconspot/beaconchain/mock"
)

func Test_combined_GetAttestationData(t *testing.T) {
	type fields struct {
		beaconChains []beaconchain.BeaconChain
	}
	type args struct {
		ctx   context.Context
		slot  types.Slot
		index types.CommitteeIndex
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
						GetAttestationDataFn: func(ctx context.Context, slot types.Slot, index types.CommitteeIndex) (*ethpb.AttestationData, error) {
							time.Sleep(time.Second / 2)
							return &ethpb.AttestationData{
								Slot: 1,
							}, nil
						},
					},
					&mock.BeaconChain{
						GetAttestationDataFn: func(ctx context.Context, slot types.Slot, index types.CommitteeIndex) (*ethpb.AttestationData, error) {
							time.Sleep(time.Second / 4)
							return &ethpb.AttestationData{
								Slot: 2,
							}, nil
						},
					},
					&mock.BeaconChain{
						GetAttestationDataFn: func(ctx context.Context, slot types.Slot, index types.CommitteeIndex) (*ethpb.AttestationData, error) {
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
						GetAttestationDataFn: func(ctx context.Context, slot types.Slot, index types.CommitteeIndex) (*ethpb.AttestationData, error) {
							return nil, errors.New("test error")
						},
					},
					&mock.BeaconChain{
						GetAttestationDataFn: func(ctx context.Context, slot types.Slot, index types.CommitteeIndex) (*ethpb.AttestationData, error) {
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
						GetAttestationDataFn: func(ctx context.Context, slot types.Slot, index types.CommitteeIndex) (*ethpb.AttestationData, error) {
							return nil, errors.New("test error 1")
						},
					},
					&mock.BeaconChain{
						GetAttestationDataFn: func(ctx context.Context, slot types.Slot, index types.CommitteeIndex) (*ethpb.AttestationData, error) {
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

func Test_combined_GetBlock(t *testing.T) {
	type fields struct {
		beaconChains []beaconchain.BeaconChain
	}
	type args struct {
		ctx          context.Context
		slot         types.Slot
		randaoReveal []byte
		graffiti     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ethpb.BeaconBlock
		wantErr bool
	}{
		{
			name: "return faster response",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						GetBlockFn: func(ctx context.Context, slot types.Slot, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error) {
							time.Sleep(time.Second / 2)
							return &ethpb.BeaconBlock{
								Slot: 1,
							}, nil
						},
					},
					&mock.BeaconChain{
						GetBlockFn: func(ctx context.Context, slot types.Slot, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error) {
							time.Sleep(time.Second / 4)
							return &ethpb.BeaconBlock{
								Slot: 2,
							}, nil
						},
					},
					&mock.BeaconChain{
						GetBlockFn: func(ctx context.Context, slot types.Slot, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error) {
							return &ethpb.BeaconBlock{
								Slot: 3,
							}, nil
						},
					},
				},
			},
			want: &ethpb.BeaconBlock{
				Slot: 3,
			},
		},
		{
			name: "return successful response",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						GetBlockFn: func(ctx context.Context, slot types.Slot, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error) {
							return nil, errors.New("test error")
						},
					},
					&mock.BeaconChain{
						GetBlockFn: func(ctx context.Context, slot types.Slot, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error) {
							return &ethpb.BeaconBlock{
								Slot: 2,
							}, nil
						},
					},
				},
			},
			want: &ethpb.BeaconBlock{
				Slot: 2,
			},
		},
		{
			name: "return error from all nodes",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						GetBlockFn: func(ctx context.Context, slot types.Slot, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error) {
							return nil, errors.New("test error 1")
						},
					},
					&mock.BeaconChain{
						GetBlockFn: func(ctx context.Context, slot types.Slot, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error) {
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
			got, err := c.GetBlock(tt.args.ctx, tt.args.slot, tt.args.randaoReveal, tt.args.graffiti)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlock() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combined_ProposeBlock(t *testing.T) {
	type fields struct {
		beaconChains []beaconchain.BeaconChain
	}
	type args struct {
		ctx       context.Context
		signature []byte
		block     *ethpb.BeaconBlock
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
						ProposeBlockFn: func(ctx context.Context, signature []byte, block *ethpb.BeaconBlock) error {
							return nil
						},
					},
					&mock.BeaconChain{
						ProposeBlockFn: func(ctx context.Context, signature []byte, block *ethpb.BeaconBlock) error {
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
						ProposeBlockFn: func(ctx context.Context, signature []byte, block *ethpb.BeaconBlock) error {
							return errors.New("test error")
						},
					},
					&mock.BeaconChain{
						ProposeBlockFn: func(ctx context.Context, signature []byte, block *ethpb.BeaconBlock) error {
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
						ProposeBlockFn: func(ctx context.Context, signature []byte, block *ethpb.BeaconBlock) error {
							return errors.New("test error")
						},
					},
					&mock.BeaconChain{
						ProposeBlockFn: func(ctx context.Context, signature []byte, block *ethpb.BeaconBlock) error {
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
			if err := c.ProposeBlock(tt.args.ctx, tt.args.signature, tt.args.block); (err != nil) != tt.wantErr {
				t.Errorf("ProposeBlock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_combined_GetAggregateSelectionProof(t *testing.T) {
	type fields struct {
		beaconChains []beaconchain.BeaconChain
	}
	type args struct {
		ctx            context.Context
		slot           types.Slot
		committeeIndex types.CommitteeIndex
		publicKey      []byte
		sig            []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ethpb.AggregateAttestationAndProof
		wantErr bool
	}{
		{
			name: "return faster response",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						GetAggregateSelectionProofFn: func(ctx context.Context, slot types.Slot, committeeIndex types.CommitteeIndex, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error) {
							time.Sleep(time.Second / 2)
							return &ethpb.AggregateAttestationAndProof{
								AggregatorIndex: 1,
							}, nil
						},
					},
					&mock.BeaconChain{
						GetAggregateSelectionProofFn: func(ctx context.Context, slot types.Slot, committeeIndex types.CommitteeIndex, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error) {
							time.Sleep(time.Second / 4)
							return &ethpb.AggregateAttestationAndProof{
								AggregatorIndex: 2,
							}, nil
						},
					},
					&mock.BeaconChain{
						GetAggregateSelectionProofFn: func(ctx context.Context, slot types.Slot, committeeIndex types.CommitteeIndex, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error) {
							return &ethpb.AggregateAttestationAndProof{
								AggregatorIndex: 3,
							}, nil
						},
					},
				},
			},
			want: &ethpb.AggregateAttestationAndProof{
				AggregatorIndex: 3,
			},
		},
		{
			name: "return successful response",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						GetAggregateSelectionProofFn: func(ctx context.Context, slot types.Slot, committeeIndex types.CommitteeIndex, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error) {
							return nil, errors.New("test error")
						},
					},
					&mock.BeaconChain{
						GetAggregateSelectionProofFn: func(ctx context.Context, slot types.Slot, committeeIndex types.CommitteeIndex, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error) {
							return &ethpb.AggregateAttestationAndProof{
								AggregatorIndex: 2,
							}, nil
						},
					},
				},
			},
			want: &ethpb.AggregateAttestationAndProof{
				AggregatorIndex: 2,
			},
		},
		{
			name: "return error from all nodes",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						GetAggregateSelectionProofFn: func(ctx context.Context, slot types.Slot, committeeIndex types.CommitteeIndex, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error) {
							return nil, errors.New("test error 1")
						},
					},
					&mock.BeaconChain{
						GetAggregateSelectionProofFn: func(ctx context.Context, slot types.Slot, committeeIndex types.CommitteeIndex, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error) {
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
			got, err := c.GetAggregateSelectionProof(tt.args.ctx, tt.args.slot, tt.args.committeeIndex, tt.args.publicKey, tt.args.sig)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAggregateSelectionProof() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAggregateSelectionProof() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combined_SubmitSignedAggregateSelectionProof(t *testing.T) {
	type fields struct {
		beaconChains []beaconchain.BeaconChain
	}
	type args struct {
		ctx       context.Context
		signature []byte
		message   *ethpb.AggregateAttestationAndProof
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
						SubmitSignedAggregateSelectionProofFn: func(ctx context.Context, signature []byte, message *ethpb.AggregateAttestationAndProof) error {
							return nil
						},
					},
					&mock.BeaconChain{
						SubmitSignedAggregateSelectionProofFn: func(ctx context.Context, signature []byte, message *ethpb.AggregateAttestationAndProof) error {
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
						SubmitSignedAggregateSelectionProofFn: func(ctx context.Context, signature []byte, message *ethpb.AggregateAttestationAndProof) error {
							return errors.New("test error")
						},
					},
					&mock.BeaconChain{
						SubmitSignedAggregateSelectionProofFn: func(ctx context.Context, signature []byte, message *ethpb.AggregateAttestationAndProof) error {
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
						SubmitSignedAggregateSelectionProofFn: func(ctx context.Context, signature []byte, message *ethpb.AggregateAttestationAndProof) error {
							return errors.New("test error")
						},
					},
					&mock.BeaconChain{
						SubmitSignedAggregateSelectionProofFn: func(ctx context.Context, signature []byte, message *ethpb.AggregateAttestationAndProof) error {
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
			if err := c.SubmitSignedAggregateSelectionProof(tt.args.ctx, tt.args.signature, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("SubmitSignedAggregateSelectionProof() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_combined_SubnetsSubscribe(t *testing.T) {
	type fields struct {
		beaconChains []beaconchain.BeaconChain
	}
	type args struct {
		ctx           context.Context
		subscriptions []beaconchain.SubnetSubscription
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
						SubnetsSubscribeFn: func(ctx context.Context, subscriptions []beaconchain.SubnetSubscription) error {
							return nil
						},
					},
					&mock.BeaconChain{
						SubnetsSubscribeFn: func(ctx context.Context, subscriptions []beaconchain.SubnetSubscription) error {
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
						SubnetsSubscribeFn: func(ctx context.Context, subscriptions []beaconchain.SubnetSubscription) error {
							return errors.New("test error")
						},
					},
					&mock.BeaconChain{
						SubnetsSubscribeFn: func(ctx context.Context, subscriptions []beaconchain.SubnetSubscription) error {
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
						SubnetsSubscribeFn: func(ctx context.Context, subscriptions []beaconchain.SubnetSubscription) error {
							return errors.New("test error")
						},
					},
					&mock.BeaconChain{
						SubnetsSubscribeFn: func(ctx context.Context, subscriptions []beaconchain.SubnetSubscription) error {
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
			if err := c.SubnetsSubscribe(tt.args.ctx, tt.args.subscriptions); (err != nil) != tt.wantErr {
				t.Errorf("SubnetsSubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_combined_DomainData(t *testing.T) {
	type fields struct {
		beaconChains []beaconchain.BeaconChain
	}
	type args struct {
		ctx    context.Context
		epoch  types.Epoch
		domain []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "return faster response",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						DomainDataFn: func(ctx context.Context, epoch types.Epoch, domain []byte) ([]byte, error) {
							time.Sleep(time.Second / 2)
							return []byte("1"), nil
						},
					},
					&mock.BeaconChain{
						DomainDataFn: func(ctx context.Context, epoch types.Epoch, domain []byte) ([]byte, error) {
							time.Sleep(time.Second / 4)
							return []byte("2"), nil
						},
					},
					&mock.BeaconChain{
						DomainDataFn: func(ctx context.Context, epoch types.Epoch, domain []byte) ([]byte, error) {
							return []byte("3"), nil
						},
					},
				},
			},
			want: []byte("3"),
		},
		{
			name: "return successful response",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						DomainDataFn: func(ctx context.Context, epoch types.Epoch, domain []byte) ([]byte, error) {
							return nil, errors.New("test error")
						},
					},
					&mock.BeaconChain{
						DomainDataFn: func(ctx context.Context, epoch types.Epoch, domain []byte) ([]byte, error) {
							return []byte("2"), nil
						},
					},
				},
			},
			want: []byte("2"),
		},
		{
			name: "return error from all nodes",
			fields: fields{
				beaconChains: []beaconchain.BeaconChain{
					&mock.BeaconChain{
						DomainDataFn: func(ctx context.Context, epoch types.Epoch, domain []byte) ([]byte, error) {
							return nil, errors.New("test error 1")
						},
					},
					&mock.BeaconChain{
						DomainDataFn: func(ctx context.Context, epoch types.Epoch, domain []byte) ([]byte, error) {
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
			got, err := c.DomainData(tt.args.ctx, tt.args.epoch, tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("DomainData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DomainData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
