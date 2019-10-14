import { Request, Response } from 'express';

import { getSongs, getSong, createSong, updateSong, deleteSong } from './service';

export async function handleGetSongs(req: Request, res: Response): Promise<void> {
  try {
    const songs = await getSongs();
    res.send(songs);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleGetSong(req: Request, res: Response): Promise<void> {
  try {
    const song = await getSong(req.params.id);
    res.send(song);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleCreateSong(req: Request, res: Response): Promise<void> {
  try {
    const createdSong = await createSong(req.body);
    res.send(createdSong);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleUpdateSong(req: Request, res: Response): Promise<void> {
  try {
    const updatedSong = await updateSong(req.params.id, req.body);
    res.send(updatedSong);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleDeleteSong(req: Request, res: Response): Promise<void> {
  try {
    const deletedSong = await deleteSong(req.params.id);
    res.send(deletedSong);
  } catch (error) {
    res.status(500).send(error);
  }
}
