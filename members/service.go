package members

import (
	"errors"
	"mime/multipart"

	"github.com/AndreyGalchevski/strident-api/db"
	"github.com/AndreyGalchevski/strident-api/images"
)

func getMembers() ([]*db.Member, error) {
	members, err := db.GetDB().Members.List(nil)

	if err != nil {
		return members, err
	}

	return members, nil
}

func getMemberByID(id string) (*db.Member, error) {
	member, err := db.GetDB().Members.Retrieve(id)

	if err != nil {
		return member, err
	}

	return member, nil
}

func createMember(params db.Member, image multipart.File) (string, error) {
	imageURL, uploadImageErr := images.UploadImage("members", image)

	if uploadImageErr != nil {
		return "", uploadImageErr
	}

	params.Image = imageURL

	newMemberID, createErr := db.GetDB().Members.Create(&params)

	if createErr != nil {
		deleteImageErr := images.DeleteImage(imageURL)

		if deleteImageErr != nil {
			return "", deleteImageErr
		}

		return "", createErr
	}

	return newMemberID, nil
}

func updateMember(memberID string, params db.Member, image multipart.File) (bool, error) {
	if image != nil {
		memberToUpdate, err := db.GetDB().Members.Retrieve(memberID)

		if err != nil {
			return false, err
		}

		err = images.DeleteImage(memberToUpdate.Image)

		if err != nil {
			return false, errors.New("failed to delete the old member image")
		}

		imageURL, err := images.UploadImage("members", image)

		if err != nil {
			return false, errors.New("failed to upload the new member image")
		}

		params.Image = imageURL
	}

	ok, err := db.GetDB().Members.Update(memberID, &params)

	if err != nil {
		if params.Image != "" {
			err := images.DeleteImage(params.Image)

			if err != nil {
				return false, errors.New("failed to delete new member image")
			}
		}

		return false, err
	}

	if !ok {
		return false, nil
	}

	return true, nil
}

func deleteMember(memberD string) (bool, error) {
	memberToDelete, err := db.GetDB().Members.Retrieve(memberD)

	if err != nil {
		return false, err
	}

	ok, err := db.GetDB().Members.Delete(memberD)

	if err != nil {
		return false, err
	}

	if !ok {
		return false, nil
	}

	err = images.DeleteImage(memberToDelete.Image)

	if err != nil {
		return false, errors.New("failed to delete the member image")
	}

	return true, nil
}
