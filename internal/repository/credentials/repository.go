package credentials

import (
	"context"
	"github.com/arxon31/gophkeep/internal/repository/credentials/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	mongo *mongo.Database
}

func New(mongo *mongo.Database) *repository {
	return &repository{mongo: mongo}
}

func (r *repository) SaveCredentials(ctx context.Context, creds *model.Credentials) error {
	coll := r.mongo.Collection(creds.User)

	bsonCard := bson.M{"meta": creds.Meta, "username_hash": creds.UserNameHash, "username_salt": creds.UserNameSalt, "password_hash": creds.PasswordHash, "password_salt": creds.PasswordSalt}

	_, err := coll.InsertOne(ctx, bsonCard)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetCredentials(ctx context.Context, req *model.GetCredentials) (*model.Credentials, error) {
	coll := r.mongo.Collection(req.User)

	var creds *model.Credentials

	err := coll.FindOne(ctx, bson.M{"meta": req.Meta}).Decode(&creds)
	if err != nil {
		return nil, err
	}

	return creds, nil
}
