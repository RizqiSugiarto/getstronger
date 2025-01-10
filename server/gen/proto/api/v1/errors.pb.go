// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: api/v1/errors.proto

package apiv1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Error int32

const (
	Error_ERROR_UNSPECIFIED            Error = 0
	Error_ERROR_EMAIL_NOT_VERIFIED     Error = 1
	Error_ERROR_PASSWORDS_DO_NOT_MATCH Error = 2
)

// Enum value maps for Error.
var (
	Error_name = map[int32]string{
		0: "ERROR_UNSPECIFIED",
		1: "ERROR_EMAIL_NOT_VERIFIED",
		2: "ERROR_PASSWORDS_DO_NOT_MATCH",
	}
	Error_value = map[string]int32{
		"ERROR_UNSPECIFIED":            0,
		"ERROR_EMAIL_NOT_VERIFIED":     1,
		"ERROR_PASSWORDS_DO_NOT_MATCH": 2,
	}
)

func (x Error) Enum() *Error {
	p := new(Error)
	*p = x
	return p
}

func (x Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_errors_proto_enumTypes[0].Descriptor()
}

func (Error) Type() protoreflect.EnumType {
	return &file_api_v1_errors_proto_enumTypes[0]
}

func (x Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error.Descriptor instead.
func (Error) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_errors_proto_rawDescGZIP(), []int{0}
}

type ErrorDetail struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Error         Error                  `protobuf:"varint,1,opt,name=error,proto3,enum=api.v1.Error" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ErrorDetail) Reset() {
	*x = ErrorDetail{}
	mi := &file_api_v1_errors_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrorDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetail) ProtoMessage() {}

func (x *ErrorDetail) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_errors_proto_msgTypes[0]
	if x != nil {
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
	return file_api_v1_errors_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorDetail) GetError() Error {
	if x != nil {
		return x.Error
	}
	return Error_ERROR_UNSPECIFIED
}

var File_api_v1_errors_proto protoreflect.FileDescriptor

var file_api_v1_errors_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x32, 0x0a,
	0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2a, 0x5e, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x20, 0x0a, 0x1c, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52,
	0x44, 0x53, 0x5f, 0x44, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10,
	0x02, 0x42, 0x8f, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x42, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x6c, 0x73,
	0x73, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x65, 0x72, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41,
	0x58, 0x58, 0xaa, 0x02, 0x06, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_errors_proto_rawDescOnce sync.Once
	file_api_v1_errors_proto_rawDescData = file_api_v1_errors_proto_rawDesc
)

func file_api_v1_errors_proto_rawDescGZIP() []byte {
	file_api_v1_errors_proto_rawDescOnce.Do(func() {
		file_api_v1_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_errors_proto_rawDescData)
	})
	return file_api_v1_errors_proto_rawDescData
}

var file_api_v1_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1_errors_proto_goTypes = []any{
	(Error)(0),          // 0: api.v1.Error
	(*ErrorDetail)(nil), // 1: api.v1.ErrorDetail
}
var file_api_v1_errors_proto_depIdxs = []int32{
	0, // 0: api.v1.ErrorDetail.error:type_name -> api.v1.Error
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_errors_proto_init() }
func file_api_v1_errors_proto_init() {
	if File_api_v1_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_errors_proto_goTypes,
		DependencyIndexes: file_api_v1_errors_proto_depIdxs,
		EnumInfos:         file_api_v1_errors_proto_enumTypes,
		MessageInfos:      file_api_v1_errors_proto_msgTypes,
	}.Build()
	File_api_v1_errors_proto = out.File
	file_api_v1_errors_proto_rawDesc = nil
	file_api_v1_errors_proto_goTypes = nil
	file_api_v1_errors_proto_depIdxs = nil
}
