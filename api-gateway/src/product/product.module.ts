import { Module } from '@nestjs/common';
import { ProductService } from './product.service';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { PRODUCT_PACKAGE_NAME as Shop } from 'pb/product/shop_service';
import { PRODUCT_PACKAGE_NAME as Product } from 'pb/product/product_service';
import { ProductController } from './product.controller';
import { ShopController } from './shop.controller';
import { AuthModule } from 'src/auth/auth.module';
import { CacheModule } from '@nestjs/cache-manager';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'PRODUCT_SERVICE',
        transport: Transport.GRPC,
        options: {
          package: Product,
          protoPath: join(__dirname, 'proto/product_service.proto'),
          url: '0.0.0.0:50051',
        },
      },
      {
        name: 'SHOP_SERVICE',
        transport: Transport.GRPC,
        options: {
          package: Shop,
          protoPath: join(__dirname, 'proto/shop_service.proto'),
          url: '0.0.0.0:50051',
        },
      },
    ]),
    CacheModule.register(),
    AuthModule,
  ],
  providers: [ProductService],
  controllers: [ProductController, ShopController],
})
export class ProductModule {}
