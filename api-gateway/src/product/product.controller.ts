import {
  Body,
  Controller,
  Delete,
  Get,
  HttpStatus,
  InternalServerErrorException,
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
import { ProductService } from './product.service';
import {
  HttpExceptionFilter,
  RpcToHttpExceptionFilter,
} from 'src/exceptions/http-exception.filter';
import { FileInterceptor } from '@nestjs/platform-express';
import {
  CreateProductRequest,
  DeleteProductRequest,
  GetProductByIdRequest,
  ProductsByShopRequest,
  UpdateProductRequest,
} from 'pb/product/product_service';
import {
  addProductDTO,
  addProductParamsDTO,
  updateProductBodyDTO,
  updateProductParamsDTO,
} from './dto/product.dto';
import { lastValueFrom } from 'rxjs';
import { AuthGuard } from 'src/auth/auth.guard';
//import { Any } from 'pb/google/protobuf/any';
//import { AnyFieldType } from 'google-protobuf';

@Controller({ path: 'product', version: '1' })
@UseFilters(RpcToHttpExceptionFilter, HttpExceptionFilter)
@UseGuards(AuthGuard)
export class ProductController {
  constructor(private readonly productService: ProductService) {}

  @Post(':id')
  @UseInterceptors(FileInterceptor('image'))
  async addProduct(
    @Param() params: addProductParamsDTO,
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
    @Body() body: addProductDTO,
    @Res() response: any,
  ) {
    const request: CreateProductRequest = {
      category: body.category,
      productName: body.productname,
      description: body.description,
      brand: body.brand,
      image: {
        buffer: file.buffer,
        originalname: file.originalname,
        fieldname: file.fieldname,
        size: file.size,
        mimetype: file.mimetype,
        encoding: file.encoding,
      },
      countInStock: body.countinstock,
      price: body.price,
      currency: body.currency,
      shop: params.id,
      rating: body.rating,
      isFeatured: body.isfeatured,
    };

    const product = await this.productService.addProduct(request);

    response.status(201).json({
      message: 'success',
      data: product,
    });
  }

  @Get(':id')
  @UseFilters(RpcToHttpExceptionFilter, HttpExceptionFilter)
  async getProductById(@Param() params: any, @Res() response: any) {
    const request: GetProductByIdRequest = {
      id: params.id,
    };

    const product = await this.productService.getProductByID(request);

    response.status(200).json({
      message: 'success',
      data: product,
    });
  }

  @Get('shop/:id')
  //@UseFilters(RpcToHttpExceptionFilter, HttpExceptionFilter)
  async getProductsByShop(
    @Param() params: any,
    @Query() query: any,
    @Res() response: any,
  ) {
    const queryString = JSON.stringify(query);
    console.log(query, queryString);

    const request: ProductsByShopRequest = {
      id: params.id,
      queryString: queryString,
    };

    const res = this.productService.getProductsByShop(request);

    const products = await lastValueFrom(res);

    response.status(200).json({
      message: 'success',
      data: products,
    });
  }

  @Patch(':id')
  @UseFilters(RpcToHttpExceptionFilter, HttpExceptionFilter)
  async updateProduct(
    @Param() params: updateProductParamsDTO,
    @Body() body: updateProductBodyDTO,
    @Res() response: any,
  ) {
    const request: UpdateProductRequest = {
      id: params.id,
      category: body.category,
      productName: body.productName,
      description: body.description,
      brand: body.brand,
      countInStock: body.countInStock,
      price: body.price,
      currency: body.currency,
      rating: body.countInStock,
      isFeatured: body.isFeatured,
    };

    const res = this.productService.updateProduct(request);

    const product = await lastValueFrom(res);

    response.status(200).json({
      message: 'success',
      data: product,
    });
  }

  @Delete()
  @UseFilters(RpcToHttpExceptionFilter, HttpExceptionFilter)
  async deleteProduct(@Param() params: any, @Res() response: any) {
    const request: DeleteProductRequest = {
      id: params.id,
    };

    const res = this.productService.deleteProduct(request);

    await lastValueFrom(res);

    response.status(200).json({
      message: 'success',
    });
  }
}
