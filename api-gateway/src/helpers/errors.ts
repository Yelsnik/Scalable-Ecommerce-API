import { HttpException, HttpStatus } from '@nestjs/common';
import { HTTP_CODE_METADATA } from '@nestjs/common/constants';
import { json } from 'stream/consumers';

export class GRPCToHTTPExceptions {
  constructor(err: any, response: any) {
    this.handleGrpcError(err, response);
  }

  private handleGrpcError(exception: any, response: any) {
    if (exception) {
      if (exception.code) {
        const code = this.mapGrpcCodeToHttpStatus(exception.code);

        return response.status(code).json({
          statusCode: code,
          message: exception.details,
          stack: exception.stack,
          // cause: exception.cause,
          timestamp: new Date().toISOString(),
          // path: request.url,
        });
      }

      return response.status(HttpStatus.INTERNAL_SERVER_ERROR).json({
        statusCode: HttpStatus.INTERNAL_SERVER_ERROR,
        message: exception.details,
        stack: exception.stack,
        // cause: exception.cause,
        timestamp: new Date().toISOString(),
        // path: request.url,
      });
    } else if (exception instanceof TypeError) {
      console.log(exception.name);
      return response.status(HttpStatus.INTERNAL_SERVER_ERROR).json({
        statusCode: HttpStatus.INTERNAL_SERVER_ERROR,
        message: exception.message,
        stack: exception.stack,
        //cause: exception.cause,
        timestamp: new Date().toISOString(),
        // path: request.url,
      });
    } else if (exception instanceof SyntaxError) {
      console.log(exception.name);
      return response.status(HttpStatus.INTERNAL_SERVER_ERROR).json({
        statusCode: HttpStatus.INTERNAL_SERVER_ERROR,
        message: exception.message,
        stack: exception.stack,
        // cause: exception.cause,
        timestamp: new Date().toISOString(),
        // path: request.url,
      });
    } else {
      return response.status(exception.code).json({
        statusCode: exception.code,
        message: exception.message,
        //stack: exception.stack,
        // cause: exception.cause,
        timestamp: new Date().toISOString(),
        // path: request.url,
      });
    }
  }

  private mapGrpcCodeToHttpStatus(grpcCode: number): HttpStatus {
    switch (grpcCode) {
      case 2: // UNKNOWN
        return HttpStatus.AMBIGUOUS;
      case 3: // INVALID_ARGUMENT
        return HttpStatus.BAD_REQUEST;
      case 5: // NOT_FOUND
        return HttpStatus.NOT_FOUND;
      case 6: // ALREADY EXISTS
        return HttpStatus.CONFLICT;
      case 7: // PERMISSION DENIED
        return HttpStatus.FORBIDDEN;
      case 13: // INTERNAL
        return HttpStatus.INTERNAL_SERVER_ERROR;
      case 14: // UNAVAILABLE
        return HttpStatus.SERVICE_UNAVAILABLE;
      case 16: // UNAUTHENTICATED
        return HttpStatus.UNAUTHORIZED;
      default:
        return HttpStatus.INTERNAL_SERVER_ERROR;
    }
  }
}
