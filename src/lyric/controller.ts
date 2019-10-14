import { Request, Response } from 'express';

import { getLyrics, getLyric, createLyric, updateLyric, deleteLyric } from './service';

export async function handleGetLyrics(req: Request, res: Response): Promise<void> {
  try {
    const lyrics = await getLyrics();
    res.send(lyrics);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleGetLyric(req: Request, res: Response): Promise<void> {
  try {
    const lyric = await getLyric(req.params.id);
    res.send(lyric);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleCreateLyric(req: Request, res: Response): Promise<void> {
  try {
    const createdLyric = await createLyric(req.body);
    res.send(createdLyric);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleUpdateLyric(req: Request, res: Response): Promise<void> {
  try {
    const updatedLyric = await updateLyric(req.params.id, req.body);
    res.send(updatedLyric);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleDeleteLyric(req: Request, res: Response): Promise<void> {
  try {
    const deletedLyric = await deleteLyric(req.params.id);
    res.send(deletedLyric);
  } catch (error) {
    res.status(500).send(error);
  }
}
