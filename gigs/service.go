package gigs

import (
	"errors"
	"mime/multipart"

	"github.com/AndreyGalchevski/strident-api/db"
	"github.com/AndreyGalchevski/strident-api/images"
)

func getGigs() ([]*db.Gig, error) {
	gigs, err := db.GetDB().Gigs.List()

	if err != nil {
		return gigs, err
	}

	return gigs, nil
}

func getGigByID(id string) (*db.Gig, error) {
	gig, err := db.GetDB().Gigs.Retrieve(id)

	if err != nil {
		return gig, err
	}

	return gig, nil
}

func createGig(params db.Gig, image multipart.File) (string, error) {
	imageURL, uploadImageErr := images.UploadImage("gigs", image)

	if uploadImageErr != nil {
		return "", uploadImageErr
	}

	params.Image = imageURL

	newGigID, createErr := db.GetDB().Gigs.Create(&params)

	if createErr != nil {
		deleteImageErr := images.DeleteImage(imageURL)

		if deleteImageErr != nil {
			return "", deleteImageErr
		}

		return "", createErr
	}

	return newGigID, nil
}

func updateGig(gigID string, params db.Gig, image multipart.File) (bool, error) {

	if image != nil {
		gigToUpdate, err := db.GetDB().Gigs.Retrieve(gigID)

		if err != nil {
			return false, err
		}

		err = images.DeleteImage(gigToUpdate.Image)

		if err != nil {
			return false, errors.New("failed to delete the old gig image")
		}

		imageURL, err := images.UploadImage("gigs", image)

		if err != nil {
			return false, errors.New("failed to upload the new gig image")
		}

		params.Image = imageURL
	}

	ok, err := db.GetDB().Gigs.Update(gigID, &params)

	if err != nil {
		if params.Image != "" {
			err := images.DeleteImage(params.Image)

			if err != nil {
				return false, errors.New("failed to delete new gig image")
			}
		}

		return false, err
	}

	if !ok {
		return false, nil
	}

	return true, nil
}

func deleteGig(gigID string) (bool, error) {
	gigToDelete, err := db.GetDB().Gigs.Retrieve(gigID)

	if err != nil {
		return false, err
	}

	ok, err := db.GetDB().Gigs.Delete(gigID)

	if err != nil {
		return false, err
	}

	if !ok {
		return false, nil
	}

	err = images.DeleteImage(gigToDelete.Image)

	if err != nil {
		return false, errors.New("failed to delete the gig image")
	}

	return true, nil
}
