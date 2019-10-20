import { GigDocument, GigModel, GigDTO } from './entity';

export async function getGigs(): Promise<GigDocument[]> {
  const sortBy = { date: -1 };
  const gigs = await GigModel.find().sort(sortBy);
  return gigs;
}

export async function getGig(id: string): Promise<GigDocument> {
  const filter = { _id: id };
  const gigs = await GigModel.findOne(filter);
  return gigs;
}

export async function createGig(data: GigDTO): Promise<GigDocument> {
  const createdGig = await GigModel.create(data);
  return createdGig;
}

export async function updateGig(id: string, data: GigDTO): Promise<GigDocument> {
  const filter = { _id: id };
  const options = { new: true, upsert: true };
  const updatedGig = await GigModel.findOneAndUpdate(filter, data, options);
  return updatedGig;
}

export async function deleteGig(id: string): Promise<GigDocument> {
  const filter = { _id: id };
  const deletedGig = await GigModel.findOneAndDelete(filter);
  return deletedGig;
}
