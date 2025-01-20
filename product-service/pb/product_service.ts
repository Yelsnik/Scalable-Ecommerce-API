// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               v5.29.2
// source: product_service.proto

/* eslint-disable */
import { GrpcMethod, GrpcStreamMethod } from "@nestjs/microservices";
import { Observable } from "rxjs";
import { File } from "./file";
import { Timestamp } from "./google/protobuf/timestamp";

export const protobufPackage = "product";

export interface Product {
  id: string;
  category: string;
  productName: string;
  description: string;
  brand: string;
  imageName: string;
  countInStock: number;
  price: number;
  currency: string;
  shop: string;
  rating: number;
  isFeatured: boolean;
  updatedAt: Timestamp | undefined;
  createdAt: Timestamp | undefined;
}

export interface CreateProductRequest {
  category: string;
  productName: string;
  description: string;
  brand: string;
  image: File | undefined;
  countInStock: number;
  price: number;
  currency: string;
  shop: string;
  rating: number;
  isFeatured: boolean;
}

export interface ProductResponse {
  product: Product | undefined;
}

export interface GetProductByIdRequest {
  id: string;
  user?: string | undefined;
}

export interface ProductsByShopRequest {
  queryString: string;
  id: string;
}

export interface GetProductsByShopResponse {
  product: Product[];
}

export interface UpdateProductRequest {
  id: string;
  category?: string | undefined;
  productName?: string | undefined;
  description?: string | undefined;
  brand?: string | undefined;
  image?: File | undefined;
  countInStock?: number | undefined;
  price?: number | undefined;
  currency?: string | undefined;
  rating?: number | undefined;
  isFeatured?: boolean | undefined;
}

export interface DeleteProductRequest {
  id: string;
}

export interface EmptyRes {
}

export const PRODUCT_PACKAGE_NAME = "product";

export interface ProductServiceClient {
  addProduct(request: CreateProductRequest): Observable<ProductResponse>;

  getProductById(request: GetProductByIdRequest): Observable<ProductResponse>;

  getProductsByShop(request: ProductsByShopRequest): Observable<GetProductsByShopResponse>;

  updateProduct(request: UpdateProductRequest): Observable<ProductResponse>;

  deleteProduct(request: DeleteProductRequest): Observable<EmptyRes>;
}

export interface ProductServiceController {
  addProduct(request: CreateProductRequest): Promise<ProductResponse> | Observable<ProductResponse> | ProductResponse;

  getProductById(
    request: GetProductByIdRequest,
  ): Promise<ProductResponse> | Observable<ProductResponse> | ProductResponse;

  getProductsByShop(
    request: ProductsByShopRequest,
  ): Promise<GetProductsByShopResponse> | Observable<GetProductsByShopResponse> | GetProductsByShopResponse;

  updateProduct(
    request: UpdateProductRequest,
  ): Promise<ProductResponse> | Observable<ProductResponse> | ProductResponse;

  deleteProduct(request: DeleteProductRequest): Promise<EmptyRes> | Observable<EmptyRes> | EmptyRes;
}

export function ProductServiceControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods: string[] = [
      "addProduct",
      "getProductById",
      "getProductsByShop",
      "updateProduct",
      "deleteProduct",
    ];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcMethod("ProductService", method)(constructor.prototype[method], method, descriptor);
    }
    const grpcStreamMethods: string[] = [];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcStreamMethod("ProductService", method)(constructor.prototype[method], method, descriptor);
    }
  };
}

export const PRODUCT_SERVICE_NAME = "ProductService";