import { Document, Schema, model } from 'mongoose';

export interface Video extends Document {
  name: string;
  url: string;
  date: Date;
}

const VideoSchema = new Schema(
  {
    name: String,
    url: String,
    date: Date,
  },
  { collection: 'videos' },
);

export default model<Video>('VideoModel', VideoSchema);
