import {
  Catch,
  ArgumentsHost,
  RpcExceptionFilter,
  NotFoundException,
} from '@nestjs/common';
import { RpcException } from '@nestjs/microservices';
import mongoose from 'mongoose';
import { GrpcInternalException } from 'nestjs-grpc-exceptions';
import { Observable, throwError } from 'rxjs';

@Catch()
export class AllExceptionFilter implements RpcExceptionFilter {
  catch(exception: any, host: ArgumentsHost): Observable<any> {
    const details = this.getExceptionDetails(exception);

    return throwError(() => ({
      code: details.code,
      message: details.message,
      details: details.details,
    }));
  }

  private getExceptionDetails(exception: any) {
    // Handle Mongoose Validation Errors
    if (exception instanceof mongoose.Error.ValidationError) {
      console.log(exception.name);
      const errors = Object.values(exception.errors)
        .map((err) => err.message)
        .join(', ');

      return {
        code: 3,
        message: 'Validation failed',
        details: errors,
      };
    }

    // Handle Mongoose Cast Errors (invalid ObjectId etc)
    if (exception instanceof mongoose.Error.CastError) {
      console.log(exception.name);
      return {
        code: 3,
        message: 'Invalid data format',
        details: exception.message,
      };
    }

    // Handle Type Errors
    if (exception instanceof TypeError) {
      console.log(exception);
      return {
        code: 13,
        message: 'Type error occurred',
        details: exception.message,
      };
    }

    if (exception instanceof SyntaxError) {
      console.log(exception);
      return {
        code: 3,
        message: 'Type error occurred',
        details: exception.message,
      };
    }

    // Handle MongoDB Duplicate Key Errors
    if (exception.code === 11000) {
      console.log(exception.name);
      return {
        code: 3,
        message: 'Duplicate entry',
        details: Object.keys(exception.keyPattern).join(', '),
      };
    }
  }
}

@Catch(RpcException)
export class RpcExceptionFilters implements RpcExceptionFilter<RpcException> {
  catch(exception: RpcException, host: ArgumentsHost): Observable<any> {
    console.log(exception.name);
    return throwError(() => exception.getError());
  }
}
