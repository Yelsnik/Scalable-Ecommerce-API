import { Catch, ArgumentsHost, RpcExceptionFilter } from '@nestjs/common';
import { RpcException } from '@nestjs/microservices';
import { Error as MongooseError } from 'mongoose';
import { GrpcInvalidArgumentException } from 'nestjs-grpc-exceptions';
import { Observable, throwError } from 'rxjs';

@Catch(MongooseError.CastError)
export class MongoExceptionFilter
  implements RpcExceptionFilter<MongooseError.CastError>
{
  catch(
    exception: MongooseError.CastError,
    host: ArgumentsHost,
  ): Observable<any> {
    console.log(exception.name);

    if (exception.name === 'CastError') {
      return throwError(
        () => new GrpcInvalidArgumentException('Invalid argument'),
      );
    }
  }
}
