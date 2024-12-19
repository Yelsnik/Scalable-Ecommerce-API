import { Module } from '@nestjs/common';
import { ShopController } from './shop.controller';
import { ShopService } from './shop.service';
import { Shop, ShopSchema } from './shop.schema';
import { MongooseModule } from '@nestjs/mongoose';
import { ConfigModule } from '@nestjs/config';
import { GridFsModule } from 'src/grid-fs/grid-fs.module';

@Module({
  imports: [
    MongooseModule.forFeature([{ name: Shop.name, schema: ShopSchema }]),
    ConfigModule.forRoot({ envFilePath: 'app.env', isGlobal: true }),
    GridFsModule
  ],
  controllers: [ShopController],
  providers: [ShopService]
})
export class ShopModule {}
