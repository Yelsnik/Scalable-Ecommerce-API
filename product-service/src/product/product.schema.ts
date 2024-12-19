import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import mongoose, { HydratedDocument, ObjectId } from 'mongoose';
import { Shop } from 'src/shop/shop.schema';


export type ProductDocument = HydratedDocument<Product>

@Schema()
export class Product {
    @Prop({
        trim: true,
        required: [true, "category is required"]
        })
    category: string

    @Prop({
        trim: true,
        required: [true, "product name is required"]
    })
    productName: string

    @Prop()
    description: string

    @Prop()
    brand: string

    @Prop()
    imageID: ObjectId

    @Prop()
    image: string

    @Prop()
    countInStock: number

    @Prop({
        required: [true, "price is required"]
    })
    price: number

    @Prop({
        required: [true, "currency is required"]
    })
    currency: string

    @Prop({ type: mongoose.Schema.Types.ObjectId, ref: 'Shop' })
    shop: Shop

    @Prop()
    rating: number

    @Prop()
    isFeatured: boolean

    @Prop()
    updatedAt: Date

    @Prop({
        type: Date,
        default: Date.now
    })
    createdAt: Date

}

export const ProductSchema = SchemaFactory.createForClass(Product);