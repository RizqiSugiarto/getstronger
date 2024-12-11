// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file api/v1/routine_service.proto (package api.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_api_v1_options } from "./options_pb";
import type { Exercise, PaginationRequest, PaginationResponse } from "./shared_pb";
import { file_api_v1_shared } from "./shared_pb";
import { file_buf_validate_validate } from "../../buf/validate/validate_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file api/v1/routine_service.proto.
 */
export const file_api_v1_routine_service: GenFile = /*@__PURE__*/
  fileDesc("ChxhcGkvdjEvcm91dGluZV9zZXJ2aWNlLnByb3RvEgZhcGkudjEiTQoUQ3JlYXRlUm91dGluZVJlcXVlc3QSFQoEbmFtZRgBIAEoCUIHukgEcgIQARIeCgxleGVyY2lzZV9pZHMYAiADKAlCCLpIBZIBAggBIiMKFUNyZWF0ZVJvdXRpbmVSZXNwb25zZRIKCgJpZBgBIAEoCSIpChFHZXRSb3V0aW5lUmVxdWVzdBIUCgJpZBgBIAEoCUIIukgFcgOwAQEiNgoSR2V0Um91dGluZVJlc3BvbnNlEiAKB3JvdXRpbmUYASABKAsyDy5hcGkudjEuUm91dGluZSJAChRVcGRhdGVSb3V0aW5lUmVxdWVzdBIoCgdyb3V0aW5lGAEgASgLMg8uYXBpLnYxLlJvdXRpbmVCBrpIA8gBASI5ChVVcGRhdGVSb3V0aW5lUmVzcG9uc2USIAoHcm91dGluZRgBIAEoCzIPLmFwaS52MS5Sb3V0aW5lIiwKFERlbGV0ZVJvdXRpbmVSZXF1ZXN0EhQKAmlkGAEgASgJQgi6SAVyA7ABASIXChVEZWxldGVSb3V0aW5lUmVzcG9uc2UiWgoTTGlzdFJvdXRpbmVzUmVxdWVzdBIMCgRuYW1lGAEgASgJEjUKCnBhZ2luYXRpb24YAiABKAsyGS5hcGkudjEuUGFnaW5hdGlvblJlcXVlc3RCBrpIA8gBASJpChRMaXN0Um91dGluZXNSZXNwb25zZRIhCghyb3V0aW5lcxgBIAMoCzIPLmFwaS52MS5Sb3V0aW5lEi4KCnBhZ2luYXRpb24YAiABKAsyGi5hcGkudjEuUGFnaW5hdGlvblJlc3BvbnNlIlEKEkFkZEV4ZXJjaXNlUmVxdWVzdBIcCgpyb3V0aW5lX2lkGAEgASgJQgi6SAVyA7ABARIdCgtleGVyY2lzZV9pZBgCIAEoCUIIukgFcgOwAQEiFQoTQWRkRXhlcmNpc2VSZXNwb25zZSJUChVSZW1vdmVFeGVyY2lzZVJlcXVlc3QSHAoKcm91dGluZV9pZBgBIAEoCUIIukgFcgOwAQESHQoLZXhlcmNpc2VfaWQYAiABKAlCCLpIBXIDsAEBIhgKFlJlbW92ZUV4ZXJjaXNlUmVzcG9uc2UiWgoaVXBkYXRlRXhlcmNpc2VPcmRlclJlcXVlc3QSHAoKcm91dGluZV9pZBgBIAEoCUIIukgFcgOwAQESHgoMZXhlcmNpc2VfaWRzGAIgAygJQgi6SAWSAQIIASIdChtVcGRhdGVFeGVyY2lzZU9yZGVyUmVzcG9uc2UiZQoHUm91dGluZRIUCgJpZBgBIAEoCUIIukgFcgOwAQESFQoEbmFtZRgCIAEoCUIHukgEcgIQARItCglleGVyY2lzZXMYAyADKAsyEC5hcGkudjEuRXhlcmNpc2VCCLpIBZIBAggBMrMFCg5Sb3V0aW5lU2VydmljZRJSCg1DcmVhdGVSb3V0aW5lEhwuYXBpLnYxLkNyZWF0ZVJvdXRpbmVSZXF1ZXN0Gh0uYXBpLnYxLkNyZWF0ZVJvdXRpbmVSZXNwb25zZSIEiLUYARJJCgpHZXRSb3V0aW5lEhkuYXBpLnYxLkdldFJvdXRpbmVSZXF1ZXN0GhouYXBpLnYxLkdldFJvdXRpbmVSZXNwb25zZSIEiLUYARJSCg1VcGRhdGVSb3V0aW5lEhwuYXBpLnYxLlVwZGF0ZVJvdXRpbmVSZXF1ZXN0Gh0uYXBpLnYxLlVwZGF0ZVJvdXRpbmVSZXNwb25zZSIEiLUYARJSCg1EZWxldGVSb3V0aW5lEhwuYXBpLnYxLkRlbGV0ZVJvdXRpbmVSZXF1ZXN0Gh0uYXBpLnYxLkRlbGV0ZVJvdXRpbmVSZXNwb25zZSIEiLUYARJPCgxMaXN0Um91dGluZXMSGy5hcGkudjEuTGlzdFJvdXRpbmVzUmVxdWVzdBocLmFwaS52MS5MaXN0Um91dGluZXNSZXNwb25zZSIEiLUYARJMCgtBZGRFeGVyY2lzZRIaLmFwaS52MS5BZGRFeGVyY2lzZVJlcXVlc3QaGy5hcGkudjEuQWRkRXhlcmNpc2VSZXNwb25zZSIEiLUYARJVCg5SZW1vdmVFeGVyY2lzZRIdLmFwaS52MS5SZW1vdmVFeGVyY2lzZVJlcXVlc3QaHi5hcGkudjEuUmVtb3ZlRXhlcmNpc2VSZXNwb25zZSIEiLUYARJkChNVcGRhdGVFeGVyY2lzZU9yZGVyEiIuYXBpLnYxLlVwZGF0ZUV4ZXJjaXNlT3JkZXJSZXF1ZXN0GiMuYXBpLnYxLlVwZGF0ZUV4ZXJjaXNlT3JkZXJSZXNwb25zZSIEiLUYAUKXAQoKY29tLmFwaS52MUITUm91dGluZVNlcnZpY2VQcm90b1ABWjtnaXRodWIuY29tL2NybHNzbi9nZXRzdHJvbmdlci9zZXJ2ZXIvcGtnL3Byb3RvL2FwaS92MTthcGl2MaICA0FYWKoCBkFwaS5WMcoCBkFwaVxWMeICEkFwaVxWMVxHUEJNZXRhZGF0YeoCB0FwaTo6VjFiBnByb3RvMw", [file_api_v1_options, file_api_v1_shared, file_buf_validate_validate]);

