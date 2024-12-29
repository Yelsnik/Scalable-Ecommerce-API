import { IsNumber, IsString } from 'class-validator';

export class AddtoCartBodyDTO {
  @IsNumber()
  quantity: number;
}

export class AddToCartParamsDTO {
  @IsString()
  id: string;
}

export class GetCartItemsByCartParamsDTO {
  @IsString()
  id: string;
}

export class GetCartItemByIDParamsDTO {
  @IsString()
  id: string;
}

export class UpdateCartItemDTO {
  @IsString()
  id: string;
}

export class UpdateCartItemBodyDTO {
  @IsNumber()
  quantity: number;
}

export class RemoveCartItemDTO {
  @IsString()
  id: string;
}
