import { Module } from '@nestjs/common';
import { ProductController } from './product.controller';
import { ProductService } from './product.service';
import { Product, ProductSchema } from './product.schema';
import { MongooseModule } from '@nestjs/mongoose';
import { GridFsModule } from 'src/grid-fs/grid-fs.module';
import { ConfigModule } from '@nestjs/config';
import { GridFsService } from 'src/grid-fs/grid-fs.service';
import { Shop, ShopSchema } from 'src/shop/shop.schema';

@Module({
  imports: [
    MongooseModule.forFeature([{ name: Product.name, schema: ProductSchema }]),
    MongooseModule.forFeature([{ name: Shop.name, schema: ShopSchema }]),
    ConfigModule.forRoot({ envFilePath: 'app.env', isGlobal: true }),
    GridFsModule
  ],
  controllers: [ProductController],
  providers: [ProductService, GridFsService]
})
export class ProductModule {}
