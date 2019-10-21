import { SongDocument, SongModel, SongDTO } from './entity';

export async function getSongs(): Promise<SongDocument[]> {
  const songs = await SongModel.find();
  return songs;
}

export async function getSong(id: string): Promise<SongDocument> {
  const filter = { _id: id };
  const song = await SongModel.findOne(filter);
  return song;
}

export async function createSong(data: SongDTO): Promise<SongDocument> {
  const createdSong = await SongModel.create(data);
  return createdSong;
}

export async function updateSong(id: string, data: SongDTO): Promise<SongDocument> {
  const filter = { _id: id };
  const options = { new: true, upsert: true };
  const updatedSong = await SongModel.findOneAndUpdate(filter, data, options);
  return updatedSong;
}

export async function deleteSong(id: string): Promise<SongDocument> {
  const filter = { _id: id };
  const deletedSong = await SongModel.findOneAndDelete(filter);
  return deletedSong;
}
