import { Request, Response } from 'express';

import { login } from './service';

export async function handleLogin(req: Request, res: Response): Promise<void> {
  try {
    const token = await login(req.body.username, req.body.password);
    res.send({ token });
  } catch (error) {
    switch (error.message) {
      case 'USER_NOT_FOUND':
        res.status(404).send();
        return;
      case 'WRONG_PASSWORD':
        res.status(400).send();
        return;
      default:
        res.status(500).send();
    }
  }
}
