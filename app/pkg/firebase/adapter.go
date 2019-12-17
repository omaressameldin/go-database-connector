package firebase

import (
	"fmt"
	"log"

	"github.com/omaressameldin/go-database-connector/app/pkg/database"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

// Create Firebase implementation of database connector create
func (f *Firebase) Create(
	validators []database.Validator,
	key string,
	data interface{},
) error {
	return addToFirebase(
		f.Collection,
		key,
		validators,
		func() error {
			_, err := f.client.Collection(f.Collection).Doc(key).Set(
				context.Background(),
				data,
			)
			return err
		},
	)
}

// Update Firebase implementation of database connector update
func (f *Firebase) Update(
	validators []database.Validator,
	key string,
	data []database.Updated,
) error {
	fmt.Println(generateFirestoreUpdate(data))
	return addToFirebase(
		f.Collection,
		key,
		validators,
		func() error {
			_, err := f.client.Collection(f.Collection).Doc(key).Update(
				context.Background(),
				generateFirestoreUpdate(data),
			)
			return err
		},
	)
}

// Read Firebase implementation of database connector read
func (f *Firebase) Read(key string, model interface{}) error {
	var err error
	docSnap, err := f.client.Collection(f.Collection).Doc(key).Get(context.Background())
	if err == nil {
		err = docSnap.DataTo(model)
	}
	if err != nil {
		return database.GenerateJsonError(database.ValidationError{
			Field:   "FIREBASE",
			Message: err.Error(),
		})
	}

	return nil
}

// Delete Firebase implementation of database connector delete
func (f *Firebase) Delete(key string) error {
	_, err := f.client.Collection(f.Collection).Doc(key).Delete(context.Background())
	if err != nil {
		return database.GenerateJsonError(database.ValidationError{
			Field:   "FIREBASE",
			Message: err.Error(),
		})
	}

	return nil
}

// ReadAll Firebase implementation of database connector read all
func (f *Firebase) ReadAll(
	genRefFn func() interface{},
	appendFn func(interface{}),
	filters []database.Filter,
) error {

	query := f.client.Collection(f.Collection).Query
	for _, f := range filters {

		query = query.Where(f.Field, string(f.Operator), f.Value)
	}

	docs := query.Documents(context.Background())

	for {
		docSnap, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return database.GenerateJsonError(database.ValidationError{
				Field:   "FIREBASE",
				Message: err.Error(),
			})
		}

		recordRef := genRefFn()
		err = docSnap.DataTo(recordRef)
		if err != nil {
			return database.GenerateJsonError(database.ValidationError{
				Field:   "FIREBASE",
				Message: err.Error(),
			})
		}
		appendFn(recordRef)
	}
	return nil
}

func (f *Firebase) Authenticate(token string) error {
	ctx := context.Background()
	client, err := f.app.Auth(ctx)
	if err != nil {
		return err
	}

	payload, err := client.VerifyIDToken(ctx, token)
	if err != nil {
		return err
	}

	log.Println(payload)
	return nil
}
