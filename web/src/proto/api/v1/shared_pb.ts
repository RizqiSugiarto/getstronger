// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file api/v1/shared.proto (package api.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_buf_validate_validate } from "../../buf/validate/validate_pb";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file api/v1/shared.proto.
 */
export const file_api_v1_shared: GenFile = /*@__PURE__*/
  fileDesc("ChNhcGkvdjEvc2hhcmVkLnByb3RvEgZhcGkudjEiXwoMRXhlcmNpc2VTZXRzEioKCGV4ZXJjaXNlGAEgASgLMhAuYXBpLnYxLkV4ZXJjaXNlQga6SAPIAQESIwoEc2V0cxgCIAMoCzILLmFwaS52MS5TZXRCCLpIBZIBAggBIk4KCEV4ZXJjaXNlEhQKAmlkGAEgASgJQgi6SAVyA7ABARIPCgd1c2VyX2lkGAIgASgJEgwKBG5hbWUYAyABKAkSDQoFbGFiZWwYBCABKAkiUwoDU2V0Eg4KBndlaWdodBgBIAEoARIVCgRyZXBzGAIgASgFQge6SAQaAigBEiUKCG1ldGFkYXRhGAMgASgLMhMuYXBpLnYxLk1ldGFkYXRhU2V0IlsKC01ldGFkYXRhU2V0EhwKCndvcmtvdXRfaWQYASABKAlCCLpIBXIDsAEBEi4KCmNyZWF0ZWRfYXQYAiABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wImcKBFVzZXISFAoCaWQYASABKAlCCLpIBXIDsAEBEhsKCmZpcnN0X25hbWUYAiABKAlCB7pIBHICEAESGgoJbGFzdF9uYW1lGAMgASgJQge6SARyAhABEhAKCGZvbGxvd2VkGAQgASgIIkYKEVBhZ2luYXRpb25SZXF1ZXN0Eh0KCnBhZ2VfbGltaXQYASABKAVCCbpIBhoEGGQoARISCgpwYWdlX3Rva2VuGAIgASgMIi0KElBhZ2luYXRpb25SZXNwb25zZRIXCg9uZXh0X3BhZ2VfdG9rZW4YASABKAxCjwEKCmNvbS5hcGkudjFCC1NoYXJlZFByb3RvUAFaO2dpdGh1Yi5jb20vY3Jsc3NuL2dldHN0cm9uZ2VyL3NlcnZlci9wa2cvcHJvdG8vYXBpL3YxO2FwaXYxogIDQVhYqgIGQXBpLlYxygIGQXBpXFYx4gISQXBpXFYxXEdQQk1ldGFkYXRh6gIHQXBpOjpWMWIGcHJvdG8z", [file_buf_validate_validate, file_google_protobuf_timestamp]);

/**
 * @generated from message api.v1.ExerciseSets
 */
export type ExerciseSets = Message<"api.v1.ExerciseSets"> & {
  /**
   * @generated from field: api.v1.Exercise exercise = 1;
   */
  exercise?: Exercise;

  /**
   * @generated from field: repeated api.v1.Set sets = 2;
   */
  sets: Set[];
};

/**
 * Describes the message api.v1.ExerciseSets.
 * Use `create(ExerciseSetsSchema)` to create a new message.
 */
export const ExerciseSetsSchema: GenMessage<ExerciseSets> = /*@__PURE__*/
  messageDesc(file_api_v1_shared, 0);

/**
 * @generated from message api.v1.Exercise
 */
export type Exercise = Message<"api.v1.Exercise"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string user_id = 2;
   */
  userId: string;

  /**
   * @generated from field: string name = 3;
   */
  name: string;

  /**
   * @generated from field: string label = 4;
   */
  label: string;
};

/**
 * Describes the message api.v1.Exercise.
 * Use `create(ExerciseSchema)` to create a new message.
 */
export const ExerciseSchema: GenMessage<Exercise> = /*@__PURE__*/
  messageDesc(file_api_v1_shared, 1);

/**
 * @generated from message api.v1.Set
 */
export type Set = Message<"api.v1.Set"> & {
  /**
   * The weight can be less than zero.
   *
   * @generated from field: double weight = 1;
   */
  weight: number;

  /**
   * @generated from field: int32 reps = 2;
   */
  reps: number;

  /**
   * @generated from field: api.v1.MetadataSet metadata = 3;
   */
  metadata?: MetadataSet;
};

/**
 * Describes the message api.v1.Set.
 * Use `create(SetSchema)` to create a new message.
 */
export const SetSchema: GenMessage<Set> = /*@__PURE__*/
  messageDesc(file_api_v1_shared, 2);

/**
 * @generated from message api.v1.MetadataSet
 */
export type MetadataSet = Message<"api.v1.MetadataSet"> & {
  /**
   * @generated from field: string workout_id = 1;
   */
  workoutId: string;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 2;
   */
  createdAt?: Timestamp;
};

/**
 * Describes the message api.v1.MetadataSet.
 * Use `create(MetadataSetSchema)` to create a new message.
 */
export const MetadataSetSchema: GenMessage<MetadataSet> = /*@__PURE__*/
  messageDesc(file_api_v1_shared, 3);

/**
 * @generated from message api.v1.User
 */
export type User = Message<"api.v1.User"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string first_name = 2;
   */
  firstName: string;

  /**
   * @generated from field: string last_name = 3;
   */
  lastName: string;

  /**
   * @generated from field: bool followed = 4;
   */
  followed: boolean;
};

/**
 * Describes the message api.v1.User.
 * Use `create(UserSchema)` to create a new message.
 */
export const UserSchema: GenMessage<User> = /*@__PURE__*/
  messageDesc(file_api_v1_shared, 4);

/**
 * @generated from message api.v1.PaginationRequest
 */
export type PaginationRequest = Message<"api.v1.PaginationRequest"> & {
  /**
   * @generated from field: int32 page_limit = 1;
   */
  pageLimit: number;

  /**
   * @generated from field: bytes page_token = 2;
   */
  pageToken: Uint8Array;
};

/**
 * Describes the message api.v1.PaginationRequest.
 * Use `create(PaginationRequestSchema)` to create a new message.
 */
export const PaginationRequestSchema: GenMessage<PaginationRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_shared, 5);

/**
 * @generated from message api.v1.PaginationResponse
 */
export type PaginationResponse = Message<"api.v1.PaginationResponse"> & {
  /**
   * @generated from field: bytes next_page_token = 1;
   */
  nextPageToken: Uint8Array;
};

/**
 * Describes the message api.v1.PaginationResponse.
 * Use `create(PaginationResponseSchema)` to create a new message.
 */
export const PaginationResponseSchema: GenMessage<PaginationResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_shared, 6);

