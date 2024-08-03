package text

import (
	"context"

	"github.com/arxon31/gophkeep/internal/models/requests"
	"github.com/arxon31/gophkeep/internal/models/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Database
}

func NewDocsRepo(db *mongo.Database) *repository {
	return &repository{db: db}
}

func (r *repository) SaveCredentials(ctx context.Context, dto requests.SaveCredentialsDTO) error {
	coll := r.db.Collection(dto.User)

	creds := bson.M{"meta": dto.Meta, "username": dto.UserName, "password": dto.Password}

	_, err := coll.InsertOne(ctx, creds)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) SaveBankCredentials(ctx context.Context, dto requests.SaveBankCredentialsDTO) error {
	coll := r.db.Collection(dto.User)

	creds := bson.M{"meta": dto.Meta, "card_number": dto.CardNumber, "owner": dto.Owner, "cvv": dto.CVV}

	_, err := coll.InsertOne(ctx, creds)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) SaveFileS3URL(ctx context.Context, dto requests.SaveFileS3URLDTO) error {
	coll := r.db.Collection(dto.User)

	url := bson.M{"meta": dto.Meta, "file_s3_url": dto.URL}

	_, err := coll.InsertOne(ctx, url)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetCredentials(ctx context.Context, dto requests.GetByMetaDTO) (resp *responses.GetCredentialsResponseDTO, err error) {
	coll := r.db.Collection(dto.User)

	meta := bson.M{"meta": dto.Meta}

	res := coll.FindOne(ctx, meta)

	err = res.Decode(&resp)
	if err != nil {
		return nil, err
	}

	return
}

func (r *repository) GetBankCredentials(ctx context.Context, dto requests.GetByMetaDTO) (resp *responses.GetBankCredentialsResponseDTO, err error) {
	coll := r.db.Collection(dto.User)

	meta := bson.M{"meta": dto.Meta}

	res := coll.FindOne(ctx, meta)

	err = res.Decode(&resp)
	if err != nil {
		return nil, err
	}

	return
}

func (r *repository) GetFileS3URL(ctx context.Context, dto requests.GetByMetaDTO) (resp *responses.GetS3FileURLDTO, err error) {
	coll := r.db.Collection(dto.User)

	meta := bson.M{"meta": dto.Meta}

	res := coll.FindOne(ctx, meta)

	err = res.Decode(&resp)
	if err != nil {
		return nil, err
	}

	return
}