/**
 * @generated from message api.v1.CreateRoutineRequest
 */
export type CreateRoutineRequest = Message<"api.v1.CreateRoutineRequest"> & {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: repeated string exercise_ids = 2;
   */
  exerciseIds: string[];
};

/**
 * Describes the message api.v1.CreateRoutineRequest.
 * Use `create(CreateRoutineRequestSchema)` to create a new message.
 */
export const CreateRoutineRequestSchema: GenMessage<CreateRoutineRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 0);

/**
 * @generated from message api.v1.CreateRoutineResponse
 */
export type CreateRoutineResponse = Message<"api.v1.CreateRoutineResponse"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message api.v1.CreateRoutineResponse.
 * Use `create(CreateRoutineResponseSchema)` to create a new message.
 */
export const CreateRoutineResponseSchema: GenMessage<CreateRoutineResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 1);

/**
 * @generated from message api.v1.GetRoutineRequest
 */
export type GetRoutineRequest = Message<"api.v1.GetRoutineRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message api.v1.GetRoutineRequest.
 * Use `create(GetRoutineRequestSchema)` to create a new message.
 */
export const GetRoutineRequestSchema: GenMessage<GetRoutineRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 2);

/**
 * @generated from message api.v1.GetRoutineResponse
 */
export type GetRoutineResponse = Message<"api.v1.GetRoutineResponse"> & {
  /**
   * @generated from field: api.v1.Routine routine = 1;
   */
  routine?: Routine;
};

/**
 * Describes the message api.v1.GetRoutineResponse.
 * Use `create(GetRoutineResponseSchema)` to create a new message.
 */
export const GetRoutineResponseSchema: GenMessage<GetRoutineResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 3);

/**
 * @generated from message api.v1.UpdateRoutineRequest
 */
