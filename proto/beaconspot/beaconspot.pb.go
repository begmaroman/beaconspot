// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: github.com/begmaroman/beaconspot/proto/beaconspot/beaconspot.proto

package beaconspotproto

import (
	context "context"
	reflect "reflect"
	sync "sync"

	health "github.com/begmaroman/beaconspot/proto/health"
	status "github.com/begmaroman/beaconspot/proto/status"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	v1alpha1 "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// GetAttestation operation
type GetAttestationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot           uint64 `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	CommitteeIndex uint64 `protobuf:"varint,2,opt,name=committee_index,json=committeeIndex,proto3" json:"committee_index,omitempty"`
}

func (x *GetAttestationRequest) Reset() {
	*x = GetAttestationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttestationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttestationRequest) ProtoMessage() {}

func (x *GetAttestationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttestationRequest.ProtoReflect.Descriptor instead.
func (*GetAttestationRequest) Descriptor() ([]byte, []int) {
	return file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDescGZIP(), []int{0}
}

func (x *GetAttestationRequest) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *GetAttestationRequest) GetCommitteeIndex() uint64 {
	if x != nil {
		return x.CommitteeIndex
	}
	return 0
}

type GetAttestationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//	*GetAttestationResponse_Error
	//	*GetAttestationResponse_AttestationData
	Result isGetAttestationResponse_Result `protobuf_oneof:"result"`
}

func (x *GetAttestationResponse) Reset() {
	*x = GetAttestationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttestationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttestationResponse) ProtoMessage() {}

func (x *GetAttestationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttestationResponse.ProtoReflect.Descriptor instead.
func (*GetAttestationResponse) Descriptor() ([]byte, []int) {
	return file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDescGZIP(), []int{1}
}

func (m *GetAttestationResponse) GetResult() isGetAttestationResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *GetAttestationResponse) GetError() *status.Status {
	if x, ok := x.GetResult().(*GetAttestationResponse_Error); ok {
		return x.Error
	}
	return nil
}

func (x *GetAttestationResponse) GetAttestationData() *v1alpha1.AttestationData {
	if x, ok := x.GetResult().(*GetAttestationResponse_AttestationData); ok {
		return x.AttestationData
	}
	return nil
}

type isGetAttestationResponse_Result interface {
	isGetAttestationResponse_Result()
}

type GetAttestationResponse_Error struct {
	Error *status.Status `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type GetAttestationResponse_AttestationData struct {
	AttestationData *v1alpha1.AttestationData `protobuf:"bytes,2,opt,name=attestation_data,json=attestationData,proto3,oneof"`
}

func (*GetAttestationResponse_Error) isGetAttestationResponse_Result() {}

func (*GetAttestationResponse_AttestationData) isGetAttestationResponse_Result() {}

// ProposeAttestation operation
type ProposeAttestationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttestationData *v1alpha1.AttestationData `protobuf:"bytes,1,opt,name=attestation_data,json=attestationData,proto3" json:"attestation_data,omitempty"`
	AggregationBits []byte                    `protobuf:"bytes,2,opt,name=aggregation_bits,json=aggregationBits,proto3" json:"aggregation_bits,omitempty"`
	Signature       []byte                    `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ProposeAttestationRequest) Reset() {
	*x = ProposeAttestationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposeAttestationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposeAttestationRequest) ProtoMessage() {}

func (x *ProposeAttestationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposeAttestationRequest.ProtoReflect.Descriptor instead.
func (*ProposeAttestationRequest) Descriptor() ([]byte, []int) {
	return file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDescGZIP(), []int{2}
}

func (x *ProposeAttestationRequest) GetAttestationData() *v1alpha1.AttestationData {
	if x != nil {
		return x.AttestationData
	}
	return nil
}

func (x *ProposeAttestationRequest) GetAggregationBits() []byte {
	if x != nil {
		return x.AggregationBits
	}
	return nil
}

func (x *ProposeAttestationRequest) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ProposeAttestationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//	*ProposeAttestationResponse_Error
	//	*ProposeAttestationResponse_Empty
	Result isProposeAttestationResponse_Result `protobuf_oneof:"result"`
}

func (x *ProposeAttestationResponse) Reset() {
	*x = ProposeAttestationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposeAttestationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposeAttestationResponse) ProtoMessage() {}

func (x *ProposeAttestationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposeAttestationResponse.ProtoReflect.Descriptor instead.
func (*ProposeAttestationResponse) Descriptor() ([]byte, []int) {
	return file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDescGZIP(), []int{3}
}

func (m *ProposeAttestationResponse) GetResult() isProposeAttestationResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ProposeAttestationResponse) GetError() *status.Status {
	if x, ok := x.GetResult().(*ProposeAttestationResponse_Error); ok {
		return x.Error
	}
	return nil
}

func (x *ProposeAttestationResponse) GetEmpty() *empty.Empty {
	if x, ok := x.GetResult().(*ProposeAttestationResponse_Empty); ok {
		return x.Empty
	}
	return nil
}

type isProposeAttestationResponse_Result interface {
	isProposeAttestationResponse_Result()
}

type ProposeAttestationResponse_Error struct {
	Error *status.Status `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type ProposeAttestationResponse_Empty struct {
	Empty *empty.Empty `protobuf:"bytes,2,opt,name=empty,proto3,oneof"`
}

func (*ProposeAttestationResponse_Error) isProposeAttestationResponse_Result() {}

func (*ProposeAttestationResponse_Empty) isProposeAttestationResponse_Result() {}

var File_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto protoreflect.FileDescriptor

var file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x67,
	0x6d, 0x61, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x73, 0x70,
	0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x73,
	0x70, 0x6f, 0x74, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x65, 0x67, 0x6d, 0x61, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x2f, 0x62, 0x65, 0x61,
	0x63, 0x6f, 0x6e, 0x73, 0x70, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65,
	0x67, 0x6d, 0x61, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x73,
	0x70, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x22, 0x98, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x53, 0x0a, 0x10, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x48, 0x00, 0x52, 0x0f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xb7,
	0x01, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x10,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0f,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x29, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62,
	0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x77, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0xd3, 0x02, 0x0a, 0x11, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x53, 0x70, 0x6f, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5d, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x12,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x61, 0x70, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x67, 0x6d, 0x61, 0x72, 0x6f, 0x6d, 0x61, 0x6e,
	0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x73, 0x70, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x73, 0x70, 0x6f, 0x74, 0x3b, 0x62, 0x65, 0x61,
	0x63, 0x6f, 0x6e, 0x73, 0x70, 0x6f, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDescOnce sync.Once
	file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDescData = file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDesc
)

