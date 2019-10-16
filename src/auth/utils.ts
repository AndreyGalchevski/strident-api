import { sign, verify } from 'jsonwebtoken';

import { UserDocument } from './entity';

export function createToken(user: UserDocument): string {
  const payload = {
    id: user._id,
    username: user.username,
  };

  const token = sign(payload, process.env.JWT_KEY, { expiresIn: 3600 });

  return token;
}

export function verifyToken(token: string): string | object {
  const decoded = verify(token, process.env.JWT_KEY);
  return decoded;
}
