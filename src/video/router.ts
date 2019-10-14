import { Router } from 'express';

import {
  handleGetVideos,
  handleGetVideo,
  handleCreateVideo,
  handleUpdateVideo,
  handleDeleteVideo,
} from './controller';
import authMiddleware from '../auth/middleware';

const videoRouter = Router();

videoRouter.get('/', handleGetVideos);
videoRouter.get('/:id', handleGetVideo);
videoRouter.post('/', authMiddleware, handleCreateVideo);
videoRouter.put('/:id', authMiddleware, handleUpdateVideo);
videoRouter.delete('/:id', authMiddleware, handleDeleteVideo);

export default videoRouter;
