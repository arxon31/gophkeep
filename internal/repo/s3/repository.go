package s3

import (
	"context"
	"time"

	"github.com/arxon31/gophkeep/internal/models"
	"github.com/minio/minio-go/v7"
)

const (
	URLExpiration = time.Hour * 24 * 365
)

type repository struct {
	client *minio.Client
}

func NewS3Repo(client *minio.Client) *repository {
	return &repository{client: client}
}

func (r *repository) SaveFile(ctx context.Context, file *models.FileDTO) (url string, err error) {
	bucketName, err := r.createBucketIfNotExists(ctx, file.User)
	if err != nil {
		return "", err
	}

	_, err = r.client.PutObject(ctx, bucketName, file.Name, &file.Data, int64(file.Data.Cap()), minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}

	u, err := r.client.PresignedGetObject(ctx, bucketName, file.Name, URLExpiration, nil)
	if err != nil {
		return "", err
	}

	return u.String(), nil

}

func (r *repository) createBucketIfNotExists(ctx context.Context, bucketName string) (string, error) {
	exists, err := r.client.BucketExists(ctx, bucketName)
	if err != nil {
		return "", err
	}

	if !exists {
		err := r.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return "", err
		}
		return bucketName, nil
	}

	return bucketName, nil
}
