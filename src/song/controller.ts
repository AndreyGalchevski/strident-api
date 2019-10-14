import { Request, Response } from 'express';

import { getSongs, getSong, createSong, updateSong, deleteSong } from './service';
import { toDTO } from './DTO';

export async function handleGetSongs(req: Request, res: Response): Promise<void> {
  try {
    const songs = await getSongs();
    const songDTOs = songs.map(video => toDTO(video));
    res.send(songDTOs);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleGetSong(req: Request, res: Response): Promise<void> {
  try {
    const song = await getSong(req.params.id);
    const songDTO = toDTO(song);
    res.send(songDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleCreateSong(req: Request, res: Response): Promise<void> {
  try {
    const createdSong = await createSong(req.body);
    const createdSongDTO = toDTO(createdSong);
    res.send(createdSongDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleUpdateSong(req: Request, res: Response): Promise<void> {
  try {
    const updatedSong = await updateSong(req.params.id, req.body);
    const updatedSongDTO = toDTO(updatedSong);
    res.send(updatedSongDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleDeleteSong(req: Request, res: Response): Promise<void> {
  try {
    const deletedSong = await deleteSong(req.params.id);
    const deletedSongDTO = toDTO(deletedSong);
    res.send(deletedSongDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}
