// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/v1/shared.proto

package apiv1

import (
	reflect "reflect"
	sync "sync"

	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExerciseSets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exercise *Exercise `protobuf:"bytes,1,opt,name=exercise,proto3" json:"exercise,omitempty"`
	Sets     []*Set    `protobuf:"bytes,2,rep,name=sets,proto3" json:"sets,omitempty"`
}

func (x *ExerciseSets) Reset() {
	*x = ExerciseSets{}
	mi := &file_api_v1_shared_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExerciseSets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExerciseSets) ProtoMessage() {}

func (x *ExerciseSets) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_shared_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExerciseSets.ProtoReflect.Descriptor instead.
func (*ExerciseSets) Descriptor() ([]byte, []int) {
	return file_api_v1_shared_proto_rawDescGZIP(), []int{0}
}

func (x *ExerciseSets) GetExercise() *Exercise {
	if x != nil {
		return x.Exercise
	}
	return nil
}

func (x *ExerciseSets) GetSets() []*Set {
	if x != nil {
		return x.Sets
	}
	return nil
}

type Exercise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Label  string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Exercise) Reset() {
	*x = Exercise{}
	mi := &file_api_v1_shared_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Exercise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exercise) ProtoMessage() {}

func (x *Exercise) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_shared_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exercise.ProtoReflect.Descriptor instead.
func (*Exercise) Descriptor() ([]byte, []int) {
	return file_api_v1_shared_proto_rawDescGZIP(), []int{1}
}

func (x *Exercise) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Exercise) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Exercise) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Exercise) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weight   float64      `protobuf:"fixed64,1,opt,name=weight,proto3" json:"weight,omitempty"` // The weight can be less than zero.
	Reps     int32        `protobuf:"varint,2,opt,name=reps,proto3" json:"reps,omitempty"`
	Metadata *MetadataSet `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Set) Reset() {
	*x = Set{}
	mi := &file_api_v1_shared_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_shared_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_api_v1_shared_proto_rawDescGZIP(), []int{2}
}

func (x *Set) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Set) GetReps() int32 {
	if x != nil {
		return x.Reps
	}
	return 0
}

func (x *Set) GetMetadata() *MetadataSet {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type MetadataSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkoutId string                 `protobuf:"bytes,1,opt,name=workout_id,json=workoutId,proto3" json:"workout_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *MetadataSet) Reset() {
	*x = MetadataSet{}
	mi := &file_api_v1_shared_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetadataSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataSet) ProtoMessage() {}

func (x *MetadataSet) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_shared_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataSet.ProtoReflect.Descriptor instead.
func (*MetadataSet) Descriptor() ([]byte, []int) {
	return file_api_v1_shared_proto_rawDescGZIP(), []int{3}
}

func (x *MetadataSet) GetWorkoutId() string {
	if x != nil {
		return x.WorkoutId
	}
	return ""
}

func (x *MetadataSet) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Followed  bool   `protobuf:"varint,4,opt,name=followed,proto3" json:"followed,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_api_v1_shared_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_shared_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_v1_shared_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetFollowed() bool {
	if x != nil {
		return x.Followed
	}
	return false
}

type PaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageLimit int32  `protobuf:"varint,1,opt,name=page_limit,json=pageLimit,proto3" json:"page_limit,omitempty"`
	PageToken []byte `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	mi := &file_api_v1_shared_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_shared_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_shared_proto_rawDescGZIP(), []int{5}
}

func (x *PaginationRequest) GetPageLimit() int32 {
	if x != nil {
		return x.PageLimit
	}
	return 0
}

func (x *PaginationRequest) GetPageToken() []byte {
	if x != nil {
		return x.PageToken
	}
	return nil
}

type PaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalResults  int64  `protobuf:"varint,1,opt,name=total_results,json=totalResults,proto3" json:"total_results,omitempty"`
	NextPageToken []byte `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *PaginationResponse) Reset() {
	*x = PaginationResponse{}
	mi := &file_api_v1_shared_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResponse) ProtoMessage() {}

func (x *PaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_shared_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationResponse.ProtoReflect.Descriptor instead.
func (*PaginationResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_shared_proto_rawDescGZIP(), []int{6}
}

func (x *PaginationResponse) GetTotalResults() int64 {
	if x != nil {
		return x.TotalResults
	}
	return 0
}

func (x *PaginationResponse) GetNextPageToken() []byte {
	if x != nil {
		return x.NextPageToken
	}
	return nil
}

var File_api_v1_shared_proto protoreflect.FileDescriptor

var file_api_v1_shared_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x0c, 0x45,
	0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x08, 0x65,
	0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x04, 0x73, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x08, 0xba, 0x48,
	0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x04, 0x73, 0x65, 0x74, 0x73, 0x22, 0x67, 0x0a, 0x08,
	0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x6b, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a, 0x02, 0x28, 0x01, 0x52, 0x04, 0x72, 0x65, 0x70,
	0x73, 0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x71, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65,
	0x74, 0x12, 0x27, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52,
	0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x22, 0x5c, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xba, 0x48, 0x06,
	0x1a, 0x04, 0x18, 0x64, 0x28, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x61, 0x0a, 0x12, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x42, 0x8c, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x42, 0x0b, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72,
	0x6c, 0x73, 0x73, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x65, 0x72,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58,
	0x58, 0xaa, 0x02, 0x06, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x41, 0x70, 0x69, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_shared_proto_rawDescOnce sync.Once
	file_api_v1_shared_proto_rawDescData = file_api_v1_shared_proto_rawDesc
)

func file_api_v1_shared_proto_rawDescGZIP() []byte {
	file_api_v1_shared_proto_rawDescOnce.Do(func() {
		file_api_v1_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_shared_proto_rawDescData)
	})
	return file_api_v1_shared_proto_rawDescData
}

var file_api_v1_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_v1_shared_proto_goTypes = []any{
	(*ExerciseSets)(nil),          // 0: api.v1.ExerciseSets
	(*Exercise)(nil),              // 1: api.v1.Exercise
	(*Set)(nil),                   // 2: api.v1.Set
	(*MetadataSet)(nil),           // 3: api.v1.MetadataSet
	(*User)(nil),                  // 4: api.v1.User
	(*PaginationRequest)(nil),     // 5: api.v1.PaginationRequest
	(*PaginationResponse)(nil),    // 6: api.v1.PaginationResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_api_v1_shared_proto_depIdxs = []int32{
	1, // 0: api.v1.ExerciseSets.exercise:type_name -> api.v1.Exercise
	2, // 1: api.v1.ExerciseSets.sets:type_name -> api.v1.Set
	3, // 2: api.v1.Set.metadata:type_name -> api.v1.MetadataSet
	7, // 3: api.v1.MetadataSet.created_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1_shared_proto_init() }
func file_api_v1_shared_proto_init() {
	if File_api_v1_shared_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_shared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_shared_proto_goTypes,
		DependencyIndexes: file_api_v1_shared_proto_depIdxs,
		MessageInfos:      file_api_v1_shared_proto_msgTypes,
	}.Build()
	File_api_v1_shared_proto = out.File
	file_api_v1_shared_proto_rawDesc = nil
	file_api_v1_shared_proto_goTypes = nil
	file_api_v1_shared_proto_depIdxs = nil
}
