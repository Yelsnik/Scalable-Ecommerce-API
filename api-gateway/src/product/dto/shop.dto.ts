import { IsNumber, IsOptional, IsString } from 'class-validator';

export class AddShopFileDTO {
  buffer: Buffer;

  @IsString()
  originalname: string;

  @IsString()
  fieldname: string;

  @IsNumber()
  size: number;

  @IsString()
  mimetype: string;

  @IsString()
  encoding: string;
}

export class shopBodyDTO {
  @IsString()
  shopname: string;

  @IsString()
  description: string;

  @IsString()
  shopOwner: string;
}

export class getShopByIDParam {
  @IsString()
  id: string;
}

export class getShopByOwnerParam {
  @IsString()
  id: string;
}

export class updateShopByIDParam {
  @IsString()
  id: string;
}

export class updateShopBodyDTO {
  @IsOptional()
  @IsString()
  name: string;

  @IsOptional()
  @IsString()
  description: string;
}

export class deleteShopParam {
  @IsString()
  id: string;
}
