// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/v1/users.proto

package apiv1

import (
	reflect "reflect"
	sync "sync"

	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_api_v1_users_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	mi := &file_api_v1_users_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type FollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowId string `protobuf:"bytes,1,opt,name=follow_id,json=followId,proto3" json:"follow_id,omitempty"`
}

func (x *FollowRequest) Reset() {
	*x = FollowRequest{}
	mi := &file_api_v1_users_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequest) ProtoMessage() {}

func (x *FollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequest.ProtoReflect.Descriptor instead.
func (*FollowRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{2}
}

func (x *FollowRequest) GetFollowId() string {
	if x != nil {
		return x.FollowId
	}
	return ""
}

type FollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FollowResponse) Reset() {
	*x = FollowResponse{}
	mi := &file_api_v1_users_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowResponse) ProtoMessage() {}

func (x *FollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowResponse.ProtoReflect.Descriptor instead.
func (*FollowResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{3}
}

type UnfollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnfollowId string `protobuf:"bytes,1,opt,name=unfollow_id,json=unfollowId,proto3" json:"unfollow_id,omitempty"`
}

func (x *UnfollowRequest) Reset() {
	*x = UnfollowRequest{}
	mi := &file_api_v1_users_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnfollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfollowRequest) ProtoMessage() {}

func (x *UnfollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfollowRequest.ProtoReflect.Descriptor instead.
func (*UnfollowRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{4}
}

func (x *UnfollowRequest) GetUnfollowId() string {
	if x != nil {
		return x.UnfollowId
	}
	return ""
}

type UnfollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnfollowResponse) Reset() {
	*x = UnfollowResponse{}
	mi := &file_api_v1_users_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnfollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfollowResponse) ProtoMessage() {}

func (x *UnfollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfollowResponse.ProtoReflect.Descriptor instead.
func (*UnfollowResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{5}
}

type ListFollowersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerId string `protobuf:"bytes,1,opt,name=follower_id,json=followerId,proto3" json:"follower_id,omitempty"`
}

func (x *ListFollowersRequest) Reset() {
	*x = ListFollowersRequest{}
	mi := &file_api_v1_users_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFollowersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowersRequest) ProtoMessage() {}

func (x *ListFollowersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowersRequest.ProtoReflect.Descriptor instead.
func (*ListFollowersRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{6}
}

func (x *ListFollowersRequest) GetFollowerId() string {
	if x != nil {
		return x.FollowerId
	}
	return ""
}

type ListFollowersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followers []*User `protobuf:"bytes,1,rep,name=followers,proto3" json:"followers,omitempty"`
}

func (x *ListFollowersResponse) Reset() {
	*x = ListFollowersResponse{}
	mi := &file_api_v1_users_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFollowersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowersResponse) ProtoMessage() {}

func (x *ListFollowersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowersResponse.ProtoReflect.Descriptor instead.
func (*ListFollowersResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{7}
}

func (x *ListFollowersResponse) GetFollowers() []*User {
	if x != nil {
		return x.Followers
	}
	return nil
}

type ListFolloweesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FolloweeId string `protobuf:"bytes,1,opt,name=followee_id,json=followeeId,proto3" json:"followee_id,omitempty"`
}

func (x *ListFolloweesRequest) Reset() {
	*x = ListFolloweesRequest{}
	mi := &file_api_v1_users_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFolloweesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFolloweesRequest) ProtoMessage() {}

func (x *ListFolloweesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFolloweesRequest.ProtoReflect.Descriptor instead.
func (*ListFolloweesRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{8}
}

func (x *ListFolloweesRequest) GetFolloweeId() string {
	if x != nil {
		return x.FolloweeId
	}
	return ""
}

type ListFolloweesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followees []*User `protobuf:"bytes,1,rep,name=followees,proto3" json:"followees,omitempty"`
}

func (x *ListFolloweesResponse) Reset() {
	*x = ListFolloweesResponse{}
	mi := &file_api_v1_users_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFolloweesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFolloweesResponse) ProtoMessage() {}

func (x *ListFolloweesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFolloweesResponse.ProtoReflect.Descriptor instead.
func (*ListFolloweesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{9}
}

func (x *ListFolloweesResponse) GetFollowees() []*User {
	if x != nil {
		return x.Followees
	}
	return nil
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query      string             `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Pagination *PaginationRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	mi := &file_api_v1_users_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{10}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetPagination() *PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users      []*User             `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Pagination *PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	mi := &file_api_v1_users_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{11}
}

func (x *SearchResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *SearchResponse) GetPagination() *PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListNotificationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnreadOnly bool               `protobuf:"varint,1,opt,name=unread_only,json=unreadOnly,proto3" json:"unread_only,omitempty"`
	MarkAsRead bool               `protobuf:"varint,2,opt,name=mark_as_read,json=markAsRead,proto3" json:"mark_as_read,omitempty"`
	Pagination *PaginationRequest `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListNotificationsRequest) Reset() {
	*x = ListNotificationsRequest{}
	mi := &file_api_v1_users_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotificationsRequest) ProtoMessage() {}

func (x *ListNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotificationsRequest.ProtoReflect.Descriptor instead.
func (*ListNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{12}
}

func (x *ListNotificationsRequest) GetUnreadOnly() bool {
	if x != nil {
		return x.UnreadOnly
	}
	return false
}

func (x *ListNotificationsRequest) GetMarkAsRead() bool {
	if x != nil {
		return x.MarkAsRead
	}
	return false
}

func (x *ListNotificationsRequest) GetPagination() *PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListNotificationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notifications []*Notification     `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
	Pagination    *PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListNotificationsResponse) Reset() {
	*x = ListNotificationsResponse{}
	mi := &file_api_v1_users_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotificationsResponse) ProtoMessage() {}