export type UpdateRoutineRequest = Message<"api.v1.UpdateRoutineRequest"> & {
  /**
   * @generated from field: api.v1.Routine routine = 1;
   */
  routine?: Routine;
};

/**
 * Describes the message api.v1.UpdateRoutineRequest.
 * Use `create(UpdateRoutineRequestSchema)` to create a new message.
 */
export const UpdateRoutineRequestSchema: GenMessage<UpdateRoutineRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 4);

/**
 * @generated from message api.v1.UpdateRoutineResponse
 */
export type UpdateRoutineResponse = Message<"api.v1.UpdateRoutineResponse"> & {
  /**
   * @generated from field: api.v1.Routine routine = 1;
   */
  routine?: Routine;
};

/**
 * Describes the message api.v1.UpdateRoutineResponse.
 * Use `create(UpdateRoutineResponseSchema)` to create a new message.
 */
export const UpdateRoutineResponseSchema: GenMessage<UpdateRoutineResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 5);

/**
 * @generated from message api.v1.DeleteRoutineRequest
 */
export type DeleteRoutineRequest = Message<"api.v1.DeleteRoutineRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message api.v1.DeleteRoutineRequest.
 * Use `create(DeleteRoutineRequestSchema)` to create a new message.
 */
export const DeleteRoutineRequestSchema: GenMessage<DeleteRoutineRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 6);

/**
 * @generated from message api.v1.DeleteRoutineResponse
 */
export type DeleteRoutineResponse = Message<"api.v1.DeleteRoutineResponse"> & {
};

/**
 * Describes the message api.v1.DeleteRoutineResponse.
 * Use `create(DeleteRoutineResponseSchema)` to create a new message.
 */
export const DeleteRoutineResponseSchema: GenMessage<DeleteRoutineResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 7);

/**
 * @generated from message api.v1.ListRoutinesRequest
 */
export type ListRoutinesRequest = Message<"api.v1.ListRoutinesRequest"> & {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: api.v1.PaginationRequest pagination = 2;
   */
  pagination?: PaginationRequest;
};

/**
 * Describes the message api.v1.ListRoutinesRequest.
 * Use `create(ListRoutinesRequestSchema)` to create a new message.
 */
export const ListRoutinesRequestSchema: GenMessage<ListRoutinesRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 8);

/**
 * @generated from message api.v1.ListRoutinesResponse
 */
export type ListRoutinesResponse = Message<"api.v1.ListRoutinesResponse"> & {
  /**
   * @generated from field: repeated api.v1.Routine routines = 1;
   */
  routines: Routine[];

  /**
   * @generated from field: api.v1.PaginationResponse pagination = 2;
   */
  pagination?: PaginationResponse;
};

/**
 * Describes the message api.v1.ListRoutinesResponse.
 * Use `create(ListRoutinesResponseSchema)` to create a new message.
 */
export const ListRoutinesResponseSchema: GenMessage<ListRoutinesResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 9);

/**
 * @generated from message api.v1.AddExerciseRequest
 */
export type AddExerciseRequest = Message<"api.v1.AddExerciseRequest"> & {
  /**
   * @generated from field: string routine_id = 1;
   */
  routineId: string;

  /**
   * @generated from field: string exercise_id = 2;
   */
  exerciseId: string;
};

/**
 * Describes the message api.v1.AddExerciseRequest.
 * Use `create(AddExerciseRequestSchema)` to create a new message.
 */
export const AddExerciseRequestSchema: GenMessage<AddExerciseRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 10);

/**
 * @generated from message api.v1.AddExerciseResponse
 */
export type AddExerciseResponse = Message<"api.v1.AddExerciseResponse"> & {
};

/**
 * Describes the message api.v1.AddExerciseResponse.
 * Use `create(AddExerciseResponseSchema)` to create a new message.
 */
export const AddExerciseResponseSchema: GenMessage<AddExerciseResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 11);

/**
 * @generated from message api.v1.RemoveExerciseRequest
 */
export type RemoveExerciseRequest = Message<"api.v1.RemoveExerciseRequest"> & {
  /**
   * @generated from field: string routine_id = 1;
   */
  routineId: string;

  /**
   * @generated from field: string exercise_id = 2;
   */
  exerciseId: string;
};

/**
 * Describes the message api.v1.RemoveExerciseRequest.
 * Use `create(RemoveExerciseRequestSchema)` to create a new message.
 */
