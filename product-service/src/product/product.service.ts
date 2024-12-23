import {
  BadRequestException,
  Injectable,
  InternalServerErrorException,
} from '@nestjs/common';
import { Model, Types } from 'mongoose';
import { GridFsService } from 'src/grid-fs/grid-fs.service';
import { Product } from './product.schema';
import { InjectModel } from '@nestjs/mongoose';
import {
  CreateProductRequest,
  ProductResponse,
  Product as P,
  GetProductByIdRequest,
} from 'pb/product_service';
import { Shop } from 'src/shop/shop.schema';
import { Time } from 'src/interface/interface';

@Injectable()
export class ProductService {
  constructor(
    @InjectModel(Product.name) private productModel: Model<Product>,
    @InjectModel(Shop.name) private shopModel: Model<Product>,
    private gridFs: GridFsService,
  ) {}

  // add product
  async AddProduct(request: CreateProductRequest): Promise<ProductResponse> {
    const buffer = Buffer.from(request.image.buffer);
    const file = this.gridFs.loadImage(request.image, buffer);

    const id = await this.gridFs.uploadFile(file, 'shop-image');

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
      isFeatured: request.isFeatured,
    });

    const response: P = {
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
      createdAt: new Time(product.createdAt),
    };

    return {
      product: response,
    };
  }

  async GetProductByID(
    request: GetProductByIdRequest,
  ): Promise<ProductResponse> {
    const product = await this.productModel.findById(request.id).exec();

    if (!product) {
      throw new InternalServerErrorException('product does not exist');
    }

    const shop = await this.shopModel.findOne(product.shop).exec();

    if (!shop) {
      throw new InternalServerErrorException('shop does not exist');
    }

    const id = shop._id.toString();

    const res: P = {
      id: product.id,
      category: product.category,
      productName: product.productName,
      description: product.description,
      brand: product.brand,
      imageName: product.image,
      countInStock: product.countInStock,
      price: product.price,
      currency: product.currency,
      shop: id,
      rating: product.rating,
      isFeatured: product.isFeatured,
      updatedAt: new Time(product.updatedAt),
      createdAt: new Time(product.createdAt),
    };

    return {
      product: res,
    };
  }
}
