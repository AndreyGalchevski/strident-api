import { VideoDocument, VideoModel, VideoDTO } from './entity';

export async function getVideos(): Promise<VideoDocument[]> {
  const videos = await VideoModel.find();
  return videos;
}

export async function getVideo(id: string): Promise<VideoDocument> {
  const filter = { _id: id };
  const videos = await VideoModel.findOne(filter);
  return videos;
}

export async function createVideo(data: VideoDTO): Promise<VideoDocument> {
  const createdVideo = await VideoModel.create(data);
  return createdVideo;
}

export async function updateVideo(id: string, data: VideoDTO): Promise<VideoDocument> {
  const filter = { _id: id };
  const options = { new: true, upsert: true };
  const updatedVideo = await VideoModel.findOneAndUpdate(filter, data, options);
  return updatedVideo;
}

export async function deleteVideo(id: string): Promise<VideoDocument> {
  const filter = { _id: id };
  const deletedVideo = await VideoModel.findOneAndDelete(filter);
  return deletedVideo;
}
