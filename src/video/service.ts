import VideoModel, { Video } from './model';
import { VideoDTO } from './DTO';

export async function getVideos(): Promise<Video[]> {
  const videos = await VideoModel.find();
  return videos;
}

export async function getVideo(id: string): Promise<Video> {
  const filter = { _id: id };
  const videos = await VideoModel.findOne(filter);
  return videos;
}

export async function createVideo(data: VideoDTO): Promise<Video> {
  const createdVideo = await VideoModel.create(data);
  return createdVideo;
}

export async function updateVideo(id: string, data: VideoDTO): Promise<Video> {
  const filter = { _id: id };
  const options = { new: true, upsert: true };
  const updatedVideo = await VideoModel.findOneAndUpdate(filter, data, options);
  return updatedVideo;
}

export async function deleteVideo(id: string): Promise<Video> {
  const filter = { _id: id };
  const deletedVideo = await VideoModel.findOneAndDelete(filter);
  return deletedVideo;
}
