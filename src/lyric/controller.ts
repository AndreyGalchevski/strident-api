import { Request, Response } from 'express';

import { getLyrics, getLyric, createLyric, updateLyric, deleteLyric } from './service';
import { toDTO } from './entity';

export async function handleGetLyrics(req: Request, res: Response): Promise<void> {
  try {
    const lyrics = await getLyrics();
    const lyricDTOs = lyrics.map(lyric => toDTO(lyric));
    res.send(lyricDTOs);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleGetLyric(req: Request, res: Response): Promise<void> {
  try {
    const lyric = await getLyric(req.params.id);
    const lyricDTO = toDTO(lyric);
    res.send(lyricDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleCreateLyric(req: Request, res: Response): Promise<void> {
  try {
    const createdLyric = await createLyric(req.body);
    const createdLyricDTO = toDTO(createdLyric);
    res.send(createdLyricDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleUpdateLyric(req: Request, res: Response): Promise<void> {
  try {
    const updatedLyric = await updateLyric(req.params.id, req.body);
    const updatedLyricDTO = toDTO(updatedLyric);
    res.send(updatedLyricDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleDeleteLyric(req: Request, res: Response): Promise<void> {
  try {
    const deletedLyric = await deleteLyric(req.params.id);
    const deletedLyricDTO = toDTO(deletedLyric);
    res.send(deletedLyricDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}
