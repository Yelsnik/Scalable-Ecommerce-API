import { Injectable, InternalServerErrorException } from '@nestjs/common';
import { Model } from 'mongoose';
import { ProductService as PS } from 'pb/product_service';
import { CreateProductRequest, CreateProductResponse } from 'pb/rpc_create_product';
import { GridFsDto } from 'src/grid-fs/dto/grid-fs.dto';
import { GridFsService } from 'src/grid-fs/grid-fs.service';
import { Product } from './product.schema';
import { InjectModel } from '@nestjs/mongoose';
import { Product as P } from 'pb/product';

@Injectable()
export class ProductService implements PS {
     constructor(@InjectModel(Product.name) private productModel: Model<Product>, private gridFs: GridFsService){}
    
   async  AddProduct(request: CreateProductRequest): Promise<CreateProductResponse> {

       try {
        const buffer = Buffer.from(request.image.buffer);
        const file = this.gridFs.loadImage(request.image, buffer)
 
         const result = await this.gridFs.uploadFile(file, "shop-image")
 
         const product = await this.productModel.create({
             category: request.category,
             productName: request.productName,
             description: request.description,
             brand: request.brand,
             imageID: result.id,
             image: result.filename,
             countInStock: request.countInStock,
             price: request.price,
             currency: request.currency,
             shop: request.shop,
             rating: request.rating,
             isFeatured: request.isFeatured
         })
 
         const response: P  = {
             id: product.id,
             category: product.category,
             productName: product.productName,
             description: product.description,
             brand: product.brand,
             image: product.image,
             countInStock: product.countInStock,
             price: product.price,
             currency: product.currency,
             shop: request.shop,
             rating: product.rating,
             isFeatured: product.isFeatured,
             updatedAt: product.updatedAt,
             createdAt: product.createdAt
 
         }
 
         return {
             message: "product added successfully!",
             product: response
         }
       }catch (err){
        throw new InternalServerErrorException("error adding product", err)
       }
    }
}
