import { Test, TestingModule } from '@nestjs/testing';
import { GridFsService } from './grid-fs.service';

describe('GridFsService', () => {
  let service: GridFsService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [GridFsService],
    }).compile();

    service = module.get<GridFsService>(GridFsService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
