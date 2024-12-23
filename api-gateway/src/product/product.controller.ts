import { Controller, UseFilters } from '@nestjs/common';

import { ProductService } from './product.service';

import { HttpExceptionFilter } from 'src/exceptions/http-exception.filter';

@Controller('product')
@UseFilters(HttpExceptionFilter)
export class ProductController {
  constructor(private readonly productService: ProductService) {}
}
