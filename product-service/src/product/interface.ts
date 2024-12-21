import { Timestamp } from "pb/google/protobuf/timestamp";

export class Time implements Timestamp{
    seconds: number;
    nanos: number;
    
    constructor(time: Date){
        const mill = time.getTime()
        this.seconds = Math.floor(mill / 1000);
        this.nanos = (mill % 1000) * 1_000_000;
    }
}