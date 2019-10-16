import { LyricDocument, LyricModel, LyricDTO } from './entity';

export async function getLyrics(): Promise<LyricDocument[]> {
  const lyrics = await LyricModel.find();
  return lyrics;
}

export async function getLyric(id: string): Promise<LyricDocument> {
  const filter = { _id: id };
  const lyrics = await LyricModel.findOne(filter);
  return lyrics;
}

export async function createLyric(data: LyricDTO): Promise<LyricDocument> {
  const createdLyric = await LyricModel.create(data);
  return createdLyric;
}

export async function updateLyric(id: string, data: LyricDTO): Promise<LyricDocument> {
  const filter = { _id: id };
  const options = { new: true, upsert: true };
  const updatedLyric = await LyricModel.findOneAndUpdate(filter, data, options);
  return updatedLyric;
}

export async function deleteLyric(id: string): Promise<LyricDocument> {
  const filter = { _id: id };
  const deletedLyric = await LyricModel.findOneAndDelete(filter);
  return deletedLyric;
}
