// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               v5.29.2
// source: shop_service.proto

/* eslint-disable */
import { GrpcMethod, GrpcStreamMethod } from "@nestjs/microservices";
import { Observable } from "rxjs";
import { File } from "./file";

export const protobufPackage = "product";

export interface CreateShopRequest {
  name: string;
  description: string;
  shopOwner: string;
  image: File | undefined;
}

export interface ShopResponse {
  shop: Shop | undefined;
}

export interface GetShopByIdRequest {
  id: string;
}

export interface GetShopsByOwnerRequest {
  id: string;
}

export interface GetShopsByOwnerResponse {
  shops: Shop[];
}

export interface UpdateShopRequest {
  id: string;
  name?: string | undefined;
  description?: string | undefined;
}

export interface DeleteShopRequest {
  id: string;
}

export interface Empty {
}

export interface Shop {
  id: string;
  name: string;
  description: string;
  imageName: string;
  shopOwner: string;
}

export const PRODUCT_PACKAGE_NAME = "product";

export interface ShopServiceClient {
  createShop(request: CreateShopRequest): Observable<ShopResponse>;

  getShopById(request: GetShopByIdRequest): Observable<ShopResponse>;

  getShopsByOwner(request: GetShopsByOwnerRequest): Observable<GetShopsByOwnerResponse>;

  updateShop(request: UpdateShopRequest): Observable<ShopResponse>;

  deleteShop(request: DeleteShopRequest): Observable<Empty>;
}

export interface ShopServiceController {
  createShop(request: CreateShopRequest): Promise<ShopResponse> | Observable<ShopResponse> | ShopResponse;

  getShopById(request: GetShopByIdRequest): Promise<ShopResponse> | Observable<ShopResponse> | ShopResponse;

  getShopsByOwner(
    request: GetShopsByOwnerRequest,
  ): Promise<GetShopsByOwnerResponse> | Observable<GetShopsByOwnerResponse> | GetShopsByOwnerResponse;

  updateShop(request: UpdateShopRequest): Promise<ShopResponse> | Observable<ShopResponse> | ShopResponse;

  deleteShop(request: DeleteShopRequest): Promise<Empty> | Observable<Empty> | Empty;
}

export function ShopServiceControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods: string[] = ["createShop", "getShopById", "getShopsByOwner", "updateShop", "deleteShop"];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcMethod("ShopService", method)(constructor.prototype[method], method, descriptor);
    }
    const grpcStreamMethods: string[] = [];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcStreamMethod("ShopService", method)(constructor.prototype[method], method, descriptor);
    }
  };
}

export const SHOP_SERVICE_NAME = "ShopService";
