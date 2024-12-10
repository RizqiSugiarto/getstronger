// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/v1/workouts.proto

package apiv1

import (
	reflect "reflect"
	sync "sync"

	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateWorkoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutineId    string                 `protobuf:"bytes,1,opt,name=routine_id,json=routineId,proto3" json:"routine_id,omitempty"`
	ExerciseSets []*ExerciseSets        `protobuf:"bytes,2,rep,name=exercise_sets,json=exerciseSets,proto3" json:"exercise_sets,omitempty"`
	StartedAt    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
}

func (x *CreateWorkoutRequest) Reset() {
	*x = CreateWorkoutRequest{}
	mi := &file_api_v1_workouts_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWorkoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkoutRequest) ProtoMessage() {}

func (x *CreateWorkoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_workouts_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkoutRequest.ProtoReflect.Descriptor instead.
func (*CreateWorkoutRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_workouts_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWorkoutRequest) GetRoutineId() string {
	if x != nil {
		return x.RoutineId
	}
	return ""
}

func (x *CreateWorkoutRequest) GetExerciseSets() []*ExerciseSets {
	if x != nil {
		return x.ExerciseSets
	}
	return nil
}

func (x *CreateWorkoutRequest) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *CreateWorkoutRequest) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

type CreateWorkoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkoutId string `protobuf:"bytes,1,opt,name=workout_id,json=workoutId,proto3" json:"workout_id,omitempty"`
}

func (x *CreateWorkoutResponse) Reset() {
	*x = CreateWorkoutResponse{}
	mi := &file_api_v1_workouts_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWorkoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkoutResponse) ProtoMessage() {}

func (x *CreateWorkoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_workouts_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkoutResponse.ProtoReflect.Descriptor instead.
func (*CreateWorkoutResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_workouts_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWorkoutResponse) GetWorkoutId() string {
	if x != nil {
		return x.WorkoutId
	}
	return ""
}

type ListWorkoutsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds   []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	PageSize  int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken []byte   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListWorkoutsRequest) Reset() {
	*x = ListWorkoutsRequest{}
	mi := &file_api_v1_workouts_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListWorkoutsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkoutsRequest) ProtoMessage() {}

