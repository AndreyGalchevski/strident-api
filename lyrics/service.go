package lyrics

import (
	"github.com/AndreyGalchevski/strident-api/db"
)

func getLyrics() ([]*db.Lyric, error) {
	lyrics, err := db.GetDB().Lyrics.List()

	if err != nil {
		return lyrics, err
	}

	return lyrics, nil
}

func getLyricByID(id string) (*db.Lyric, error) {
	lyric, err := db.GetDB().Lyrics.Retrieve(id)

	if err != nil {
		return lyric, err
	}

	return lyric, nil
}

func createLyric(params db.Lyric) (string, error) {
	newLyricID, err := db.GetDB().Lyrics.Create(&params)

	if err != nil {
		return "", err
	}

	return newLyricID, nil
}

func updateLyric(lyricID string, params db.Lyric) (bool, error) {
	ok, err := db.GetDB().Lyrics.Update(lyricID, &params)

	if err != nil {
		return false, err
	}

	return ok, nil
}

func deleteLyric(lyricID string) (bool, error) {
	ok, err := db.GetDB().Lyrics.Delete(lyricID)

	if err != nil {
		return false, err
	}

	return ok, nil
}
