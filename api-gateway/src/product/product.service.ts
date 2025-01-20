import { CACHE_MANAGER, CacheStore } from '@nestjs/cache-manager';
import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import {
  CreateProductRequest,
  DeleteProductRequest,
  GetProductByIdRequest,
  ProductsByShopRequest,
  ProductResponse,
  ProductServiceClient,
  UpdateProductRequest,
  GetProductsByShopResponse,
} from 'pb/product/product_service';
import {
  CreateShopRequest,
  DeleteShopRequest,
  GetShopByIdRequest,
  GetShopsByOwnerRequest,
  GetShopsByOwnerResponse,
  ShopResponse,
  ShopServiceClient,
  UpdateShopRequest,
} from 'pb/product/shop_service';
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

    await this.cacheManager.set(shop.shop.id, shop, 60000);

    const cachedData = await this.cacheManager.get<ShopResponse>(shop.shop.id);
    console.log('cached_data', cachedData);

    return cachedData;
  }

  // get shop by id
  async getShopById(request: GetShopByIdRequest) {
    const cachedData = await this.cacheManager.get<ShopResponse>(request.id);

    if (!cachedData) {
      const res = this.shopService.getShopById(request);

      const shop = await lastValueFrom(res);

      await this.cacheManager.set(shop.shop.id, shop, 60000);

      return shop;
    }

    return cachedData;
  }

  // get shops by owner
  async getShopsByOwner(request: GetShopsByOwnerRequest) {
    if (request.queryString === '{}') {
      const cachedData = await this.cacheManager.get<GetShopsByOwnerResponse>(
        request.id,
      );

      if (!cachedData) {
        const res = this.shopService.getShopsByOwner(request);

        const shops = await lastValueFrom(res);
        await this.cacheManager.set(request.id, shops);

        return shops;
      }

      return cachedData;
    }

    const cachedData = await this.cacheManager.get<GetShopsByOwnerResponse>(
      request.queryString,
    );

    if (!cachedData) {
      const res = this.shopService.getShopsByOwner(request);

      const shops = await lastValueFrom(res);
      await this.cacheManager.set(request.queryString, shops);

      return shops;
    }

    return cachedData;
  }

  // update shop
  async updateShop(request: UpdateShopRequest) {
    const res = this.shopService.updateShop(request);

    const shop = await lastValueFrom(res);

    // get cahce if it exists
    const cachedData = await this.cacheManager.get<ShopResponse>(request.id);
    if (cachedData) {
      await this.cacheManager.del(request.id);

      await this.cacheManager.set(shop.shop.id, shop, 60000);
    } else if (!cachedData) {
      await this.cacheManager.set(shop.shop.id, shop, 60000);
    }

    return shop;
  }

  // delete shop
  async deleteShop(request: DeleteShopRequest) {
    await this.cacheManager.del(request.id);
    const res = this.shopService.deleteShop(request);

    return await lastValueFrom(res);
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

    await this.cacheManager.set(product.product.id, product, 60000);

    const cachedData = await this.cacheManager.get<ProductResponse>(
      product.product.id,
    );
    console.log('2', cachedData);

    return cachedData;
  }

  // get product by id
  async getProductByID(request: GetProductByIdRequest) {
    const cachedData = await this.cacheManager.get<ProductResponse>(request.id);

    if (!cachedData) {
      const res = this.productService.getProductById(request);

      const product = await lastValueFrom(res);

      await this.cacheManager.set(product.product.id, product);

      return product;
    }

    return cachedData;
  }

  // get products by shop
  async getProductsByShop(request: ProductsByShopRequest) {
    if (request.queryString === '{}') {
      const cachedData = await this.cacheManager.get<GetProductsByShopResponse>(
        request.id,
      );
      console.log(cachedData);

      if (!cachedData) {
        // get products from the product grpc service
        const res = this.productService.getProductsByShop(request);
        const products = await lastValueFrom(res);

        // cache the product
        await this.cacheManager.set(request.id, products, 60000);
        return products;
      }

      return cachedData;
    }

    const cachedData = await this.cacheManager.get<GetProductsByShopResponse>(
      request.queryString,
    );
    console.log(cachedData);

    if (!cachedData) {
      // get products from the product grpc service
      const res = this.productService.getProductsByShop(request);
      const products = await lastValueFrom(res);

      // cache the products
      await this.cacheManager.set(request.queryString, products, 60000);
      return products;
    }

    return cachedData;
  }

  // update product
  async updateProduct(request: UpdateProductRequest) {
    const res = this.productService.updateProduct(request);

    const product = await lastValueFrom(res);

    const cachedData = await this.cacheManager.get<ProductResponse>(request.id);
    if (cachedData) {
      await this.cacheManager.del(request.id);

      await this.cacheManager.set(product.product.id, product, 60000);
    } else if (!cachedData) {
      await this.cacheManager.set(product.product.id, product, 60000);
    }

    return product;
  }

  // delete product
  async deleteProduct(request: DeleteProductRequest) {
    await this.cacheManager.del(request.id);

    const res = this.productService.deleteProduct(request);
    return await lastValueFrom(res);
  }
}
