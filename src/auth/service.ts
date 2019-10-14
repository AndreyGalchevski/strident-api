import UserModel from './model';
import { createToken } from './utils';

export async function login(username: string, password: string): Promise<string> {
  const filter = { username };
  const user = await UserModel.findOne(filter);

  if (!user) {
    throw new Error('USER_NOT_FOUND');
  } else if (user.password !== password) {
    throw new Error('WRONG_PASSWORD');
  }

  const token = createToken(user);
  return token;
}
