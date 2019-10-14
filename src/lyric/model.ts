import { Document, Schema, model } from 'mongoose';

export interface Lyric extends Document {
  name: string;
  text: string;
}

const LyricSchema = new Schema(
  {
    name: String,
    text: String,
  },
  { collection: 'lyrics' },
);

export default model<Lyric>('LyricModel', LyricSchema);
