import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { ReflectionService } from '@grpc/reflection';

declare const module: any

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  
  app.connectMicroservice<MicroserviceOptions>({
    transport: Transport.GRPC,
    options: {
      onLoadPackageDefinition: (pkg, server) => {
        new ReflectionService(pkg).addToServer(server);
      },
      package: "product",
      protoPath: [
        join(__dirname,  "proto/product_service.proto"), 
        join(__dirname,  "proto/shop_service.proto"),
        join(__dirname,  "proto/file.proto")
      ],
      url: "0.0.0.0:50051"
    }
  })

  await app.startAllMicroservices()
  await app.listen(process.env.PORT ?? 4000);

  if (module.hot) {
    module.hot.accept();
    module.hot.dispose(() => app.close());
  }
}
bootstrap();
