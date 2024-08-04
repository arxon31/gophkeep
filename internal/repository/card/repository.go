package card

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/arxon31/gophkeep/internal/repository"
	"github.com/arxon31/gophkeep/internal/repository/card/model"
)

type repo struct {
	mongo *mongo.Database
}

func New(mongo *mongo.Database) *repo {
	return &repo{mongo: mongo}
}

func (r *repo) SaveCard(ctx context.Context, card *model.Card) error {
	coll := r.mongo.Collection(card.User)

	_, err := coll.InsertOne(ctx, card)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetCard(ctx context.Context, req *model.GetCard) (*model.Card, error) {
	coll := r.mongo.Collection(req.User)

	var card *model.Card

	err := coll.FindOne(ctx, bson.M{"meta": req.Meta}).Decode(&card)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, repository.ErrNotFound
		default:
			return nil, err
		}
	}

	return card, nil
}
