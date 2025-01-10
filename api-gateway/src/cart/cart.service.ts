import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import {
  AddtoCartRequest,
  CartItemResponse,
  CartItemServiceClient,
  CartTxResult,
  GetCartItemByIDRequest,
  GetCartItemsByCartRequest,
  GetCartItemsByCartResponse,
  RemoveCartTxRequest,
  RemoveCartTxResult,
  UpdateCartTxRequest,
} from 'pb/cart-service/cart_item_service';
import { CartServiceClient } from 'pb/cart-service/cart_service';
import { Observable } from 'rxjs';
import {
  AddtoCartBodyDTO,
  AddToCartParamsDTO,
  GetCartItemByIDParamsDTO,
  GetCartItemsByCartParamsDTO,
  RemoveCartItemDTO,
  UpdateCartItemBodyDTO,
  UpdateCartItemDTO,
} from './dto/cart-item.dto';

@Injectable()
export class CartService implements OnModuleInit {
  cartItemService: CartItemServiceClient;
  cartService: CartServiceClient;

  constructor(
    @Inject('CART_ITEM_SERVICE') private cartItemClient: ClientGrpc,
    @Inject('CART_SERVICE') private cartClient: ClientGrpc,
  ) {}

  onModuleInit() {
    this.cartItemService =
      this.cartItemClient.getService<CartItemServiceClient>('CartItemService');
    this.cartService =
      this.cartClient.getService<CartServiceClient>('CartService');
  }

  addToCart(
    req: any,
    params: AddToCartParamsDTO,
    body: AddtoCartBodyDTO,
  ): Observable<CartTxResult> {
    console.log(req.user, req.user.user.id);

    const userId = req.user.user.id;

    const request: AddtoCartRequest = {
      productId: params.id,
      quantity: body.quantity,
      userId: userId,
    };

    return this.cartItemService.addToCartTx(request);
  }

  getCartItemsByCart(
    params: GetCartItemsByCartParamsDTO,
  ): Observable<GetCartItemsByCartResponse> {
    const request: GetCartItemsByCartRequest = {
      id: params.id,
    };

    return this.cartItemService.getCartItemsByCart(request);
  }

  getCartItemById(
    params: GetCartItemByIDParamsDTO,
  ): Observable<CartItemResponse> {
    const request: GetCartItemByIDRequest = {
      id: params.id,
    };

    return this.cartItemService.getCartItem(request);
  }

  updateCart(
    params: UpdateCartItemDTO,
    body: UpdateCartItemBodyDTO,
  ): Observable<CartTxResult> {
    const request: UpdateCartTxRequest = {
      id: params.id,
      quantity: body.quantity,
    };

    return this.cartItemService.updateCartTx(request);
  }

  remove(params: RemoveCartItemDTO): Observable<RemoveCartTxResult> {
    const request: RemoveCartTxRequest = {
      id: params.id,
    };

    return this.cartItemService.removeCartTx(request);
  }
}