func (x *ListWorkoutsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_workouts_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkoutsRequest.ProtoReflect.Descriptor instead.
func (*ListWorkoutsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_workouts_proto_rawDescGZIP(), []int{2}
}

func (x *ListWorkoutsRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *ListWorkoutsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListWorkoutsRequest) GetPageToken() []byte {
	if x != nil {
		return x.PageToken
	}
	return nil
}

type ListWorkoutsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workouts      []*Workout `protobuf:"bytes,1,rep,name=workouts,proto3" json:"workouts,omitempty"`
	NextPageToken []byte     `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListWorkoutsResponse) Reset() {
	*x = ListWorkoutsResponse{}
	mi := &file_api_v1_workouts_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListWorkoutsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkoutsResponse) ProtoMessage() {}

func (x *ListWorkoutsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_workouts_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkoutsResponse.ProtoReflect.Descriptor instead.
func (*ListWorkoutsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_workouts_proto_rawDescGZIP(), []int{3}
}

func (x *ListWorkoutsResponse) GetWorkouts() []*Workout {
	if x != nil {
		return x.Workouts
	}
	return nil
}

func (x *ListWorkoutsResponse) GetNextPageToken() []byte {
	if x != nil {
		return x.NextPageToken
	}
	return nil
}

type GetWorkoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetWorkoutRequest) Reset() {
	*x = GetWorkoutRequest{}
	mi := &file_api_v1_workouts_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorkoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkoutRequest) ProtoMessage() {}

func (x *GetWorkoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_workouts_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkoutRequest.ProtoReflect.Descriptor instead.
func (*GetWorkoutRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_workouts_proto_rawDescGZIP(), []int{4}
}

func (x *GetWorkoutRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetWorkoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workout *Workout `protobuf:"bytes,1,opt,name=workout,proto3" json:"workout,omitempty"`
}

func (x *GetWorkoutResponse) Reset() {
	*x = GetWorkoutResponse{}
	mi := &file_api_v1_workouts_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorkoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkoutResponse) ProtoMessage() {}

func (x *GetWorkoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_workouts_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkoutResponse.ProtoReflect.Descriptor instead.
func (*GetWorkoutResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_workouts_proto_rawDescGZIP(), []int{5}
}

func (x *GetWorkoutResponse) GetWorkout() *Workout {
	if x != nil {
		return x.Workout
	}
	return nil
}

type DeleteWorkoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteWorkoutRequest) Reset() {
	*x = DeleteWorkoutRequest{}
	mi := &file_api_v1_workouts_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteWorkoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkoutRequest) ProtoMessage() {}

func (x *DeleteWorkoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_workouts_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkoutRequest.ProtoReflect.Descriptor instead.
func (*DeleteWorkoutRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_workouts_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteWorkoutRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteWorkoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteWorkoutResponse) Reset() {
	*x = DeleteWorkoutResponse{}
	mi := &file_api_v1_workouts_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteWorkoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkoutResponse) ProtoMessage() {}

func (x *DeleteWorkoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_workouts_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkoutResponse.ProtoReflect.Descriptor instead.
func (*DeleteWorkoutResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_workouts_proto_rawDescGZIP(), []int{7}
}

type PostCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkoutId string `protobuf:"bytes,1,opt,name=workout_id,json=workoutId,proto3" json:"workout_id,omitempty"`
	Comment   string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *PostCommentRequest) Reset() {
	*x = PostCommentRequest{}
	mi := &file_api_v1_workouts_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCommentRequest) ProtoMessage() {}

func (x *PostCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_workouts_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCommentRequest.ProtoReflect.Descriptor instead.
func (*PostCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_workouts_proto_rawDescGZIP(), []int{8}
}

func (x *PostCommentRequest) GetWorkoutId() string {
	if x != nil {
		return x.WorkoutId
	}
	return ""
}

func (x *PostCommentRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type PostCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *WorkoutComment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *PostCommentResponse) Reset() {
	*x = PostCommentResponse{}
	mi := &file_api_v1_workouts_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCommentResponse) ProtoMessage() {}

func (x *PostCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_workouts_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCommentResponse.ProtoReflect.Descriptor instead.
func (*PostCommentResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_workouts_proto_rawDescGZIP(), []int{9}
}

func (x *PostCommentResponse) GetComment() *WorkoutComment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type UpdateWorkoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workout *Workout `protobuf:"bytes,1,opt,name=workout,proto3" json:"workout,omitempty"`
}

func (x *UpdateWorkoutRequest) Reset() {
	*x = UpdateWorkoutRequest{}
	mi := &file_api_v1_workouts_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateWorkoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkoutRequest) ProtoMessage() {}

func (x *UpdateWorkoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_workouts_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkoutRequest.ProtoReflect.Descriptor instead.
func (*UpdateWorkoutRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_workouts_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateWorkoutRequest) GetWorkout() *Workout {
	if x != nil {
		return x.Workout
	}
	return nil
}

type UpdateWorkoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateWorkoutResponse) Reset() {
	*x = UpdateWorkoutResponse{}
	mi := &file_api_v1_workouts_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateWorkoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkoutResponse) ProtoMessage() {}

func (x *UpdateWorkoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_workouts_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkoutResponse.ProtoReflect.Descriptor instead.
func (*UpdateWorkoutResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_workouts_proto_rawDescGZIP(), []int{11}
}

type Workout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	User         *User                  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	ExerciseSets []*ExerciseSets        `protobuf:"bytes,4,rep,name=exercise_sets,json=exerciseSets,proto3" json:"exercise_sets,omitempty"`
	Comments     []*WorkoutComment      `protobuf:"bytes,5,rep,name=comments,proto3" json:"comments,omitempty"`
	StartedAt    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
}

func (x *Workout) Reset() {
	*x = Workout{}
	mi := &file_api_v1_workouts_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Workout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workout) ProtoMessage() {}

func (x *Workout) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_workouts_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workout.ProtoReflect.Descriptor instead.
func (*Workout) Descriptor() ([]byte, []int) {
	return file_api_v1_workouts_proto_rawDescGZIP(), []int{12}
}

func (x *Workout) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Workout) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Workout) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Workout) GetExerciseSets() []*ExerciseSets {
	if x != nil {
		return x.ExerciseSets
	}
	return nil
}

func (x *Workout) GetComments() []*WorkoutComment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *Workout) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Workout) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

type WorkoutComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User      *User                  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Comment   string                 `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *WorkoutComment) Reset() {
	*x = WorkoutComment{}
	mi := &file_api_v1_workouts_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkoutComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkoutComment) ProtoMessage() {}

