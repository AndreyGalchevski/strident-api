import { Router } from 'express';

import {
  handleGetSongs,
  handleGetSong,
  handleCreateSong,
  handleUpdateSong,
  handleDeleteSong,
} from './controller';
import authMiddleware from '../auth/middleware';

const songRouter = Router();

songRouter.get('/', handleGetSongs);
songRouter.get('/:id', handleGetSong);
songRouter.post('/', authMiddleware, handleCreateSong);
songRouter.put('/:id', authMiddleware, handleUpdateSong);
songRouter.delete('/:id', authMiddleware, handleDeleteSong);

export default songRouter;
