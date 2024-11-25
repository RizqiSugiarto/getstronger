// @generated by protoc-gen-connect-es v1.6.1 with parameter "target=ts"
// @generated from file api/v1/exercise.proto (package api.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateExerciseRequest, CreateExerciseResponse, DeleteExerciseRequest, DeleteExerciseResponse, GetExerciseRequest, GetExerciseResponse, GetPersonalBestsRequest, GetPersonalBestsResponse, GetPreviousWorkoutSetsRequest, GetPreviousWorkoutSetsResponse, ListExercisesRequest, ListExercisesResponse, UpdateExerciseRequest, UpdateExerciseResponse } from "./exercise_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service api.v1.ExerciseService
 */
export const ExerciseService = {
  typeName: "api.v1.ExerciseService",
  methods: {
    /**
     * @generated from rpc api.v1.ExerciseService.Create
     */
    create: {
      name: "Create",
      I: CreateExerciseRequest,
      O: CreateExerciseResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.ExerciseService.Get
     */
    get: {
      name: "Get",
      I: GetExerciseRequest,
      O: GetExerciseResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.ExerciseService.Update
     */
    update: {
      name: "Update",
      I: UpdateExerciseRequest,
      O: UpdateExerciseResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.ExerciseService.Delete
     */
    delete: {
      name: "Delete",
      I: DeleteExerciseRequest,
      O: DeleteExerciseResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.ExerciseService.List
     */
    list: {
      name: "List",
      I: ListExercisesRequest,
      O: ListExercisesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.ExerciseService.GetPreviousWorkoutSets
     */
    getPreviousWorkoutSets: {
      name: "GetPreviousWorkoutSets",
      I: GetPreviousWorkoutSetsRequest,
      O: GetPreviousWorkoutSetsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.v1.ExerciseService.GetPersonalBests
     */
    getPersonalBests: {
      name: "GetPersonalBests",
      I: GetPersonalBestsRequest,
      O: GetPersonalBestsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

