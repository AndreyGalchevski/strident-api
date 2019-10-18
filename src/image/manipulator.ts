import { parse, resolve } from 'path';
import sharp from 'sharp';

export async function convertToWebP(imagePath: string): Promise<string> {
  const imageName = parse(imagePath).name;
  const convertedImagePath = `${resolve(__dirname)}/${imageName}_ng.webp`;

  await sharp(imagePath)
    .webp()
    .toFile(convertedImagePath);

  return convertedImagePath;
}
