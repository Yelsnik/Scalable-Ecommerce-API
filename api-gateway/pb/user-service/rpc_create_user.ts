// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               v5.29.2
// source: user-service/rpc_create_user.proto

/* eslint-disable */
import { User } from "./user";

export const protobufPackage = "pb";

export interface CreateUserRequest {
  name: string;
  email: string;
  password: string;
  role: string;
}

export interface CreateUserResponse {
  user: User | undefined;
}

export const PB_PACKAGE_NAME = "pb";