func (x *ListNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotificationsResponse.ProtoReflect.Descriptor instead.
func (*ListNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{13}
}

func (x *ListNotificationsResponse) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

func (x *ListNotificationsResponse) GetPagination() *PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// DEBT: This should be a timestamp but the client is not able to parse it.
	NotifiedAtUnix int64 `protobuf:"varint,2,opt,name=notified_at_unix,json=notifiedAtUnix,proto3" json:"notified_at_unix,omitempty"`
	// Types that are assignable to Type:
	//
	//	*Notification_WorkoutComment_
	Type isNotification_Type `protobuf_oneof:"type"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	mi := &file_api_v1_users_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{14}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetNotifiedAtUnix() int64 {
	if x != nil {
		return x.NotifiedAtUnix
	}
	return 0
}

func (m *Notification) GetType() isNotification_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Notification) GetWorkoutComment() *Notification_WorkoutComment {
	if x, ok := x.GetType().(*Notification_WorkoutComment_); ok {
		return x.WorkoutComment
	}
	return nil
}

type isNotification_Type interface {
	isNotification_Type()
}

type Notification_WorkoutComment_ struct {
	WorkoutComment *Notification_WorkoutComment `protobuf:"bytes,3,opt,name=workout_comment,json=workoutComment,proto3,oneof"`
}

func (*Notification_WorkoutComment_) isNotification_Type() {}

type Notification_WorkoutComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actor   *User    `protobuf:"bytes,1,opt,name=actor,proto3" json:"actor,omitempty"`
	Workout *Workout `protobuf:"bytes,2,opt,name=workout,proto3" json:"workout,omitempty"`
}

func (x *Notification_WorkoutComment) Reset() {
	*x = Notification_WorkoutComment{}
	mi := &file_api_v1_users_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification_WorkoutComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification_WorkoutComment) ProtoMessage() {}

func (x *Notification_WorkoutComment) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_users_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification_WorkoutComment.ProtoReflect.Descriptor instead.
func (*Notification_WorkoutComment) Descriptor() ([]byte, []int) {
	return file_api_v1_users_proto_rawDescGZIP(), []int{14, 0}
}

func (x *Notification_WorkoutComment) GetActor() *User {
	if x != nil {
		return x.Actor
	}
	return nil
}

func (x *Notification_WorkoutComment) GetWorkout() *Workout {
	if x != nil {
		return x.Workout
	}
	return nil
}

var File_api_v1_users_proto protoreflect.FileDescriptor

var file_api_v1_users_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x36, 0x0a,
	0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x08, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x0a, 0x0f, 0x55, 0x6e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x0b, 0x75, 0x6e,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0a, 0x75, 0x6e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0x0a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x15,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x22, 0x41, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x0b, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x65, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a,
	0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x73, 0x22, 0x71, 0x0a, 0x0d, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02,
	0x10, 0x03, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x41, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x70, 0x0a, 0x0e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa0,
	0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75,
	0x6e, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0c,
	0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x61, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x41,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x93, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x81, 0x02, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x55, 0x6e,
	0x69, 0x78, 0x12, 0x4e, 0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x1a, 0x5f, 0x0a, 0x0e, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b,
	0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b,
	0x6f, 0x75, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x32, 0x96, 0x04, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x3d, 0x0a, 0x06, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x43, 0x0a, 0x08, 0x55, 0x6e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x52, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x1c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18,
	0x01, 0x12, 0x52, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x65, 0x73, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04,
	0x88, 0xb5, 0x18, 0x01, 0x12, 0x5e, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04,
	0x88, 0xb5, 0x18, 0x01, 0x42, 0x8b, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x42, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
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
	file_api_v1_users_proto_rawDescOnce sync.Once
	file_api_v1_users_proto_rawDescData = file_api_v1_users_proto_rawDesc
)

func file_api_v1_users_proto_rawDescGZIP() []byte {
	file_api_v1_users_proto_rawDescOnce.Do(func() {
		file_api_v1_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_users_proto_rawDescData)
	})
	return file_api_v1_users_proto_rawDescData
}

var file_api_v1_users_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_api_v1_users_proto_goTypes = []any{
	(*GetUserRequest)(nil),              // 0: api.v1.GetUserRequest
	(*GetUserResponse)(nil),             // 1: api.v1.GetUserResponse
	(*FollowRequest)(nil),               // 2: api.v1.FollowRequest
	(*FollowResponse)(nil),              // 3: api.v1.FollowResponse
	(*UnfollowRequest)(nil),             // 4: api.v1.UnfollowRequest
	(*UnfollowResponse)(nil),            // 5: api.v1.UnfollowResponse
	(*ListFollowersRequest)(nil),        // 6: api.v1.ListFollowersRequest
	(*ListFollowersResponse)(nil),       // 7: api.v1.ListFollowersResponse
	(*ListFolloweesRequest)(nil),        // 8: api.v1.ListFolloweesRequest
	(*ListFolloweesResponse)(nil),       // 9: api.v1.ListFolloweesResponse
	(*SearchRequest)(nil),               // 10: api.v1.SearchRequest
	(*SearchResponse)(nil),              // 11: api.v1.SearchResponse
	(*ListNotificationsRequest)(nil),    // 12: api.v1.ListNotificationsRequest
	(*ListNotificationsResponse)(nil),   // 13: api.v1.ListNotificationsResponse
	(*Notification)(nil),                // 14: api.v1.Notification
	(*Notification_WorkoutComment)(nil), // 15: api.v1.Notification.WorkoutComment
	(*User)(nil),                        // 16: api.v1.User
	(*PaginationRequest)(nil),           // 17: api.v1.PaginationRequest
	(*PaginationResponse)(nil),          // 18: api.v1.PaginationResponse
	(*Workout)(nil),                     // 19: api.v1.Workout
}
var file_api_v1_users_proto_depIdxs = []int32{
	16, // 0: api.v1.GetUserResponse.user:type_name -> api.v1.User
	16, // 1: api.v1.ListFollowersResponse.followers:type_name -> api.v1.User
	16, // 2: api.v1.ListFolloweesResponse.followees:type_name -> api.v1.User
	17, // 3: api.v1.SearchRequest.pagination:type_name -> api.v1.PaginationRequest
	16, // 4: api.v1.SearchResponse.users:type_name -> api.v1.User
	18, // 5: api.v1.SearchResponse.pagination:type_name -> api.v1.PaginationResponse
	17, // 6: api.v1.ListNotificationsRequest.pagination:type_name -> api.v1.PaginationRequest
	14, // 7: api.v1.ListNotificationsResponse.notifications:type_name -> api.v1.Notification
	18, // 8: api.v1.ListNotificationsResponse.pagination:type_name -> api.v1.PaginationResponse
	15, // 9: api.v1.Notification.workout_comment:type_name -> api.v1.Notification.WorkoutComment
	16, // 10: api.v1.Notification.WorkoutComment.actor:type_name -> api.v1.User
	19, // 11: api.v1.Notification.WorkoutComment.workout:type_name -> api.v1.Workout
	0,  // 12: api.v1.UserService.Get:input_type -> api.v1.GetUserRequest
	2,  // 13: api.v1.UserService.Follow:input_type -> api.v1.FollowRequest
	4,  // 14: api.v1.UserService.Unfollow:input_type -> api.v1.UnfollowRequest
	6,  // 15: api.v1.UserService.ListFollowers:input_type -> api.v1.ListFollowersRequest
	8,  // 16: api.v1.UserService.ListFollowees:input_type -> api.v1.ListFolloweesRequest
	10, // 17: api.v1.UserService.Search:input_type -> api.v1.SearchRequest
	12, // 18: api.v1.UserService.ListNotifications:input_type -> api.v1.ListNotificationsRequest
	1,  // 19: api.v1.UserService.Get:output_type -> api.v1.GetUserResponse
	3,  // 20: api.v1.UserService.Follow:output_type -> api.v1.FollowResponse
	5,  // 21: api.v1.UserService.Unfollow:output_type -> api.v1.UnfollowResponse
	7,  // 22: api.v1.UserService.ListFollowers:output_type -> api.v1.ListFollowersResponse
	9,  // 23: api.v1.UserService.ListFollowees:output_type -> api.v1.ListFolloweesResponse
	11, // 24: api.v1.UserService.Search:output_type -> api.v1.SearchResponse
	13, // 25: api.v1.UserService.ListNotifications:output_type -> api.v1.ListNotificationsResponse
	19, // [19:26] is the sub-list for method output_type
	12, // [12:19] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_api_v1_users_proto_init() }
func file_api_v1_users_proto_init() {
	if File_api_v1_users_proto != nil {
		return
	}
	file_api_v1_options_proto_init()
	file_api_v1_shared_proto_init()
	file_api_v1_workouts_proto_init()
	file_api_v1_users_proto_msgTypes[14].OneofWrappers = []any{
		(*Notification_WorkoutComment_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_users_proto_goTypes,
		DependencyIndexes: file_api_v1_users_proto_depIdxs,
		MessageInfos:      file_api_v1_users_proto_msgTypes,
	}.Build()
	File_api_v1_users_proto = out.File
	file_api_v1_users_proto_rawDesc = nil
	file_api_v1_users_proto_goTypes = nil
	file_api_v1_users_proto_depIdxs = nil
}
