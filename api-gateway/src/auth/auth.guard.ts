import {
  CanActivate,
  ExecutionContext,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common';
import { lastValueFrom, Observable } from 'rxjs';
import { AuthService } from './auth.service';
import { GetUserByIdRequest, VerifyTokenRequest } from 'pb/auth_service';

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(private authService: AuthService) {}

  async canActivate(context: ExecutionContext) {
    const request = context.switchToHttp().getRequest();
    let t = '';

    if (
      request.headers.authorization &&
      request.headers.authorization.startsWith('Bearer')
    ) {
      t = request.headers.authorization.split(' ')[1];
    } else {
      throw new UnauthorizedException('Please login!');
    }

    if (!t) {
      throw new UnauthorizedException('Please login!');
    }

    const reqToken: VerifyTokenRequest = {
      token: t,
    };

    const decode = this.authService.verifyToken(reqToken);

    const payload = await lastValueFrom(decode);

    const reqUser: GetUserByIdRequest = {
      id: payload.payload.userId,
    };

    const observeUser = this.authService.getUserByID(reqUser);

    const user = await lastValueFrom(observeUser);

    if (!user) {
      throw new UnauthorizedException('user does not exist');
    }

    request.user = user;

    return true;
  }
}
