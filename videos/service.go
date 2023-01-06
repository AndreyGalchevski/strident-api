package videos

import (
	"github.com/AndreyGalchevski/strident-api/db"
)

func getVideos() ([]*db.Video, error) {
	videos, err := db.GetDB().Videos.List()

	if err != nil {
		return videos, err
	}

	return videos, nil
}

func getVideoByID(id string) (*db.Video, error) {
	video, err := db.GetDB().Videos.Retrieve(id)

	if err != nil {
		return video, err
	}

	return video, nil
}

func createVideo(params db.Video) (string, error) {
	newVideoID, err := db.GetDB().Videos.Create(&params)

	if err != nil {
		return "", err
	}

	return newVideoID, nil
}

func updateVideo(videoID string, params db.Video) (bool, error) {
	ok, err := db.GetDB().Videos.Update(videoID, &params)

	if err != nil {
		return false, err
	}

	return ok, nil
}

func deleteVideo(videoID string) (bool, error) {
	ok, err := db.GetDB().Videos.Delete(videoID)

	if err != nil {
		return false, err
	}

	return ok, nil
}
