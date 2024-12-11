// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file api/v1/auth_service.proto (package api.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_buf_validate_validate } from "../../buf/validate/validate_pb";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file api/v1/auth_service.proto.
 */
export const file_api_v1_auth_service: GenFile = /*@__PURE__*/
  fileDesc("ChlhcGkvdjEvYXV0aF9zZXJ2aWNlLnByb3RvEgZhcGkudjEimgEKDVNpZ251cFJlcXVlc3QSFgoFZW1haWwYASABKAlCB7pIBHICYAESGQoIcGFzc3dvcmQYAiABKAlCB7pIBHICEAYSHQoVcGFzc3dvcmRfY29uZmlybWF0aW9uGAMgASgJEhsKCmZpcnN0X25hbWUYBCABKAlCB7pIBHICEAESGgoJbGFzdF9uYW1lGAUgASgJQge6SARyAhABIhAKDlNpZ251cFJlc3BvbnNlIkEKDExvZ2luUmVxdWVzdBIWCgVlbWFpbBgBIAEoCUIHukgEcgJgARIZCghwYXNzd29yZBgCIAEoCUIHukgEcgIQASIlCg1Mb2dpblJlc3BvbnNlEhQKDGFjY2Vzc190b2tlbhgBIAEoCSIVChNSZWZyZXNoVG9rZW5SZXF1ZXN0IiwKFFJlZnJlc2hUb2tlblJlc3BvbnNlEhQKDGFjY2Vzc190b2tlbhgBIAEoCSIPCg1Mb2dvdXRSZXF1ZXN0IhAKDkxvZ291dFJlc3BvbnNlIiwKElZlcmlmeUVtYWlsUmVxdWVzdBIWCgV0b2tlbhgBIAEoCUIHukgEcgIQASIVChNWZXJpZnlFbWFpbFJlc3BvbnNlIi4KFFJlc2V0UGFzc3dvcmRSZXF1ZXN0EhYKBWVtYWlsGAEgASgJQge6SARyAmABIhcKFVJlc2V0UGFzc3dvcmRSZXNwb25zZSJqChVVcGRhdGVQYXNzd29yZFJlcXVlc3QSFwoFdG9rZW4YASABKAlCCLpIBXIDsAEBEhkKCHBhc3N3b3JkGAIgASgJQge6SARyAhAGEh0KFXBhc3N3b3JkX2NvbmZpcm1hdGlvbhgDIAEoCSIYChZVcGRhdGVQYXNzd29yZFJlc3BvbnNlMvUDCgtBdXRoU2VydmljZRI5CgZTaWdudXASFS5hcGkudjEuU2lnbnVwUmVxdWVzdBoWLmFwaS52MS5TaWdudXBSZXNwb25zZSIAEjYKBUxvZ2luEhQuYXBpLnYxLkxvZ2luUmVxdWVzdBoVLmFwaS52MS5Mb2dpblJlc3BvbnNlIgASSwoMUmVmcmVzaFRva2VuEhsuYXBpLnYxLlJlZnJlc2hUb2tlblJlcXVlc3QaHC5hcGkudjEuUmVmcmVzaFRva2VuUmVzcG9uc2UiABI5CgZMb2dvdXQSFS5hcGkudjEuTG9nb3V0UmVxdWVzdBoWLmFwaS52MS5Mb2dvdXRSZXNwb25zZSIAEkgKC1ZlcmlmeUVtYWlsEhouYXBpLnYxLlZlcmlmeUVtYWlsUmVxdWVzdBobLmFwaS52MS5WZXJpZnlFbWFpbFJlc3BvbnNlIgASTgoNUmVzZXRQYXNzd29yZBIcLmFwaS52MS5SZXNldFBhc3N3b3JkUmVxdWVzdBodLmFwaS52MS5SZXNldFBhc3N3b3JkUmVzcG9uc2UiABJRCg5VcGRhdGVQYXNzd29yZBIdLmFwaS52MS5VcGRhdGVQYXNzd29yZFJlcXVlc3QaHi5hcGkudjEuVXBkYXRlUGFzc3dvcmRSZXNwb25zZSIAQpQBCgpjb20uYXBpLnYxQhBBdXRoU2VydmljZVByb3RvUAFaO2dpdGh1Yi5jb20vY3Jsc3NuL2dldHN0cm9uZ2VyL3NlcnZlci9wa2cvcHJvdG8vYXBpL3YxO2FwaXYxogIDQVhYqgIGQXBpLlYxygIGQXBpXFYx4gISQXBpXFYxXEdQQk1ldGFkYXRh6gIHQXBpOjpWMWIGcHJvdG8z", [file_buf_validate_validate, file_google_protobuf_timestamp]);

