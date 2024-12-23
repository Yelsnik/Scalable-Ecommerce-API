import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { ProductServiceClient } from 'pb/product_service';
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
}
