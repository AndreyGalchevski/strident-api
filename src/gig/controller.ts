import { Request, Response } from 'express';

import { getGigs, getGig, createGig, updateGig, deleteGig } from './service';

export async function handleGetGigs(req: Request, res: Response): Promise<void> {
  try {
    const gigs = await getGigs();
    res.send(gigs);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleGetGig(req: Request, res: Response): Promise<void> {
  try {
    const gig = await getGig(req.params.id);
    res.send(gig);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleCreateGig(req: Request, res: Response): Promise<void> {
  try {
    const createdGig = await createGig(req.body);
    res.send(createdGig);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleUpdateGig(req: Request, res: Response): Promise<void> {
  try {
    const updatedGig = await updateGig(req.params.id, req.body);
    res.send(updatedGig);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleDeleteGig(req: Request, res: Response): Promise<void> {
  try {
    const deletedGig = await deleteGig(req.params.id);
    res.send(deletedGig);
  } catch (error) {
    res.status(500).send(error);
  }
}
