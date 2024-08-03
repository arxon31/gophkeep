package mngconn

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(ctx context.Context, uri string, db string) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	go func() {
		disconnectCtx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				client.Disconnect(disconnectCtx)
			}
		}
	}()

	return client.Database(db), nil
}
