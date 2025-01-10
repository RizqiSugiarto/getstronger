// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/workout_service.proto

package apiv1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"
	v1 "github.com/crlssn/getstronger/server/gen/proto/api/v1"
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
	// WorkoutServiceCreateWorkoutProcedure is the fully-qualified name of the WorkoutService's
	// CreateWorkout RPC.
	WorkoutServiceCreateWorkoutProcedure = "/api.v1.WorkoutService/CreateWorkout"
	// WorkoutServiceGetWorkoutProcedure is the fully-qualified name of the WorkoutService's GetWorkout
	// RPC.
	WorkoutServiceGetWorkoutProcedure = "/api.v1.WorkoutService/GetWorkout"
	// WorkoutServiceListWorkoutsProcedure is the fully-qualified name of the WorkoutService's
	// ListWorkouts RPC.
	WorkoutServiceListWorkoutsProcedure = "/api.v1.WorkoutService/ListWorkouts"
	// WorkoutServiceDeleteWorkoutProcedure is the fully-qualified name of the WorkoutService's
	// DeleteWorkout RPC.
	WorkoutServiceDeleteWorkoutProcedure = "/api.v1.WorkoutService/DeleteWorkout"
	// WorkoutServicePostCommentProcedure is the fully-qualified name of the WorkoutService's
	// PostComment RPC.
	WorkoutServicePostCommentProcedure = "/api.v1.WorkoutService/PostComment"
	// WorkoutServiceUpdateWorkoutProcedure is the fully-qualified name of the WorkoutService's
	// UpdateWorkout RPC.
	WorkoutServiceUpdateWorkoutProcedure = "/api.v1.WorkoutService/UpdateWorkout"
)

// WorkoutServiceClient is a client for the api.v1.WorkoutService service.
type WorkoutServiceClient interface {
	CreateWorkout(context.Context, *connect.Request[v1.CreateWorkoutRequest]) (*connect.Response[v1.CreateWorkoutResponse], error)
	GetWorkout(context.Context, *connect.Request[v1.GetWorkoutRequest]) (*connect.Response[v1.GetWorkoutResponse], error)
	ListWorkouts(context.Context, *connect.Request[v1.ListWorkoutsRequest]) (*connect.Response[v1.ListWorkoutsResponse], error)
	DeleteWorkout(context.Context, *connect.Request[v1.DeleteWorkoutRequest]) (*connect.Response[v1.DeleteWorkoutResponse], error)
	PostComment(context.Context, *connect.Request[v1.PostCommentRequest]) (*connect.Response[v1.PostCommentResponse], error)
	UpdateWorkout(context.Context, *connect.Request[v1.UpdateWorkoutRequest]) (*connect.Response[v1.UpdateWorkoutResponse], error)
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
	workoutServiceMethods := v1.File_api_v1_workout_service_proto.Services().ByName("WorkoutService").Methods()
	return &workoutServiceClient{
		createWorkout: connect.NewClient[v1.CreateWorkoutRequest, v1.CreateWorkoutResponse](
			httpClient,
			baseURL+WorkoutServiceCreateWorkoutProcedure,
			connect.WithSchema(workoutServiceMethods.ByName("CreateWorkout")),
			connect.WithClientOptions(opts...),
		),
		getWorkout: connect.NewClient[v1.GetWorkoutRequest, v1.GetWorkoutResponse](
			httpClient,
			baseURL+WorkoutServiceGetWorkoutProcedure,
			connect.WithSchema(workoutServiceMethods.ByName("GetWorkout")),
			connect.WithClientOptions(opts...),
		),
		listWorkouts: connect.NewClient[v1.ListWorkoutsRequest, v1.ListWorkoutsResponse](
			httpClient,
			baseURL+WorkoutServiceListWorkoutsProcedure,
			connect.WithSchema(workoutServiceMethods.ByName("ListWorkouts")),
			connect.WithClientOptions(opts...),
		),
		deleteWorkout: connect.NewClient[v1.DeleteWorkoutRequest, v1.DeleteWorkoutResponse](
			httpClient,
			baseURL+WorkoutServiceDeleteWorkoutProcedure,
			connect.WithSchema(workoutServiceMethods.ByName("DeleteWorkout")),
			connect.WithClientOptions(opts...),
		),
		postComment: connect.NewClient[v1.PostCommentRequest, v1.PostCommentResponse](
			httpClient,
			baseURL+WorkoutServicePostCommentProcedure,
			connect.WithSchema(workoutServiceMethods.ByName("PostComment")),
			connect.WithClientOptions(opts...),
		),
		updateWorkout: connect.NewClient[v1.UpdateWorkoutRequest, v1.UpdateWorkoutResponse](
			httpClient,
			baseURL+WorkoutServiceUpdateWorkoutProcedure,
			connect.WithSchema(workoutServiceMethods.ByName("UpdateWorkout")),
			connect.WithClientOptions(opts...),
		),
	}
}

