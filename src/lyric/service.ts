import LyricModel, { Lyric } from './model';

export async function getLyrics(): Promise<Lyric[]> {
  const lyrics = await LyricModel.find();
  return lyrics;
}

export async function getLyric(id: string): Promise<Lyric> {
  const filter = { _id: id };
  const lyrics = await LyricModel.findOne(filter);
  return lyrics;
}

export async function createLyric(data: Lyric): Promise<Lyric> {
  const createdLyric = await LyricModel.create(data);
  return createdLyric;
}

export async function updateLyric(id: string, data: Lyric): Promise<Lyric> {
  const filter = { _id: id };
  const options = { new: true, upsert: true };
  const updatedLyric = await LyricModel.findOneAndUpdate(filter, data, options);
  return updatedLyric;
}

export async function deleteLyric(id: string): Promise<Lyric> {
  const filter = { _id: id };
  const deletedLyric = await LyricModel.findOneAndDelete(filter);
  return deletedLyric;
}
