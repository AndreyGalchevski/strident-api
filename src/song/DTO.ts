import { Song } from './model';

export interface SongDTO {
  id?: string;
  name: string;
  album: string;
  url: string;
}

export function toDTO(doc: Song): SongDTO {
  const video: SongDTO = {
    id: String(doc._id),
    name: doc.name,
    album: doc.album,
    url: doc.url,
  };

  return video;
}
