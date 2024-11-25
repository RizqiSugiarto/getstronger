// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/v1/exercise.proto

package apiv1

import (
	reflect "reflect"
	sync "sync"

	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateExerciseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *CreateExerciseRequest) Reset() {
	*x = CreateExerciseRequest{}
	mi := &file_api_v1_exercise_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateExerciseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExerciseRequest) ProtoMessage() {}

func (x *CreateExerciseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_exercise_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExerciseRequest.ProtoReflect.Descriptor instead.
func (*CreateExerciseRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_exercise_proto_rawDescGZIP(), []int{0}
}

func (x *CreateExerciseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateExerciseRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type CreateExerciseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateExerciseResponse) Reset() {
	*x = CreateExerciseResponse{}
	mi := &file_api_v1_exercise_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateExerciseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExerciseResponse) ProtoMessage() {}

func (x *CreateExerciseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_exercise_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExerciseResponse.ProtoReflect.Descriptor instead.
func (*CreateExerciseResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_exercise_proto_rawDescGZIP(), []int{1}
}

func (x *CreateExerciseResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetExerciseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetExerciseRequest) Reset() {
	*x = GetExerciseRequest{}
	mi := &file_api_v1_exercise_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetExerciseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExerciseRequest) ProtoMessage() {}

func (x *GetExerciseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_exercise_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExerciseRequest.ProtoReflect.Descriptor instead.
func (*GetExerciseRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_exercise_proto_rawDescGZIP(), []int{2}
}

func (x *GetExerciseRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetExerciseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exercise *Exercise `protobuf:"bytes,1,opt,name=exercise,proto3" json:"exercise,omitempty"`
}

func (x *GetExerciseResponse) Reset() {
	*x = GetExerciseResponse{}
	mi := &file_api_v1_exercise_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetExerciseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExerciseResponse) ProtoMessage() {}

func (x *GetExerciseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_exercise_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExerciseResponse.ProtoReflect.Descriptor instead.
func (*GetExerciseResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_exercise_proto_rawDescGZIP(), []int{3}
}

func (x *GetExerciseResponse) GetExercise() *Exercise {
	if x != nil {
		return x.Exercise
	}
	return nil
}

type UpdateExerciseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exercise   *Exercise              `protobuf:"bytes,1,opt,name=exercise,proto3" json:"exercise,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateExerciseRequest) Reset() {
	*x = UpdateExerciseRequest{}
	mi := &file_api_v1_exercise_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateExerciseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExerciseRequest) ProtoMessage() {}

func (x *UpdateExerciseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_exercise_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExerciseRequest.ProtoReflect.Descriptor instead.
func (*UpdateExerciseRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_exercise_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateExerciseRequest) GetExercise() *Exercise {
	if x != nil {
		return x.Exercise
	}
	return nil
}

func (x *UpdateExerciseRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type UpdateExerciseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exercise *Exercise `protobuf:"bytes,1,opt,name=exercise,proto3" json:"exercise,omitempty"`
}

func (x *UpdateExerciseResponse) Reset() {
	*x = UpdateExerciseResponse{}
	mi := &file_api_v1_exercise_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateExerciseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExerciseResponse) ProtoMessage() {}

func (x *UpdateExerciseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_exercise_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExerciseResponse.ProtoReflect.Descriptor instead.
func (*UpdateExerciseResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_exercise_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateExerciseResponse) GetExercise() *Exercise {
	if x != nil {
		return x.Exercise
	}
	return nil
}

type DeleteExerciseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteExerciseRequest) Reset() {
	*x = DeleteExerciseRequest{}
	mi := &file_api_v1_exercise_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteExerciseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExerciseRequest) ProtoMessage() {}

func (x *DeleteExerciseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_exercise_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExerciseRequest.ProtoReflect.Descriptor instead.
func (*DeleteExerciseRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_exercise_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteExerciseRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteExerciseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteExerciseResponse) Reset() {
	*x = DeleteExerciseResponse{}
	mi := &file_api_v1_exercise_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteExerciseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExerciseResponse) ProtoMessage() {}

func (x *DeleteExerciseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_exercise_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExerciseResponse.ProtoReflect.Descriptor instead.
func (*DeleteExerciseResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_exercise_proto_rawDescGZIP(), []int{7}
}

type ListExercisesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ExerciseIds []string `protobuf:"bytes,2,rep,name=exercise_ids,json=exerciseIds,proto3" json:"exercise_ids,omitempty"`
	PageSize    int32    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken   []byte   `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListExercisesRequest) Reset() {
	*x = ListExercisesRequest{}
	mi := &file_api_v1_exercise_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListExercisesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExercisesRequest) ProtoMessage() {}

func (x *ListExercisesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_exercise_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExercisesRequest.ProtoReflect.Descriptor instead.
func (*ListExercisesRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_exercise_proto_rawDescGZIP(), []int{8}
}

func (x *ListExercisesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListExercisesRequest) GetExerciseIds() []string {
	if x != nil {
		return x.ExerciseIds
	}
	return nil
}

func (x *ListExercisesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListExercisesRequest) GetPageToken() []byte {
	if x != nil {
		return x.PageToken
	}
	return nil
}

type ListExercisesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exercises     []*Exercise `protobuf:"bytes,1,rep,name=exercises,proto3" json:"exercises,omitempty"`
	NextPageToken []byte      `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListExercisesResponse) Reset() {
	*x = ListExercisesResponse{}
	mi := &file_api_v1_exercise_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListExercisesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExercisesResponse) ProtoMessage() {}

func (x *ListExercisesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_exercise_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExercisesResponse.ProtoReflect.Descriptor instead.
func (*ListExercisesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_exercise_proto_rawDescGZIP(), []int{9}
}

func (x *ListExercisesResponse) GetExercises() []*Exercise {
	if x != nil {
		return x.Exercises
	}
	return nil
}

func (x *ListExercisesResponse) GetNextPageToken() []byte {
	if x != nil {
		return x.NextPageToken
	}
	return nil
}

type GetPreviousWorkoutSetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExerciseIds []string `protobuf:"bytes,1,rep,name=exercise_ids,json=exerciseIds,proto3" json:"exercise_ids,omitempty"`
}

func (x *GetPreviousWorkoutSetsRequest) Reset() {
	*x = GetPreviousWorkoutSetsRequest{}
	mi := &file_api_v1_exercise_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPreviousWorkoutSetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPreviousWorkoutSetsRequest) ProtoMessage() {}

func (x *GetPreviousWorkoutSetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_exercise_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPreviousWorkoutSetsRequest.ProtoReflect.Descriptor instead.
func (*GetPreviousWorkoutSetsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_exercise_proto_rawDescGZIP(), []int{10}
}

func (x *GetPreviousWorkoutSetsRequest) GetExerciseIds() []string {
	if x != nil {
		return x.ExerciseIds
	}
	return nil
}

type GetPreviousWorkoutSetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExerciseSets []*ExerciseSets `protobuf:"bytes,1,rep,name=exercise_sets,json=exerciseSets,proto3" json:"exercise_sets,omitempty"`
}

func (x *GetPreviousWorkoutSetsResponse) Reset() {
	*x = GetPreviousWorkoutSetsResponse{}
	mi := &file_api_v1_exercise_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPreviousWorkoutSetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPreviousWorkoutSetsResponse) ProtoMessage() {}

func (x *GetPreviousWorkoutSetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_exercise_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPreviousWorkoutSetsResponse.ProtoReflect.Descriptor instead.
func (*GetPreviousWorkoutSetsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_exercise_proto_rawDescGZIP(), []int{11}
}

func (x *GetPreviousWorkoutSetsResponse) GetExerciseSets() []*ExerciseSets {
	if x != nil {
		return x.ExerciseSets
	}
	return nil
}

type Exercise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Label string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Exercise) Reset() {
	*x = Exercise{}
	mi := &file_api_v1_exercise_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Exercise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exercise) ProtoMessage() {}

func (x *Exercise) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_exercise_proto_msgTypes[12]
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
	return file_api_v1_exercise_proto_rawDescGZIP(), []int{12}
}

func (x *Exercise) GetId() string {
	if x != nil {
		return x.Id
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

var File_api_v1_exercise_proto protoreflect.FileDescriptor

var file_api_v1_exercise_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x28, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x43, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x08, 0x65, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x34, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x65, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x73, 0x6b, 0x22, 0x46, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08,
	0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x52, 0x08, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x22, 0x31, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x18, 0x0a,
	0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0x92,
	0x01, 0x07, 0x22, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x49, 0x64, 0x73, 0x12, 0x26, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xba, 0x48, 0x06, 0x1a, 0x04,
	0x18, 0x64, 0x28, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6f, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x09, 0x65, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x51,
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x57, 0x6f, 0x72,
	0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x49, 0x64,
	0x73, 0x22, 0x5b, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x5f,
	0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73,
	0x52, 0x0c, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73, 0x22, 0x44,
	0x0a, 0x08, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x32, 0xfe, 0x03, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x4d, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x4d, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x49, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x6d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x74, 0x73,
	0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x57, 0x6f, 0x72, 0x6b,
	0x6f, 0x75, 0x74, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x04, 0x88, 0xb5, 0x18, 0x01, 0x42, 0x8e, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x72, 0x6c, 0x73, 0x73, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x73, 0x74, 0x72, 0x6f,
	0x6e, 0x67, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x06, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x41,
	0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_exercise_proto_rawDescOnce sync.Once
	file_api_v1_exercise_proto_rawDescData = file_api_v1_exercise_proto_rawDesc
)

func file_api_v1_exercise_proto_rawDescGZIP() []byte {
	file_api_v1_exercise_proto_rawDescOnce.Do(func() {
		file_api_v1_exercise_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_exercise_proto_rawDescData)
	})
	return file_api_v1_exercise_proto_rawDescData
}

var file_api_v1_exercise_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_v1_exercise_proto_goTypes = []any{
	(*CreateExerciseRequest)(nil),          // 0: api.v1.CreateExerciseRequest
	(*CreateExerciseResponse)(nil),         // 1: api.v1.CreateExerciseResponse
	(*GetExerciseRequest)(nil),             // 2: api.v1.GetExerciseRequest
	(*GetExerciseResponse)(nil),            // 3: api.v1.GetExerciseResponse
	(*UpdateExerciseRequest)(nil),          // 4: api.v1.UpdateExerciseRequest
	(*UpdateExerciseResponse)(nil),         // 5: api.v1.UpdateExerciseResponse
	(*DeleteExerciseRequest)(nil),          // 6: api.v1.DeleteExerciseRequest
	(*DeleteExerciseResponse)(nil),         // 7: api.v1.DeleteExerciseResponse
	(*ListExercisesRequest)(nil),           // 8: api.v1.ListExercisesRequest
	(*ListExercisesResponse)(nil),          // 9: api.v1.ListExercisesResponse
	(*GetPreviousWorkoutSetsRequest)(nil),  // 10: api.v1.GetPreviousWorkoutSetsRequest
	(*GetPreviousWorkoutSetsResponse)(nil), // 11: api.v1.GetPreviousWorkoutSetsResponse
	(*Exercise)(nil),                       // 12: api.v1.Exercise
	(*fieldmaskpb.FieldMask)(nil),          // 13: google.protobuf.FieldMask
	(*ExerciseSets)(nil),                   // 14: api.v1.ExerciseSets
}
var file_api_v1_exercise_proto_depIdxs = []int32{
	12, // 0: api.v1.GetExerciseResponse.exercise:type_name -> api.v1.Exercise
	12, // 1: api.v1.UpdateExerciseRequest.exercise:type_name -> api.v1.Exercise
	13, // 2: api.v1.UpdateExerciseRequest.update_mask:type_name -> google.protobuf.FieldMask
	12, // 3: api.v1.UpdateExerciseResponse.exercise:type_name -> api.v1.Exercise
	12, // 4: api.v1.ListExercisesResponse.exercises:type_name -> api.v1.Exercise
	14, // 5: api.v1.GetPreviousWorkoutSetsResponse.exercise_sets:type_name -> api.v1.ExerciseSets
	0,  // 6: api.v1.ExerciseService.Create:input_type -> api.v1.CreateExerciseRequest
	2,  // 7: api.v1.ExerciseService.Get:input_type -> api.v1.GetExerciseRequest
	4,  // 8: api.v1.ExerciseService.Update:input_type -> api.v1.UpdateExerciseRequest
	6,  // 9: api.v1.ExerciseService.Delete:input_type -> api.v1.DeleteExerciseRequest
	8,  // 10: api.v1.ExerciseService.List:input_type -> api.v1.ListExercisesRequest
	10, // 11: api.v1.ExerciseService.GetPreviousWorkoutSets:input_type -> api.v1.GetPreviousWorkoutSetsRequest
	1,  // 12: api.v1.ExerciseService.Create:output_type -> api.v1.CreateExerciseResponse
	3,  // 13: api.v1.ExerciseService.Get:output_type -> api.v1.GetExerciseResponse
	5,  // 14: api.v1.ExerciseService.Update:output_type -> api.v1.UpdateExerciseResponse
	7,  // 15: api.v1.ExerciseService.Delete:output_type -> api.v1.DeleteExerciseResponse
	9,  // 16: api.v1.ExerciseService.List:output_type -> api.v1.ListExercisesResponse
	11, // 17: api.v1.ExerciseService.GetPreviousWorkoutSets:output_type -> api.v1.GetPreviousWorkoutSetsResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_v1_exercise_proto_init() }
func file_api_v1_exercise_proto_init() {
	if File_api_v1_exercise_proto != nil {
		return
	}
	file_api_v1_shared_proto_init()
	file_api_v1_options_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_exercise_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_exercise_proto_goTypes,
		DependencyIndexes: file_api_v1_exercise_proto_depIdxs,
		MessageInfos:      file_api_v1_exercise_proto_msgTypes,
	}.Build()
	File_api_v1_exercise_proto = out.File
	file_api_v1_exercise_proto_rawDesc = nil
	file_api_v1_exercise_proto_goTypes = nil
	file_api_v1_exercise_proto_depIdxs = nil
}
