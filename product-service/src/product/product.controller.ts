import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { ProductService } from './product.service';
import { CreateProductRequest, CreateProductResponse } from 'pb/rpc_create_product';

@Controller('product')
export class ProductController {
    constructor(private readonly productService: ProductService) {}

    @GrpcMethod('ProductService', 'AddProduct')
    createShop(request: CreateProductRequest): Promise<CreateProductResponse>{
        return this.productService.AddProduct(request)
    }
}
