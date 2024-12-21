import { Injectable, InternalServerErrorException } from '@nestjs/common';
import { Model } from 'mongoose';
import { GridFsService } from 'src/grid-fs/grid-fs.service';
import { Product } from './product.schema';
import { InjectModel } from '@nestjs/mongoose';
import { Time} from './interface';
import { CreateProductRequest, ProductResponse, Product as P } from 'pb/product_service';

@Injectable()
export class ProductService  {
     constructor(@InjectModel(Product.name) private productModel: Model<Product>, private gridFs: GridFsService){}
   
     async  AddProduct(request: CreateProductRequest): Promise<ProductResponse> {
        try {
            const buffer = Buffer.from(request.image.buffer);
            const file = this.gridFs.loadImage(request.image, buffer)
     
             const id = await this.gridFs.uploadFile(file, "shop-image")
     
             const product = await this.productModel.create({
                 category: request.category,
                 productName: request.productName,
                 description: request.description,
                 brand: request.brand,
                 imageID: id,
                 image: request.image.originalname,
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
                 imageName: product.image,
                 countInStock: product.countInStock,
                 price: product.price,
                 currency: product.currency,
                 shop: request.shop,
                 rating: product.rating,
                 isFeatured: product.isFeatured,
                 updatedAt: new Time(product.updatedAt),
                 createdAt: new Time(product.createdAt)
     
             }
     
             return {
                 product: response
             }
           }catch (err){
            throw new InternalServerErrorException("error adding product", err)
           }
    }
 
}
