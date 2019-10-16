import { Document, Schema, model } from 'mongoose';

export interface SongDTO {
  id?: string;
  name: string;
  album: string;
  url: string;
}

export type SongDocument = SongDTO & Document;

const SongSchema = new Schema(
  {
    name: String,
    url: String,
    album: String,
  },
  { collection: 'songs' },
);

export const SongModel = model<SongDocument>('SongModel', SongSchema);

export function toDTO(doc: SongDocument): SongDTO {
  const video: SongDTO = {
    id: String(doc._id),
    name: doc.name,
    album: doc.album,
    url: doc.url,
  };

  return video;
}
