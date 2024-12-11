// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file api/v1/user_service.proto (package api.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_api_v1_options } from "./options_pb";
import type { PaginationRequest, PaginationResponse, User } from "./shared_pb";
import { file_api_v1_shared } from "./shared_pb";
import { file_api_v1_workout_service } from "./workout_service_pb";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import { file_buf_validate_validate } from "../../buf/validate/validate_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file api/v1/user_service.proto.
 */
export const file_api_v1_user_service: GenFile = /*@__PURE__*/
  fileDesc("ChlhcGkvdjEvdXNlcl9zZXJ2aWNlLnByb3RvEgZhcGkudjEiJgoOR2V0VXNlclJlcXVlc3QSFAoCaWQYASABKAlCCLpIBXIDsAEBIi0KD0dldFVzZXJSZXNwb25zZRIaCgR1c2VyGAEgASgLMgwuYXBpLnYxLlVzZXIiMAoRRm9sbG93VXNlclJlcXVlc3QSGwoJZm9sbG93X2lkGAEgASgJQgi6SAVyA7ABASIUChJGb2xsb3dVc2VyUmVzcG9uc2UiNAoTVW5mb2xsb3dVc2VyUmVxdWVzdBIdCgt1bmZvbGxvd19pZBgBIAEoCUIIukgFcgOwAQEiFgoUVW5mb2xsb3dVc2VyUmVzcG9uc2UiNQoUTGlzdEZvbGxvd2Vyc1JlcXVlc3QSHQoLZm9sbG93ZXJfaWQYASABKAlCCLpIBXIDsAEBIjgKFUxpc3RGb2xsb3dlcnNSZXNwb25zZRIfCglmb2xsb3dlcnMYASADKAsyDC5hcGkudjEuVXNlciI1ChRMaXN0Rm9sbG93ZWVzUmVxdWVzdBIdCgtmb2xsb3dlZV9pZBgBIAEoCUIIukgFcgOwAQEiOAoVTGlzdEZvbGxvd2Vlc1Jlc3BvbnNlEh8KCWZvbGxvd2VlcxgBIAMoCzIMLmFwaS52MS5Vc2VyImMKElNlYXJjaFVzZXJzUmVxdWVzdBIWCgVxdWVyeRgBIAEoCUIHukgEcgIQAxI1CgpwYWdpbmF0aW9uGAIgASgLMhkuYXBpLnYxLlBhZ2luYXRpb25SZXF1ZXN0Qga6SAPIAQEiYgoTU2VhcmNoVXNlcnNSZXNwb25zZRIbCgV1c2VycxgBIAMoCzIMLmFwaS52MS5Vc2VyEi4KCnBhZ2luYXRpb24YAiABKAsyGi5hcGkudjEuUGFnaW5hdGlvblJlc3BvbnNlMuEDCgtVc2VyU2VydmljZRJACgdHZXRVc2VyEhYuYXBpLnYxLkdldFVzZXJSZXF1ZXN0GhcuYXBpLnYxLkdldFVzZXJSZXNwb25zZSIEiLUYARJJCgpGb2xsb3dVc2VyEhkuYXBpLnYxLkZvbGxvd1VzZXJSZXF1ZXN0GhouYXBpLnYxLkZvbGxvd1VzZXJSZXNwb25zZSIEiLUYARJPCgxVbmZvbGxvd1VzZXISGy5hcGkudjEuVW5mb2xsb3dVc2VyUmVxdWVzdBocLmFwaS52MS5VbmZvbGxvd1VzZXJSZXNwb25zZSIEiLUYARJSCg1MaXN0Rm9sbG93ZXJzEhwuYXBpLnYxLkxpc3RGb2xsb3dlcnNSZXF1ZXN0Gh0uYXBpLnYxLkxpc3RGb2xsb3dlcnNSZXNwb25zZSIEiLUYARJSCg1MaXN0Rm9sbG93ZWVzEhwuYXBpLnYxLkxpc3RGb2xsb3dlZXNSZXF1ZXN0Gh0uYXBpLnYxLkxpc3RGb2xsb3dlZXNSZXNwb25zZSIEiLUYARJMCgtTZWFyY2hVc2VycxIaLmFwaS52MS5TZWFyY2hVc2Vyc1JlcXVlc3QaGy5hcGkudjEuU2VhcmNoVXNlcnNSZXNwb25zZSIEiLUYAUKUAQoKY29tLmFwaS52MUIQVXNlclNlcnZpY2VQcm90b1ABWjtnaXRodWIuY29tL2NybHNzbi9nZXRzdHJvbmdlci9zZXJ2ZXIvcGtnL3Byb3RvL2FwaS92MTthcGl2MaICA0FYWKoCBkFwaS5WMcoCBkFwaVxWMeICEkFwaVxWMVxHUEJNZXRhZGF0YeoCB0FwaTo6VjFiBnByb3RvMw", [file_api_v1_options, file_api_v1_shared, file_api_v1_workout_service, file_google_protobuf_timestamp, file_buf_validate_validate]);

