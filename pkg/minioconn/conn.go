package minioconn

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func New(uri, user, password string) (*minio.Client, error) {
	return minio.New(uri, &minio.Options{
		Creds:  credentials.NewStaticV4(user, password, ""),
		Secure: false,
	})
}
