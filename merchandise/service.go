package merchandise

import (
	"errors"
	"mime/multipart"

	"github.com/AndreyGalchevski/strident-api/db"
	"github.com/AndreyGalchevski/strident-api/images"
)

func getMerchandise() ([]*db.Merchandise, error) {
	merchandise, err := db.GetDB().Merchandise.List()

	if err != nil {
		return merchandise, err
	}

	return merchandise, nil
}

func getMerchandiseByID(id string) (*db.Merchandise, error) {
	merchandise, err := db.GetDB().Merchandise.Retrieve(id)

	if err != nil {
		return merchandise, err
	}

	return merchandise, nil
}

func createMerchandise(params db.Merchandise, image multipart.File) (string, error) {
	imageURL, uploadImageErr := images.UploadImage("merchandise", image)

	if uploadImageErr != nil {
		return "", uploadImageErr
	}

	params.Image = imageURL

	newMerchandiseID, createErr := db.GetDB().Merchandise.Create(&params)

	if createErr != nil {
		deleteImageErr := images.DeleteImage(imageURL)

		if deleteImageErr != nil {
			return "", deleteImageErr
		}

		return "", createErr
	}

	return newMerchandiseID, nil
}

func updateMerchandise(merchandiseID string, params db.Merchandise, image multipart.File) (bool, error) {
	if image != nil {
		merchandiseToUpdate, err := db.GetDB().Merchandise.Retrieve(merchandiseID)

		if err != nil {
			return false, err
		}

		err = images.DeleteImage(merchandiseToUpdate.Image)

		if err != nil {
			return false, errors.New("failed to delete the old merchandise image")
		}

		imageURL, err := images.UploadImage("merchandise", image)

		if err != nil {
			return false, errors.New("failed to upload the new merchandise image")
		}

		params.Image = imageURL
	}

	ok, err := db.GetDB().Merchandise.Update(merchandiseID, &params)

	if err != nil {
		if params.Image != "" {
			err := images.DeleteImage(params.Image)

			if err != nil {
				return false, errors.New("failed to delete new merchandise image")
			}
		}

		return false, err
	}

	if !ok {
		return false, nil
	}

	return true, nil
}

func deleteMerchandise(merchandiseID string) (bool, error) {
	merchandiseToDelete, err := db.GetDB().Merchandise.Retrieve(merchandiseID)

	if err != nil {
		return false, err
	}

	ok, err := db.GetDB().Merchandise.Delete(merchandiseID)

	if err != nil {
		return false, err
	}

	if !ok {
		return false, nil
	}

	err = images.DeleteImage(merchandiseToDelete.Image)

	if err != nil {
		return false, errors.New("failed to delete the merchandise image")
	}

	return true, nil
}
