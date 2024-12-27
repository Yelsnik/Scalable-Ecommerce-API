import { Timestamp } from 'pb/google/protobuf/timestamp';

export class TimestampUtils {
  /**
   * Converts a JavaScript Date to a Protobuf Timestamp
   */
  static dateToTimestamp(date: Date): Timestamp {
    const seconds = Math.floor(date.getTime() / 1000);
    const nanos = (date.getTime() % 1000) * 1000000;

    return {
      seconds: seconds,
      nanos: nanos,
    };
  }

  /**
   * Converts a Protobuf Timestamp to a JavaScript Date
   */
  static timestampToDate(timestamp: Timestamp): Date {
    const millis = timestamp.seconds * 1000 + timestamp.nanos / 1000000;
    return new Date(millis);
  }
}
