import { Document, Schema, model } from 'mongoose';

export interface Gig extends Document {
  venue: string;
  address: string;
  date: Date;
  fbEvent: string;
  image: string;
}

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

export default model<Gig>('GigModel', GigSchema);
