import { Router } from 'express';

import { handleUploadImage, handleDestroyImage } from './controller';
import authMiddleware from '../auth/middleware';

const imageRouter = Router();

imageRouter.post('/', authMiddleware, handleUploadImage);
imageRouter.delete('/', authMiddleware, handleDestroyImage);

export default imageRouter;
