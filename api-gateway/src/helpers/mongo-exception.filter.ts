import {
    ExceptionFilter,
    Catch,
    ArgumentsHost,
    HttpStatus,
    RpcExceptionFilter,
  } from '@nestjs/common';
  import { Request, Response } from 'express';
  import { MongooseError, CastError } from 'mongoose';
  import { MongoError } from 'mongodb';
  
  @Catch(MongooseError)
  export class MongoExceptionFilter implements ExceptionFilter {
    catch(exception: MongooseError, host: ArgumentsHost) {
      if (exception.name === 'CastError') {
        const ctx = host.switchToHttp();
        const response = ctx.getResponse<Response>();
        response.statusCode = HttpStatus.BAD_REQUEST;
        response.json({
          statusCode: HttpStatus.BAD_REQUEST,
          timestamp: new Date().toISOString(),
          message: `Invalid ${exception.message
            .split(' ')
            .slice(11, 12)
            .join('')}: ${exception.message.split(' ').slice(6, 7).join('')}`,
        });
      } else if (exception.name === 'ValidationError') {
        const ctx = host.switchToHttp();
        const response = ctx.getResponse<Response>();
        response.statusCode = HttpStatus.BAD_REQUEST;
        response.json({
          statusCode: HttpStatus.BAD_REQUEST,
          timestamp: new Date().toISOString(),
          message: exception.message,
        });
      }
    }
  }
  
  @Catch(MongoError)
  export class MongoDbExceptionFilter implements ExceptionFilter {
    catch(exception: MongoError, host: ArgumentsHost) {
      console.log(exception, exception.name, exception.code);
      if (exception.code === 11000) {
        const ctx = host.switchToHttp();
        const response = ctx.getResponse<Response>();
        response.statusCode = HttpStatus.BAD_REQUEST;
        response.json({
          statusCode: HttpStatus.BAD_REQUEST,
          timestamp: new Date().toISOString(),
          message: `Duplicate key error at ${exception.message
            .split(' ')
            .slice(10)
            .join(' ')}. Please input a unique ${exception.message
            .split(' ')
            .slice(11, 12)
            .join(' ')
            .split('')
            .slice(0, 5)
            .join('')}`,
        });
      }
    }
  }
  