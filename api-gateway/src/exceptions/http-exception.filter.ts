import {
    ExceptionFilter,
    Catch,
    ArgumentsHost,
    HttpException,
  } from '@nestjs/common';
  import { Request, Response } from 'express';
  
  @Catch(HttpException)
  export class HttpExceptionFilter implements ExceptionFilter {
    catch(exception: HttpException, host: ArgumentsHost) {
      const ctx = host.switchToHttp();
      const response = ctx.getResponse<Response>();
      const request = ctx.getRequest<Request>();
      const status = exception.getStatus();
  
      if (process.env.NODE_ENV === 'development'){
        response.status(status).json({
          statusCode: status,
          message: exception.message,
          stack: exception.stack,
         // cause: exception.cause,
          timestamp: new Date().toISOString(),
          // path: request.url,
        });
      }

      response.status(status).json({
        statusCode: status,
          message: exception.message,
          timestamp: new Date().toISOString(),
          // path: request.url,
      })
    }
  }
  
  /*
  @Catch(UnauthorizedException)
  export class UnauthorizedExceptionFilter implements ExceptionFilter {
    catch(exception: UnauthorizedException, host: ArgumentsHost): void {
      const ctx = host.switchToHttp();
      const response = ctx.getResponse<Response>();
      let status = exception.getStatus();
  
      response.status(status).json({
        status: status,
        timestamp: new Date().toISOString(),
        message: exception.message,
      });
    }
  }
  */
  