import { Document, Schema, model } from 'mongoose';

export interface VideoDTO {
  id?: string;
  name: string;
  url: string;
  date: Date;
}

export type VideoDocument = VideoDTO & Document;

const VideoSchema = new Schema(
  {
    name: String,
    url: String,
    date: Date,
  },
  { collection: 'videos' },
);

export const VideoModel = model<VideoDocument>('VideoModel', VideoSchema);

export function toDTO(doc: VideoDocument): VideoDTO {
  const video: VideoDTO = {
    id: String(doc._id),
    name: doc.name,
    url: doc.url,
    date: doc.date,
  };

  return video;
}
