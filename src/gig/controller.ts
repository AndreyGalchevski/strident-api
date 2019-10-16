import { Request, Response } from 'express';

import { getGigs, getGig, createGig, updateGig, deleteGig } from './service';
import { toDTO } from './entity';

export async function handleGetGigs(req: Request, res: Response): Promise<void> {
  try {
    const gigs = await getGigs();
    const gigDTOs = gigs.map(gig => toDTO(gig));
    res.send(gigDTOs);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleGetGig(req: Request, res: Response): Promise<void> {
  try {
    const gig = await getGig(req.params.id);
    const gigDTO = toDTO(gig);
    res.send(gigDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleCreateGig(req: Request, res: Response): Promise<void> {
  try {
    const createdGig = await createGig(req.body);
    const createdGigDTO = toDTO(createdGig);
    res.send(createdGigDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleUpdateGig(req: Request, res: Response): Promise<void> {
  try {
    const updatedGig = await updateGig(req.params.id, req.body);
    const updatedGigDTO = toDTO(updatedGig);
    res.send(updatedGigDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleDeleteGig(req: Request, res: Response): Promise<void> {
  try {
    const deletedGig = await deleteGig(req.params.id);
    const deletedGigDTO = toDTO(deletedGig);
    res.send(deletedGigDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}
