import { Controller, UseFilters } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { ShopService } from './shop.service';
import {
  CreateShopRequest,
  DeleteShopRequest,
  Empty,
  GetShopByIdRequest,
  GetShopsByOwnerRequest,
  GetShopsByOwnerResponse,
  ShopResponse,
  ShopServiceController,
  UpdateShopRequest,
} from 'pb/shop_service';
import { Observable } from 'rxjs';
import { MongoExceptionFilter } from 'src/helpers/mongo-exception.filter';

@Controller('shop')
@UseFilters(new MongoExceptionFilter())
export class ShopController implements ShopServiceController {
  constructor(private readonly shopService: ShopService) {}

  @GrpcMethod('ShopService', 'CreateShop')
  createShop(
    request: CreateShopRequest,
  ): Promise<ShopResponse> | Observable<ShopResponse> | ShopResponse {
    console.log(request);
    return this.shopService.createShop(request);
  }

  @GrpcMethod('ShopService', 'GetShopByID')
  getShopById(
    request: GetShopByIdRequest,
  ): Promise<ShopResponse> | Observable<ShopResponse> | ShopResponse {
    return this.shopService.getShopById(request);
  }

  @GrpcMethod('ShopService', 'GetShopsByOwner')
  getShopsByOwner(
    request: GetShopsByOwnerRequest,
  ):
    | Promise<GetShopsByOwnerResponse>
    | Observable<GetShopsByOwnerResponse>
    | GetShopsByOwnerResponse {
    return this.shopService.getShopsByOwner(request);
  }

  @GrpcMethod('ShopService', 'UpdateShop')
  updateShop(
    request: UpdateShopRequest,
  ): Promise<ShopResponse> | Observable<ShopResponse> | ShopResponse {
    return this.shopService.updateShop(request);
  }

  @GrpcMethod('ShopService', 'DeleteShop')
  deleteShop(
    request: DeleteShopRequest,
  ): Promise<Empty> | Observable<Empty> | Empty {
    return this.shopService.deleteShop(request);
  }
}
