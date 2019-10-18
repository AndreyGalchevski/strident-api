import cloudinary from 'cloudinary';
import { promises as fs } from 'fs';

function configureCloudinary(): void {
  // @ts-ignore
  cloudinary.config({
    cloud_name: process.env.CLOUDINARY_NAME,
    api_key: process.env.CLOUDINARY_API_KEY,
    api_secret: process.env.CLOUDINARY_API_SECRET,
  });
}

export async function uploadImage(filePath: string, publicID: string): Promise<string> {
  configureCloudinary();
  // @ts-ignore
  const result = await cloudinary.v2.uploader.upload(filePath, {
    public_id: publicID,
  });
  fs.unlink(filePath);
  return result.secure_url;
}

export async function destroyImage(imageUrl: string, folderName: string): Promise<unknown> {
  configureCloudinary();
  const publicId = `${folderName}/${imageUrl.substr(-24, 20)}`;
  // @ts-ignore
  const response = await cloudinary.v2.uploader.destroy(publicId);
  return response;
}
