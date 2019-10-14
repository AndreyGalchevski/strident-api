import { Router } from 'express';

import { handleLogin } from './controller';

const authRouter = Router();

authRouter.post('/login', handleLogin);

export default authRouter;
