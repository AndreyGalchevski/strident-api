import { Request, Response } from 'express';

import { getVideos, getVideo, createVideo, updateVideo, deleteVideo } from './service';
import { toDTO } from './DTO';

export async function handleGetVideos(req: Request, res: Response): Promise<void> {
  try {
    const videos = await getVideos();
    const videoDTOs = videos.map(video => toDTO(video));
    res.send(videoDTOs);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleGetVideo(req: Request, res: Response): Promise<void> {
  try {
    const video = await getVideo(req.params.id);
    const videoDTO = toDTO(video);
    res.send(videoDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleCreateVideo(req: Request, res: Response): Promise<void> {
  try {
    const createdVideo = await createVideo(req.body);
    const createdVideoDTO = toDTO(createdVideo);
    res.send(createdVideoDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleUpdateVideo(req: Request, res: Response): Promise<void> {
  try {
    const updatedVideo = await updateVideo(req.params.id, req.body);
    const updatedVideoDTO = toDTO(updatedVideo);
    res.send(updatedVideoDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleDeleteVideo(req: Request, res: Response): Promise<void> {
  try {
    const deletedVideo = await deleteVideo(req.params.id);
    const deletedVideoDTO = toDTO(deletedVideo);
    res.send(deletedVideoDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}
