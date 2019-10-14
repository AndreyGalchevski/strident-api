import { Request, Response } from 'express';

import { getVideos, getVideo, createVideo, updateVideo, deleteVideo } from './service';

export async function handleGetVideos(req: Request, res: Response): Promise<void> {
  try {
    const videos = await getVideos();
    res.send(videos);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleGetVideo(req: Request, res: Response): Promise<void> {
  try {
    const video = await getVideo(req.params.id);
    res.send(video);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleCreateVideo(req: Request, res: Response): Promise<void> {
  try {
    const createdVideo = await createVideo(req.body);
    res.send(createdVideo);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleUpdateVideo(req: Request, res: Response): Promise<void> {
  try {
    const updatedVideo = await updateVideo(req.params.id, req.body);
    res.send(updatedVideo);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleDeleteVideo(req: Request, res: Response): Promise<void> {
  try {
    const deletedVideo = await deleteVideo(req.params.id);
    res.send(deletedVideo);
  } catch (error) {
    res.status(500).send(error);
  }
}
