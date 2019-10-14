import { Document, Schema, model } from 'mongoose';

export interface Member extends Document {
  name: string;
  instrument: string;
  info: string;
  image: string;
}

const MemberSchema = new Schema(
  {
    name: String,
    instrument: String,
    info: String,
    image: String,
  },
  { collection: 'members' },
);

export default model<Member>('MemberModel', MemberSchema);
