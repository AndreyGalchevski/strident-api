package songs

import (
	"github.com/AndreyGalchevski/strident-api/db"
)

func getSongs() ([]*db.Song, error) {
	songs, err := db.GetDB().Songs.List(nil)

	if err != nil {
		return songs, err
	}

	return songs, nil
}

func getSongByID(id string) (*db.Song, error) {
	song, err := db.GetDB().Songs.Retrieve(id)

	if err != nil {
		return song, err
	}

	return song, nil
}

func createSong(params db.Song) (string, error) {
	newSongID, err := db.GetDB().Songs.Create(&params)

	if err != nil {
		return "", err
	}

	return newSongID, nil
}

func updateSong(songID string, params db.Song) (bool, error) {
	ok, err := db.GetDB().Songs.Update(songID, &params)

	if err != nil {
		return false, err
	}

	return ok, nil
}

func deleteSong(songID string) (bool, error) {
	ok, err := db.GetDB().Songs.Delete(songID)

	if err != nil {
		return false, err
	}

	return ok, nil
}
