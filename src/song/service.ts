import SongModel, { Song } from './model';

export async function getSongs(): Promise<Song[]> {
  const songs = await SongModel.find();
  return songs;
}

export async function getSong(id: string): Promise<Song> {
  const filter = { _id: id };
  const songs = await SongModel.findOne(filter);
  return songs;
}

export async function createSong(data: Song): Promise<Song> {
  const createdSong = await SongModel.create(data);
  return createdSong;
}

export async function updateSong(id: string, data: Song): Promise<Song> {
  const filter = { _id: id };
  const options = { new: true, upsert: true };
  const updatedSong = await SongModel.findOneAndUpdate(filter, data, options);
  return updatedSong;
}

export async function deleteSong(id: string): Promise<Song> {
  const filter = { _id: id };
  const deletedSong = await SongModel.findOneAndDelete(filter);
  return deletedSong;
}
