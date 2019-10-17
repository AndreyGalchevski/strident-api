const imageService = require('./imageService');
const { GoodResponse, BadResponse } = require('../../utils/httpResponses');

async function uploadImage(req, res) {
  const files = Object.values(req.files);
  const numberOfFiles = files.length;
  const folderName = `lappa/${req.query.folder}`;
  if (numberOfFiles === 1) {
    const filePath = files[0].path;
    const imageUrl = await imageService.uploadSingleImage(filePath, folderName);
    return res.send(GoodResponse(imageUrl));
  }
  if (numberOfFiles > 1) {
    const filePaths = files.map(file => file.path);
    const imageUrls = await imageService.uploadMultipleImages(filePaths, folderName);
    return res.send(GoodResponse(imageUrls));
  }
  return res.status(400).send(BadResponse({ error: 'At least one image is required' }));
}

async function destroyImage(req, res) {
  const numberOfImages = req.body.imageUrls.length;
  const folderName = `lappa/${req.query.folder}`;
  if (numberOfImages === 1) {
    const destroyedImage = await imageService.destroySingleImage(req.body.imageUrls[0], folderName);
    return res.send(GoodResponse(destroyedImage));
  }
  if (numberOfImages > 1) {
    const destroyedImages = await imageService.destroyMultipleImages(
      req.body.imageUrls,
      folderName,
    );
    return res.send(GoodResponse(destroyedImages));
  }
  return res.status(400).send(BadResponse({ error: 'At least one image is required' }));
}

module.exports = {
  uploadImage,
  destroyImage,
};
