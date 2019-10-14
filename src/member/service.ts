import MemberModel, { Member } from './model';
import { MemberDTO } from './DTO';

export async function getMembers(): Promise<Member[]> {
  const members = await MemberModel.find();
  return members;
}

export async function getMember(id: string): Promise<Member> {
  const filter = { _id: id };
  const members = await MemberModel.findOne(filter);
  return members;
}

export async function createMember(data: MemberDTO): Promise<Member> {
  const createdMember = await MemberModel.create(data);
  return createdMember;
}

export async function updateMember(id: string, data: MemberDTO): Promise<Member> {
  const filter = { _id: id };
  const options = { new: true, upsert: true };
  const updatedMember = await MemberModel.findOneAndUpdate(filter, data, options);
  return updatedMember;
}

export async function deleteMember(id: string): Promise<Member> {
  const filter = { _id: id };
  const deletedMember = await MemberModel.findOneAndDelete(filter);
  return deletedMember;
}
