import { Timestamp } from 'pb/google/protobuf/timestamp';

export class TimestampUtils {
  /**
   * Converts a JavaScript Date to a Protobuf Timestamp
   */
  static dateToTimestamp(date: Date): Timestamp | undefined {
    if (!date) {
      return undefined;
    }
    return {
      seconds: Math.floor(date.getTime() / 1000) as number,
      nanos: ((date.getTime() % 1000) * 1000000) as number,
    };
  }

  /**
   * Converts a Protobuf Timestamp to a JavaScript Date
   */
  static timestampToDate(timestamp: Timestamp): Date {
    return new Date(
      (timestamp.seconds as number) * 1000 +
        Math.floor((timestamp.nanos as number) / 1000000),
    );
  }
}
