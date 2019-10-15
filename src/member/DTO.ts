import { Member } from './model';

export interface MemberDTO {
  id?: string;
  name: string;
  instrument: string;
  info: string;
  image: string;
  imageNG: string;
}

export function toDTO(doc: Member): MemberDTO {
  const member: MemberDTO = {
    id: String(doc._id),
    name: doc.name,
    instrument: doc.instrument,
    info: doc.info,
    image: doc.image,
    imageNG: doc.imageNG,
  };

  return member;
}
