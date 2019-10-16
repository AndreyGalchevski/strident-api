import { Document, Schema, model } from 'mongoose';

export interface LyricDTO {
  id?: string;
  name: string;
  text: string;
}

export type LyricDocument = LyricDTO & Document;

const LyricSchema = new Schema(
  {
    name: String,
    text: String,
  },
  { collection: 'lyrics' },
);

export const LyricModel = model<LyricDocument>('LyricModel', LyricSchema);

export function toDTO(doc: LyricDocument): LyricDTO {
  const lyric: LyricDTO = {
    id: String(doc._id),
    name: doc.name,
    text: doc.text,
  };

  return lyric;
}
