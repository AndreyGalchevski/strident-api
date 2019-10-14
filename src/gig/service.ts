import GigModel, { Gig } from './model';
import { GigDTO } from './DTO';

export async function getGigs(): Promise<Gig[]> {
  const gigs = await GigModel.find();
  return gigs;
}

export async function getGig(id: string): Promise<Gig> {
  const filter = { _id: id };
  const gigs = await GigModel.findOne(filter);
  return gigs;
}

export async function createGig(data: GigDTO): Promise<Gig> {
  const createdGig = await GigModel.create(data);
  return createdGig;
}

export async function updateGig(id: string, data: GigDTO): Promise<Gig> {
  const filter = { _id: id };
  const options = { new: true, upsert: true };
  const updatedGig = await GigModel.findOneAndUpdate(filter, data, options);
  return updatedGig;
}

export async function deleteGig(id: string): Promise<Gig> {
  const filter = { _id: id };
  const deletedGig = await GigModel.findOneAndDelete(filter);
  return deletedGig;
}
