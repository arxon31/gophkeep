package attachment

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/arxon31/gophkeep/internal/repository"
	"github.com/arxon31/gophkeep/internal/repository/attachment/model"
)

type repo struct {
	s3    *minio.Client
	mongo *mongo.Database
}

func New(s3 *minio.Client, mongo *mongo.Database) *repo {
	return &repo{s3: s3, mongo: mongo}
}

func (r *repo) SaveAttachment(ctx context.Context, attachmentInfo *model.Attachment) error {
	bucketName, err := r.createBucketIfNotExists(ctx, attachmentInfo.User)
	if err != nil {
		return err
	}

	_, err = r.s3.PutObject(ctx, bucketName, attachmentInfo.Name, attachmentInfo.Content, int64(attachmentInfo.Content.Cap()), minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	coll := r.mongo.Collection(attachmentInfo.User)

	info := bson.M{"meta": attachmentInfo.Meta, "attachment_name": attachmentInfo.Name}

	_, err = coll.InsertOne(ctx, info)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetAttachment(ctx context.Context, req *model.GetAttachment) (*model.Attachment, error) {
	exists, err := r.isBucketExists(ctx, req.User)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("%w:%s", ErrBucketNotExists, req.User)
	}

	coll := r.mongo.Collection(req.User)

	filter := bson.M{"meta": req.Meta}

	res := coll.FindOne(ctx, filter)

	var attachment *model.Attachment

	err = res.Decode(&attachment)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, repository.ErrNotFound
		default:
			return nil, err
		}
	}

	obj, err := r.s3.GetObject(ctx, req.User, attachment.Name, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	stat, err := obj.Stat()
	if err != nil {
		return nil, err
	}

	_, err = io.CopyN(attachment.Content, obj, stat.Size)
	if err != nil {
		return nil, err
	}

	return attachment, nil
}

func (r *repo) createBucketIfNotExists(ctx context.Context, bucketName string) (string, error) {
	exists, err := r.isBucketExists(ctx, bucketName)
	if err != nil {
		return "", err
	}

	if !exists {
		err = r.s3.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return "", err
		}
		return bucketName, nil
	}

	return bucketName, nil
}

func (r *repo) isBucketExists(ctx context.Context, bucketName string) (bool, error) {
	return r.s3.BucketExists(ctx, bucketName)
}
