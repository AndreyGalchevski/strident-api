import { Document, Schema, model } from 'mongoose';

export interface Video extends Document {
  name: string;
  url: string;
}

const VideoSchema = new Schema(
  {
    name: String,
    url: String,
  },
  { collection: 'videos' },
);

export default model<Video>('VideoModel', VideoSchema);
