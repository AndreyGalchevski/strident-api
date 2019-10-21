import { Document, Schema, model } from 'mongoose';

export interface MerchandiseDTO {
  id?: string;
  name: string;
  type: string;
  price: number;
  url: string;
  image: string;
  imageNG: string;
}

export type MerchandiseDocument = MerchandiseDTO & Document;

const MerchandiseSchema = new Schema(
  {
    name: String,
    type: String,
    price: Number,
    url: String,
    image: String,
    imageNG: String,
  },
  { collection: 'merchandises' },
);

export const MerchandiseModel = model<MerchandiseDocument>('MerchandiseModel', MerchandiseSchema);

export function toDTO(doc: MerchandiseDocument): MerchandiseDTO {
  const merchandise: MerchandiseDTO = {
    id: String(doc._id),
    name: doc.name,
    type: doc.type,
    price: doc.price,
    url: doc.url,
    image: doc.image,
    imageNG: doc.imageNG,
  };

  return merchandise;
}
