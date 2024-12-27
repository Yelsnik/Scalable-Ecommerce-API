// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               v5.29.2
// source: auth_service.proto

/* eslint-disable */
import { GrpcMethod, GrpcStreamMethod } from "@nestjs/microservices";
import { Observable } from "rxjs";
import { Duration } from "./google/protobuf/duration";
import { Timestamp } from "./google/protobuf/timestamp";
import { User } from "./user";

export const protobufPackage = "pb";

export interface Payload {
  id: string;
  userId: string;
  role: string;
  issuedAt: Timestamp | undefined;
  expiredAt: Timestamp | undefined;
}

export interface GetUserByIdRequest {
  id: string;
}

export interface GetUserByIdResponse {
  user: User | undefined;
}

export interface GetUserByEmailRequest {
  email: string;
}

export interface GetUserByEmailResponse {
  user: User | undefined;
}

export interface CreateTokenRequest {
  userId: string;
  role: string;
  duration: Duration | undefined;
}

export interface CreateTokenResponse {
  token: string;
  payload: Payload | undefined;
}

export interface VerifyTokenRequest {
  token: string;
}

export interface VerifyTokenResponse {
  payload: Payload | undefined;
}

export const PB_PACKAGE_NAME = "pb";

export interface AuthServiceClient {
  getUserById(request: GetUserByIdRequest): Observable<GetUserByIdResponse>;

  getUserByEmail(request: GetUserByEmailRequest): Observable<GetUserByEmailResponse>;

  createToken(request: CreateTokenRequest): Observable<CreateTokenResponse>;

  verifyToken(request: VerifyTokenRequest): Observable<VerifyTokenResponse>;
}

export interface AuthServiceController {
  getUserById(
    request: GetUserByIdRequest,
  ): Promise<GetUserByIdResponse> | Observable<GetUserByIdResponse> | GetUserByIdResponse;

  getUserByEmail(
    request: GetUserByEmailRequest,
  ): Promise<GetUserByEmailResponse> | Observable<GetUserByEmailResponse> | GetUserByEmailResponse;

  createToken(
    request: CreateTokenRequest,
  ): Promise<CreateTokenResponse> | Observable<CreateTokenResponse> | CreateTokenResponse;

  verifyToken(
    request: VerifyTokenRequest,
  ): Promise<VerifyTokenResponse> | Observable<VerifyTokenResponse> | VerifyTokenResponse;
}

export function AuthServiceControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods: string[] = ["getUserById", "getUserByEmail", "createToken", "verifyToken"];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcMethod("AuthService", method)(constructor.prototype[method], method, descriptor);
    }
    const grpcStreamMethods: string[] = [];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcStreamMethod("AuthService", method)(constructor.prototype[method], method, descriptor);
    }
  };
}

export const AUTH_SERVICE_NAME = "AuthService";
