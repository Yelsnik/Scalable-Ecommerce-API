import { Test, TestingModule } from '@nestjs/testing';
import { CartController } from './cart.controller';
import { CartService } from './cart.service';
import { of } from 'rxjs';

describe('CartController', () => {
  let controller: CartController;
  let service: CartService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [CartController],
      providers: [
        {
          provide: CartService,
          useValue: {
            addToCart: jest.fn().mockReturnValue(of({})),
            getCartItemsByCart: jest.fn().mockReturnValue(of([])),
            getCartItemById: jest.fn().mockReturnValue(of({})),
            updateCart: jest.fn().mockReturnValue(of({})),
            remove: jest.fn().mockReturnValue(of({})),
          },
        },
      ],
    }).compile();

    controller = module.get<CartController>(CartController);
    service = module.get<CartService>(CartService);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
