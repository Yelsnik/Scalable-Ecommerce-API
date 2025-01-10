import {
  ExceptionFilter,
  Catch,
  ArgumentsHost,
  HttpException,
  HttpStatus,
} from '@nestjs/common';
import { RpcException } from '@nestjs/microservices';
import { Request, Response } from 'express';
import { RpcExceptionFilter } from './rpc-exception.filter';
import { GrpcToHttpInterceptor } from 'nestjs-grpc-exceptions';
import { GRPCToHTTPExceptions } from 'src/helpers/errors';

@Catch()
export class RpcToHttpExceptionFilter implements ExceptionFilter {
  catch(exception: any, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    // const request = ctx.getRequest<Request>();

    return new GRPCToHTTPExceptions(exception, response);
  }
}

@Catch(HttpException)
export class HttpExceptionFilter implements ExceptionFilter {
  catch(exception: HttpException, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    let status = exception.getStatus();

    return response.status(status).json({
      status: status,
      timestamp: new Date().toISOString(),
      message: exception.message,
    });
  }
}
