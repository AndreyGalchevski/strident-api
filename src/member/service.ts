import { MemberDocument, MemberModel, MemberDTO } from './entity';

export async function getMembers(): Promise<MemberDocument[]> {
  const members = await MemberModel.find();
  return members;
}

export async function getMember(id: string): Promise<MemberDocument> {
  const filter = { _id: id };
  const members = await MemberModel.findOne(filter);
  return members;
}

export async function createMember(data: MemberDTO): Promise<MemberDocument> {
  const createdMember = await MemberModel.create(data);
  return createdMember;
}

export async function updateMember(id: string, data: MemberDTO): Promise<MemberDocument> {
  const filter = { _id: id };
  const options = { new: true, upsert: true };
  const updatedMember = await MemberModel.findOneAndUpdate(filter, data, options);
  return updatedMember;
}

export async function deleteMember(id: string): Promise<MemberDocument> {
  const filter = { _id: id };
  const deletedMember = await MemberModel.findOneAndDelete(filter);
  return deletedMember;
}