func (x *WorkoutComment) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_workouts_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkoutComment.ProtoReflect.Descriptor instead.
func (*WorkoutComment) Descriptor() ([]byte, []int) {
	return file_api_v1_workouts_proto_rawDescGZIP(), []int{13}
}

func (x *WorkoutComment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WorkoutComment) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *WorkoutComment) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *WorkoutComment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_api_v1_workouts_proto protoreflect.FileDescriptor

var file_api_v1_workouts_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x14, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x02, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01,
	0x01, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x0d,
	0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73, 0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01,
	0x02, 0x08, 0x01, 0x52, 0x0c, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x53, 0x65, 0x74,
	0x73, 0x12, 0x41, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x43, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x22, 0x36, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x49,
	0x64, 0x22, 0x88, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0f, 0xba, 0x48, 0x0c,
	0x92, 0x01, 0x09, 0x08, 0x01, 0x22, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x26, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xba, 0x48, 0x06, 0x1a, 0x04, 0x18,
	0x64, 0x28, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6b, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2d, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74,
	0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x22, 0x30, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba,
	0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x0a, 0x12, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0a, 0x77, 0x6f,
	0x72, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75,
	0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x47, 0x0a, 0x13, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x49, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x6f,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0xe3, 0x02, 0x0a, 0x07, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x12,
	0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05,
	0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x43, 0x0a, 0x0d, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73, 0x42, 0x08, 0xba,
	0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x53, 0x65, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x43, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x22, 0xba, 0x01, 0x0a, 0x0e, 0x57, 0x6f,
	0x72, 0x6b, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0,
	0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x21, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xd9, 0x03, 0x0a, 0x0e, 0x57, 0x6f, 0x72, 0x6b, 0x6f,
	0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x42, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x47, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x6f, 0x75, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88,
	0xb5, 0x18, 0x01, 0x12, 0x4b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01,
	0x12, 0x4c, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x52,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x12,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5,
	0x18, 0x01, 0x42, 0x8e, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x42, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x72, 0x6c, 0x73, 0x73, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41,
	0x58, 0x58, 0xaa, 0x02, 0x06, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_workouts_proto_rawDescOnce sync.Once
	file_api_v1_workouts_proto_rawDescData = file_api_v1_workouts_proto_rawDesc
)

func file_api_v1_workouts_proto_rawDescGZIP() []byte {
	file_api_v1_workouts_proto_rawDescOnce.Do(func() {
		file_api_v1_workouts_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_workouts_proto_rawDescData)
	})
	return file_api_v1_workouts_proto_rawDescData
}

