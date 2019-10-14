import { Router } from 'express';

import {
  handleGetLyrics,
  handleGetLyric,
  handleCreateLyric,
  handleUpdateLyric,
  handleDeleteLyric,
} from './controller';
import authMiddleware from '../auth/middleware';

const lyricRouter = Router();

lyricRouter.get('/', handleGetLyrics);
lyricRouter.get('/:id', handleGetLyric);
lyricRouter.post('/', authMiddleware, handleCreateLyric);
lyricRouter.put('/:id', authMiddleware, handleUpdateLyric);
lyricRouter.delete('/:id', authMiddleware, handleDeleteLyric);

export default lyricRouter;
