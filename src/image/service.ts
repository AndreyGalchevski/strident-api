const cloudinaryUtil = require('../../utils/cloudinary');
const { resizeImage } = require('../../utils/imageManipulator');
const { deleteFile } = require('../../utils/general');

export async function uploadImage(filePath: string, folderName) {
  const resizedImagePath = await resizeImage(filePath);
  const imageUrl = await cloudinaryUtil.uploadImage(resizedImagePath, folderName);
  deleteFile(resizedImagePath);
  return imageUrl;
}

export async function destroyImage(imageUrl, folderName) {
  const response = await cloudinaryUtil.destroyImage(imageUrl, folderName);
  return response;
}
