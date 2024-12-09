// @generated by protoc-gen-es v2.2.2 with parameter "target=ts"
// @generated from file api/v1/exercise.proto (package api.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_api_v1_options } from "./options_pb";
import type { Exercise, ExerciseSets, PaginationRequest, PaginationResponse, Set } from "./shared_pb";
import { file_api_v1_shared } from "./shared_pb";
import type { FieldMask } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_field_mask, file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import { file_buf_validate_validate } from "../../buf/validate/validate_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file api/v1/exercise.proto.
 */
export const file_api_v1_exercise: GenFile = /*@__PURE__*/
  fileDesc("ChVhcGkvdjEvZXhlcmNpc2UucHJvdG8SBmFwaS52MSI9ChVDcmVhdGVFeGVyY2lzZVJlcXVlc3QSFQoEbmFtZRgBIAEoCUIHukgEcgIQARINCgVsYWJlbBgCIAEoCSIkChZDcmVhdGVFeGVyY2lzZVJlc3BvbnNlEgoKAmlkGAEgASgJIioKEkdldEV4ZXJjaXNlUmVxdWVzdBIUCgJpZBgBIAEoCUIIukgFcgOwAQEiOQoTR2V0RXhlcmNpc2VSZXNwb25zZRIiCghleGVyY2lzZRgBIAEoCzIQLmFwaS52MS5FeGVyY2lzZSJ0ChVVcGRhdGVFeGVyY2lzZVJlcXVlc3QSKgoIZXhlcmNpc2UYASABKAsyEC5hcGkudjEuRXhlcmNpc2VCBrpIA8gBARIvCgt1cGRhdGVfbWFzaxgCIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5GaWVsZE1hc2siPAoWVXBkYXRlRXhlcmNpc2VSZXNwb25zZRIiCghleGVyY2lzZRgBIAEoCzIQLmFwaS52MS5FeGVyY2lzZSItChVEZWxldGVFeGVyY2lzZVJlcXVlc3QSFAoCaWQYASABKAlCCLpIBXIDsAEBIhgKFkRlbGV0ZUV4ZXJjaXNlUmVzcG9uc2UiewoUTGlzdEV4ZXJjaXNlc1JlcXVlc3QSDAoEbmFtZRgBIAEoCRIjCgxleGVyY2lzZV9pZHMYAiADKAlCDbpICpIBByIFcgOwAQESHAoJcGFnZV9zaXplGAMgASgFQgm6SAYaBBhkKAESEgoKcGFnZV90b2tlbhgEIAEoDCJVChVMaXN0RXhlcmNpc2VzUmVzcG9uc2USIwoJZXhlcmNpc2VzGAEgAygLMhAuYXBpLnYxLkV4ZXJjaXNlEhcKD25leHRfcGFnZV90b2tlbhgCIAEoDCJECh1HZXRQcmV2aW91c1dvcmtvdXRTZXRzUmVxdWVzdBIjCgxleGVyY2lzZV9pZHMYASADKAlCDbpICpIBByIFcgOwAQEiTQoeR2V0UHJldmlvdXNXb3Jrb3V0U2V0c1Jlc3BvbnNlEisKDWV4ZXJjaXNlX3NldHMYASADKAsyFC5hcGkudjEuRXhlcmNpc2VTZXRzIjQKF0dldFBlcnNvbmFsQmVzdHNSZXF1ZXN0EhkKB3VzZXJfaWQYASABKAlCCLpIBXIDsAEBIkgKGEdldFBlcnNvbmFsQmVzdHNSZXNwb25zZRIsCg5wZXJzb25hbF9iZXN0cxgBIAMoCzIULmFwaS52MS5QZXJzb25hbEJlc3QiXwoPTGlzdFNldHNSZXF1ZXN0Eh0KC2V4ZXJjaXNlX2lkGAEgASgJQgi6SAVyA7ABARItCgpwYWdpbmF0aW9uGAIgASgLMhkuYXBpLnYxLlBhZ2luYXRpb25SZXF1ZXN0Il0KEExpc3RTZXRzUmVzcG9uc2USGQoEc2V0cxgBIAMoCzILLmFwaS52MS5TZXQSLgoKcGFnaW5hdGlvbhgCIAEoCzIaLmFwaS52MS5QYWdpbmF0aW9uUmVzcG9uc2UiXAoMUGVyc29uYWxCZXN0EioKCGV4ZXJjaXNlGAEgASgLMhAuYXBpLnYxLkV4ZXJjaXNlQga6SAPIAQESIAoDc2V0GAIgASgLMgsuYXBpLnYxLlNldEIGukgDyAEBMqAFCg9FeGVyY2lzZVNlcnZpY2USTQoGQ3JlYXRlEh0uYXBpLnYxLkNyZWF0ZUV4ZXJjaXNlUmVxdWVzdBoeLmFwaS52MS5DcmVhdGVFeGVyY2lzZVJlc3BvbnNlIgSItRgBEkQKA0dldBIaLmFwaS52MS5HZXRFeGVyY2lzZVJlcXVlc3QaGy5hcGkudjEuR2V0RXhlcmNpc2VSZXNwb25zZSIEiLUYARJNCgZVcGRhdGUSHS5hcGkudjEuVXBkYXRlRXhlcmNpc2VSZXF1ZXN0Gh4uYXBpLnYxLlVwZGF0ZUV4ZXJjaXNlUmVzcG9uc2UiBIi1GAESTQoGRGVsZXRlEh0uYXBpLnYxLkRlbGV0ZUV4ZXJjaXNlUmVxdWVzdBoeLmFwaS52MS5EZWxldGVFeGVyY2lzZVJlc3BvbnNlIgSItRgBEkkKBExpc3QSHC5hcGkudjEuTGlzdEV4ZXJjaXNlc1JlcXVlc3QaHS5hcGkudjEuTGlzdEV4ZXJjaXNlc1Jlc3BvbnNlIgSItRgBEm0KFkdldFByZXZpb3VzV29ya291dFNldHMSJS5hcGkudjEuR2V0UHJldmlvdXNXb3Jrb3V0U2V0c1JlcXVlc3QaJi5hcGkudjEuR2V0UHJldmlvdXNXb3Jrb3V0U2V0c1Jlc3BvbnNlIgSItRgBElsKEEdldFBlcnNvbmFsQmVzdHMSHy5hcGkudjEuR2V0UGVyc29uYWxCZXN0c1JlcXVlc3QaIC5hcGkudjEuR2V0UGVyc29uYWxCZXN0c1Jlc3BvbnNlIgSItRgBEkMKCExpc3RTZXRzEhcuYXBpLnYxLkxpc3RTZXRzUmVxdWVzdBoYLmFwaS52MS5MaXN0U2V0c1Jlc3BvbnNlIgSItRgBQo4BCgpjb20uYXBpLnYxQg1FeGVyY2lzZVByb3RvUAFaOGdpdGh1Yi5jb20vY3Jsc3NuL2dldHN0cm9uZ2VyL3NlcnZlci9wa2cvcGIvYXBpL3YxO2FwaXYxogIDQVhYqgIGQXBpLlYxygIGQXBpXFYx4gISQXBpXFYxXEdQQk1ldGFkYXRh6gIHQXBpOjpWMWIGcHJvdG8z", [file_api_v1_options, file_api_v1_shared, file_google_protobuf_field_mask, file_google_protobuf_timestamp, file_buf_validate_validate]);

