import { IsNumber, IsString } from "class-validator"

export class GridFsDto {
  id: any
    
  buffer: Buffer

  @IsString()
  originalname: string

  @IsString()
  filename: string

  @IsNumber()
  size: number

  @IsNumber()
  chunksize: number

  @IsString()
  mimetype: string

  @IsString()
  bucketname: string

  @IsString()
  md5: string

  @IsString()
  contentType: string

  metadata: any
}
