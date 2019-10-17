import cloudinary from 'cloudinary';

// @ts-ignore
cloudinary.config({
  cloud_name: process.env.CLOUDINARY_NAME,
  api_key: process.env.CLOUDINARY_API_KEY,
  api_secret: process.env.CLOUDINARY_API_SECRET,
});

export async function uploadImage(imagePath: string, folderName: string): Promise<string> {
  // @ts-ignore
  const image = await cloudinary.v2.uploader.upload(imagePath, {
    folder: folderName,
  });
  return image.secure_url;
}

export async function destroyImage(imageUrl: string, folderName: string): Promise<unknown> {
  const publicId = `${folderName}/${imageUrl.substr(-24, 20)}`;
  // @ts-ignore
  const response = await cloudinary.v2.uploader.destroy(publicId);
  return response;
}