/**
 * @generated from message api.v1.GetUserRequest
 */
export type GetUserRequest = Message<"api.v1.GetUserRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message api.v1.GetUserRequest.
 * Use `create(GetUserRequestSchema)` to create a new message.
 */
export const GetUserRequestSchema: GenMessage<GetUserRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_user_service, 0);

/**
 * @generated from message api.v1.GetUserResponse
 */
export type GetUserResponse = Message<"api.v1.GetUserResponse"> & {
  /**
   * @generated from field: api.v1.User user = 1;
   */
  user?: User;
};

/**
 * Describes the message api.v1.GetUserResponse.
 * Use `create(GetUserResponseSchema)` to create a new message.
 */
export const GetUserResponseSchema: GenMessage<GetUserResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_user_service, 1);

/**
 * @generated from message api.v1.FollowUserRequest
 */
export type FollowUserRequest = Message<"api.v1.FollowUserRequest"> & {
  /**
   * @generated from field: string follow_id = 1;
   */
  followId: string;
};

/**
 * Describes the message api.v1.FollowUserRequest.
 * Use `create(FollowUserRequestSchema)` to create a new message.
 */
export const FollowUserRequestSchema: GenMessage<FollowUserRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_user_service, 2);

/**
 * @generated from message api.v1.FollowUserResponse
 */
export type FollowUserResponse = Message<"api.v1.FollowUserResponse"> & {
};

/**
 * Describes the message api.v1.FollowUserResponse.
 * Use `create(FollowUserResponseSchema)` to create a new message.
 */
export const FollowUserResponseSchema: GenMessage<FollowUserResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_user_service, 3);

/**
 * @generated from message api.v1.UnfollowUserRequest
 */
export type UnfollowUserRequest = Message<"api.v1.UnfollowUserRequest"> & {
  /**
   * @generated from field: string unfollow_id = 1;
   */
  unfollowId: string;
};

/**
 * Describes the message api.v1.UnfollowUserRequest.
 * Use `create(UnfollowUserRequestSchema)` to create a new message.
 */
export const UnfollowUserRequestSchema: GenMessage<UnfollowUserRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_user_service, 4);

/**
 * @generated from message api.v1.UnfollowUserResponse
 */
export type UnfollowUserResponse = Message<"api.v1.UnfollowUserResponse"> & {
};

/**
 * Describes the message api.v1.UnfollowUserResponse.
 * Use `create(UnfollowUserResponseSchema)` to create a new message.
 */
export const UnfollowUserResponseSchema: GenMessage<UnfollowUserResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_user_service, 5);

/**
 * @generated from message api.v1.ListFollowersRequest
 */
export type ListFollowersRequest = Message<"api.v1.ListFollowersRequest"> & {
  /**
   * @generated from field: string follower_id = 1;
   */
  followerId: string;
};

/**
 * Describes the message api.v1.ListFollowersRequest.
 * Use `create(ListFollowersRequestSchema)` to create a new message.
 */
export const ListFollowersRequestSchema: GenMessage<ListFollowersRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_user_service, 6);

/**
 * @generated from message api.v1.ListFollowersResponse
 */
