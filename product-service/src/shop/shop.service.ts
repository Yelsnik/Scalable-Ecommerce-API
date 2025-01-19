import { Injectable } from '@nestjs/common';
import { Shop } from './shop.schema';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import {
  CreateShopRequest,
  Shop as S,
  GetShopByIdRequest,
  GetShopsByOwnerRequest,
  GetShopsByOwnerResponse,
  ShopResponse,
  UpdateShopRequest,
  DeleteShopRequest,
  Empty,
} from 'pb/shop_service';
import { GridFsService } from 'src/grid-fs/grid-fs.service';
import {
  GrpcAbortedException,
  GrpcNotFoundException,
} from 'nestjs-grpc-exceptions';
import { ApiFeatures } from 'src/helpers/apiFeatures';

@Injectable()
export class ShopService {
  constructor(
    @InjectModel(Shop.name) private shopModel: Model<Shop>,
    private gridFs: GridFsService,
  ) {}

  // create shop
  async createShop(request: CreateShopRequest): Promise<ShopResponse> {
    const buffer = Buffer.from(request.image.buffer);

    const file = this.gridFs.loadImage(request.image, buffer);

    const id = await this.gridFs.uploadFile(file, 'shop-image');
    console.log(id);

    const shop = await this.shopModel.create({
      name: request.name,
      description: request.description,
      imageID: id,
      image: request.image.originalname,
      shopOwner: request.shopOwner,
    });

    if (!shop) {
      throw new GrpcAbortedException('failed to create shop');
    }

    const res: ShopResponse = {
      shop: {
        name: shop.name,
        description: shop.description,
        shopOwner: shop.shopOwner,
        id: shop.id,
        imageName: shop.image,
      },
    };

    return res;
  }

  // get shop by id
  async getShopById(request: GetShopByIdRequest): Promise<ShopResponse> {
    const shop = await this.shopModel.findById(request.id).exec();
    // check if shop exists
    if (!shop || shop === undefined) {
      throw new GrpcNotFoundException('Shop not found');
    }

    const response: S = {
      name: shop.name,
      description: shop.description,
      shopOwner: shop.shopOwner,
      id: shop.id,
      imageName: shop.image,
    };

    const res: ShopResponse = {
      shop: response,
    };
    return res;
  }

  // get shops by owner id
  async getShopsByOwner(
    request: GetShopsByOwnerRequest,
  ): Promise<GetShopsByOwnerResponse> {
    let queryString: string;
    let features: ApiFeatures;

    if (request.queryString) {
      queryString = JSON.parse(request.queryString);
      features = new ApiFeatures(
        this.shopModel.find({ shopOwner: request.id }),
        queryString,
      )
        .filter()
        .sort()
        .limit()
        .pagination();
    } else {
      features = new ApiFeatures(
        this.shopModel.find({ shopOwner: request.id }),
        {},
      )
        .filter()
        .sort()
        .limit()
        .pagination();
    }

    const shops = await features.query;
    console.log(typeof shops, shops);

    if (!shops) {
      throw new GrpcNotFoundException('no shops belonging to you exist');
    }

    const res: S[] = shops.map((shop: any) => {
      let response: S = {
        id: shop.id,
        name: shop.name,
        description: shop.description,
        imageName: shop.image,
        shopOwner: shop.shopOwner,
      };

      return response;
    });

    const result: GetShopsByOwnerResponse = {
      shops: res,
    };

    return result;
  }

  // update shop
  async updateShop(request: UpdateShopRequest): Promise<ShopResponse> {
    const body = {
      name: request.name,
      description: request.description,
    };

    const shop = await this.shopModel
      .findByIdAndUpdate(request.id, body)
      .setOptions({ overwrite: true, new: true });

    if (!shop) {
      throw new GrpcNotFoundException(
        'shop does not exist or has been deleted',
      );
    }

    const res: S = {
      id: shop.id,
      name: shop.name,
      description: shop.description,
      imageName: shop.image,
      shopOwner: shop.shopOwner,
    };

    const response: ShopResponse = {
      shop: res,
    };

    return response;
  }

  async deleteShop(request: DeleteShopRequest): Promise<Empty> {
    return await this.shopModel.findByIdAndDelete(request.id);
  }
}
