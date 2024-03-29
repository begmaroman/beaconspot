// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: github.com/begmaroman/beaconspot/proto/common/error_response.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status  string         `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Details []*ErrorDetail `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
	Message string         `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Error) GetDetails() []*ErrorDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ErrorDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields  []string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorDetail) Reset() {
	*x = ErrorDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetail) ProtoMessage() {}

func (x *ErrorDetail) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetail.ProtoReflect.Descriptor instead.
func (*ErrorDetail) Descriptor() ([]byte, []int) {
	return file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorDetail) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *ErrorDetail) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_github_com_begmaroman_beaconspot_proto_common_error_response_proto protoreflect.FileDescriptor

var file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x67,
	0x6d, 0x61, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x73, 0x70,
	0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x75, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3f, 0x0a, 0x0b, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x67, 0x6d, 0x61, 0x72,
	0x6f, 0x6d, 0x61, 0x6e, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x73, 0x70, 0x6f, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_rawDescOnce sync.Once
	file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_rawDescData = file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_rawDesc
)

func file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_rawDescGZIP() []byte {
	file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_rawDescOnce.Do(func() {
		file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_rawDescData)
	})
	return file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_rawDescData
}

var file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_goTypes = []interface{}{
	(*ErrorResponse)(nil), // 0: ErrorResponse
	(*Error)(nil),         // 1: Error
	(*ErrorDetail)(nil),   // 2: ErrorDetail
}
var file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_depIdxs = []int32{
	1, // 0: ErrorResponse.error:type_name -> Error
	2, // 1: Error.details:type_name -> ErrorDetail
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_init() }
func file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_init() {
	if File_github_com_begmaroman_beaconspot_proto_common_error_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
		file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetail); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_goTypes,
		DependencyIndexes: file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_depIdxs,
		MessageInfos:      file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_msgTypes,
	}.Build()
	File_github_com_begmaroman_beaconspot_proto_common_error_response_proto = out.File
	file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_rawDesc = nil
	file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_goTypes = nil
	file_github_com_begmaroman_beaconspot_proto_common_error_response_proto_depIdxs = nil
}
