import { Document, Schema, model } from 'mongoose';

export interface Song extends Document {
  name: string;
  url: string;
  album: string;
}

const SongSchema = new Schema(
  {
    name: String,
    url: String,
    album: String,
  },
  { collection: 'songs' },
);

export default model<Song>('SongModel', SongSchema);
