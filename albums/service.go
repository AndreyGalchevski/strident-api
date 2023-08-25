package albums

import (
	"github.com/AndreyGalchevski/strident-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getAlbums() ([]*db.Album, error) {
	opts := options.Find().SetSort(bson.D{{Key: "year", Value: -1}})

	albums, err := db.GetDB().Albums.List(opts)

	if err != nil {
		return albums, err
	}

	return albums, nil
}

func getAlbumByID(id string) (*db.Album, error) {
	album, err := db.GetDB().Albums.Retrieve(id)

	if err != nil {
		return album, err
	}

	return album, nil
}

func createAlbum(params db.Album) (string, error) {
	newAlbumID, err := db.GetDB().Albums.Create(&params)

	if err != nil {
		return "", err
	}

	return newAlbumID, nil
}

func updateAlbum(id string, params db.Album) (bool, error) {
	ok, err := db.GetDB().Albums.Update(id, &params)

	if err != nil {
		return false, err
	}

	return ok, nil
}

func deleteAlbum(songID string) (bool, error) {
	ok, err := db.GetDB().Albums.Delete(songID)

	if err != nil {
		return false, err
	}

	return ok, nil
}
