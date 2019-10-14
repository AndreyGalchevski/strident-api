import { Router } from 'express';

import {
  handleGetGigs,
  handleCreateGig,
  handleGetGig,
  handleUpdateGig,
  handleDeleteGig,
} from './controller';
import authMiddleware from '../auth/middleware';

const gigRouter = Router();

gigRouter.get('/', handleGetGigs);
gigRouter.get('/:id', handleGetGig);
gigRouter.post('/', authMiddleware, handleCreateGig);
gigRouter.put('/:id', authMiddleware, handleUpdateGig);
gigRouter.delete('/:id', authMiddleware, handleDeleteGig);

export default gigRouter;
