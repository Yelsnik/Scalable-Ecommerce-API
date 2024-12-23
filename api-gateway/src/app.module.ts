import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ProductModule } from './product/product.module';
import { ShopController } from './shop/shop.controller';

@Module({
  imports: [ProductModule],
  controllers: [AppController, ShopController],
  providers: [AppService],
})
export class AppModule {}
