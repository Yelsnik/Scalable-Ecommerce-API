import { IsNumber, IsString } from 'class-validator';

export class AddtoCartBodyDTO {
  @IsNumber()
  quantity: number;
}

export class AddToCartParamsDTO {
  @IsString()
  id: string;
}
