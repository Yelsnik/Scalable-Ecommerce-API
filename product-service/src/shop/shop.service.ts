import { Injectable } from '@nestjs/common';
import { Shop } from './shop.schema';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { ShopService as SH } from 'pb/shop_service';
import { CreateShopRequest, CreateShopResponse } from 'pb/rpc_create_shop';


@Injectable()
export class ShopService implements SH {
    constructor(@InjectModel(Shop.name) private shopModel: Model<Shop>){}


    CreateShop(request: CreateShopRequest): Promise<CreateShopResponse> {
        throw new Error('Method not implemented.');
    }
}