export const RemoveExerciseRequestSchema: GenMessage<RemoveExerciseRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 12);

/**
 * @generated from message api.v1.RemoveExerciseResponse
 */
export type RemoveExerciseResponse = Message<"api.v1.RemoveExerciseResponse"> & {
};

/**
 * Describes the message api.v1.RemoveExerciseResponse.
 * Use `create(RemoveExerciseResponseSchema)` to create a new message.
 */
export const RemoveExerciseResponseSchema: GenMessage<RemoveExerciseResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 13);

/**
 * @generated from message api.v1.UpdateExerciseOrderRequest
 */
export type UpdateExerciseOrderRequest = Message<"api.v1.UpdateExerciseOrderRequest"> & {
  /**
   * @generated from field: string routine_id = 1;
   */
  routineId: string;

  /**
   * @generated from field: repeated string exercise_ids = 2;
   */
  exerciseIds: string[];
};

/**
 * Describes the message api.v1.UpdateExerciseOrderRequest.
 * Use `create(UpdateExerciseOrderRequestSchema)` to create a new message.
 */
export const UpdateExerciseOrderRequestSchema: GenMessage<UpdateExerciseOrderRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 14);

/**
 * @generated from message api.v1.UpdateExerciseOrderResponse
 */
export type UpdateExerciseOrderResponse = Message<"api.v1.UpdateExerciseOrderResponse"> & {
};

/**
 * Describes the message api.v1.UpdateExerciseOrderResponse.
 * Use `create(UpdateExerciseOrderResponseSchema)` to create a new message.
 */
export const UpdateExerciseOrderResponseSchema: GenMessage<UpdateExerciseOrderResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 15);

/**
 * @generated from message api.v1.Routine
 */
export type Routine = Message<"api.v1.Routine"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * @generated from field: repeated api.v1.Exercise exercises = 3;
   */
  exercises: Exercise[];
};

/**
 * Describes the message api.v1.Routine.
 * Use `create(RoutineSchema)` to create a new message.
 */
export const RoutineSchema: GenMessage<Routine> = /*@__PURE__*/
  messageDesc(file_api_v1_routine_service, 16);

/**
 * @generated from service api.v1.RoutineService
 */
export const RoutineService: GenService<{
  /**
   * @generated from rpc api.v1.RoutineService.CreateRoutine
   */
  createRoutine: {
    methodKind: "unary";
    input: typeof CreateRoutineRequestSchema;
    output: typeof CreateRoutineResponseSchema;
  },
  /**
   * @generated from rpc api.v1.RoutineService.GetRoutine
   */
  getRoutine: {
    methodKind: "unary";
    input: typeof GetRoutineRequestSchema;
    output: typeof GetRoutineResponseSchema;
  },
  /**
   * @generated from rpc api.v1.RoutineService.UpdateRoutine
   */
  updateRoutine: {
    methodKind: "unary";
    input: typeof UpdateRoutineRequestSchema;
    output: typeof UpdateRoutineResponseSchema;
  },
  /**
   * @generated from rpc api.v1.RoutineService.DeleteRoutine
   */
  deleteRoutine: {
    methodKind: "unary";
    input: typeof DeleteRoutineRequestSchema;
    output: typeof DeleteRoutineResponseSchema;
  },
  /**
   * @generated from rpc api.v1.RoutineService.ListRoutines
   */
  listRoutines: {
    methodKind: "unary";
    input: typeof ListRoutinesRequestSchema;
    output: typeof ListRoutinesResponseSchema;
  },
  /**
   * @generated from rpc api.v1.RoutineService.AddExercise
   */
  addExercise: {
    methodKind: "unary";
    input: typeof AddExerciseRequestSchema;
    output: typeof AddExerciseResponseSchema;
  },
  /**
   * @generated from rpc api.v1.RoutineService.RemoveExercise
   */
  removeExercise: {
    methodKind: "unary";
    input: typeof RemoveExerciseRequestSchema;
    output: typeof RemoveExerciseResponseSchema;
  },
  /**
   * @generated from rpc api.v1.RoutineService.UpdateExerciseOrder
   */
  updateExerciseOrder: {
    methodKind: "unary";
    input: typeof UpdateExerciseOrderRequestSchema;
    output: typeof UpdateExerciseOrderResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_api_v1_routine_service, 0);

