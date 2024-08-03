package card

import (
	"context"
	"github.com/arxon31/gophkeep/internal/repository/card/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	mongo *mongo.Database
}

func New(mongo *mongo.Database) *repository {
	return &repository{mongo: mongo}
}

func (r *repository) SaveCard(ctx context.Context, card *model.Card) error {
	coll := r.mongo.Collection(card.User)

	bsonCard := bson.M{"meta": card.Meta, "owner": card.Owner, "number_hash": card.NumberHash, "number_salt": card.NumberSalt, "cvv_hash": card.CVVHash, "cvv_salt": card.CVVSalt}

	_, err := coll.InsertOne(ctx, bsonCard)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetCard(ctx context.Context, req *model.GetCard) (*model.Card, error) {
	coll := r.mongo.Collection(req.User)

	var card *model.Card

	err := coll.FindOne(ctx, bson.M{"meta": req.Meta}).Decode(&card)
	if err != nil {
		return nil, err
	}

	return card, nil
}