/**
 * @generated from message api.v1.CreateExerciseRequest
 */
export type CreateExerciseRequest = Message<"api.v1.CreateExerciseRequest"> & {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string label = 2;
   */
  label: string;
};

/**
 * Describes the message api.v1.CreateExerciseRequest.
 * Use `create(CreateExerciseRequestSchema)` to create a new message.
 */
export const CreateExerciseRequestSchema: GenMessage<CreateExerciseRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 0);

/**
 * @generated from message api.v1.CreateExerciseResponse
 */
export type CreateExerciseResponse = Message<"api.v1.CreateExerciseResponse"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message api.v1.CreateExerciseResponse.
 * Use `create(CreateExerciseResponseSchema)` to create a new message.
 */
export const CreateExerciseResponseSchema: GenMessage<CreateExerciseResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 1);

/**
 * @generated from message api.v1.GetExerciseRequest
 */
export type GetExerciseRequest = Message<"api.v1.GetExerciseRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message api.v1.GetExerciseRequest.
 * Use `create(GetExerciseRequestSchema)` to create a new message.
 */
export const GetExerciseRequestSchema: GenMessage<GetExerciseRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 2);

/**
 * @generated from message api.v1.GetExerciseResponse
 */
export type GetExerciseResponse = Message<"api.v1.GetExerciseResponse"> & {
  /**
   * @generated from field: api.v1.Exercise exercise = 1;
   */
  exercise?: Exercise;
};

/**
 * Describes the message api.v1.GetExerciseResponse.
 * Use `create(GetExerciseResponseSchema)` to create a new message.
 */
export const GetExerciseResponseSchema: GenMessage<GetExerciseResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 3);

/**
 * @generated from message api.v1.UpdateExerciseRequest
 */
