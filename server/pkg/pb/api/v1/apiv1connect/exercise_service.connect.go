// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/exercise_service.proto

package apiv1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"
	v1 "github.com/crlssn/getstronger/server/pkg/pb/api/v1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ExerciseServiceName is the fully-qualified name of the ExerciseService service.
	ExerciseServiceName = "api.v1.ExerciseService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ExerciseServiceCreateExerciseProcedure is the fully-qualified name of the ExerciseService's
	// CreateExercise RPC.
	ExerciseServiceCreateExerciseProcedure = "/api.v1.ExerciseService/CreateExercise"
	// ExerciseServiceGetExerciseProcedure is the fully-qualified name of the ExerciseService's
	// GetExercise RPC.
	ExerciseServiceGetExerciseProcedure = "/api.v1.ExerciseService/GetExercise"
	// ExerciseServiceUpdateExerciseProcedure is the fully-qualified name of the ExerciseService's
	// UpdateExercise RPC.
	ExerciseServiceUpdateExerciseProcedure = "/api.v1.ExerciseService/UpdateExercise"
	// ExerciseServiceDeleteExerciseProcedure is the fully-qualified name of the ExerciseService's
	// DeleteExercise RPC.
	ExerciseServiceDeleteExerciseProcedure = "/api.v1.ExerciseService/DeleteExercise"
	// ExerciseServiceListExercisesProcedure is the fully-qualified name of the ExerciseService's
	// ListExercises RPC.
	ExerciseServiceListExercisesProcedure = "/api.v1.ExerciseService/ListExercises"
	// ExerciseServiceGetPreviousWorkoutSetsProcedure is the fully-qualified name of the
	// ExerciseService's GetPreviousWorkoutSets RPC.
	ExerciseServiceGetPreviousWorkoutSetsProcedure = "/api.v1.ExerciseService/GetPreviousWorkoutSets"
	// ExerciseServiceGetPersonalBestsProcedure is the fully-qualified name of the ExerciseService's
	// GetPersonalBests RPC.
	ExerciseServiceGetPersonalBestsProcedure = "/api.v1.ExerciseService/GetPersonalBests"
	// ExerciseServiceListSetsProcedure is the fully-qualified name of the ExerciseService's ListSets
	// RPC.
	ExerciseServiceListSetsProcedure = "/api.v1.ExerciseService/ListSets"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	exerciseServiceServiceDescriptor                      = v1.File_api_v1_exercise_service_proto.Services().ByName("ExerciseService")
	exerciseServiceCreateExerciseMethodDescriptor         = exerciseServiceServiceDescriptor.Methods().ByName("CreateExercise")
	exerciseServiceGetExerciseMethodDescriptor            = exerciseServiceServiceDescriptor.Methods().ByName("GetExercise")
	exerciseServiceUpdateExerciseMethodDescriptor         = exerciseServiceServiceDescriptor.Methods().ByName("UpdateExercise")
	exerciseServiceDeleteExerciseMethodDescriptor         = exerciseServiceServiceDescriptor.Methods().ByName("DeleteExercise")
	exerciseServiceListExercisesMethodDescriptor          = exerciseServiceServiceDescriptor.Methods().ByName("ListExercises")
	exerciseServiceGetPreviousWorkoutSetsMethodDescriptor = exerciseServiceServiceDescriptor.Methods().ByName("GetPreviousWorkoutSets")
	exerciseServiceGetPersonalBestsMethodDescriptor       = exerciseServiceServiceDescriptor.Methods().ByName("GetPersonalBests")
	exerciseServiceListSetsMethodDescriptor               = exerciseServiceServiceDescriptor.Methods().ByName("ListSets")
)

// ExerciseServiceClient is a client for the api.v1.ExerciseService service.
type ExerciseServiceClient interface {
	CreateExercise(context.Context, *connect.Request[v1.CreateExerciseRequest]) (*connect.Response[v1.CreateExerciseResponse], error)
	GetExercise(context.Context, *connect.Request[v1.GetExerciseRequest]) (*connect.Response[v1.GetExerciseResponse], error)
	UpdateExercise(context.Context, *connect.Request[v1.UpdateExerciseRequest]) (*connect.Response[v1.UpdateExerciseResponse], error)
	DeleteExercise(context.Context, *connect.Request[v1.DeleteExerciseRequest]) (*connect.Response[v1.DeleteExerciseResponse], error)
	ListExercises(context.Context, *connect.Request[v1.ListExercisesRequest]) (*connect.Response[v1.ListExercisesResponse], error)
	GetPreviousWorkoutSets(context.Context, *connect.Request[v1.GetPreviousWorkoutSetsRequest]) (*connect.Response[v1.GetPreviousWorkoutSetsResponse], error)
	GetPersonalBests(context.Context, *connect.Request[v1.GetPersonalBestsRequest]) (*connect.Response[v1.GetPersonalBestsResponse], error)
	ListSets(context.Context, *connect.Request[v1.ListSetsRequest]) (*connect.Response[v1.ListSetsResponse], error)
}

