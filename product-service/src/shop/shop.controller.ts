import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { CreateShopRequest, CreateShopResponse } from 'pb/rpc_create_shop';
import { ShopService } from './shop.service';

@Controller('shop')
export class ShopController {
    constructor(private readonly shopService: ShopService) {}

    @GrpcMethod('ShopService', 'CreateShop')
    createShop(request: CreateShopRequest): Promise<CreateShopResponse>{
        return this.shopService.CreateShop(request)
    }
}
