package images

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func uploadImage(folderName string, file multipart.File) (string, error) {
	cld, err := cloudinary.New()

	if err != nil {
		return "", err
	}

	ctx := context.Background()

	path := fmt.Sprintf("%s/%s/%s", "strident", os.Getenv("APP_ENV"), folderName)

	uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: path})

	if err != nil {
		return "", nil
	}

	return uploadResult.SecureURL, nil
}

func extractPublicID(imageURL string) string {
	splitStr := strings.Split(imageURL, "/")
	relevantPart := splitStr[len(splitStr)-4:]
	relevantPart[3] = strings.TrimSuffix(relevantPart[3], filepath.Ext(relevantPart[3]))
	return strings.Join(relevantPart, "/")
}

func deleteImage(imageURL string) error {
	cld, err := cloudinary.New()

	if err != nil {
		return err
	}

	ctx := context.Background()

	publicID := extractPublicID(imageURL)

	destroyResult, err := cld.Upload.Destroy(
		ctx,
		uploader.DestroyParams{PublicID: publicID})

	if err != nil {
		return err
	}

	if destroyResult.Result == "not found" {
		return errors.New("image not found")
	}

	return nil
}
