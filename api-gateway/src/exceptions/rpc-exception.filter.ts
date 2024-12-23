import { ArgumentsHost, Catch, ExceptionFilter } from '@nestjs/common';
import { RpcException } from '@nestjs/microservices';
import { Response } from 'express';

@Catch(RpcException)
export class RpcExceptionFilter implements ExceptionFilter {
  catch(exception: RpcException, host: ArgumentsHost) {
    const error: any = exception.getError();
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const request = ctx.getRequest<Request>();
    
    response
      .status(error.statusCode)
      .json(
        {
            statusCode: error.statusCode,
            message: exception.message,
            stack: exception.stack,
            timestamp: new Date().toISOString(),
            path: request.url,
          }
      );
  }
}