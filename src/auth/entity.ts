import { Document, Schema, model } from 'mongoose';

export interface UserDTO {
  id?: string;
  username: string;
  password: string;
}

export type UserDocument = UserDTO & Document;

const UserSchema = new Schema(
  {
    username: String,
    password: String,
  },
  { collection: 'users' },
);

export const UserModel = model<UserDocument>('UserModel', UserSchema);

export function toDTO(doc: UserDocument): UserDTO {
  const user: UserDTO = {
    id: String(doc._id),
    username: doc.username,
    password: doc.password,
  };

  return user;
}
