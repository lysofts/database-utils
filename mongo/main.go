package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	databaseutils "github.com/lysofts/database-utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	UserCollectionName = "users"
)

type Database struct {
	mongo.Client
	URL  string
	Name string
}

func New() *Database {
	url := databaseutils.GetEnv("AUTH_DATABASE_URL")
	name := databaseutils.GetEnv("AUTH_DATABASE_NAME")

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		Client: *client,
		URL:    url,
		Name:   name,
	}
}

//Create creates an object in database
func (d *Database) Create(ctx context.Context, collectionName string, data interface{}) (*mongo.InsertOneResult, error) {
	collection := d.Database(d.Name).Collection(collectionName)

	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("creation error: %v", err)
	}

	return result, nil
}

//Get retrieves data from the database
func (d *Database) Get(ctx context.Context, collectionName string, filter bson.M) (*mongo.Cursor, error) {
	collection := d.Database(d.Name).Collection(collectionName)

	result, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get error: %v", err)
	}

	return result, nil
}

//Update updates the filtered result using provided data
func (d *Database) Update(ctx context.Context, collectionName string, filter bson.M, data interface{}) (*mongo.UpdateResult, error) {

	collection := d.Database(d.Name).Collection(collectionName)

	updateData := bson.M{"$set": data}

	result, err := collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return nil, fmt.Errorf("update error: %v", err)
	}

	return result, nil
}

//Delete deletes all records matching the filter inside the collection
func (d *Database) Delete(ctx context.Context, collectionName string, filer bson.M) (*mongo.DeleteResult, error) {

	collection := d.Database(d.Name).Collection(collectionName)

	result, err := collection.DeleteMany(ctx, filer)
	if err != nil {
		return nil, fmt.Errorf("delete error: %v", err)
	}

	return result, nil
}
