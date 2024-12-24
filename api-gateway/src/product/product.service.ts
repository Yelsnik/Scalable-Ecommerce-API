import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import {
  CreateProductRequest,
  DeleteProductRequest,
  GetProductByIdRequest,
  GetProductsByShopRequest,
  ProductServiceClient,
  UpdateProductRequest,
} from 'pb/product_service';
import {
  CreateShopRequest,
  DeleteShopRequest,
  GetShopByIdRequest,
  GetShopsByOwnerRequest,
  GetShopsByOwnerResponse,
  ShopResponse,
  ShopServiceClient,
  UpdateShopRequest,
} from 'pb/shop_service';
import { Observable } from 'rxjs';

@Injectable()
export class ProductService implements OnModuleInit {
  productService: ProductServiceClient;
  shopService: ShopServiceClient;

  constructor(
    @Inject('PRODUCT_SERVICE') private productClient: ClientGrpc,
    @Inject('SHOP_SERVICE') private shopClient: ClientGrpc,
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
  createShop(request: CreateShopRequest): Observable<ShopResponse> {
    const res = this.shopService.createShop(request);
    return res;
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

  addProduct(request: CreateProductRequest) {
    return this.productService.addProduct(request);
  }

  getProductByID(request: GetProductByIdRequest) {
    return this.productService.getProductById(request);
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
