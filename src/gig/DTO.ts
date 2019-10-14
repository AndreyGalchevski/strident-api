import { Gig } from './model';

export interface GigDTO {
  id?: string;
  venue: string;
  address: string;
  date: Date;
  fbEvent: string;
  image: string;
}

export function toDTO(doc: Gig): GigDTO {
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
