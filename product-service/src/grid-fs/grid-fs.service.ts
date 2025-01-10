import { Injectable, OnModuleInit } from '@nestjs/common';
import { MongoClient, GridFSBucket, GridFSBucketWriteStream } from 'mongodb';
import { Readable } from 'stream';
import { GridFsDto } from './dto/grid-fs.dto';
import { File } from 'pb/product-service/file';

@Injectable()
export class GridFsService implements OnModuleInit {
  //private bucket: GridFSBucket
  private db: any;

  private async initDb() {
    const client = new MongoClient(process.env.DATABASE);
    await client.connect();
    this.db = client.db(process.env.DB_NAME);
    console.log('Database connected');
  }

  async onModuleInit() {
    await this.initDb();
  }

  /**
   * Dynamically creates and returns a GridFSBucket instance with the specified bucket name.
   * @param bucketName - The name of the bucket.
   * @returns GridFSBucket instance.
   */
  private getBucket(bucketName: string): GridFSBucket {
    if (!this.db) {
      throw new Error('Database connection not initialized');
    }
    return new GridFSBucket(this.db, { bucketName });
  }

  /**
   * Uploads a file to the specified bucket.
   * @param file - File data to be uploaded.
   * @param bucketName - The name of the bucket.
   * @returns Promise with the upload result.
   */

  async uploadFile(file: GridFsDto, bucketName: string): Promise<string> {
    const stream = Readable.from(file.buffer);
    const bucket = this.getBucket(bucketName);

    return new Promise((resolve, reject) => {
      const uploadStream = bucket.openUploadStream(file.originalname, {
        metadata: {
          mimetype: file.mimetype,
          size: file.size,
        },
      });
      stream
        .pipe(uploadStream)
        .on('error', (error) => reject(error))
        .on('finish', () => resolve(uploadStream.id.toString()));
    });
  }

  loadImage(image: File, buffer: Buffer) {
    const file: GridFsDto = {
      buffer: buffer,
      originalname: image.originalname,
      fieldname: image.fieldname,
      size: image.size,
      mimetype: image.mimetype,
      encoding: image.encoding,
    };

    return file;
  }
}
