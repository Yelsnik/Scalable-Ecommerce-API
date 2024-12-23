import {
  Body,
  Controller,
  Delete,
  Get,
  HttpException,
  InternalServerErrorException,
  Param,
  Patch,
  Post,
  Res,
  UploadedFile,
  UseFilters,
  UseInterceptors,
} from '@nestjs/common';
import { FileInterceptor } from '@nestjs/platform-express';
import { File } from 'pb/file';
import {
  AddShopFileDTO,
  deleteShopParam,
  getShopByIDParam,
  shopBodyDTO,
  updateShopBodyDTO,
  updateShopByIDParam,
} from './dto/shop.dto';
import { ProductService } from './product.service';
import {
  CreateShopRequest,
  DeleteShopRequest,
  GetShopByIdRequest,
  GetShopsByOwnerRequest,
  ShopServiceClient,
  UpdateShopRequest,
} from 'pb/shop_service';
import { Error } from 'mongoose';
import { lastValueFrom } from 'rxjs';
import { RpcExceptionFilter } from 'src/exceptions/rpc-exception.filter';
import { RpcException } from '@nestjs/microservices';
import { HttpExceptionFilter } from 'src/exceptions/http-exception.filter';
import { GrpcToHttpInterceptor } from 'nestjs-grpc-exceptions';
import { GRPCToHTTPExceptions } from 'src/helpers/errors';
import { MongoExceptionFilter } from 'src/helpers/mongo-exception.filter';

@Controller('product')
@UseFilters(HttpExceptionFilter)
export class ProductController {
  constructor(private readonly productService: ProductService) {}

  @Post('add-shop')
  @UseInterceptors(FileInterceptor('image'))
  async createShop(
    @UploadedFile() file: AddShopFileDTO,
    @Body() data: shopBodyDTO,
    @Res() response: any,
  ) {
    try {
      const fileData: File = {
        buffer: file.buffer,
        originalname: file.originalname,
        fieldname: file.fieldname,
        size: file.size,
        mimetype: file.mimetype,
        encoding: file.encoding,
      };
      const req: CreateShopRequest = {
        name: data.shopname,
        description: data.description,
        shopOwner: data.shopOwner,
        image: fileData,
      };

      const result = this.productService.createShop(req);

      const shop = await lastValueFrom(result);

      return response.status(201).json({
        message: 'success',
        data: shop,
      });
    } catch (err) {
      new GRPCToHTTPExceptions(err);
    }
  }

  @Get('shop/:id')
  async getShopByID(@Param() params: getShopByIDParam, @Res() response: any) {
    try {
      const req: GetShopByIdRequest = {
        id: params.id,
      };
      const shop = this.productService.getShopById(req);

      const result = await lastValueFrom(shop);

      return response.status(201).json({
        message: 'success',
        data: result,
      });
    } catch (err) {
      new GRPCToHTTPExceptions(err);
    }
  }

  @Get('shops/:ownerid')
  async getShopsByOwner(@Param() params: any, @Res() response: any) {
    try {
      const request: GetShopsByOwnerRequest = {
        id: params.id,
      };
      const res = this.productService.getShopsByOwner(request);

      const shops = await lastValueFrom(res);

      return response.status(201).json({
        message: 'success',
        data: shops,
      });
    } catch (err) {
      new GRPCToHTTPExceptions(err);
    }
  }

  @Patch('shops/:id')
  async updateShop(
    @Param() params: updateShopByIDParam,
    @Body() body: updateShopBodyDTO,
    @Res() response: any,
  ) {
    try {
      let request: UpdateShopRequest;
      if (body.name !== undefined) {
        request = {
          id: params.id,
          name: body.name,
        };
      }

      if (body.description !== undefined) {
        request = {
          id: params.id,
          description: body.description,
        };
      }

      const res = this.productService.updateShop(request);
      const shop = await lastValueFrom(res);

      return response.status(201).json({
        message: 'success',
        data: shop,
      });
    } catch (err) {
      new GRPCToHTTPExceptions(err);
    }
  }

  @Delete('shop/:id')
  async deleteShop(@Param() params: deleteShopParam, @Res() response: any) {
    try {
      const request: DeleteShopRequest = {
        id: params.id,
      };

      const res = this.productService.deleteShop(request);

      await lastValueFrom(res);

      return response.status(201).json({
        message: 'successfully deleted shop',
      });
    } catch (err) {
      new GRPCToHTTPExceptions(err);
    }
  }
}
