import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.connectMicroservice<MicroserviceOptions>({
    transport: Transport.GRPC,
    options: {
      package: "pb",
      protoPath: join(__dirname,  "../proto/shop_service.proto"),
      url: "0.0.0.0:50051"
    }
  })

  await app.startAllMicroservices()
  await app.listen(process.env.PORT ?? 4000);
}
bootstrap();
