import { Request, Response } from 'express';

import {
  getMerchandises,
  getMerchandise,
  createMerchandise,
  updateMerchandise,
  deleteMerchandise,
} from './service';
import { toDTO } from './entity';

export async function handleGetMerchandises(req: Request, res: Response): Promise<void> {
  try {
    const merchandises = await getMerchandises();
    const merchandiseDTOs = merchandises.map(merchandise => toDTO(merchandise));
    res.send(merchandiseDTOs);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleGetMerchandise(req: Request, res: Response): Promise<void> {
  try {
    const merchandise = await getMerchandise(req.params.id);
    const merchandiseDTO = toDTO(merchandise);
    res.send(merchandiseDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleCreateMerchandise(req: Request, res: Response): Promise<void> {
  try {
    const createdMerchandise = await createMerchandise(req.body);
    const createdMerchandiseDTO = toDTO(createdMerchandise);
    res.send(createdMerchandiseDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleUpdateMerchandise(req: Request, res: Response): Promise<void> {
  try {
    const updatedMerchandise = await updateMerchandise(req.params.id, req.body);
    const updatedMerchandiseDTO = toDTO(updatedMerchandise);
    res.send(updatedMerchandiseDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleDeleteMerchandise(req: Request, res: Response): Promise<void> {
  try {
    const deletedMerchandise = await deleteMerchandise(req.params.id);
    const deletedMerchandiseDTO = toDTO(deletedMerchandise);
    res.send(deletedMerchandiseDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}
