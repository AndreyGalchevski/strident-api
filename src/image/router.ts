import { Router } from 'express';

import { handleGetLyrics, handleGetLyric } from './controller';
import authMiddleware from '../auth/middleware';

const lyricRouter = Router();

lyricRouter.post('/', authMiddleware, handleGetLyrics);
lyricRouter.delete('/', authMiddleware, handleGetLyric);

export default lyricRouter;