// NewExerciseServiceClient constructs a client for the api.v1.ExerciseService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewExerciseServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ExerciseServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &exerciseServiceClient{
		createExercise: connect.NewClient[v1.CreateExerciseRequest, v1.CreateExerciseResponse](
			httpClient,
			baseURL+ExerciseServiceCreateExerciseProcedure,
			connect.WithSchema(exerciseServiceCreateExerciseMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getExercise: connect.NewClient[v1.GetExerciseRequest, v1.GetExerciseResponse](
			httpClient,
			baseURL+ExerciseServiceGetExerciseProcedure,
			connect.WithSchema(exerciseServiceGetExerciseMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateExercise: connect.NewClient[v1.UpdateExerciseRequest, v1.UpdateExerciseResponse](
			httpClient,
			baseURL+ExerciseServiceUpdateExerciseProcedure,
			connect.WithSchema(exerciseServiceUpdateExerciseMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteExercise: connect.NewClient[v1.DeleteExerciseRequest, v1.DeleteExerciseResponse](
			httpClient,
			baseURL+ExerciseServiceDeleteExerciseProcedure,
			connect.WithSchema(exerciseServiceDeleteExerciseMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listExercises: connect.NewClient[v1.ListExercisesRequest, v1.ListExercisesResponse](
			httpClient,
			baseURL+ExerciseServiceListExercisesProcedure,
			connect.WithSchema(exerciseServiceListExercisesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getPreviousWorkoutSets: connect.NewClient[v1.GetPreviousWorkoutSetsRequest, v1.GetPreviousWorkoutSetsResponse](
			httpClient,
			baseURL+ExerciseServiceGetPreviousWorkoutSetsProcedure,
			connect.WithSchema(exerciseServiceGetPreviousWorkoutSetsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getPersonalBests: connect.NewClient[v1.GetPersonalBestsRequest, v1.GetPersonalBestsResponse](
			httpClient,
			baseURL+ExerciseServiceGetPersonalBestsProcedure,
			connect.WithSchema(exerciseServiceGetPersonalBestsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listSets: connect.NewClient[v1.ListSetsRequest, v1.ListSetsResponse](
			httpClient,
			baseURL+ExerciseServiceListSetsProcedure,
			connect.WithSchema(exerciseServiceListSetsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// exerciseServiceClient implements ExerciseServiceClient.
type exerciseServiceClient struct {
	createExercise         *connect.Client[v1.CreateExerciseRequest, v1.CreateExerciseResponse]
	getExercise            *connect.Client[v1.GetExerciseRequest, v1.GetExerciseResponse]
	updateExercise         *connect.Client[v1.UpdateExerciseRequest, v1.UpdateExerciseResponse]
	deleteExercise         *connect.Client[v1.DeleteExerciseRequest, v1.DeleteExerciseResponse]
	listExercises          *connect.Client[v1.ListExercisesRequest, v1.ListExercisesResponse]
	getPreviousWorkoutSets *connect.Client[v1.GetPreviousWorkoutSetsRequest, v1.GetPreviousWorkoutSetsResponse]
	getPersonalBests       *connect.Client[v1.GetPersonalBestsRequest, v1.GetPersonalBestsResponse]
	listSets               *connect.Client[v1.ListSetsRequest, v1.ListSetsResponse]
}

// CreateExercise calls api.v1.ExerciseService.CreateExercise.
func (c *exerciseServiceClient) CreateExercise(ctx context.Context, req *connect.Request[v1.CreateExerciseRequest]) (*connect.Response[v1.CreateExerciseResponse], error) {
	return c.createExercise.CallUnary(ctx, req)
}

// GetExercise calls api.v1.ExerciseService.GetExercise.
func (c *exerciseServiceClient) GetExercise(ctx context.Context, req *connect.Request[v1.GetExerciseRequest]) (*connect.Response[v1.GetExerciseResponse], error) {
	return c.getExercise.CallUnary(ctx, req)
}

// UpdateExercise calls api.v1.ExerciseService.UpdateExercise.
func (c *exerciseServiceClient) UpdateExercise(ctx context.Context, req *connect.Request[v1.UpdateExerciseRequest]) (*connect.Response[v1.UpdateExerciseResponse], error) {
	return c.updateExercise.CallUnary(ctx, req)
}

// DeleteExercise calls api.v1.ExerciseService.DeleteExercise.
func (c *exerciseServiceClient) DeleteExercise(ctx context.Context, req *connect.Request[v1.DeleteExerciseRequest]) (*connect.Response[v1.DeleteExerciseResponse], error) {
	return c.deleteExercise.CallUnary(ctx, req)
}

// ListExercises calls api.v1.ExerciseService.ListExercises.
func (c *exerciseServiceClient) ListExercises(ctx context.Context, req *connect.Request[v1.ListExercisesRequest]) (*connect.Response[v1.ListExercisesResponse], error) {
	return c.listExercises.CallUnary(ctx, req)
}

// GetPreviousWorkoutSets calls api.v1.ExerciseService.GetPreviousWorkoutSets.
func (c *exerciseServiceClient) GetPreviousWorkoutSets(ctx context.Context, req *connect.Request[v1.GetPreviousWorkoutSetsRequest]) (*connect.Response[v1.GetPreviousWorkoutSetsResponse], error) {
	return c.getPreviousWorkoutSets.CallUnary(ctx, req)
}

// GetPersonalBests calls api.v1.ExerciseService.GetPersonalBests.
func (c *exerciseServiceClient) GetPersonalBests(ctx context.Context, req *connect.Request[v1.GetPersonalBestsRequest]) (*connect.Response[v1.GetPersonalBestsResponse], error) {
	return c.getPersonalBests.CallUnary(ctx, req)
}

// ListSets calls api.v1.ExerciseService.ListSets.
func (c *exerciseServiceClient) ListSets(ctx context.Context, req *connect.Request[v1.ListSetsRequest]) (*connect.Response[v1.ListSetsResponse], error) {
	return c.listSets.CallUnary(ctx, req)
}

// ExerciseServiceHandler is an implementation of the api.v1.ExerciseService service.
type ExerciseServiceHandler interface {
	CreateExercise(context.Context, *connect.Request[v1.CreateExerciseRequest]) (*connect.Response[v1.CreateExerciseResponse], error)
	GetExercise(context.Context, *connect.Request[v1.GetExerciseRequest]) (*connect.Response[v1.GetExerciseResponse], error)
	UpdateExercise(context.Context, *connect.Request[v1.UpdateExerciseRequest]) (*connect.Response[v1.UpdateExerciseResponse], error)
	DeleteExercise(context.Context, *connect.Request[v1.DeleteExerciseRequest]) (*connect.Response[v1.DeleteExerciseResponse], error)
	ListExercises(context.Context, *connect.Request[v1.ListExercisesRequest]) (*connect.Response[v1.ListExercisesResponse], error)
	GetPreviousWorkoutSets(context.Context, *connect.Request[v1.GetPreviousWorkoutSetsRequest]) (*connect.Response[v1.GetPreviousWorkoutSetsResponse], error)
	GetPersonalBests(context.Context, *connect.Request[v1.GetPersonalBestsRequest]) (*connect.Response[v1.GetPersonalBestsResponse], error)
	ListSets(context.Context, *connect.Request[v1.ListSetsRequest]) (*connect.Response[v1.ListSetsResponse], error)
}

// NewExerciseServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewExerciseServiceHandler(svc ExerciseServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	exerciseServiceCreateExerciseHandler := connect.NewUnaryHandler(
		ExerciseServiceCreateExerciseProcedure,
		svc.CreateExercise,
		connect.WithSchema(exerciseServiceCreateExerciseMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	exerciseServiceGetExerciseHandler := connect.NewUnaryHandler(
		ExerciseServiceGetExerciseProcedure,
		svc.GetExercise,
		connect.WithSchema(exerciseServiceGetExerciseMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	exerciseServiceUpdateExerciseHandler := connect.NewUnaryHandler(
		ExerciseServiceUpdateExerciseProcedure,
		svc.UpdateExercise,
		connect.WithSchema(exerciseServiceUpdateExerciseMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	exerciseServiceDeleteExerciseHandler := connect.NewUnaryHandler(
		ExerciseServiceDeleteExerciseProcedure,
		svc.DeleteExercise,
		connect.WithSchema(exerciseServiceDeleteExerciseMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	exerciseServiceListExercisesHandler := connect.NewUnaryHandler(
		ExerciseServiceListExercisesProcedure,
		svc.ListExercises,
		connect.WithSchema(exerciseServiceListExercisesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	exerciseServiceGetPreviousWorkoutSetsHandler := connect.NewUnaryHandler(
		ExerciseServiceGetPreviousWorkoutSetsProcedure,
		svc.GetPreviousWorkoutSets,
		connect.WithSchema(exerciseServiceGetPreviousWorkoutSetsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	exerciseServiceGetPersonalBestsHandler := connect.NewUnaryHandler(
		ExerciseServiceGetPersonalBestsProcedure,
		svc.GetPersonalBests,
		connect.WithSchema(exerciseServiceGetPersonalBestsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	exerciseServiceListSetsHandler := connect.NewUnaryHandler(
		ExerciseServiceListSetsProcedure,
		svc.ListSets,
		connect.WithSchema(exerciseServiceListSetsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.ExerciseService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ExerciseServiceCreateExerciseProcedure:
			exerciseServiceCreateExerciseHandler.ServeHTTP(w, r)
		case ExerciseServiceGetExerciseProcedure:
			exerciseServiceGetExerciseHandler.ServeHTTP(w, r)
		case ExerciseServiceUpdateExerciseProcedure:
			exerciseServiceUpdateExerciseHandler.ServeHTTP(w, r)
		case ExerciseServiceDeleteExerciseProcedure:
			exerciseServiceDeleteExerciseHandler.ServeHTTP(w, r)
		case ExerciseServiceListExercisesProcedure:
			exerciseServiceListExercisesHandler.ServeHTTP(w, r)
		case ExerciseServiceGetPreviousWorkoutSetsProcedure:
			exerciseServiceGetPreviousWorkoutSetsHandler.ServeHTTP(w, r)
		case ExerciseServiceGetPersonalBestsProcedure:
			exerciseServiceGetPersonalBestsHandler.ServeHTTP(w, r)
		case ExerciseServiceListSetsProcedure:
			exerciseServiceListSetsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedExerciseServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedExerciseServiceHandler struct{}

func (UnimplementedExerciseServiceHandler) CreateExercise(context.Context, *connect.Request[v1.CreateExerciseRequest]) (*connect.Response[v1.CreateExerciseResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ExerciseService.CreateExercise is not implemented"))
}

func (UnimplementedExerciseServiceHandler) GetExercise(context.Context, *connect.Request[v1.GetExerciseRequest]) (*connect.Response[v1.GetExerciseResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ExerciseService.GetExercise is not implemented"))
}

func (UnimplementedExerciseServiceHandler) UpdateExercise(context.Context, *connect.Request[v1.UpdateExerciseRequest]) (*connect.Response[v1.UpdateExerciseResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ExerciseService.UpdateExercise is not implemented"))
}

func (UnimplementedExerciseServiceHandler) DeleteExercise(context.Context, *connect.Request[v1.DeleteExerciseRequest]) (*connect.Response[v1.DeleteExerciseResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ExerciseService.DeleteExercise is not implemented"))
}

func (UnimplementedExerciseServiceHandler) ListExercises(context.Context, *connect.Request[v1.ListExercisesRequest]) (*connect.Response[v1.ListExercisesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ExerciseService.ListExercises is not implemented"))
}

func (UnimplementedExerciseServiceHandler) GetPreviousWorkoutSets(context.Context, *connect.Request[v1.GetPreviousWorkoutSetsRequest]) (*connect.Response[v1.GetPreviousWorkoutSetsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ExerciseService.GetPreviousWorkoutSets is not implemented"))
}

func (UnimplementedExerciseServiceHandler) GetPersonalBests(context.Context, *connect.Request[v1.GetPersonalBestsRequest]) (*connect.Response[v1.GetPersonalBestsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ExerciseService.GetPersonalBests is not implemented"))
}

func (UnimplementedExerciseServiceHandler) ListSets(context.Context, *connect.Request[v1.ListSetsRequest]) (*connect.Response[v1.ListSetsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ExerciseService.ListSets is not implemented"))
}
