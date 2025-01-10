// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               v5.29.2
// source: cart-service/cart_service.proto

/* eslint-disable */
import { GrpcMethod, GrpcStreamMethod } from "@nestjs/microservices";
import { Observable } from "rxjs";

export const protobufPackage = "cart";

export interface CreateCartRequest {
  userId: string;
  totalPrice: number;
}

export interface CartResponse {
  id: string;
  userId: string;
  totalPrice: number;
}

export interface GetCartByUserIDRequest {
  id: string;
}

export interface GetCartByIDRequest {
  id: string;
}

export interface UpdateCartRequest {
  id: string;
  totalPrice: number;
}

export const CART_PACKAGE_NAME = "cart";

export interface CartServiceClient {
  createCart(request: CreateCartRequest): Observable<CartResponse>;

  getCartByUserId(request: GetCartByUserIDRequest): Observable<CartResponse>;

  getCart(request: GetCartByIDRequest): Observable<CartResponse>;

  updateCart(request: UpdateCartRequest): Observable<CartResponse>;
}

export interface CartServiceController {
  createCart(request: CreateCartRequest): Promise<CartResponse> | Observable<CartResponse> | CartResponse;

  getCartByUserId(request: GetCartByUserIDRequest): Promise<CartResponse> | Observable<CartResponse> | CartResponse;

  getCart(request: GetCartByIDRequest): Promise<CartResponse> | Observable<CartResponse> | CartResponse;

  updateCart(request: UpdateCartRequest): Promise<CartResponse> | Observable<CartResponse> | CartResponse;
}

export function CartServiceControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods: string[] = ["createCart", "getCartByUserId", "getCart", "updateCart"];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcMethod("CartService", method)(constructor.prototype[method], method, descriptor);
    }
    const grpcStreamMethods: string[] = [];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcStreamMethod("CartService", method)(constructor.prototype[method], method, descriptor);
    }
  };
}

export const CART_SERVICE_NAME = "CartService";