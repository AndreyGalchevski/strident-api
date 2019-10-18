import { Document, Schema, model } from 'mongoose';

export interface MemberDTO {
  id?: string;
  name: string;
  instrument: string;
  info: string;
  image: string;
  imageNG: string;
  order: number;
}

export type MemberDocument = MemberDTO & Document;

const MemberSchema = new Schema(
  {
    name: String,
    instrument: String,
    info: String,
    image: String,
    imageNG: String,
    order: Number,
  },
  { collection: 'members' },
);

export const MemberModel = model<MemberDocument>('MemberModel', MemberSchema);

export function toDTO(doc: MemberDocument): MemberDTO {
  const member: MemberDTO = {
    id: String(doc._id),
    name: doc.name,
    instrument: doc.instrument,
    info: doc.info,
    image: doc.image,
    imageNG: doc.imageNG,
    order: doc.order,
  };

  return member;
}
