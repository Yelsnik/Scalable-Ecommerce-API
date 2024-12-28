import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ProductModule } from './product/product.module';
import { AuthModule } from './auth/auth.module';
import { CartModule } from './cart/cart.module';

@Module({
  imports: [ProductModule, AuthModule, CartModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