func file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDescGZIP() []byte {
	file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDescOnce.Do(func() {
		file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDescData)
	})
	return file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDescData
}

var file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_goTypes = []interface{}{
	(*GetAttestationRequest)(nil),      // 0: grpcapiproto.GetAttestationRequest
	(*GetAttestationResponse)(nil),     // 1: grpcapiproto.GetAttestationResponse
	(*ProposeAttestationRequest)(nil),  // 2: grpcapiproto.ProposeAttestationRequest
	(*ProposeAttestationResponse)(nil), // 3: grpcapiproto.ProposeAttestationResponse
	(*status.Status)(nil),              // 4: Status
	(*v1alpha1.AttestationData)(nil),   // 5: ethereum.eth.v1alpha1.AttestationData
	(*empty.Empty)(nil),                // 6: google.protobuf.Empty
	(*health.HealthResponse)(nil),      // 7: health.HealthResponse
}
var file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_depIdxs = []int32{
	4, // 0: grpcapiproto.GetAttestationResponse.error:type_name -> Status
	5, // 1: grpcapiproto.GetAttestationResponse.attestation_data:type_name -> ethereum.eth.v1alpha1.AttestationData
	5, // 2: grpcapiproto.ProposeAttestationRequest.attestation_data:type_name -> ethereum.eth.v1alpha1.AttestationData
	4, // 3: grpcapiproto.ProposeAttestationResponse.error:type_name -> Status
	6, // 4: grpcapiproto.ProposeAttestationResponse.empty:type_name -> google.protobuf.Empty
	6, // 5: grpcapiproto.BeaconSpotService.Health:input_type -> google.protobuf.Empty
	6, // 6: grpcapiproto.BeaconSpotService.Ping:input_type -> google.protobuf.Empty
	0, // 7: grpcapiproto.BeaconSpotService.GetAttestation:input_type -> grpcapiproto.GetAttestationRequest
	2, // 8: grpcapiproto.BeaconSpotService.ProposeAttestation:input_type -> grpcapiproto.ProposeAttestationRequest
	7, // 9: grpcapiproto.BeaconSpotService.Health:output_type -> health.HealthResponse
	6, // 10: grpcapiproto.BeaconSpotService.Ping:output_type -> google.protobuf.Empty
	1, // 11: grpcapiproto.BeaconSpotService.GetAttestation:output_type -> grpcapiproto.GetAttestationResponse
	3, // 12: grpcapiproto.BeaconSpotService.ProposeAttestation:output_type -> grpcapiproto.ProposeAttestationResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_init() }
func file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_init() {
	if File_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttestationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttestationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposeAttestationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposeAttestationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GetAttestationResponse_Error)(nil),
		(*GetAttestationResponse_AttestationData)(nil),
	}
	file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ProposeAttestationResponse_Error)(nil),
		(*ProposeAttestationResponse_Empty)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_goTypes,
		DependencyIndexes: file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_depIdxs,
		MessageInfos:      file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_msgTypes,
	}.Build()
	File_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto = out.File
	file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_rawDesc = nil
	file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_goTypes = nil
	file_github_com_begmaroman_beaconspot_proto_beaconspot_beaconspot_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// BeaconSpotServiceClient is the client API for BeaconSpotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BeaconSpotServiceClient interface {
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*health.HealthResponse, error)
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	// Attester
	GetAttestation(ctx context.Context, in *GetAttestationRequest, opts ...grpc.CallOption) (*GetAttestationResponse, error)
	ProposeAttestation(ctx context.Context, in *ProposeAttestationRequest, opts ...grpc.CallOption) (*ProposeAttestationResponse, error)
}