export type ListFollowersResponse = Message<"api.v1.ListFollowersResponse"> & {
  /**
   * @generated from field: repeated api.v1.User followers = 1;
   */
  followers: User[];
};

/**
 * Describes the message api.v1.ListFollowersResponse.
 * Use `create(ListFollowersResponseSchema)` to create a new message.
 */
export const ListFollowersResponseSchema: GenMessage<ListFollowersResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_user_service, 7);

/**
 * @generated from message api.v1.ListFolloweesRequest
 */
export type ListFolloweesRequest = Message<"api.v1.ListFolloweesRequest"> & {
  /**
   * @generated from field: string followee_id = 1;
   */
  followeeId: string;
};

/**
 * Describes the message api.v1.ListFolloweesRequest.
 * Use `create(ListFolloweesRequestSchema)` to create a new message.
 */
export const ListFolloweesRequestSchema: GenMessage<ListFolloweesRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_user_service, 8);

/**
 * @generated from message api.v1.ListFolloweesResponse
 */
export type ListFolloweesResponse = Message<"api.v1.ListFolloweesResponse"> & {
  /**
   * @generated from field: repeated api.v1.User followees = 1;
   */
  followees: User[];
};

/**
 * Describes the message api.v1.ListFolloweesResponse.
 * Use `create(ListFolloweesResponseSchema)` to create a new message.
 */
export const ListFolloweesResponseSchema: GenMessage<ListFolloweesResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_user_service, 9);

/**
 * @generated from message api.v1.SearchUsersRequest
 */
export type SearchUsersRequest = Message<"api.v1.SearchUsersRequest"> & {
  /**
   * @generated from field: string query = 1;
   */
  query: string;

  /**
   * @generated from field: api.v1.PaginationRequest pagination = 2;
   */
  pagination?: PaginationRequest;
};

/**
 * Describes the message api.v1.SearchUsersRequest.
 * Use `create(SearchUsersRequestSchema)` to create a new message.
 */
export const SearchUsersRequestSchema: GenMessage<SearchUsersRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_user_service, 10);

/**
 * @generated from message api.v1.SearchUsersResponse
 */
export type SearchUsersResponse = Message<"api.v1.SearchUsersResponse"> & {
  /**
   * @generated from field: repeated api.v1.User users = 1;
   */
  users: User[];

  /**
   * @generated from field: api.v1.PaginationResponse pagination = 2;
   */
  pagination?: PaginationResponse;
};

/**
 * Describes the message api.v1.SearchUsersResponse.
 * Use `create(SearchUsersResponseSchema)` to create a new message.
 */
export const SearchUsersResponseSchema: GenMessage<SearchUsersResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_user_service, 11);

/**
 * @generated from service api.v1.UserService
 */
export const UserService: GenService<{
  /**
   * @generated from rpc api.v1.UserService.GetUser
   */
  getUser: {
    methodKind: "unary";
    input: typeof GetUserRequestSchema;
    output: typeof GetUserResponseSchema;
  },
  /**
   * @generated from rpc api.v1.UserService.FollowUser
   */
  followUser: {
    methodKind: "unary";
    input: typeof FollowUserRequestSchema;
    output: typeof FollowUserResponseSchema;
  },
  /**
   * @generated from rpc api.v1.UserService.UnfollowUser
   */
  unfollowUser: {
    methodKind: "unary";
    input: typeof UnfollowUserRequestSchema;
    output: typeof UnfollowUserResponseSchema;
  },
  /**
   * @generated from rpc api.v1.UserService.ListFollowers
   */
  listFollowers: {
    methodKind: "unary";
    input: typeof ListFollowersRequestSchema;
    output: typeof ListFollowersResponseSchema;
  },
  /**
   * @generated from rpc api.v1.UserService.ListFollowees
   */
  listFollowees: {
    methodKind: "unary";
    input: typeof ListFolloweesRequestSchema;
    output: typeof ListFolloweesResponseSchema;
  },
  /**
   * @generated from rpc api.v1.UserService.SearchUsers
   */
  searchUsers: {
    methodKind: "unary";
    input: typeof SearchUsersRequestSchema;
    output: typeof SearchUsersResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_api_v1_user_service, 0);

