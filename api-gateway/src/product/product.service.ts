import { CACHE_MANAGER } from '@nestjs/cache-manager';
import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import {
  CreateProductRequest,
  DeleteProductRequest,
  GetProductByIdRequest,
  GetProductsByShopRequest,
  ProductResponse,
  ProductServiceClient,
  UpdateProductRequest,
} from 'pb/product-service/product_service';
import {
  CreateShopRequest,
  DeleteShopRequest,
  GetShopByIdRequest,
  GetShopsByOwnerRequest,
  GetShopsByOwnerResponse,
  ShopResponse,
  ShopServiceClient,
  UpdateShopRequest,
} from 'pb/product-service/shop_service';
import { lastValueFrom, Observable } from 'rxjs';
import { Cache } from 'cache-manager';

@Injectable()
export class ProductService implements OnModuleInit {
  productService: ProductServiceClient;
  shopService: ShopServiceClient;

  constructor(
    @Inject('PRODUCT_SERVICE') private productClient: ClientGrpc,
    @Inject('SHOP_SERVICE') private shopClient: ClientGrpc,
    @Inject(CACHE_MANAGER) private cacheManager: Cache,
  ) {}

  onModuleInit() {
    this.productService =
      this.productClient.getService<ProductServiceClient>('ProductService');
    this.shopService =
      this.shopClient.getService<ShopServiceClient>('ShopService');
  }

  /**
   * Service file for calling grpc shop service
   * Create shop
   * Get shop by id
   * Get shops by owner
   * Update shop by id
   * Delete shop
   */
  async createShop(request: CreateShopRequest) {
    const res = this.shopService.createShop(request);

    const shop = await lastValueFrom(res);

    const cachedData = await this.cacheManager.set(shop.shop.id, shop);

    return cachedData;
  }

  getShopById(request: GetShopByIdRequest): Observable<ShopResponse> {
    const shop = this.shopService.getShopById(request);
    return shop;
  }

  getShopsByOwner(
    request: GetShopsByOwnerRequest,
  ): Observable<GetShopsByOwnerResponse> {
    const shops = this.shopService.getShopsByOwner(request);
    return shops;
  }

  updateShop(request: UpdateShopRequest): Observable<ShopResponse> {
    const shop = this.shopService.updateShop(request);
    return shop;
  }

  deleteShop(request: DeleteShopRequest) {
    return this.shopService.deleteShop(request);
  }

  /**
   * Service file for calling grpc product service
   * Add product
   * Get product by id
   * Get products by shop
   * Update product by id
   * Delete product
   */

  async addProduct(request: CreateProductRequest) {
    const res = this.productService.addProduct(request);

    const product = await lastValueFrom(res);

    const cachedData = await this.cacheManager.set(product.product.id, product);

    return cachedData;
  }

  async getProductByID(request: GetProductByIdRequest) {
    const res = this.productService.getProductById(request);

    const product = await lastValueFrom(res);

    const cachedData = await this.cacheManager.get<ProductResponse>(
      product.product.id,
    );

    if (!cachedData) {
      return product;
    }

    return cachedData;
  }

  getProductsByShop(request: GetProductsByShopRequest) {
    return this.productService.getProductsByShop(request);
  }

  updateProduct(request: UpdateProductRequest) {
    return this.productService.updateProduct(request);
  }

  deleteProduct(request: DeleteProductRequest) {
    return this.productService.deleteProduct(request);
  }
}
