import { Video } from './model';

export interface VideoDTO {
  id?: string;
  name: string;
  url: string;
  date: Date;
}

export function toDTO(doc: Video): VideoDTO {
  const video: VideoDTO = {
    id: String(doc._id),
    name: doc.name,
    url: doc.url,
    date: doc.date,
  };

  return video;
}
