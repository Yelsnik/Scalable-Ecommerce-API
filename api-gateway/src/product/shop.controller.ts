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
  Query,
  Req,
  Res,
  UploadedFile,
  UseFilters,
  UseGuards,
  UseInterceptors,
} from '@nestjs/common';
import { FileInterceptor } from '@nestjs/platform-express';
import { File } from 'pb/product/file';
import {
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
  UpdateShopRequest,
} from 'pb/product/shop_service';
import { lastValueFrom } from 'rxjs';
import {
  HttpExceptionFilter,
  RpcToHttpExceptionFilter,
} from 'src/exceptions/http-exception.filter';
import { AuthGuard } from 'src/auth/auth.guard';

@Controller({ path: 'shop', version: '1' })
@UseGuards(AuthGuard)
@UseFilters(RpcToHttpExceptionFilter, HttpExceptionFilter)
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
    @Req() request: any,
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
      shopOwner: request.user.user.id,
      image: fileData,
    };

    const shop = await this.productService.createShop(req);

    return response.status(201).json({
      message: 'success',
      data: shop,
    });
  }

  @Get(':id')
  async getShopByID(@Param() params: getShopByIDParam, @Res() response: any) {
    const req: GetShopByIdRequest = {
      id: params.id,
    };
    const shop = await this.productService.getShopById(req);

    return response.status(200).json({
      message: 'success',
      data: shop,
    });
  }

  @Get('user/:id')
  async getShopsByOwner(
    @Param() params: any,
    @Query() query: any,
    @Res() response: any,
  ) {
    const queryString = JSON.stringify(query);
    console.log(query, queryString);

    const request: GetShopsByOwnerRequest = {
      id: params.id,
      queryString: queryString,
    };

    const shops = await this.productService.getShopsByOwner(request);

    return response.status(200).json({
      message: 'success',
      data: shops,
    });
  }

  @Patch(':id')
  async updateShop(
    @Param() params: updateShopByIDParam,
    @Body() body: updateShopBodyDTO,
    @Res() response: any,
  ) {
    const request: UpdateShopRequest = {
      id: params.id,
      name: body.name,
      description: body.description,
    };

    const shop = await this.productService.updateShop(request);

    return response.status(200).json({
      message: 'success',
      data: shop,
    });
  }

  @Delete(':id')
  async deleteShop(@Param() params: deleteShopParam, @Res() response: any) {
    const request: DeleteShopRequest = {
      id: params.id,
    };

    await this.productService.deleteShop(request);

    return response.status(200).json({
      message: 'successfully deleted shop',
    });
  }
}