/**
 * @generated from message api.v1.SignupRequest
 */
export type SignupRequest = Message<"api.v1.SignupRequest"> & {
  /**
   * @generated from field: string email = 1;
   */
  email: string;

  /**
   * @generated from field: string password = 2;
   */
  password: string;

  /**
   * @generated from field: string password_confirmation = 3;
   */
  passwordConfirmation: string;

  /**
   * @generated from field: string first_name = 4;
   */
  firstName: string;

  /**
   * @generated from field: string last_name = 5;
   */
  lastName: string;
};

/**
 * Describes the message api.v1.SignupRequest.
 * Use `create(SignupRequestSchema)` to create a new message.
 */
export const SignupRequestSchema: GenMessage<SignupRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_auth_service, 0);

/**
 * @generated from message api.v1.SignupResponse
 */
export type SignupResponse = Message<"api.v1.SignupResponse"> & {
};

/**
 * Describes the message api.v1.SignupResponse.
 * Use `create(SignupResponseSchema)` to create a new message.
 */
export const SignupResponseSchema: GenMessage<SignupResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_auth_service, 1);

/**
 * @generated from message api.v1.LoginRequest
 */
export type LoginRequest = Message<"api.v1.LoginRequest"> & {
  /**
   * @generated from field: string email = 1;
   */
  email: string;

  /**
   * @generated from field: string password = 2;
   */
  password: string;
};

/**
 * Describes the message api.v1.LoginRequest.
 * Use `create(LoginRequestSchema)` to create a new message.
 */
export const LoginRequestSchema: GenMessage<LoginRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_auth_service, 2);

/**
 * @generated from message api.v1.LoginResponse
 */
export type LoginResponse = Message<"api.v1.LoginResponse"> & {
  /**
   * @generated from field: string access_token = 1;
   */
  accessToken: string;
};

/**
 * Describes the message api.v1.LoginResponse.
 * Use `create(LoginResponseSchema)` to create a new message.
 */
export const LoginResponseSchema: GenMessage<LoginResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_auth_service, 3);

/**
 * @generated from message api.v1.RefreshTokenRequest
 */
export type RefreshTokenRequest = Message<"api.v1.RefreshTokenRequest"> & {
};

/**
 * Describes the message api.v1.RefreshTokenRequest.
 * Use `create(RefreshTokenRequestSchema)` to create a new message.
 */
export const RefreshTokenRequestSchema: GenMessage<RefreshTokenRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_auth_service, 4);

/**
 * @generated from message api.v1.RefreshTokenResponse
 */
export type RefreshTokenResponse = Message<"api.v1.RefreshTokenResponse"> & {
  /**
   * @generated from field: string access_token = 1;
   */
  accessToken: string;
};

/**
 * Describes the message api.v1.RefreshTokenResponse.
 * Use `create(RefreshTokenResponseSchema)` to create a new message.
 */
export const RefreshTokenResponseSchema: GenMessage<RefreshTokenResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_auth_service, 5);

/**
 * @generated from message api.v1.LogoutRequest
 */
export type LogoutRequest = Message<"api.v1.LogoutRequest"> & {
};

/**
 * Describes the message api.v1.LogoutRequest.
 * Use `create(LogoutRequestSchema)` to create a new message.
 */
export const LogoutRequestSchema: GenMessage<LogoutRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_auth_service, 6);

/**
 * @generated from message api.v1.LogoutResponse
 */
export type LogoutResponse = Message<"api.v1.LogoutResponse"> & {
};

/**
 * Describes the message api.v1.LogoutResponse.
 * Use `create(LogoutResponseSchema)` to create a new message.
 */
export const LogoutResponseSchema: GenMessage<LogoutResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_auth_service, 7);

/**
 * @generated from message api.v1.VerifyEmailRequest
 */
export type VerifyEmailRequest = Message<"api.v1.VerifyEmailRequest"> & {
  /**
   * @generated from field: string token = 1;
   */
  token: string;
};

/**
 * Describes the message api.v1.VerifyEmailRequest.
 * Use `create(VerifyEmailRequestSchema)` to create a new message.
 */
