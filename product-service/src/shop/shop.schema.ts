import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { HydratedDocument, ObjectId, Types } from 'mongoose';


export type ShopDocument = HydratedDocument<Shop>

@Schema()
export class Shop {
    @Prop({
        trim: true,
        required: [true, "name is required"]
        })
    name: string

    @Prop({
        trim: true
    })
    description: string

    @Prop({type: Types.ObjectId})
    imageID: Types.ObjectId

    @Prop()
    image: string

    @Prop()
    shopOwner: string

    @Prop()
    updatedAt: Date

    @Prop({
        type: Date,
        default: Date.now
    })
    createdAt: Date

}

export const ShopSchema = SchemaFactory.createForClass(Shop);