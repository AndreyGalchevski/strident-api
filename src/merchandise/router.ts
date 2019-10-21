import { Router } from 'express';

import {
  handleGetMerchandises,
  handleGetMerchandise,
  handleCreateMerchandise,
  handleUpdateMerchandise,
  handleDeleteMerchandise,
} from './controller';
import authMiddleware from '../auth/middleware';

const merchandiseRouter = Router();

merchandiseRouter.get('/', handleGetMerchandises);
merchandiseRouter.get('/:id', handleGetMerchandise);
merchandiseRouter.post('/', authMiddleware, handleCreateMerchandise);
merchandiseRouter.put('/:id', authMiddleware, handleUpdateMerchandise);
merchandiseRouter.delete('/:id', authMiddleware, handleDeleteMerchandise);

export default merchandiseRouter;
