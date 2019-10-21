import { MerchandiseDocument, MerchandiseModel, MerchandiseDTO } from './entity';

export async function getMerchandises(): Promise<MerchandiseDocument[]> {
  const merchandises = await MerchandiseModel.find();
  return merchandises;
}

export async function getMerchandise(id: string): Promise<MerchandiseDocument> {
  const filter = { _id: id };
  const merchandise = await MerchandiseModel.findOne(filter);
  return merchandise;
}

export async function createMerchandise(data: MerchandiseDTO): Promise<MerchandiseDocument> {
  const createdMerchandise = await MerchandiseModel.create(data);
  return createdMerchandise;
}

export async function updateMerchandise(
  id: string,
  data: MerchandiseDTO,
): Promise<MerchandiseDocument> {
  const filter = { _id: id };
  const options = { new: true, upsert: true };
  const updatedMerchandise = await MerchandiseModel.findOneAndUpdate(filter, data, options);
  return updatedMerchandise;
}

export async function deleteMerchandise(id: string): Promise<MerchandiseDocument> {
  const filter = { _id: id };
  const deletedMerchandise = await MerchandiseModel.findOneAndDelete(filter);
  return deletedMerchandise;
}
