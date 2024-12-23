import { Timestamp } from "pb/google/protobuf/timestamp";

export class Time implements Timestamp {
    seconds: number;
    nanos: number;

    constructor(time: Date){
        const secondsSinceEpoch = Math.floor(time.getTime() / 1000);

        const [hrSeconds, nanoseconds] = process.hrtime();
        this.seconds = secondsSinceEpoch + hrSeconds,
        this.nanos = nanoseconds
    }
}