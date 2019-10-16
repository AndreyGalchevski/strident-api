import { Document, Schema, model } from 'mongoose';

export interface GigDTO {
  id?: string;
  venue: string;
  address: string;
  date: Date;
  fbEvent: string;
  image: string;
}

export type GigDocument = GigDTO & Document;

const GigSchema = new Schema(
  {
    venue: String,
    address: String,
    date: Date,
    fbEvent: String,
    image: String,
  },
  { collection: 'gigs' },
);

export const GigModel = model<GigDocument>('GigModel', GigSchema);

export function toDTO(doc: GigDocument): GigDTO {
  const gig: GigDTO = {
    id: String(doc._id),
    venue: doc.venue,
    address: doc.address,
    date: doc.date,
    fbEvent: doc.fbEvent,
    image: doc.image,
  };

  return gig;
}
