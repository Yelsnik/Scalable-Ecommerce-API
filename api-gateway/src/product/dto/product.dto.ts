import {
  IsBoolean,
  IsLowercase,
  IsNumber,
  IsOptional,
  IsString,
} from 'class-validator';

export class addProductDTO {
  @IsString()
  category: string;

  @IsString()
  productname: string;

  @IsString()
  description: string;

  @IsString()
  brand: string;

  @IsNumber()
  countinstock: number;

  @IsNumber()
  price: number;

  @IsLowercase()
  @IsString()
  currency: string;

  @IsOptional()
  @IsNumber()
  rating?: number;

  @IsOptional()
  @IsBoolean()
  isfeatured?: boolean;
}

export class addProductParamsDTO {
  @IsString()
  id: string;
}

export class updateProductParamsDTO {
  @IsString()
  id: string;
}

export class updateProductBodyDTO {
  @IsOptional()
  @IsString()
  category?: string;

  @IsOptional()
  @IsString()
  productName?: string;

  @IsOptional()
  @IsString()
  description?: string;

  @IsOptional()
  @IsString()
  brand?: string;

  image?: Express.Multer.File;

  @IsOptional()
  @IsNumber()
  countInStock?: number;

  @IsOptional()
  @IsNumber()
  price?: number;

  @IsOptional()
  @IsString()
  currency?: string;

  @IsOptional()
  @IsNumber()
  rating?: number;

  @IsOptional()
  @IsBoolean()
  isFeatured?: boolean;
}
