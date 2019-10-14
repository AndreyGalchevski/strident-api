import { Request, Response } from 'express';

import { getMembers, getMember, createMember, updateMember, deleteMember } from './service';

export async function handleGetMembers(req: Request, res: Response): Promise<void> {
  try {
    const members = await getMembers();
    res.send(members);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleGetMember(req: Request, res: Response): Promise<void> {
  try {
    const member = await getMember(req.params.id);
    res.send(member);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleCreateMember(req: Request, res: Response): Promise<void> {
  try {
    const createdMember = await createMember(req.body);
    res.send(createdMember);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleUpdateMember(req: Request, res: Response): Promise<void> {
  try {
    const updatedMember = await updateMember(req.params.id, req.body);
    res.send(updatedMember);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleDeleteMember(req: Request, res: Response): Promise<void> {
  try {
    const deletedMember = await deleteMember(req.params.id);
    res.send(deletedMember);
  } catch (error) {
    res.status(500).send(error);
  }
}
