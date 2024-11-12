// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file api/v1/exercise.proto (package api.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { FieldMask, Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message api.v1.CreateExerciseRequest
 */
export class CreateExerciseRequest extends Message<CreateExerciseRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: string label = 2;
   */
  label = "";

  /**
   * @generated from field: api.v1.RestBetweenSets rest_between_sets = 3;
   */
  restBetweenSets?: RestBetweenSets;

  constructor(data?: PartialMessage<CreateExerciseRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.CreateExerciseRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "label", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "rest_between_sets", kind: "message", T: RestBetweenSets },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateExerciseRequest {
    return new CreateExerciseRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateExerciseRequest {
    return new CreateExerciseRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateExerciseRequest {
    return new CreateExerciseRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateExerciseRequest | PlainMessage<CreateExerciseRequest> | undefined, b: CreateExerciseRequest | PlainMessage<CreateExerciseRequest> | undefined): boolean {
    return proto3.util.equals(CreateExerciseRequest, a, b);
  }
}

/**
 * @generated from message api.v1.CreateExerciseResponse
 */
export class CreateExerciseResponse extends Message<CreateExerciseResponse> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<CreateExerciseResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.CreateExerciseResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateExerciseResponse {
    return new CreateExerciseResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateExerciseResponse {
    return new CreateExerciseResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateExerciseResponse {
    return new CreateExerciseResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateExerciseResponse | PlainMessage<CreateExerciseResponse> | undefined, b: CreateExerciseResponse | PlainMessage<CreateExerciseResponse> | undefined): boolean {
    return proto3.util.equals(CreateExerciseResponse, a, b);
  }
}

/**
 * @generated from message api.v1.GetExerciseRequest
 */
export class GetExerciseRequest extends Message<GetExerciseRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<GetExerciseRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.GetExerciseRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetExerciseRequest {
    return new GetExerciseRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetExerciseRequest {
    return new GetExerciseRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetExerciseRequest {
    return new GetExerciseRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetExerciseRequest | PlainMessage<GetExerciseRequest> | undefined, b: GetExerciseRequest | PlainMessage<GetExerciseRequest> | undefined): boolean {
    return proto3.util.equals(GetExerciseRequest, a, b);
  }
}

/**
 * @generated from message api.v1.GetExerciseResponse
 */
export class GetExerciseResponse extends Message<GetExerciseResponse> {
  /**
   * @generated from field: api.v1.Exercise exercise = 1;
   */
  exercise?: Exercise;

  constructor(data?: PartialMessage<GetExerciseResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.GetExerciseResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "exercise", kind: "message", T: Exercise },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetExerciseResponse {
    return new GetExerciseResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetExerciseResponse {
    return new GetExerciseResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetExerciseResponse {
    return new GetExerciseResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetExerciseResponse | PlainMessage<GetExerciseResponse> | undefined, b: GetExerciseResponse | PlainMessage<GetExerciseResponse> | undefined): boolean {
    return proto3.util.equals(GetExerciseResponse, a, b);
  }
}

/**
 * @generated from message api.v1.UpdateExerciseRequest
 */
export class UpdateExerciseRequest extends Message<UpdateExerciseRequest> {
  /**
   * @generated from field: api.v1.Exercise exercise = 1;
   */
  exercise?: Exercise;

  /**
   * @generated from field: google.protobuf.FieldMask update_mask = 2;
   */
  updateMask?: FieldMask;

  constructor(data?: PartialMessage<UpdateExerciseRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.UpdateExerciseRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "exercise", kind: "message", T: Exercise },
    { no: 2, name: "update_mask", kind: "message", T: FieldMask },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateExerciseRequest {
    return new UpdateExerciseRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateExerciseRequest {
    return new UpdateExerciseRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateExerciseRequest {
    return new UpdateExerciseRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateExerciseRequest | PlainMessage<UpdateExerciseRequest> | undefined, b: UpdateExerciseRequest | PlainMessage<UpdateExerciseRequest> | undefined): boolean {
    return proto3.util.equals(UpdateExerciseRequest, a, b);
  }
}

/**
 * @generated from message api.v1.UpdateExerciseResponse
 */
export class UpdateExerciseResponse extends Message<UpdateExerciseResponse> {
  /**
   * @generated from field: api.v1.Exercise exercise = 1;
   */
  exercise?: Exercise;

  constructor(data?: PartialMessage<UpdateExerciseResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.UpdateExerciseResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "exercise", kind: "message", T: Exercise },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateExerciseResponse {
    return new UpdateExerciseResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateExerciseResponse {
    return new UpdateExerciseResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateExerciseResponse {
    return new UpdateExerciseResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateExerciseResponse | PlainMessage<UpdateExerciseResponse> | undefined, b: UpdateExerciseResponse | PlainMessage<UpdateExerciseResponse> | undefined): boolean {
    return proto3.util.equals(UpdateExerciseResponse, a, b);
  }
}

/**
 * @generated from message api.v1.DeleteExerciseRequest
 */
export class DeleteExerciseRequest extends Message<DeleteExerciseRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<DeleteExerciseRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.DeleteExerciseRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteExerciseRequest {
    return new DeleteExerciseRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteExerciseRequest {
    return new DeleteExerciseRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteExerciseRequest {
    return new DeleteExerciseRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteExerciseRequest | PlainMessage<DeleteExerciseRequest> | undefined, b: DeleteExerciseRequest | PlainMessage<DeleteExerciseRequest> | undefined): boolean {
    return proto3.util.equals(DeleteExerciseRequest, a, b);
  }
}

/**
 * @generated from message api.v1.DeleteExerciseResponse
 */
export class DeleteExerciseResponse extends Message<DeleteExerciseResponse> {
  constructor(data?: PartialMessage<DeleteExerciseResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.DeleteExerciseResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteExerciseResponse {
    return new DeleteExerciseResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteExerciseResponse {
    return new DeleteExerciseResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteExerciseResponse {
    return new DeleteExerciseResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteExerciseResponse | PlainMessage<DeleteExerciseResponse> | undefined, b: DeleteExerciseResponse | PlainMessage<DeleteExerciseResponse> | undefined): boolean {
    return proto3.util.equals(DeleteExerciseResponse, a, b);
  }
}

/**
 * @generated from message api.v1.ListExercisesRequest
 */
export class ListExercisesRequest extends Message<ListExercisesRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: bytes page_token = 2;
   */
  pageToken = new Uint8Array(0);

  constructor(data?: PartialMessage<ListExercisesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.ListExercisesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "page_token", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListExercisesRequest {
    return new ListExercisesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListExercisesRequest {
    return new ListExercisesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListExercisesRequest {
    return new ListExercisesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListExercisesRequest | PlainMessage<ListExercisesRequest> | undefined, b: ListExercisesRequest | PlainMessage<ListExercisesRequest> | undefined): boolean {
    return proto3.util.equals(ListExercisesRequest, a, b);
  }
}

/**
 * @generated from message api.v1.ListExercisesResponse
 */
export class ListExercisesResponse extends Message<ListExercisesResponse> {
  /**
   * @generated from field: repeated api.v1.Exercise exercises = 1;
   */
  exercises: Exercise[] = [];

  /**
   * @generated from field: bytes next_page_token = 2;
   */
  nextPageToken = new Uint8Array(0);

  constructor(data?: PartialMessage<ListExercisesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.ListExercisesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "exercises", kind: "message", T: Exercise, repeated: true },
    { no: 2, name: "next_page_token", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListExercisesResponse {
    return new ListExercisesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListExercisesResponse {
    return new ListExercisesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListExercisesResponse {
    return new ListExercisesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListExercisesResponse | PlainMessage<ListExercisesResponse> | undefined, b: ListExercisesResponse | PlainMessage<ListExercisesResponse> | undefined): boolean {
    return proto3.util.equals(ListExercisesResponse, a, b);
  }
}

/**
 * @generated from message api.v1.Exercise
 */
export class Exercise extends Message<Exercise> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string label = 3;
   */
  label = "";

  /**
   * @generated from field: api.v1.RestBetweenSets rest_between_sets = 4;
   */
  restBetweenSets?: RestBetweenSets;

  constructor(data?: PartialMessage<Exercise>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.Exercise";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "label", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "rest_between_sets", kind: "message", T: RestBetweenSets },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Exercise {
    return new Exercise().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Exercise {
    return new Exercise().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Exercise {
    return new Exercise().fromJsonString(jsonString, options);
  }

  static equals(a: Exercise | PlainMessage<Exercise> | undefined, b: Exercise | PlainMessage<Exercise> | undefined): boolean {
    return proto3.util.equals(Exercise, a, b);
  }
}

/**
 * @generated from message api.v1.RestBetweenSets
 */
export class RestBetweenSets extends Message<RestBetweenSets> {
  /**
   * @generated from field: int32 seconds = 1;
   */
  seconds = 0;

  constructor(data?: PartialMessage<RestBetweenSets>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.RestBetweenSets";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "seconds", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RestBetweenSets {
    return new RestBetweenSets().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RestBetweenSets {
    return new RestBetweenSets().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RestBetweenSets {
    return new RestBetweenSets().fromJsonString(jsonString, options);
  }

  static equals(a: RestBetweenSets | PlainMessage<RestBetweenSets> | undefined, b: RestBetweenSets | PlainMessage<RestBetweenSets> | undefined): boolean {
    return proto3.util.equals(RestBetweenSets, a, b);
  }
}

