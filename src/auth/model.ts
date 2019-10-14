import { Document, Schema, model } from 'mongoose';

export interface User extends Document {
  username: string;
  password: string;
}

const UserSchema = new Schema(
  {
    username: String,
    password: String,
  },
  { collection: 'users' },
);

export default model<User>('UserModel', UserSchema);