var file_api_v1_workouts_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_v1_workouts_proto_goTypes = []any{
	(*CreateWorkoutRequest)(nil),  // 0: api.v1.CreateWorkoutRequest
	(*CreateWorkoutResponse)(nil), // 1: api.v1.CreateWorkoutResponse
	(*ListWorkoutsRequest)(nil),   // 2: api.v1.ListWorkoutsRequest
	(*ListWorkoutsResponse)(nil),  // 3: api.v1.ListWorkoutsResponse
	(*GetWorkoutRequest)(nil),     // 4: api.v1.GetWorkoutRequest
	(*GetWorkoutResponse)(nil),    // 5: api.v1.GetWorkoutResponse
	(*DeleteWorkoutRequest)(nil),  // 6: api.v1.DeleteWorkoutRequest
	(*DeleteWorkoutResponse)(nil), // 7: api.v1.DeleteWorkoutResponse
	(*PostCommentRequest)(nil),    // 8: api.v1.PostCommentRequest
	(*PostCommentResponse)(nil),   // 9: api.v1.PostCommentResponse
	(*UpdateWorkoutRequest)(nil),  // 10: api.v1.UpdateWorkoutRequest
	(*UpdateWorkoutResponse)(nil), // 11: api.v1.UpdateWorkoutResponse
	(*Workout)(nil),               // 12: api.v1.Workout
	(*WorkoutComment)(nil),        // 13: api.v1.WorkoutComment
	(*ExerciseSets)(nil),          // 14: api.v1.ExerciseSets
	(*timestamppb.Timestamp)(nil), // 15: google.protobuf.Timestamp
	(*User)(nil),                  // 16: api.v1.User
}
var file_api_v1_workouts_proto_depIdxs = []int32{
	14, // 0: api.v1.CreateWorkoutRequest.exercise_sets:type_name -> api.v1.ExerciseSets
	15, // 1: api.v1.CreateWorkoutRequest.started_at:type_name -> google.protobuf.Timestamp
	15, // 2: api.v1.CreateWorkoutRequest.finished_at:type_name -> google.protobuf.Timestamp
	12, // 3: api.v1.ListWorkoutsResponse.workouts:type_name -> api.v1.Workout
	12, // 4: api.v1.GetWorkoutResponse.workout:type_name -> api.v1.Workout
	13, // 5: api.v1.PostCommentResponse.comment:type_name -> api.v1.WorkoutComment
	12, // 6: api.v1.UpdateWorkoutRequest.workout:type_name -> api.v1.Workout
	16, // 7: api.v1.Workout.user:type_name -> api.v1.User
	14, // 8: api.v1.Workout.exercise_sets:type_name -> api.v1.ExerciseSets
	13, // 9: api.v1.Workout.comments:type_name -> api.v1.WorkoutComment
	15, // 10: api.v1.Workout.started_at:type_name -> google.protobuf.Timestamp
	15, // 11: api.v1.Workout.finished_at:type_name -> google.protobuf.Timestamp
	16, // 12: api.v1.WorkoutComment.user:type_name -> api.v1.User
	15, // 13: api.v1.WorkoutComment.created_at:type_name -> google.protobuf.Timestamp
	0,  // 14: api.v1.WorkoutService.Create:input_type -> api.v1.CreateWorkoutRequest
	4,  // 15: api.v1.WorkoutService.Get:input_type -> api.v1.GetWorkoutRequest
	2,  // 16: api.v1.WorkoutService.List:input_type -> api.v1.ListWorkoutsRequest
	6,  // 17: api.v1.WorkoutService.Delete:input_type -> api.v1.DeleteWorkoutRequest
	8,  // 18: api.v1.WorkoutService.PostComment:input_type -> api.v1.PostCommentRequest
	10, // 19: api.v1.WorkoutService.UpdateWorkout:input_type -> api.v1.UpdateWorkoutRequest
	1,  // 20: api.v1.WorkoutService.Create:output_type -> api.v1.CreateWorkoutResponse
	5,  // 21: api.v1.WorkoutService.Get:output_type -> api.v1.GetWorkoutResponse
	3,  // 22: api.v1.WorkoutService.List:output_type -> api.v1.ListWorkoutsResponse
	7,  // 23: api.v1.WorkoutService.Delete:output_type -> api.v1.DeleteWorkoutResponse
	9,  // 24: api.v1.WorkoutService.PostComment:output_type -> api.v1.PostCommentResponse
	11, // 25: api.v1.WorkoutService.UpdateWorkout:output_type -> api.v1.UpdateWorkoutResponse
	20, // [20:26] is the sub-list for method output_type
	14, // [14:20] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_api_v1_workouts_proto_init() }
func file_api_v1_workouts_proto_init() {
	if File_api_v1_workouts_proto != nil {
		return
	}
	file_api_v1_options_proto_init()
	file_api_v1_shared_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_workouts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_workouts_proto_goTypes,
		DependencyIndexes: file_api_v1_workouts_proto_depIdxs,
		MessageInfos:      file_api_v1_workouts_proto_msgTypes,
	}.Build()
	File_api_v1_workouts_proto = out.File
	file_api_v1_workouts_proto_rawDesc = nil
	file_api_v1_workouts_proto_goTypes = nil
	file_api_v1_workouts_proto_depIdxs = nil
}
