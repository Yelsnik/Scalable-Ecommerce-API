import { HttpException, HttpStatus } from "@nestjs/common";
import { json } from "stream/consumers";

export class GRPCToHTTPExceptions {
  
    constructor(err: any){
        this.handleGrpcError(err)
    }

    private handleGrpcError(error: any): never {
          if (error){
            const httpStatus = this.mapGrpcCodeToHttpStatus(error.code);
            throw new HttpException(error.message, httpStatus);
          }
            throw new HttpException('Internal server error', HttpStatus.INTERNAL_SERVER_ERROR);
      }

    private mapGrpcCodeToHttpStatus(grpcCode: number): HttpStatus {
        switch (grpcCode) {
          case 3: // INVALID_ARGUMENT
            return HttpStatus.BAD_REQUEST;
          case 5: // NOT_FOUND
            return HttpStatus.NOT_FOUND;
          case 13: // INTERNAL
            return HttpStatus.INTERNAL_SERVER_ERROR;
          case 16: // UNAUTHENTICATED
            return HttpStatus.UNAUTHORIZED;
          case 14: // UNAVAILABLE
            return HttpStatus.SERVICE_UNAVAILABLE;
          default:
            return HttpStatus.INTERNAL_SERVER_ERROR;
        }
    }
}