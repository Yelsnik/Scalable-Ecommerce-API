import { IsString } from "class-validator"

export class ShopDto {
  @IsString()
    name: string

  @IsString()
    description: string

  @IsString()
    shopOwner: string

    updatedAt?: Date
}
