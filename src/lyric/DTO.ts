import { Lyric } from './model';

export interface LyricDTO {
  id?: string;
  name: string;
  text: string;
}

export function toDTO(doc: Lyric): LyricDTO {
  const lyric: LyricDTO = {
    id: String(doc._id),
    name: doc.name,
    text: doc.text,
  };

  return lyric;
}
