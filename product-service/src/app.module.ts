import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ProductModule } from './product/product.module';
import { MongooseModule } from '@nestjs/mongoose';
import { ConfigModule } from '@nestjs/config';
import { ShopModule } from './shop/shop.module';

@Module({
  imports: [
    ProductModule,
    ConfigModule.forRoot({ envFilePath: 'app.env', isGlobal: true }),
    MongooseModule.forRoot(process.env.DATABASE),
    ShopModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
