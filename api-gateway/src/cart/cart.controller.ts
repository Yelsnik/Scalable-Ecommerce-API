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
  Get,
  Patch,
} from '@nestjs/common';
import { CartService } from './cart.service';
import {
  AddtoCartBodyDTO,
  AddToCartParamsDTO,
  GetCartItemByIDParamsDTO,
  GetCartItemsByCartParamsDTO,
  RemoveCartItemDTO,
  UpdateCartItemBodyDTO,
  UpdateCartItemDTO,
} from './dto/cart-item.dto';
import { lastValueFrom } from 'rxjs';
import { AuthGuard } from 'src/auth/auth.guard';
import {
  RpcToHttpExceptionFilter,
  HttpExceptionFilter,
} from 'src/exceptions/http-exception.filter';

@Controller({ path: 'cart', version: '1' })
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

  @Get('items/:id')
  async getCartItemsByCart(
    @Param() params: GetCartItemsByCartParamsDTO,
    @Res() response: any,
  ) {
    const res = this.cartService.getCartItemsByCart(params);

    const carts = await lastValueFrom(res);

    response.status(200).json({
      message: 'success',
      data: carts,
    });
  }

  @Get(':id')
  async getCartItemsByID(
    @Param() params: GetCartItemByIDParamsDTO,
    @Res() response: any,
  ) {
    const res = this.cartService.getCartItemById(params);

    const cart = await lastValueFrom(res);

    response.status(200).json({
      message: 'success',
      data: cart,
    });
  }

  @Patch(':id')
  async updateCart(
    @Param() params: UpdateCartItemDTO,
    @Body() body: UpdateCartItemBodyDTO,
    @Res() response: any,
  ) {
    const res = this.cartService.updateCart(params, body);

    const cart = await lastValueFrom(res);

    response.status(200).json({
      message: 'success',
      data: cart,
    });
  }

  @Delete(':id')
  async remove(@Param() param: RemoveCartItemDTO, @Res() response: any) {
    const res = this.cartService.remove(param);

    await lastValueFrom(res);

    response.status(200).json({
      message: 'successfully deleted!',
    });
  }
}