export const VerifyEmailRequestSchema: GenMessage<VerifyEmailRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_auth_service, 8);

/**
 * @generated from message api.v1.VerifyEmailResponse
 */
export type VerifyEmailResponse = Message<"api.v1.VerifyEmailResponse"> & {
};

/**
 * Describes the message api.v1.VerifyEmailResponse.
 * Use `create(VerifyEmailResponseSchema)` to create a new message.
 */
export const VerifyEmailResponseSchema: GenMessage<VerifyEmailResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_auth_service, 9);

/**
 * @generated from message api.v1.ResetPasswordRequest
 */
export type ResetPasswordRequest = Message<"api.v1.ResetPasswordRequest"> & {
  /**
   * @generated from field: string email = 1;
   */
  email: string;
};

/**
 * Describes the message api.v1.ResetPasswordRequest.
 * Use `create(ResetPasswordRequestSchema)` to create a new message.
 */
export const ResetPasswordRequestSchema: GenMessage<ResetPasswordRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_auth_service, 10);

/**
 * @generated from message api.v1.ResetPasswordResponse
 */
export type ResetPasswordResponse = Message<"api.v1.ResetPasswordResponse"> & {
};

/**
 * Describes the message api.v1.ResetPasswordResponse.
 * Use `create(ResetPasswordResponseSchema)` to create a new message.
 */
export const ResetPasswordResponseSchema: GenMessage<ResetPasswordResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_auth_service, 11);

/**
 * @generated from message api.v1.UpdatePasswordRequest
 */
export type UpdatePasswordRequest = Message<"api.v1.UpdatePasswordRequest"> & {
  /**
   * @generated from field: string token = 1;
   */
  token: string;

  /**
   * @generated from field: string password = 2;
   */
  password: string;

  /**
   * @generated from field: string password_confirmation = 3;
   */
  passwordConfirmation: string;
};

/**
 * Describes the message api.v1.UpdatePasswordRequest.
 * Use `create(UpdatePasswordRequestSchema)` to create a new message.
 */
export const UpdatePasswordRequestSchema: GenMessage<UpdatePasswordRequest> = /*@__PURE__*/
  messageDesc(file_api_v1_auth_service, 12);

/**
 * @generated from message api.v1.UpdatePasswordResponse
 */
export type UpdatePasswordResponse = Message<"api.v1.UpdatePasswordResponse"> & {
};

/**
 * Describes the message api.v1.UpdatePasswordResponse.
 * Use `create(UpdatePasswordResponseSchema)` to create a new message.
 */
export const UpdatePasswordResponseSchema: GenMessage<UpdatePasswordResponse> = /*@__PURE__*/
  messageDesc(file_api_v1_auth_service, 13);

/**
 * @generated from service api.v1.AuthService
 */
export const AuthService: GenService<{
  /**
   * @generated from rpc api.v1.AuthService.Signup
   */
  signup: {
    methodKind: "unary";
    input: typeof SignupRequestSchema;
    output: typeof SignupResponseSchema;
  },
  /**
   * @generated from rpc api.v1.AuthService.Login
   */
  login: {
    methodKind: "unary";
    input: typeof LoginRequestSchema;
    output: typeof LoginResponseSchema;
  },
  /**
   * @generated from rpc api.v1.AuthService.RefreshToken
   */
  refreshToken: {
    methodKind: "unary";
    input: typeof RefreshTokenRequestSchema;
    output: typeof RefreshTokenResponseSchema;
  },
  /**
   * @generated from rpc api.v1.AuthService.Logout
   */
  logout: {
    methodKind: "unary";
    input: typeof LogoutRequestSchema;
    output: typeof LogoutResponseSchema;
  },
  /**
   * @generated from rpc api.v1.AuthService.VerifyEmail
   */
  verifyEmail: {
    methodKind: "unary";
    input: typeof VerifyEmailRequestSchema;
    output: typeof VerifyEmailResponseSchema;
  },
  /**
   * @generated from rpc api.v1.AuthService.ResetPassword
   */
  resetPassword: {
    methodKind: "unary";
    input: typeof ResetPasswordRequestSchema;
    output: typeof ResetPasswordResponseSchema;
  },
  /**
   * @generated from rpc api.v1.AuthService.UpdatePassword
   */
  updatePassword: {
    methodKind: "unary";
    input: typeof UpdatePasswordRequestSchema;
    output: typeof UpdatePasswordResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_api_v1_auth_service, 0);

