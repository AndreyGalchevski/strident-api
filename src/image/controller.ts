import { Request, Response } from 'express';

import { uploadImage, destroyImage } from './service';
import { convertToWebP } from './manipulator';

interface File {
  path: string;
}

interface ImageRequest extends Request {
  files: File[];
}

export async function handleUploadImage(req: ImageRequest, res: Response): Promise<void> {
  const files = Object.values(req.files);

  if (files.length !== 1) {
    res.status(400).send({ error: 'Please attach one image' });
    return;
  }

  const publicID = `strident/${req.query.folderName}/${process.env.NODE_ENV}/${req.query.fileName}`;
  const imagePath = files[0].path;
  const ngImagePath = await convertToWebP(imagePath);
  try {
    const imageURL = await uploadImage(imagePath, publicID);
    const ngImageURL = await uploadImage(ngImagePath, `${publicID}_ng`);
    res.send({ imageURL, ngImageURL });
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleDestroyImage(req: Request, res: Response): Promise<void> {
  const numberOfImages = req.body.imageUrls.length;
  const folderName = `strident/${req.query.folder}`;

  if (numberOfImages !== 1) {
    res.status(400).send({ error: 'No image was specified' });
    return;
  }

  const destroyedImage = await destroyImage(req.body.imageUrls[0], folderName);
  res.send({ destroyedImage });
}
