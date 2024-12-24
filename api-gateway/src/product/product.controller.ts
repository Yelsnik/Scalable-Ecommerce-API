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
  GetProductsByShopRequest,
  UpdateProductRequest,
} from 'pb/product_service';
import {
  addProductDTO,
  addProductParamsDTO,
  updateProductBodyDTO,
  updateProductParamsDTO,
} from './dto/product.dto';
import { lastValueFrom } from 'rxjs';

@Controller('product')
@UseFilters(RpcToHttpExceptionFilter, HttpExceptionFilter)
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

    const res = this.productService.addProduct(request);

    const product = await lastValueFrom(res);

    response.status(201).json({
      message: 'success',
      data: product,
    });
  }

  @Get(':id')
  async getProductById(@Param() params: any, @Res() response: any) {
    const request: GetProductByIdRequest = {
      id: params.id,
    };

    const res = this.productService.getProductByID(request);

    const product = await lastValueFrom(res);

    response.status(200).json({
      message: 'success',
      data: product,
    });
  }

  @Get('shop/:id')
  async getProductsByShop(@Param() params: any, @Res() response: any) {
    const request: GetProductsByShopRequest = {
      id: params.id,
    };

    const res = this.productService.getProductsByShop(request);

    const products = await lastValueFrom(res);

    response.status(200).json({
      message: 'success',
      data: products,
    });
  }

  @Patch(':id')
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
