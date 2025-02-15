// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/auth_service.proto

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
	// AuthServiceName is the fully-qualified name of the AuthService service.
	AuthServiceName = "api.v1.AuthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthServiceSignupProcedure is the fully-qualified name of the AuthService's Signup RPC.
	AuthServiceSignupProcedure = "/api.v1.AuthService/Signup"
	// AuthServiceLoginProcedure is the fully-qualified name of the AuthService's Login RPC.
	AuthServiceLoginProcedure = "/api.v1.AuthService/Login"
	// AuthServiceRefreshTokenProcedure is the fully-qualified name of the AuthService's RefreshToken
	// RPC.
	AuthServiceRefreshTokenProcedure = "/api.v1.AuthService/RefreshToken"
	// AuthServiceLogoutProcedure is the fully-qualified name of the AuthService's Logout RPC.
	AuthServiceLogoutProcedure = "/api.v1.AuthService/Logout"
	// AuthServiceVerifyEmailProcedure is the fully-qualified name of the AuthService's VerifyEmail RPC.
	AuthServiceVerifyEmailProcedure = "/api.v1.AuthService/VerifyEmail"
	// AuthServiceResetPasswordProcedure is the fully-qualified name of the AuthService's ResetPassword
	// RPC.
	AuthServiceResetPasswordProcedure = "/api.v1.AuthService/ResetPassword"
	// AuthServiceUpdatePasswordProcedure is the fully-qualified name of the AuthService's
	// UpdatePassword RPC.
	AuthServiceUpdatePasswordProcedure = "/api.v1.AuthService/UpdatePassword"
)

// AuthServiceClient is a client for the api.v1.AuthService service.
type AuthServiceClient interface {
	Signup(context.Context, *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error)
	Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
	RefreshToken(context.Context, *connect.Request[v1.RefreshTokenRequest]) (*connect.Response[v1.RefreshTokenResponse], error)
	Logout(context.Context, *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error)
	VerifyEmail(context.Context, *connect.Request[v1.VerifyEmailRequest]) (*connect.Response[v1.VerifyEmailResponse], error)
	ResetPassword(context.Context, *connect.Request[v1.ResetPasswordRequest]) (*connect.Response[v1.ResetPasswordResponse], error)
	UpdatePassword(context.Context, *connect.Request[v1.UpdatePasswordRequest]) (*connect.Response[v1.UpdatePasswordResponse], error)
}

// NewAuthServiceClient constructs a client for the api.v1.AuthService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	authServiceMethods := v1.File_api_v1_auth_service_proto.Services().ByName("AuthService").Methods()
	return &authServiceClient{
		signup: connect.NewClient[v1.SignupRequest, v1.SignupResponse](
			httpClient,
			baseURL+AuthServiceSignupProcedure,
			connect.WithSchema(authServiceMethods.ByName("Signup")),
			connect.WithClientOptions(opts...),
		),
		login: connect.NewClient[v1.LoginRequest, v1.LoginResponse](
			httpClient,
			baseURL+AuthServiceLoginProcedure,
			connect.WithSchema(authServiceMethods.ByName("Login")),
			connect.WithClientOptions(opts...),
		),
		refreshToken: connect.NewClient[v1.RefreshTokenRequest, v1.RefreshTokenResponse](
			httpClient,
			baseURL+AuthServiceRefreshTokenProcedure,
			connect.WithSchema(authServiceMethods.ByName("RefreshToken")),
			connect.WithClientOptions(opts...),
		),
		logout: connect.NewClient[v1.LogoutRequest, v1.LogoutResponse](
			httpClient,
			baseURL+AuthServiceLogoutProcedure,
			connect.WithSchema(authServiceMethods.ByName("Logout")),
			connect.WithClientOptions(opts...),
		),
		verifyEmail: connect.NewClient[v1.VerifyEmailRequest, v1.VerifyEmailResponse](
			httpClient,
			baseURL+AuthServiceVerifyEmailProcedure,
			connect.WithSchema(authServiceMethods.ByName("VerifyEmail")),
			connect.WithClientOptions(opts...),
		),
		resetPassword: connect.NewClient[v1.ResetPasswordRequest, v1.ResetPasswordResponse](
			httpClient,
			baseURL+AuthServiceResetPasswordProcedure,
			connect.WithSchema(authServiceMethods.ByName("ResetPassword")),
			connect.WithClientOptions(opts...),
		),
		updatePassword: connect.NewClient[v1.UpdatePasswordRequest, v1.UpdatePasswordResponse](
			httpClient,
			baseURL+AuthServiceUpdatePasswordProcedure,
			connect.WithSchema(authServiceMethods.ByName("UpdatePassword")),
			connect.WithClientOptions(opts...),
		),
	}
}

