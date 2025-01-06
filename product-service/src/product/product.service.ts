import { Injectable } from '@nestjs/common';
import { Model } from 'mongoose';
import { GridFsService } from 'src/grid-fs/grid-fs.service';
import { Product } from './product.schema';
import { InjectModel } from '@nestjs/mongoose';
import {
  CreateProductRequest,
  ProductResponse,
  Product as P,
  GetProductByIdRequest,
  GetProductsByShopRequest,
  GetProductsByShopResponse,
  UpdateProductRequest,
  DeleteProductRequest,
  EmptyRes,
} from 'pb/product_service';
import { Shop } from 'src/shop/shop.schema';
import { TimestampUtils } from 'src/interface/interface';
import { GrpcNotFoundException } from 'nestjs-grpc-exceptions';

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

    const id = await this.gridFs.uploadFile(file, 'product-image');

    const shop = await this.shopModel.findOne({ _id: request.shop }).exec();

    if (!shop) {
      throw new GrpcNotFoundException('shop does not exist');
    }

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

    const pi: string = product.id;

    const response: P = {
      id: pi,
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
      updatedAt: undefined,
      createdAt: TimestampUtils.dateToTimestamp(product.createdAt),
    };

    return {
      product: response,
    };
  }

  // get product by id
  async GetProductByID(
    request: GetProductByIdRequest,
  ): Promise<ProductResponse> {
    const product = await this.productModel.findById(request.id).exec();

    if (!product) {
      throw new GrpcNotFoundException('product does not exist');
    }

    const shop = await this.shopModel.findOne(product.shop).exec();

    if (!shop) {
      throw new GrpcNotFoundException('shop does not exist');
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
      updatedAt: undefined,
      createdAt: TimestampUtils.dateToTimestamp(product.createdAt),
    };

    return {
      product: res,
    };
  }

  // get product by shop
  async GetProductsByShop(
    request: GetProductsByShopRequest,
  ): Promise<GetProductsByShopResponse> {
    const products = await this.productModel.find({ shop: request.id }).exec();

    if (!products) {
      throw new GrpcNotFoundException('Products not found');
    }

    const res: P[] = products.map((product) => {
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
        shop: request.id,
        rating: product.rating,
        isFeatured: product.isFeatured,
        updatedAt: undefined,
        createdAt: TimestampUtils.dateToTimestamp(product.createdAt),
      };

      return res;
    });

    return {
      product: res,
    };
  }

  // update product
  async UpdateProduct(request: UpdateProductRequest): Promise<ProductResponse> {
    const body = {
      category: request.category,
      productName: request.productName,
      description: request.description,
      brand: request.brand,
      image: request.image,
      countInStock: request.countInStock,
      price: request.price,
      currency: request.currency,
      rating: request.rating,
      isFeatured: request.isFeatured,
    };
    const product = await this.productModel.findByIdAndUpdate(request.id, body);

    if (!product) {
      throw new GrpcNotFoundException('product not found or deleted');
    }

    const shop = await this.shopModel.findOne(product.shop).exec();

    if (!shop) {
      throw new GrpcNotFoundException('shop does not exist');
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
      updatedAt: TimestampUtils.dateToTimestamp(product.updatedAt),
      createdAt: TimestampUtils.dateToTimestamp(product.createdAt),
    };

    return {
      product: res,
    };
  }

  // delete product
  async DeleteProduct(request: DeleteProductRequest): Promise<EmptyRes> {
    return await this.productModel.findByIdAndDelete(request.id);
  }
}
