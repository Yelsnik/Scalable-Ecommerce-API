import { Module } from '@nestjs/common';
import { ProductService } from './product.service';
import { ShopController } from './shop/shop.controller';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { dirname, join } from 'path';
import { PRODUCT_PACKAGE_NAME as Shop } from 'pb/shop_service';
import { PRODUCT_PACKAGE_NAME as Product } from 'pb/product_service';
import { ProductController } from './product.controller';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'PRODUCT_SERVICE',
        transport: Transport.GRPC,
        options: {
          package: Product,
          protoPath: join(
            __dirname,
            'proto/product-service/product_service.proto',
          ),
          url: '0.0.0.0:50051',
        },
      },
      {
        name: 'SHOP_SERVICE',
        transport: Transport.GRPC,
        options: {
          package: Shop,
          protoPath: join(
            __dirname,
            'proto/product-service/shop_service.proto',
          ),
          url: '0.0.0.0:50051',
        },
      },
    ]),
  ],
  providers: [ProductService],
  controllers: [ShopController, ProductController],
})
export class ProductModule {}
