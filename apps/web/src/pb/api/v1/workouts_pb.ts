// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file api/v1/workouts.proto (package api.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message api.v1.CreateWorkoutRequest
 */
export class CreateWorkoutRequest extends Message<CreateWorkoutRequest> {
  /**
   * @generated from field: string routine_id = 1;
   */
  routineId = "";

  /**
   * @generated from field: repeated api.v1.ExerciseSets exercise_sets = 2;
   */
  exerciseSets: ExerciseSets[] = [];

  /**
   * @generated from field: google.protobuf.Timestamp finished_at = 3;
   */
  finishedAt?: Timestamp;

  constructor(data?: PartialMessage<CreateWorkoutRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.CreateWorkoutRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "routine_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "exercise_sets", kind: "message", T: ExerciseSets, repeated: true },
    { no: 3, name: "finished_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateWorkoutRequest {
    return new CreateWorkoutRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateWorkoutRequest {
    return new CreateWorkoutRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateWorkoutRequest {
    return new CreateWorkoutRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateWorkoutRequest | PlainMessage<CreateWorkoutRequest> | undefined, b: CreateWorkoutRequest | PlainMessage<CreateWorkoutRequest> | undefined): boolean {
    return proto3.util.equals(CreateWorkoutRequest, a, b);
  }
}

/**
 * @generated from message api.v1.CreateWorkoutResponse
 */
export class CreateWorkoutResponse extends Message<CreateWorkoutResponse> {
  /**
   * @generated from field: string workout_id = 1;
   */
  workoutId = "";

  constructor(data?: PartialMessage<CreateWorkoutResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.CreateWorkoutResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "workout_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateWorkoutResponse {
    return new CreateWorkoutResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateWorkoutResponse {
    return new CreateWorkoutResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateWorkoutResponse {
    return new CreateWorkoutResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateWorkoutResponse | PlainMessage<CreateWorkoutResponse> | undefined, b: CreateWorkoutResponse | PlainMessage<CreateWorkoutResponse> | undefined): boolean {
    return proto3.util.equals(CreateWorkoutResponse, a, b);
  }
}

/**
 * @generated from message api.v1.ListWorkoutsRequest
 */
export class ListWorkoutsRequest extends Message<ListWorkoutsRequest> {
  /**
   * @generated from field: int32 page_size = 1;
   */
  pageSize = 0;

  /**
   * @generated from field: bytes page_token = 2;
   */
  pageToken = new Uint8Array(0);

  constructor(data?: PartialMessage<ListWorkoutsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.ListWorkoutsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "page_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "page_token", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListWorkoutsRequest {
    return new ListWorkoutsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListWorkoutsRequest {
    return new ListWorkoutsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListWorkoutsRequest {
    return new ListWorkoutsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListWorkoutsRequest | PlainMessage<ListWorkoutsRequest> | undefined, b: ListWorkoutsRequest | PlainMessage<ListWorkoutsRequest> | undefined): boolean {
    return proto3.util.equals(ListWorkoutsRequest, a, b);
  }
}

/**
 * @generated from message api.v1.ListWorkoutsResponse
 */
export class ListWorkoutsResponse extends Message<ListWorkoutsResponse> {
  /**
   * @generated from field: repeated api.v1.Workout workouts = 1;
   */
  workouts: Workout[] = [];

  /**
   * @generated from field: bytes next_page_token = 2;
   */
  nextPageToken = new Uint8Array(0);

  constructor(data?: PartialMessage<ListWorkoutsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.ListWorkoutsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "workouts", kind: "message", T: Workout, repeated: true },
    { no: 2, name: "next_page_token", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListWorkoutsResponse {
    return new ListWorkoutsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListWorkoutsResponse {
    return new ListWorkoutsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListWorkoutsResponse {
    return new ListWorkoutsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListWorkoutsResponse | PlainMessage<ListWorkoutsResponse> | undefined, b: ListWorkoutsResponse | PlainMessage<ListWorkoutsResponse> | undefined): boolean {
    return proto3.util.equals(ListWorkoutsResponse, a, b);
  }
}

/**
 * @generated from message api.v1.GetWorkoutRequest
 */
export class GetWorkoutRequest extends Message<GetWorkoutRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<GetWorkoutRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.GetWorkoutRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetWorkoutRequest {
    return new GetWorkoutRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetWorkoutRequest {
    return new GetWorkoutRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetWorkoutRequest {
    return new GetWorkoutRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetWorkoutRequest | PlainMessage<GetWorkoutRequest> | undefined, b: GetWorkoutRequest | PlainMessage<GetWorkoutRequest> | undefined): boolean {
    return proto3.util.equals(GetWorkoutRequest, a, b);
  }
}

/**
 * @generated from message api.v1.GetWorkoutResponse
 */
export class GetWorkoutResponse extends Message<GetWorkoutResponse> {
  /**
   * @generated from field: api.v1.Workout workout = 1;
   */
  workout?: Workout;

  constructor(data?: PartialMessage<GetWorkoutResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.GetWorkoutResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "workout", kind: "message", T: Workout },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetWorkoutResponse {
    return new GetWorkoutResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetWorkoutResponse {
    return new GetWorkoutResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetWorkoutResponse {
    return new GetWorkoutResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetWorkoutResponse | PlainMessage<GetWorkoutResponse> | undefined, b: GetWorkoutResponse | PlainMessage<GetWorkoutResponse> | undefined): boolean {
    return proto3.util.equals(GetWorkoutResponse, a, b);
  }
}

/**
 * @generated from message api.v1.Workout
 */
export class Workout extends Message<Workout> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: google.protobuf.Timestamp finished_at = 3;
   */
  finishedAt?: Timestamp;

  /**
   * @generated from field: repeated api.v1.ExerciseSets exercise_sets = 4;
   */
  exerciseSets: ExerciseSets[] = [];

  constructor(data?: PartialMessage<Workout>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.Workout";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "finished_at", kind: "message", T: Timestamp },
    { no: 4, name: "exercise_sets", kind: "message", T: ExerciseSets, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Workout {
    return new Workout().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Workout {
    return new Workout().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Workout {
    return new Workout().fromJsonString(jsonString, options);
  }

  static equals(a: Workout | PlainMessage<Workout> | undefined, b: Workout | PlainMessage<Workout> | undefined): boolean {
    return proto3.util.equals(Workout, a, b);
  }
}

/**
 * @generated from message api.v1.ExerciseSets
 */
export class ExerciseSets extends Message<ExerciseSets> {
  /**
   * @generated from field: string exercise_id = 1;
   */
  exerciseId = "";

  /**
   * @generated from field: repeated api.v1.Set sets = 2;
   */
  sets: Set[] = [];

  constructor(data?: PartialMessage<ExerciseSets>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.ExerciseSets";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "exercise_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "sets", kind: "message", T: Set, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExerciseSets {
    return new ExerciseSets().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExerciseSets {
    return new ExerciseSets().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExerciseSets {
    return new ExerciseSets().fromJsonString(jsonString, options);
  }

  static equals(a: ExerciseSets | PlainMessage<ExerciseSets> | undefined, b: ExerciseSets | PlainMessage<ExerciseSets> | undefined): boolean {
    return proto3.util.equals(ExerciseSets, a, b);
  }
}

/**
 * @generated from message api.v1.Set
 */
export class Set extends Message<Set> {
  /**
   * The weight can be less than zero.
   *
   * @generated from field: double weight = 1;
   */
  weight = 0;

  /**
   * @generated from field: int32 reps = 2;
   */
  reps = 0;

  constructor(data?: PartialMessage<Set>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1.Set";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "weight", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 2, name: "reps", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Set {
    return new Set().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Set {
    return new Set().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Set {
    return new Set().fromJsonString(jsonString, options);
  }

  static equals(a: Set | PlainMessage<Set> | undefined, b: Set | PlainMessage<Set> | undefined): boolean {
    return proto3.util.equals(Set, a, b);
  }
}