export type UpdateExerciseRequest = Message<"api.v1.UpdateExerciseRequest"> & {
  /**
   * @generated from field: api.v1.Exercise exercise = 1;
   */
  exercise?: Exercise;

  /**
   * @generated from field: google.protobuf.FieldMask update_mask = 2;
   */
  updateMask?: FieldMask;
};

/**
 * Describes the message api.v1.UpdateExerciseRequest.
 * Use `create(UpdateExerciseRequestSchema)` to create a new message.
 */
export const UpdateExerciseRequestSchema: GenMessage<UpdateExerciseRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 4);

/**
 * @generated from message api.v1.UpdateExerciseResponse
 */
export type UpdateExerciseResponse = Message<"api.v1.UpdateExerciseResponse"> & {
  /**
   * @generated from field: api.v1.Exercise exercise = 1;
   */
  exercise?: Exercise;
};

/**
 * Describes the message api.v1.UpdateExerciseResponse.
 * Use `create(UpdateExerciseResponseSchema)` to create a new message.
 */
export const UpdateExerciseResponseSchema: GenMessage<UpdateExerciseResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 5);

/**
 * @generated from message api.v1.DeleteExerciseRequest
 */
export type DeleteExerciseRequest = Message<"api.v1.DeleteExerciseRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message api.v1.DeleteExerciseRequest.
 * Use `create(DeleteExerciseRequestSchema)` to create a new message.
 */
export const DeleteExerciseRequestSchema: GenMessage<DeleteExerciseRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 6);

/**
 * @generated from message api.v1.DeleteExerciseResponse
 */
export type DeleteExerciseResponse = Message<"api.v1.DeleteExerciseResponse"> & {
};

/**
 * Describes the message api.v1.DeleteExerciseResponse.
 * Use `create(DeleteExerciseResponseSchema)` to create a new message.
 */
export const DeleteExerciseResponseSchema: GenMessage<DeleteExerciseResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 7);

/**
 * @generated from message api.v1.ListExercisesRequest
 */
export type ListExercisesRequest = Message<"api.v1.ListExercisesRequest"> & {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: repeated string exercise_ids = 2;
   */
  exerciseIds: string[];

  /**
   * @generated from field: int32 page_size = 3;
   */
  pageSize: number;

  /**
   * @generated from field: bytes page_token = 4;
   */
  pageToken: Uint8Array;
};

/**
 * Describes the message api.v1.ListExercisesRequest.
 * Use `create(ListExercisesRequestSchema)` to create a new message.
 */
export const ListExercisesRequestSchema: GenMessage<ListExercisesRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 8);

/**
 * @generated from message api.v1.ListExercisesResponse
 */
export type ListExercisesResponse = Message<"api.v1.ListExercisesResponse"> & {
  /**
   * @generated from field: repeated api.v1.Exercise exercises = 1;
   */
  exercises: Exercise[];

  /**
   * @generated from field: bytes next_page_token = 2;
   */
  nextPageToken: Uint8Array;
};

/**
 * Describes the message api.v1.ListExercisesResponse.
 * Use `create(ListExercisesResponseSchema)` to create a new message.
 */
export const ListExercisesResponseSchema: GenMessage<ListExercisesResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 9);

/**
 * @generated from message api.v1.GetPreviousWorkoutSetsRequest
 */
export type GetPreviousWorkoutSetsRequest = Message<"api.v1.GetPreviousWorkoutSetsRequest"> & {
  /**
   * @generated from field: repeated string exercise_ids = 1;
   */
  exerciseIds: string[];
};

/**
 * Describes the message api.v1.GetPreviousWorkoutSetsRequest.
 * Use `create(GetPreviousWorkoutSetsRequestSchema)` to create a new message.
 */
export const GetPreviousWorkoutSetsRequestSchema: GenMessage<GetPreviousWorkoutSetsRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 10);

/**
 * @generated from message api.v1.GetPreviousWorkoutSetsResponse
 */
export type GetPreviousWorkoutSetsResponse = Message<"api.v1.GetPreviousWorkoutSetsResponse"> & {
  /**
   * @generated from field: repeated api.v1.ExerciseSets exercise_sets = 1;
   */
  exerciseSets: ExerciseSets[];
};

/**
 * Describes the message api.v1.GetPreviousWorkoutSetsResponse.
 * Use `create(GetPreviousWorkoutSetsResponseSchema)` to create a new message.
 */
export const GetPreviousWorkoutSetsResponseSchema: GenMessage<GetPreviousWorkoutSetsResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 11);

/**
 * @generated from message api.v1.GetPersonalBestsRequest
 */
export type GetPersonalBestsRequest = Message<"api.v1.GetPersonalBestsRequest"> & {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;
};

