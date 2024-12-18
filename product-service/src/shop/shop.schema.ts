import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { HydratedDocument } from 'mongoose';


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

    @Prop()
    imageID: string

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