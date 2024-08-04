package credentials

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/arxon31/gophkeep/internal/repository"
	"github.com/arxon31/gophkeep/internal/repository/credentials/dto"
)

type repo struct {
	mongo *mongo.Database
}

func New(mongo *mongo.Database) *repo {
	return &repo{mongo: mongo}
}

func (r *repo) SaveCredentials(ctx context.Context, creds *dto.Credentials) error {
	coll := r.mongo.Collection(creds.User)

	_, err := coll.InsertOne(ctx, creds)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetCredentials(ctx context.Context, req *dto.GetCredentials) (*dto.Credentials, error) {
	coll := r.mongo.Collection(req.User)

	var creds *dto.Credentials

	err := coll.FindOne(ctx, bson.M{"meta": req.Meta}).Decode(&creds)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, repository.ErrNotFound
		default:
			return nil, err
		}
	}

	return creds, nil
}
