import { Injectable } from '@nestjs/common';
import { MongoClient, GridFSBucket, GridFSBucketWriteStream } from 'mongodb';
import { Readable } from 'stream';
import { GridFsDto } from './dto/grid-fs.dto';
import { File } from 'pb/file';

@Injectable()
export class GridFsService {
    private bucket: GridFSBucket
    private db: any;

    constructor(){
        const client = new MongoClient(process.env.DATABASE);
        client.connect().then(() => {
          this.db = client.db(process.env.DB_NAME); 
        });
    }

     /**
   * Dynamically creates and returns a GridFSBucket instance with the specified bucket name.
   * @param bucketName - The name of the bucket.
   * @returns GridFSBucket instance.
   */
    getBucket(bucketName: string): GridFSBucket {
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

    async uploadFile(file: GridFsDto, bucketName: string): Promise<GridFSBucketWriteStream> {
        const stream = Readable.from(file.buffer);
        const bucket = this.getBucket(bucketName);

        return new Promise((resolve, reject) => {
            const uploadStream = this.bucket.openUploadStream(file.originalname, {
              metadata: {
                mimetype: file.mimetype,
                size: file.size,
              },
            });            
            stream.pipe(uploadStream)
              .on('error', (error) => reject(error))
              .on('finish', (result: GridFSBucketWriteStream) => resolve(result));
          });
    }

    loadImage(image: File, buffer: Buffer<Uint8Array<ArrayBufferLike>>){
      const file: GridFsDto = {
        id: image.id,
        buffer: buffer,
        originalname: image.originalname,
        filename: image.filename,
        size: image.size,
        chunksize: image.chunksize,
        mimetype: image.mimetype,
        bucketname: image.bucketname,
        md5: image.md5,
        contentType: image.contentType,
        metadata: image.metadata
       } 

       return file
    }

}
