import { Document, Schema, model } from 'mongoose';

export interface GigDTO {
  id?: string;
  name: string;
  venue: string;
  address: string;
  city: string;
  date: Date;
  fbEvent: string;
  image: string;
  imageNG: string;
}

export type GigDocument = GigDTO & Document;

const GigSchema = new Schema(
  {
    name: String,
    venue: String,
    address: String,
    city: String,
    date: Date,
    fbEvent: String,
    image: String,
    imageNG: String,
  },
  { collection: 'gigs' },
);

export const GigModel = model<GigDocument>('GigModel', GigSchema);

export function toDTO(doc: GigDocument): GigDTO {
  const gig: GigDTO = {
    id: String(doc._id),
    name: doc.name,
    venue: doc.venue,
    address: doc.address,
    city: doc.city,
    date: doc.date,
    fbEvent: doc.fbEvent,
    image: doc.image,
    imageNG: doc.imageNG,
  };

  return gig;
}
