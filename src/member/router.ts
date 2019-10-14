import { Router } from 'express';

import {
  handleGetMembers,
  handleGetMember,
  handleCreateMember,
  handleUpdateMember,
  handleDeleteMember,
} from './controller';
import authMiddleware from '../auth/middleware';

const memberRouter = Router();

memberRouter.get('/', handleGetMembers);
memberRouter.get('/:id', handleGetMember);
memberRouter.post('/', authMiddleware, handleCreateMember);
memberRouter.put('/:id', authMiddleware, handleUpdateMember);
memberRouter.delete('/:id', authMiddleware, handleDeleteMember);

export default memberRouter;
