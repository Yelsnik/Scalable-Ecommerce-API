import {
  Body,
  Controller,
  Delete,
  Get,
  HttpStatus,
  Param,
  ParseFilePipeBuilder,
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
  deleteShopParam,
  getShopByIDParam,
  shopBodyDTO,
  updateShopBodyDTO,
  updateShopByIDParam,
} from '../dto/shop.dto';
import { ProductService } from '../product.service';
import {
  CreateShopRequest,
  DeleteShopRequest,
  GetShopByIdRequest,
  GetShopsByOwnerRequest,
  UpdateShopRequest,
} from 'pb/shop_service';
import { lastValueFrom } from 'rxjs';
import { HttpExceptionFilter } from 'src/exceptions/http-exception.filter';
import { GRPCToHTTPExceptions } from 'src/helpers/errors';

@Controller('shop')
@UseFilters(HttpExceptionFilter)
export class ShopController {
  constructor(private readonly productService: ProductService) {}

  /**
   * Endpoints for Modifying shops
   * Create shop
   * Get shop by id
   * Get shops by owner
   * Update shop by id
   * Delete shop
   */
  @Post('add-shop')
  @UseInterceptors(FileInterceptor('image'))
  async createShop(
    @UploadedFile(
      new ParseFilePipeBuilder()
        .addFileTypeValidator({
          fileType: 'jpeg',
        })
        .build({
          errorHttpStatusCode: HttpStatus.UNPROCESSABLE_ENTITY,
        }),
    )
    file: Express.Multer.File,
    @Body() data: shopBodyDTO,
    @Res() response: any,
  ) {
    console.log(file);
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
  }

  @Get(':id')
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

  @Get(':ownerid')
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

  @Patch(':id')
  async updateShop(
    @Param() params: updateShopByIDParam,
    @Body() body: updateShopBodyDTO,
    @Res() response: any,
  ) {
    try {
      const request: UpdateShopRequest = {
        id: params.id,
        name: body.name,
        description: body.description,
      };

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

  @Delete(':id')
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
