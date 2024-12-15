// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file api/v1/feed_service.proto (package api.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_api_v1_options } from "./options_pb";
import type { PaginationRequest, PaginationResponse } from "./shared_pb";
import { file_api_v1_shared } from "./shared_pb";
import type { Workout } from "./workout_service_pb";
import { file_api_v1_workout_service } from "./workout_service_pb";
import { file_buf_validate_validate } from "../../buf/validate/validate_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file api/v1/feed_service.proto.
 */
export const file_api_v1_feed_service: GenFile = /*@__PURE__*/
  fileDesc("ChlhcGkvdjEvZmVlZF9zZXJ2aWNlLnByb3RvEgZhcGkudjEiZAoUTGlzdEZlZWRJdGVtc1JlcXVlc3QSFQoNZm9sbG93ZWRfb25seRgBIAEoCBI1CgpwYWdpbmF0aW9uGAIgASgLMhkuYXBpLnYxLlBhZ2luYXRpb25SZXF1ZXN0Qga6SAPIAQEiaAoVTGlzdEZlZWRJdGVtc1Jlc3BvbnNlEh8KBWl0ZW1zGAEgAygLMhAuYXBpLnYxLkZlZWRJdGVtEi4KCnBhZ2luYXRpb24YAiABKAsyGi5hcGkudjEuUGFnaW5hdGlvblJlc3BvbnNlIjYKCEZlZWRJdGVtEiIKB3dvcmtvdXQYASABKAsyDy5hcGkudjEuV29ya291dEgAQgYKBHR5cGUyYQoLRmVlZFNlcnZpY2USUgoNTGlzdEZlZWRJdGVtcxIcLmFwaS52MS5MaXN0RmVlZEl0ZW1zUmVxdWVzdBodLmFwaS52MS5MaXN0RmVlZEl0ZW1zUmVzcG9uc2UiBIi1GAFClAEKCmNvbS5hcGkudjFCEEZlZWRTZXJ2aWNlUHJvdG9QAVo7Z2l0aHViLmNvbS9jcmxzc24vZ2V0c3Ryb25nZXIvc2VydmVyL3BrZy9wcm90by9hcGkvdjE7YXBpdjGiAgNBWFiqAgZBcGkuVjHKAgZBcGlcVjHiAhJBcGlcVjFcR1BCTWV0YWRhdGHqAgdBcGk6OlYxYgZwcm90bzM", [file_api_v1_options, file_api_v1_shared, file_api_v1_workout_service, file_buf_validate_validate]);

/**
 * @generated from message api.v1.ListFeedItemsRequest
 */
export type ListFeedItemsRequest = Message<"api.v1.ListFeedItemsRequest"> & {
  /**
   * @generated from field: bool followed_only = 1;
   */
  followedOnly: boolean;

  /**
   * @generated from field: api.v1.PaginationRequest pagination = 2;
   */
  pagination?: PaginationRequest;
};

/**
 * Describes the message api.v1.ListFeedItemsRequest.
 * Use `create(ListFeedItemsRequestSchema)` to create a new message.
 */
export const ListFeedItemsRequestSchema: GenMessage<ListFeedItemsRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_feed_service, 0);

/**
 * @generated from message api.v1.ListFeedItemsResponse
 */
export type ListFeedItemsResponse = Message<"api.v1.ListFeedItemsResponse"> & {
  /**
   * @generated from field: repeated api.v1.FeedItem items = 1;
   */
  items: FeedItem[];

  /**
   * @generated from field: api.v1.PaginationResponse pagination = 2;
   */
  pagination?: PaginationResponse;
};

/**
 * Describes the message api.v1.ListFeedItemsResponse.
 * Use `create(ListFeedItemsResponseSchema)` to create a new message.
 */
export const ListFeedItemsResponseSchema: GenMessage<ListFeedItemsResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_feed_service, 1);

/**
 * @generated from message api.v1.FeedItem
 */
export type FeedItem = Message<"api.v1.FeedItem"> & {
  /**
   * @generated from oneof api.v1.FeedItem.type
   */
  type: {
    /**
     * @generated from field: api.v1.Workout workout = 1;
     */
    value: Workout;
    case: "workout";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message api.v1.FeedItem.
 * Use `create(FeedItemSchema)` to create a new message.
 */
export const FeedItemSchema: GenMessage<FeedItem> = /*@__PURE__*/
  messageDesc(file_api_v1_feed_service, 2);

/**
 * @generated from service api.v1.FeedService
 */
export const FeedService: GenService<{
  /**
   * @generated from rpc api.v1.FeedService.ListFeedItems
   */
  listFeedItems: {
    methodKind: "unary";
    input: typeof ListFeedItemsRequestSchema;
    output: typeof ListFeedItemsResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_api_v1_feed_service, 0);

