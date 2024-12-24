import { Controller, UseFilters } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { ProductService } from './product.service';
import {
  CreateProductRequest,
  DeleteProductRequest,
  EmptyRes,
  GetProductByIdRequest,
  GetProductsByShopRequest,
  GetProductsByShopResponse,
  ProductResponse,
  ProductServiceController,
  UpdateProductRequest,
} from 'pb/product_service';
import { Observable } from 'rxjs';
import {
  AllExceptionFilter,
  RpcExceptionFilters,
} from 'src/helpers/all-exception.filter';

@Controller('product')
@UseFilters(AllExceptionFilter, RpcExceptionFilters)
export class ProductController implements ProductServiceController {
  constructor(private readonly productService: ProductService) {}

  @GrpcMethod('ProductService', 'AddProduct')
  addProduct(
    request: CreateProductRequest,
  ): Promise<ProductResponse> | Observable<ProductResponse> | ProductResponse {
    return this.productService.AddProduct(request);
  }

  @GrpcMethod('ProductService', 'GetProductByID')
  getProductById(
    request: GetProductByIdRequest,
  ): Promise<ProductResponse> | Observable<ProductResponse> | ProductResponse {
    return this.productService.GetProductByID(request);
  }

  @GrpcMethod('ProductService', 'GetProductsByShop')
  getProductsByShop(
    request: GetProductsByShopRequest,
  ):
    | Promise<GetProductsByShopResponse>
    | Observable<GetProductsByShopResponse>
    | GetProductsByShopResponse {
    return this.productService.GetProductsByShop(request);
  }

  @GrpcMethod('ProductService', 'UpdateProduct')
  updateProduct(
    request: UpdateProductRequest,
  ): Promise<ProductResponse> | Observable<ProductResponse> | ProductResponse {
    return this.productService.UpdateProduct(request);
  }

  @GrpcMethod('ProductService', 'DeleteProduct')
  deleteProduct(
    request: DeleteProductRequest,
  ): Promise<EmptyRes> | Observable<EmptyRes> | EmptyRes {
    return this.productService.DeleteProduct(request);
  }
}
