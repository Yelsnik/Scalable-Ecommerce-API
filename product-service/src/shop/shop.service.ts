import { Injectable, InternalServerErrorException } from '@nestjs/common';
import { Shop } from './shop.schema';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { ShopService as SH } from 'pb/shop_service';
import { CreateShopRequest, CreateShopResponse } from 'pb/rpc_create_shop';
import { GridFsService } from 'src/grid-fs/grid-fs.service';
import { GridFsDto } from 'src/grid-fs/dto/grid-fs.dto';


@Injectable()
export class ShopService implements SH {
    constructor(@InjectModel(Shop.name) private shopModel: Model<Shop>, private gridFs: GridFsService){}

   async  CreateShop(request: CreateShopRequest): Promise<CreateShopResponse> {
       try{
        const buffer = Buffer.from(request.file.buffer);

        const file = this.gridFs.loadImage(request.file, buffer)
 
       const result = await this.gridFs.uploadFile(file, "shop-image")
 
        const shop = await this.shopModel.create({
         name: request.name,
         description: request.description,
         imageID: result.id,
         image: request.file.filename,
         shopOwner: request.shopOwner
        })
 
        const response = {
         id: shop.id,
         name: shop.name,
         description: shop.description,
         imageName: shop.image,
         shopOwner: shop.shopOwner
        }

       return {
         shop: response
       }

       }catch(err){
        throw new InternalServerErrorException("error creating shop", err)
       }
         
    } 
}
