// @generated by protoc-gen-connect-es v1.6.1 with parameter "target=ts"
// @generated from file api/v1/workouts.proto (package api.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateWorkoutRequest, CreateWorkoutResponse, DeleteWorkoutRequest, DeleteWorkoutResponse, GetWorkoutRequest, GetWorkoutResponse, ListWorkoutsRequest, ListWorkoutsResponse } from "./workouts_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service api.v1.WorkoutService
 */
export const WorkoutService = {
  typeName: "api.v1.WorkoutService",
  methods: {
    /**
     * @generated from rpc api.v1.WorkoutService.Create
     */
    create: {
      name: "Create",
      I: CreateWorkoutRequest,
      O: CreateWorkoutResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.WorkoutService.Get
     */
    get: {
      name: "Get",
      I: GetWorkoutRequest,
      O: GetWorkoutResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.WorkoutService.List
     */
    list: {
      name: "List",
      I: ListWorkoutsRequest,
      O: ListWorkoutsResponse,
      kind: MethodKind.Unary,
    },
    /**
     *  rpc Start(StartWorkoutRequest) returns (StartWorkoutResponse) {
     *    option (auth) = true;
     *  }
     *  rpc Finish(FinishWorkoutRequest) returns (FinishWorkoutResponse) {
     *    option (auth) = true;
     *  }
     *
     * @generated from rpc api.v1.WorkoutService.Delete
     */
    delete: {
      name: "Delete",
      I: DeleteWorkoutRequest,
      O: DeleteWorkoutResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

