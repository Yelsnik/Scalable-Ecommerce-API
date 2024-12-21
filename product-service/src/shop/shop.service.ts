import { Injectable, InternalServerErrorException } from '@nestjs/common';
import { Shop } from './shop.schema';
import { InjectModel } from '@nestjs/mongoose';
import { Model, Types } from 'mongoose';
import { CreateShopRequest, Shop as S, GetShopByIdRequest, GetShopsByOwnerRequest, GetShopsByOwnerResponse, ShopResponse } from 'pb/shop_service';
import { GridFsService } from 'src/grid-fs/grid-fs.service';



@Injectable()
export class ShopService {
    constructor(@InjectModel(Shop.name) private shopModel: Model<Shop>, private gridFs: GridFsService){}
  
    // create shop
  async  createShop(request: CreateShopRequest): Promise<ShopResponse> {
    const buffer = Buffer.from(request.image.buffer);

    const file = this.gridFs.loadImage(request.image, buffer)


     const id = await this.gridFs.uploadFile(file, "shop-image")
     console.log(id)


      const shop = await this.shopModel.create({
       name: request.name,
       description: request.description,
       imageID: id,
       image: request.image.originalname,
       shopOwner: request.shopOwner
      })

      const res: ShopResponse = {
        shop: {
          name: shop.name,
          description: shop.description,
          shopOwner: shop.shopOwner,
          id: shop.id,
          imageName: shop.image
        }
      }

     return res
  }

  // get shop by id
  async getShopById(request: GetShopByIdRequest): Promise<ShopResponse> {
   // const id = new Types.ObjectId(request.id)

    const shop = await this.shopModel.findById(request.id).exec()

    const response: S = {
      name: shop.name,
        description: shop.description,
        shopOwner: shop.shopOwner,
        id: shop.id,
        imageName: shop.image
    }

    const res: ShopResponse = {
      shop: response
    }

    return res

    
  }

  // get shop by owner id
  async getShopsByOwner (request: GetShopsByOwnerRequest): Promise<GetShopsByOwnerResponse>{

    const db = await this.shopModel.find({shopOwner: request.id}).exec()

    const res: S[] = db.map((shop)=>{
      let response: S = {
        id: shop.id,
        name: shop.name,
        description: shop.description,
        imageName: shop.image,
        shopOwner: shop.shopOwner
      } 

      return response
    })
    

    const result: GetShopsByOwnerResponse = {
      shops: res
    }

    return result
    
  }
   
}
