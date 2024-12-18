import { IsString } from "class-validator"

export class ShopDto {
  @IsString()
  private readonly name: string

  @IsString()
  private readonly description: string

  @IsString()
  private readonly shopOwner: string

  private readonly updatedAt?: Date
}
