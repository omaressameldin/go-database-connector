package firebase

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/storage"
	"golang.org/x/net/context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

type Firebase struct {
	Collection string
	client     *firestore.Client
	storage    *storage.Client
	app        *firebase.App
}

func StartConnection(
	jsonConfig string,
	collection string,
	bucket string,
) (*Firebase, error) {
	opt := option.WithCredentialsFile(jsonConfig)
	cfg := &firebase.Config{}
	if bucket != "" {
		cfg.StorageBucket = bucket
	}
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, cfg, opt)
	if err != nil {
		return nil, createError(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, createError(err)
	}

	storage, err := app.Storage(ctx)
	if err != nil {
		return nil, createError(err)
	}

	return &Firebase{
		collection,
		client,
		storage,
		app,
	}, nil
}

func (f *Firebase) CloseConnection() error {
	return f.client.Close()
}
