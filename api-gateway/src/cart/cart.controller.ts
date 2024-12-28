import {
  Controller,
  Post,
  Body,
  Param,
  Delete,
  Res,
  UseFilters,
  UseGuards,
  Req,
} from '@nestjs/common';
import { CartService } from './cart.service';
import { AddtoCartBodyDTO, AddToCartParamsDTO } from './dto/cart-item.dto';
import { lastValueFrom } from 'rxjs';
import { AuthGuard } from 'src/auth/auth.guard';
import {
  RpcToHttpExceptionFilter,
  HttpExceptionFilter,
} from 'src/exceptions/http-exception.filter';
import { Metadata, ServerUnaryCall } from '@grpc/grpc-js';

@Controller('cart')
@UseGuards(AuthGuard)
@UseFilters(RpcToHttpExceptionFilter, HttpExceptionFilter)
export class CartController {
  constructor(private readonly cartService: CartService) {}

  @Post('add-to-cart/:id')
  async addToCart(
    @Param() params: AddToCartParamsDTO,
    @Body() body: AddtoCartBodyDTO,
    @Res() response: any,
    @Req() req: any,
  ) {
    const res = this.cartService.addToCart(req, params, body);

    const cart = await lastValueFrom(res);

    response.status(201).json({
      message: 'success',
      data: cart,
    });
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.cartService.remove(+id);
  }
}