// authServiceClient implements AuthServiceClient.
type authServiceClient struct {
	signup         *connect.Client[v1.SignupRequest, v1.SignupResponse]
	login          *connect.Client[v1.LoginRequest, v1.LoginResponse]
	refreshToken   *connect.Client[v1.RefreshTokenRequest, v1.RefreshTokenResponse]
	logout         *connect.Client[v1.LogoutRequest, v1.LogoutResponse]
	verifyEmail    *connect.Client[v1.VerifyEmailRequest, v1.VerifyEmailResponse]
	resetPassword  *connect.Client[v1.ResetPasswordRequest, v1.ResetPasswordResponse]
	updatePassword *connect.Client[v1.UpdatePasswordRequest, v1.UpdatePasswordResponse]
}

// Signup calls api.v1.AuthService.Signup.
func (c *authServiceClient) Signup(ctx context.Context, req *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error) {
	return c.signup.CallUnary(ctx, req)
}

// Login calls api.v1.AuthService.Login.
func (c *authServiceClient) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// RefreshToken calls api.v1.AuthService.RefreshToken.
func (c *authServiceClient) RefreshToken(ctx context.Context, req *connect.Request[v1.RefreshTokenRequest]) (*connect.Response[v1.RefreshTokenResponse], error) {
	return c.refreshToken.CallUnary(ctx, req)
}

// Logout calls api.v1.AuthService.Logout.
func (c *authServiceClient) Logout(ctx context.Context, req *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error) {
	return c.logout.CallUnary(ctx, req)
}

// VerifyEmail calls api.v1.AuthService.VerifyEmail.
func (c *authServiceClient) VerifyEmail(ctx context.Context, req *connect.Request[v1.VerifyEmailRequest]) (*connect.Response[v1.VerifyEmailResponse], error) {
	return c.verifyEmail.CallUnary(ctx, req)
}

// ResetPassword calls api.v1.AuthService.ResetPassword.
func (c *authServiceClient) ResetPassword(ctx context.Context, req *connect.Request[v1.ResetPasswordRequest]) (*connect.Response[v1.ResetPasswordResponse], error) {
	return c.resetPassword.CallUnary(ctx, req)
}

// UpdatePassword calls api.v1.AuthService.UpdatePassword.
func (c *authServiceClient) UpdatePassword(ctx context.Context, req *connect.Request[v1.UpdatePasswordRequest]) (*connect.Response[v1.UpdatePasswordResponse], error) {
	return c.updatePassword.CallUnary(ctx, req)
}

// AuthServiceHandler is an implementation of the api.v1.AuthService service.
type AuthServiceHandler interface {
	Signup(context.Context, *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error)
	Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
	RefreshToken(context.Context, *connect.Request[v1.RefreshTokenRequest]) (*connect.Response[v1.RefreshTokenResponse], error)
	Logout(context.Context, *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error)
	VerifyEmail(context.Context, *connect.Request[v1.VerifyEmailRequest]) (*connect.Response[v1.VerifyEmailResponse], error)
	ResetPassword(context.Context, *connect.Request[v1.ResetPasswordRequest]) (*connect.Response[v1.ResetPasswordResponse], error)
	UpdatePassword(context.Context, *connect.Request[v1.UpdatePasswordRequest]) (*connect.Response[v1.UpdatePasswordResponse], error)
}

// NewAuthServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthServiceHandler(svc AuthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	authServiceMethods := v1.File_api_v1_auth_service_proto.Services().ByName("AuthService").Methods()
	authServiceSignupHandler := connect.NewUnaryHandler(
		AuthServiceSignupProcedure,
		svc.Signup,
		connect.WithSchema(authServiceMethods.ByName("Signup")),
		connect.WithHandlerOptions(opts...),
	)
	authServiceLoginHandler := connect.NewUnaryHandler(
		AuthServiceLoginProcedure,
		svc.Login,
		connect.WithSchema(authServiceMethods.ByName("Login")),
		connect.WithHandlerOptions(opts...),
	)
	authServiceRefreshTokenHandler := connect.NewUnaryHandler(
		AuthServiceRefreshTokenProcedure,
		svc.RefreshToken,
		connect.WithSchema(authServiceMethods.ByName("RefreshToken")),
		connect.WithHandlerOptions(opts...),
	)
	authServiceLogoutHandler := connect.NewUnaryHandler(
		AuthServiceLogoutProcedure,
		svc.Logout,
		connect.WithSchema(authServiceMethods.ByName("Logout")),
		connect.WithHandlerOptions(opts...),
	)
	authServiceVerifyEmailHandler := connect.NewUnaryHandler(
		AuthServiceVerifyEmailProcedure,
		svc.VerifyEmail,
		connect.WithSchema(authServiceMethods.ByName("VerifyEmail")),
		connect.WithHandlerOptions(opts...),
	)
	authServiceResetPasswordHandler := connect.NewUnaryHandler(
		AuthServiceResetPasswordProcedure,
		svc.ResetPassword,
		connect.WithSchema(authServiceMethods.ByName("ResetPassword")),
		connect.WithHandlerOptions(opts...),
	)
	authServiceUpdatePasswordHandler := connect.NewUnaryHandler(
		AuthServiceUpdatePasswordProcedure,
		svc.UpdatePassword,
		connect.WithSchema(authServiceMethods.ByName("UpdatePassword")),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.AuthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthServiceSignupProcedure:
			authServiceSignupHandler.ServeHTTP(w, r)
		case AuthServiceLoginProcedure:
			authServiceLoginHandler.ServeHTTP(w, r)
		case AuthServiceRefreshTokenProcedure:
			authServiceRefreshTokenHandler.ServeHTTP(w, r)
		case AuthServiceLogoutProcedure:
			authServiceLogoutHandler.ServeHTTP(w, r)
		case AuthServiceVerifyEmailProcedure:
			authServiceVerifyEmailHandler.ServeHTTP(w, r)
		case AuthServiceResetPasswordProcedure:
			authServiceResetPasswordHandler.ServeHTTP(w, r)
		case AuthServiceUpdatePasswordProcedure:
			authServiceUpdatePasswordHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthServiceHandler struct{}

func (UnimplementedAuthServiceHandler) Signup(context.Context, *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.AuthService.Signup is not implemented"))
}

func (UnimplementedAuthServiceHandler) Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.AuthService.Login is not implemented"))
}

func (UnimplementedAuthServiceHandler) RefreshToken(context.Context, *connect.Request[v1.RefreshTokenRequest]) (*connect.Response[v1.RefreshTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.AuthService.RefreshToken is not implemented"))
}

func (UnimplementedAuthServiceHandler) Logout(context.Context, *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.AuthService.Logout is not implemented"))
}

func (UnimplementedAuthServiceHandler) VerifyEmail(context.Context, *connect.Request[v1.VerifyEmailRequest]) (*connect.Response[v1.VerifyEmailResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.AuthService.VerifyEmail is not implemented"))
}

func (UnimplementedAuthServiceHandler) ResetPassword(context.Context, *connect.Request[v1.ResetPasswordRequest]) (*connect.Response[v1.ResetPasswordResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.AuthService.ResetPassword is not implemented"))
}

func (UnimplementedAuthServiceHandler) UpdatePassword(context.Context, *connect.Request[v1.UpdatePasswordRequest]) (*connect.Response[v1.UpdatePasswordResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.AuthService.UpdatePassword is not implemented"))
}
