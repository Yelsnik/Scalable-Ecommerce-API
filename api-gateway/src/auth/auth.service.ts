import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import {
  AuthServiceClient,
  CreateTokenRequest,
  GetUserByIdRequest,
  GetUserByIdResponse,
  VerifyTokenRequest,
  VerifyTokenResponse,
} from 'pb/auth_service';
import { Observable } from 'rxjs';

@Injectable()
export class AuthService implements OnModuleInit {
  authService: AuthServiceClient;

  constructor(@Inject('AUTH_SERVICE') private authClient: ClientGrpc) {}

  onModuleInit() {
    this.authService =
      this.authClient.getService<AuthServiceClient>('AuthService');
  }

  createToken(request: CreateTokenRequest) {
    return this.authService.createToken(request);
  }

  verifyToken(request: VerifyTokenRequest): Observable<VerifyTokenResponse> {
    return this.authService.verifyToken(request);
  }

  getUserByID(request: GetUserByIdRequest): Observable<GetUserByIdResponse> {
    return this.authService.getUserById(request);
  }
}
