import { Request, Response } from 'express';

import { getMembers, getMember, createMember, updateMember, deleteMember } from './service';
import { toDTO } from './entity';

export async function handleGetMembers(req: Request, res: Response): Promise<void> {
  try {
    const members = await getMembers();
    const memberDTOs = members.map(member => toDTO(member));
    res.send(memberDTOs);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleGetMember(req: Request, res: Response): Promise<void> {
  try {
    const member = await getMember(req.params.id);
    const memberDTO = toDTO(member);
    res.send(memberDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleCreateMember(req: Request, res: Response): Promise<void> {
  try {
    const createdMember = await createMember(req.body);
    const createdMemberDTO = toDTO(createdMember);
    res.send(createdMemberDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleUpdateMember(req: Request, res: Response): Promise<void> {
  try {
    const updatedMember = await updateMember(req.params.id, req.body);
    const updatedMemberDTO = toDTO(updatedMember);
    res.send(updatedMemberDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}

export async function handleDeleteMember(req: Request, res: Response): Promise<void> {
  try {
    const deletedMember = await deleteMember(req.params.id);
    const deletedMemberDTO = toDTO(deletedMember);
    res.send(deletedMemberDTO);
  } catch (error) {
    res.status(500).send(error);
  }
}