// workoutServiceClient implements WorkoutServiceClient.
type workoutServiceClient struct {
	createWorkout *connect.Client[v1.CreateWorkoutRequest, v1.CreateWorkoutResponse]
	getWorkout    *connect.Client[v1.GetWorkoutRequest, v1.GetWorkoutResponse]
	listWorkouts  *connect.Client[v1.ListWorkoutsRequest, v1.ListWorkoutsResponse]
	deleteWorkout *connect.Client[v1.DeleteWorkoutRequest, v1.DeleteWorkoutResponse]
	postComment   *connect.Client[v1.PostCommentRequest, v1.PostCommentResponse]
	updateWorkout *connect.Client[v1.UpdateWorkoutRequest, v1.UpdateWorkoutResponse]
}

// CreateWorkout calls api.v1.WorkoutService.CreateWorkout.
func (c *workoutServiceClient) CreateWorkout(ctx context.Context, req *connect.Request[v1.CreateWorkoutRequest]) (*connect.Response[v1.CreateWorkoutResponse], error) {
	return c.createWorkout.CallUnary(ctx, req)
}

// GetWorkout calls api.v1.WorkoutService.GetWorkout.
func (c *workoutServiceClient) GetWorkout(ctx context.Context, req *connect.Request[v1.GetWorkoutRequest]) (*connect.Response[v1.GetWorkoutResponse], error) {
	return c.getWorkout.CallUnary(ctx, req)
}

// ListWorkouts calls api.v1.WorkoutService.ListWorkouts.
func (c *workoutServiceClient) ListWorkouts(ctx context.Context, req *connect.Request[v1.ListWorkoutsRequest]) (*connect.Response[v1.ListWorkoutsResponse], error) {
	return c.listWorkouts.CallUnary(ctx, req)
}

// DeleteWorkout calls api.v1.WorkoutService.DeleteWorkout.
func (c *workoutServiceClient) DeleteWorkout(ctx context.Context, req *connect.Request[v1.DeleteWorkoutRequest]) (*connect.Response[v1.DeleteWorkoutResponse], error) {
	return c.deleteWorkout.CallUnary(ctx, req)
}

// PostComment calls api.v1.WorkoutService.PostComment.
func (c *workoutServiceClient) PostComment(ctx context.Context, req *connect.Request[v1.PostCommentRequest]) (*connect.Response[v1.PostCommentResponse], error) {
	return c.postComment.CallUnary(ctx, req)
}

// UpdateWorkout calls api.v1.WorkoutService.UpdateWorkout.
func (c *workoutServiceClient) UpdateWorkout(ctx context.Context, req *connect.Request[v1.UpdateWorkoutRequest]) (*connect.Response[v1.UpdateWorkoutResponse], error) {
	return c.updateWorkout.CallUnary(ctx, req)
}

// WorkoutServiceHandler is an implementation of the api.v1.WorkoutService service.
type WorkoutServiceHandler interface {
	CreateWorkout(context.Context, *connect.Request[v1.CreateWorkoutRequest]) (*connect.Response[v1.CreateWorkoutResponse], error)
	GetWorkout(context.Context, *connect.Request[v1.GetWorkoutRequest]) (*connect.Response[v1.GetWorkoutResponse], error)
	ListWorkouts(context.Context, *connect.Request[v1.ListWorkoutsRequest]) (*connect.Response[v1.ListWorkoutsResponse], error)
	DeleteWorkout(context.Context, *connect.Request[v1.DeleteWorkoutRequest]) (*connect.Response[v1.DeleteWorkoutResponse], error)
	PostComment(context.Context, *connect.Request[v1.PostCommentRequest]) (*connect.Response[v1.PostCommentResponse], error)
	UpdateWorkout(context.Context, *connect.Request[v1.UpdateWorkoutRequest]) (*connect.Response[v1.UpdateWorkoutResponse], error)
}

// NewWorkoutServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWorkoutServiceHandler(svc WorkoutServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	workoutServiceMethods := v1.File_api_v1_workout_service_proto.Services().ByName("WorkoutService").Methods()
	workoutServiceCreateWorkoutHandler := connect.NewUnaryHandler(
		WorkoutServiceCreateWorkoutProcedure,
		svc.CreateWorkout,
		connect.WithSchema(workoutServiceMethods.ByName("CreateWorkout")),
		connect.WithHandlerOptions(opts...),
	)
	workoutServiceGetWorkoutHandler := connect.NewUnaryHandler(
		WorkoutServiceGetWorkoutProcedure,
		svc.GetWorkout,
		connect.WithSchema(workoutServiceMethods.ByName("GetWorkout")),
		connect.WithHandlerOptions(opts...),
	)
	workoutServiceListWorkoutsHandler := connect.NewUnaryHandler(
		WorkoutServiceListWorkoutsProcedure,
		svc.ListWorkouts,
		connect.WithSchema(workoutServiceMethods.ByName("ListWorkouts")),
		connect.WithHandlerOptions(opts...),
	)
	workoutServiceDeleteWorkoutHandler := connect.NewUnaryHandler(
		WorkoutServiceDeleteWorkoutProcedure,
		svc.DeleteWorkout,
		connect.WithSchema(workoutServiceMethods.ByName("DeleteWorkout")),
		connect.WithHandlerOptions(opts...),
	)
	workoutServicePostCommentHandler := connect.NewUnaryHandler(
		WorkoutServicePostCommentProcedure,
		svc.PostComment,
		connect.WithSchema(workoutServiceMethods.ByName("PostComment")),
		connect.WithHandlerOptions(opts...),
	)
	workoutServiceUpdateWorkoutHandler := connect.NewUnaryHandler(
		WorkoutServiceUpdateWorkoutProcedure,
		svc.UpdateWorkout,
		connect.WithSchema(workoutServiceMethods.ByName("UpdateWorkout")),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.WorkoutService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case WorkoutServiceCreateWorkoutProcedure:
			workoutServiceCreateWorkoutHandler.ServeHTTP(w, r)
		case WorkoutServiceGetWorkoutProcedure:
			workoutServiceGetWorkoutHandler.ServeHTTP(w, r)
		case WorkoutServiceListWorkoutsProcedure:
			workoutServiceListWorkoutsHandler.ServeHTTP(w, r)
		case WorkoutServiceDeleteWorkoutProcedure:
			workoutServiceDeleteWorkoutHandler.ServeHTTP(w, r)
		case WorkoutServicePostCommentProcedure:
			workoutServicePostCommentHandler.ServeHTTP(w, r)
		case WorkoutServiceUpdateWorkoutProcedure:
			workoutServiceUpdateWorkoutHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedWorkoutServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWorkoutServiceHandler struct{}

func (UnimplementedWorkoutServiceHandler) CreateWorkout(context.Context, *connect.Request[v1.CreateWorkoutRequest]) (*connect.Response[v1.CreateWorkoutResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkoutService.CreateWorkout is not implemented"))
}

func (UnimplementedWorkoutServiceHandler) GetWorkout(context.Context, *connect.Request[v1.GetWorkoutRequest]) (*connect.Response[v1.GetWorkoutResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkoutService.GetWorkout is not implemented"))
}

func (UnimplementedWorkoutServiceHandler) ListWorkouts(context.Context, *connect.Request[v1.ListWorkoutsRequest]) (*connect.Response[v1.ListWorkoutsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkoutService.ListWorkouts is not implemented"))
}

func (UnimplementedWorkoutServiceHandler) DeleteWorkout(context.Context, *connect.Request[v1.DeleteWorkoutRequest]) (*connect.Response[v1.DeleteWorkoutResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkoutService.DeleteWorkout is not implemented"))
}

func (UnimplementedWorkoutServiceHandler) PostComment(context.Context, *connect.Request[v1.PostCommentRequest]) (*connect.Response[v1.PostCommentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkoutService.PostComment is not implemented"))
}

func (UnimplementedWorkoutServiceHandler) UpdateWorkout(context.Context, *connect.Request[v1.UpdateWorkoutRequest]) (*connect.Response[v1.UpdateWorkoutResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkoutService.UpdateWorkout is not implemented"))
}
