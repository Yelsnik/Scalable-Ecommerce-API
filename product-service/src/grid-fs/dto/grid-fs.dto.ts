import { IsNumber, IsString } from "class-validator"

export class GridFsDto {
    
  buffer: Buffer

  @IsString()
  originalname: string

  @IsString()
  fieldname: string

  @IsNumber()
  size: number

  @IsString()
  mimetype: string

  @IsString()
  encoding: string
}