/**
 * Describes the message api.v1.GetPersonalBestsRequest.
 * Use `create(GetPersonalBestsRequestSchema)` to create a new message.
 */
export const GetPersonalBestsRequestSchema: GenMessage<GetPersonalBestsRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 12);

/**
 * @generated from message api.v1.GetPersonalBestsResponse
 */
export type GetPersonalBestsResponse = Message<"api.v1.GetPersonalBestsResponse"> & {
  /**
   * @generated from field: repeated api.v1.PersonalBest personal_bests = 1;
   */
  personalBests: PersonalBest[];
};

/**
 * Describes the message api.v1.GetPersonalBestsResponse.
 * Use `create(GetPersonalBestsResponseSchema)` to create a new message.
 */
export const GetPersonalBestsResponseSchema: GenMessage<GetPersonalBestsResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 13);

/**
 * @generated from message api.v1.ListSetsRequest
 */
export type ListSetsRequest = Message<"api.v1.ListSetsRequest"> & {
  /**
   * @generated from field: string exercise_id = 1;
   */
  exerciseId: string;

  /**
   * @generated from field: api.v1.PaginationRequest pagination = 2;
   */
  pagination?: PaginationRequest;
};

/**
 * Describes the message api.v1.ListSetsRequest.
 * Use `create(ListSetsRequestSchema)` to create a new message.
 */
export const ListSetsRequestSchema: GenMessage<ListSetsRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 14);

/**
 * @generated from message api.v1.ListSetsResponse
 */
export type ListSetsResponse = Message<"api.v1.ListSetsResponse"> & {
  /**
   * @generated from field: repeated api.v1.Set sets = 1;
   */
  sets: Set[];

  /**
   * @generated from field: api.v1.PaginationResponse pagination = 2;
   */
  pagination?: PaginationResponse;
};

/**
 * Describes the message api.v1.ListSetsResponse.
 * Use `create(ListSetsResponseSchema)` to create a new message.
 */
export const ListSetsResponseSchema: GenMessage<ListSetsResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 15);

/**
 * @generated from message api.v1.PersonalBest
 */
export type PersonalBest = Message<"api.v1.PersonalBest"> & {
  /**
   * @generated from field: api.v1.Exercise exercise = 1;
   */
  exercise?: Exercise;

  /**
   * @generated from field: api.v1.Set set = 2;
   */
  set?: Set;
};

/**
 * Describes the message api.v1.PersonalBest.
 * Use `create(PersonalBestSchema)` to create a new message.
 */
export const PersonalBestSchema: GenMessage<PersonalBest> = /*@__PURE__*/
  messageDesc(file_api_v1_exercise, 16);

/**
 * @generated from service api.v1.ExerciseService
 */
export const ExerciseService: GenService<{
  /**
   * @generated from rpc api.v1.ExerciseService.Create
   */
  create: {
    methodKind: "unary";
    input: typeof CreateExerciseRequestSchema;
    output: typeof CreateExerciseResponseSchema;
  },
  /**
   * @generated from rpc api.v1.ExerciseService.Get
   */
  get: {
    methodKind: "unary";
    input: typeof GetExerciseRequestSchema;
    output: typeof GetExerciseResponseSchema;
  },
  /**
   * @generated from rpc api.v1.ExerciseService.Update
   */
  update: {
    methodKind: "unary";
    input: typeof UpdateExerciseRequestSchema;
    output: typeof UpdateExerciseResponseSchema;
  },
  /**
   * @generated from rpc api.v1.ExerciseService.Delete
   */
  delete: {
    methodKind: "unary";
    input: typeof DeleteExerciseRequestSchema;
    output: typeof DeleteExerciseResponseSchema;
  },
  /**
   * @generated from rpc api.v1.ExerciseService.List
   */
  list: {
    methodKind: "unary";
    input: typeof ListExercisesRequestSchema;
    output: typeof ListExercisesResponseSchema;
  },
  /**
   * @generated from rpc api.v1.ExerciseService.GetPreviousWorkoutSets
   */
  getPreviousWorkoutSets: {
    methodKind: "unary";
    input: typeof GetPreviousWorkoutSetsRequestSchema;
    output: typeof GetPreviousWorkoutSetsResponseSchema;
  },
  /**
   * @generated from rpc api.v1.ExerciseService.GetPersonalBests
   */
  getPersonalBests: {
    methodKind: "unary";
    input: typeof GetPersonalBestsRequestSchema;
    output: typeof GetPersonalBestsResponseSchema;
  },
  /**
   * @generated from rpc api.v1.ExerciseService.ListSets
   */
  listSets: {
    methodKind: "unary";
    input: typeof ListSetsRequestSchema;
    output: typeof ListSetsResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_api_v1_exercise, 0);