type beaconSpotServiceClient struct {
	cc *grpc.ClientConn
}

func NewBeaconSpotServiceClient(cc *grpc.ClientConn) BeaconSpotServiceClient {
	return &beaconSpotServiceClient{cc}
}

func (c *beaconSpotServiceClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*health.HealthResponse, error) {
	out := new(health.HealthResponse)
	err := c.cc.Invoke(ctx, "/grpcapiproto.BeaconSpotService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconSpotServiceClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpcapiproto.BeaconSpotService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconSpotServiceClient) GetAttestation(ctx context.Context, in *GetAttestationRequest, opts ...grpc.CallOption) (*GetAttestationResponse, error) {
	out := new(GetAttestationResponse)
	err := c.cc.Invoke(ctx, "/grpcapiproto.BeaconSpotService/GetAttestation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconSpotServiceClient) ProposeAttestation(ctx context.Context, in *ProposeAttestationRequest, opts ...grpc.CallOption) (*ProposeAttestationResponse, error) {
	out := new(ProposeAttestationResponse)
	err := c.cc.Invoke(ctx, "/grpcapiproto.BeaconSpotService/ProposeAttestation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeaconSpotServiceServer is the server API for BeaconSpotService service.
type BeaconSpotServiceServer interface {
	Health(context.Context, *empty.Empty) (*health.HealthResponse, error)
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
	// Attester
	GetAttestation(context.Context, *GetAttestationRequest) (*GetAttestationResponse, error)
	ProposeAttestation(context.Context, *ProposeAttestationRequest) (*ProposeAttestationResponse, error)
}

// UnimplementedBeaconSpotServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBeaconSpotServiceServer struct {
}

func (*UnimplementedBeaconSpotServiceServer) Health(context.Context, *empty.Empty) (*health.HealthResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedBeaconSpotServiceServer) Ping(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedBeaconSpotServiceServer) GetAttestation(context.Context, *GetAttestationRequest) (*GetAttestationResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAttestation not implemented")
}
func (*UnimplementedBeaconSpotServiceServer) ProposeAttestation(context.Context, *ProposeAttestationRequest) (*ProposeAttestationResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method ProposeAttestation not implemented")
}

func RegisterBeaconSpotServiceServer(s *grpc.Server, srv BeaconSpotServiceServer) {
	s.RegisterService(&_BeaconSpotService_serviceDesc, srv)
}

func _BeaconSpotService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconSpotServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapiproto.BeaconSpotService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconSpotServiceServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconSpotService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconSpotServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapiproto.BeaconSpotService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconSpotServiceServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconSpotService_GetAttestation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttestationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconSpotServiceServer).GetAttestation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapiproto.BeaconSpotService/GetAttestation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconSpotServiceServer).GetAttestation(ctx, req.(*GetAttestationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconSpotService_ProposeAttestation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposeAttestationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconSpotServiceServer).ProposeAttestation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapiproto.BeaconSpotService/ProposeAttestation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconSpotServiceServer).ProposeAttestation(ctx, req.(*ProposeAttestationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BeaconSpotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapiproto.BeaconSpotService",
	HandlerType: (*BeaconSpotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _BeaconSpotService_Health_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _BeaconSpotService_Ping_Handler,
		},
		{
			MethodName: "GetAttestation",
			Handler:    _BeaconSpotService_GetAttestation_Handler,
		},
		{
			MethodName: "ProposeAttestation",
			Handler:    _BeaconSpotService_ProposeAttestation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/begmaroman/beaconspot/proto/beaconspot/beaconspot.proto",
}
