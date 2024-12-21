import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { ShopService } from './shop.service';
import { CreateShopRequest, DeleteShopRequest, Empty, GetShopByIdRequest, GetShopsByOwnerRequest, GetShopsByOwnerResponse, ShopResponse, ShopServiceController, UpdateShopRequest } from 'pb/shop_service';
import { ServerUnaryCall } from '@grpc/grpc-js';
import { Observable } from 'rxjs';

@Controller('shop')
export class ShopController implements ShopServiceController   {
    constructor(private readonly shopService: ShopService) {} 
    
    @GrpcMethod('ShopService', 'CreateShop')
    createShop(request: CreateShopRequest): Promise<ShopResponse> | Observable<ShopResponse> | ShopResponse {
        console.log(request) 
        return this.shopService.createShop(request)
    }

    @GrpcMethod('ShopService', 'GetShopByID')
     getShopById(request: GetShopByIdRequest): Promise<ShopResponse> | Observable<ShopResponse> | ShopResponse {
        return this.shopService.getShopById(request)
    }
    getShopsByOwner(request: GetShopsByOwnerRequest): Promise<GetShopsByOwnerResponse> | Observable<GetShopsByOwnerResponse> | GetShopsByOwnerResponse {
       return  this.shopService.getShopsByOwner(request)
    }
    updateShop(request: UpdateShopRequest): Promise<ShopResponse> | Observable<ShopResponse> | ShopResponse {
        throw new Error('Method not implemented.');
    }
    deleteShop(request: DeleteShopRequest): Promise<Empty> | Observable<Empty> | Empty {
        throw new Error('Method not implemented.');
    }
}
