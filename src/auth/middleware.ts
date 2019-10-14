import { Request, Response, NextFunction } from 'express';

import { verifyToken } from './utils';

export default function authMiddleware(req: Request, res: Response, next: NextFunction): void {
  const token = req.headers.authorization;
  if (!token) {
    res.status(401).send('Unauthorized');
    return;
  }

  try {
    verifyToken(token.replace('Bearer ', ''));
  } catch (error) {
    res.status(401).send('Unauthorized');
    return;
  }

  next();
}
