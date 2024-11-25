// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/workouts.proto

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
	// WorkoutServiceName is the fully-qualified name of the WorkoutService service.
	WorkoutServiceName = "api.v1.WorkoutService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// WorkoutServiceCreateProcedure is the fully-qualified name of the WorkoutService's Create RPC.
	WorkoutServiceCreateProcedure = "/api.v1.WorkoutService/Create"
	// WorkoutServiceGetProcedure is the fully-qualified name of the WorkoutService's Get RPC.
	WorkoutServiceGetProcedure = "/api.v1.WorkoutService/Get"
	// WorkoutServiceListProcedure is the fully-qualified name of the WorkoutService's List RPC.
	WorkoutServiceListProcedure = "/api.v1.WorkoutService/List"
	// WorkoutServiceDeleteProcedure is the fully-qualified name of the WorkoutService's Delete RPC.
	WorkoutServiceDeleteProcedure = "/api.v1.WorkoutService/Delete"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	workoutServiceServiceDescriptor      = v1.File_api_v1_workouts_proto.Services().ByName("WorkoutService")
	workoutServiceCreateMethodDescriptor = workoutServiceServiceDescriptor.Methods().ByName("Create")
	workoutServiceGetMethodDescriptor    = workoutServiceServiceDescriptor.Methods().ByName("Get")
	workoutServiceListMethodDescriptor   = workoutServiceServiceDescriptor.Methods().ByName("List")
	workoutServiceDeleteMethodDescriptor = workoutServiceServiceDescriptor.Methods().ByName("Delete")
)

// WorkoutServiceClient is a client for the api.v1.WorkoutService service.
type WorkoutServiceClient interface {
	Create(context.Context, *connect.Request[v1.CreateWorkoutRequest]) (*connect.Response[v1.CreateWorkoutResponse], error)
	Get(context.Context, *connect.Request[v1.GetWorkoutRequest]) (*connect.Response[v1.GetWorkoutResponse], error)
	List(context.Context, *connect.Request[v1.ListWorkoutsRequest]) (*connect.Response[v1.ListWorkoutsResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteWorkoutRequest]) (*connect.Response[v1.DeleteWorkoutResponse], error)
}

// NewWorkoutServiceClient constructs a client for the api.v1.WorkoutService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWorkoutServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) WorkoutServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &workoutServiceClient{
		create: connect.NewClient[v1.CreateWorkoutRequest, v1.CreateWorkoutResponse](
			httpClient,
			baseURL+WorkoutServiceCreateProcedure,
			connect.WithSchema(workoutServiceCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		get: connect.NewClient[v1.GetWorkoutRequest, v1.GetWorkoutResponse](
			httpClient,
			baseURL+WorkoutServiceGetProcedure,
			connect.WithSchema(workoutServiceGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		list: connect.NewClient[v1.ListWorkoutsRequest, v1.ListWorkoutsResponse](
			httpClient,
			baseURL+WorkoutServiceListProcedure,
			connect.WithSchema(workoutServiceListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		delete: connect.NewClient[v1.DeleteWorkoutRequest, v1.DeleteWorkoutResponse](
			httpClient,
			baseURL+WorkoutServiceDeleteProcedure,
			connect.WithSchema(workoutServiceDeleteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// workoutServiceClient implements WorkoutServiceClient.
type workoutServiceClient struct {
	create *connect.Client[v1.CreateWorkoutRequest, v1.CreateWorkoutResponse]
	get    *connect.Client[v1.GetWorkoutRequest, v1.GetWorkoutResponse]
	list   *connect.Client[v1.ListWorkoutsRequest, v1.ListWorkoutsResponse]
	delete *connect.Client[v1.DeleteWorkoutRequest, v1.DeleteWorkoutResponse]
}

// Create calls api.v1.WorkoutService.Create.
func (c *workoutServiceClient) Create(ctx context.Context, req *connect.Request[v1.CreateWorkoutRequest]) (*connect.Response[v1.CreateWorkoutResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Get calls api.v1.WorkoutService.Get.
func (c *workoutServiceClient) Get(ctx context.Context, req *connect.Request[v1.GetWorkoutRequest]) (*connect.Response[v1.GetWorkoutResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// List calls api.v1.WorkoutService.List.
func (c *workoutServiceClient) List(ctx context.Context, req *connect.Request[v1.ListWorkoutsRequest]) (*connect.Response[v1.ListWorkoutsResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Delete calls api.v1.WorkoutService.Delete.
func (c *workoutServiceClient) Delete(ctx context.Context, req *connect.Request[v1.DeleteWorkoutRequest]) (*connect.Response[v1.DeleteWorkoutResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// WorkoutServiceHandler is an implementation of the api.v1.WorkoutService service.
type WorkoutServiceHandler interface {
	Create(context.Context, *connect.Request[v1.CreateWorkoutRequest]) (*connect.Response[v1.CreateWorkoutResponse], error)
	Get(context.Context, *connect.Request[v1.GetWorkoutRequest]) (*connect.Response[v1.GetWorkoutResponse], error)
	List(context.Context, *connect.Request[v1.ListWorkoutsRequest]) (*connect.Response[v1.ListWorkoutsResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteWorkoutRequest]) (*connect.Response[v1.DeleteWorkoutResponse], error)
}

// NewWorkoutServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWorkoutServiceHandler(svc WorkoutServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	workoutServiceCreateHandler := connect.NewUnaryHandler(
		WorkoutServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(workoutServiceCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workoutServiceGetHandler := connect.NewUnaryHandler(
		WorkoutServiceGetProcedure,
		svc.Get,
		connect.WithSchema(workoutServiceGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workoutServiceListHandler := connect.NewUnaryHandler(
		WorkoutServiceListProcedure,
		svc.List,
		connect.WithSchema(workoutServiceListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	workoutServiceDeleteHandler := connect.NewUnaryHandler(
		WorkoutServiceDeleteProcedure,
		svc.Delete,
		connect.WithSchema(workoutServiceDeleteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.WorkoutService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case WorkoutServiceCreateProcedure:
			workoutServiceCreateHandler.ServeHTTP(w, r)
		case WorkoutServiceGetProcedure:
			workoutServiceGetHandler.ServeHTTP(w, r)
		case WorkoutServiceListProcedure:
			workoutServiceListHandler.ServeHTTP(w, r)
		case WorkoutServiceDeleteProcedure:
			workoutServiceDeleteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedWorkoutServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWorkoutServiceHandler struct{}

func (UnimplementedWorkoutServiceHandler) Create(context.Context, *connect.Request[v1.CreateWorkoutRequest]) (*connect.Response[v1.CreateWorkoutResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkoutService.Create is not implemented"))
}

func (UnimplementedWorkoutServiceHandler) Get(context.Context, *connect.Request[v1.GetWorkoutRequest]) (*connect.Response[v1.GetWorkoutResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkoutService.Get is not implemented"))
}

func (UnimplementedWorkoutServiceHandler) List(context.Context, *connect.Request[v1.ListWorkoutsRequest]) (*connect.Response[v1.ListWorkoutsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkoutService.List is not implemented"))
}

func (UnimplementedWorkoutServiceHandler) Delete(context.Context, *connect.Request[v1.DeleteWorkoutRequest]) (*connect.Response[v1.DeleteWorkoutResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkoutService.Delete is not implemented"))
}
