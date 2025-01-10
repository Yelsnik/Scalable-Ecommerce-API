import { Module } from '@nestjs/common';
import { CartService } from './cart.service';
import { CartController } from './cart.controller';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { AuthModule } from 'src/auth/auth.module';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'CART_ITEM_SERVICE',
        transport: Transport.GRPC,
        options: {
          package: 'cart',
          protoPath: join(__dirname, 'proto/cart_item_service.proto'),
          url: '0.0.0.0:7070',
        },
      },
      {
        name: 'CART_SERVICE',
        transport: Transport.GRPC,
        options: {
          package: 'cart',
          protoPath: join(__dirname, 'proto/cart_service.proto'),
          url: '0.0.0.0:7070',
        },
      },
    ]),
    AuthModule,
  ],
  controllers: [CartController],
  providers: [CartService],
})
export class CartModule {}
