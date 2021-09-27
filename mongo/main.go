package mongo

import (
	"context"
	"fmt"
	"time"

	databaseutils "github.com/lysofts/database-utils"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	UserCollectionName = "test_users"
)

type Database interface {
	GetOne(ctx context.Context, collectionName string, payload interface{}) (interface{}, error)
	Create(ctx context.Context, collectionName string, payload interface{}) (interface{}, error)
	Get(ctx context.Context, collectionName string, filter bson.M) (interface{}, error)
	Update(ctx context.Context, collectionName string, filter bson.M, payload interface{}) (interface{}, error)
	Delete(ctx context.Context, collectionName string, filer bson.M) (int64, error)
}

type DatabaseImpl struct {
	mongo.Client
	URL  string
	Name string
}

func New() Database {
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

	return &DatabaseImpl{
		Client: *client,
		URL:    url,
		Name:   name,
	}
}

//GetOne finds and returns exactly one object
func (d *DatabaseImpl) GetOne(ctx context.Context, collectionName string, payload interface{}) (interface{}, error) {
	collection := d.Database(d.Name).Collection(collectionName)

	result := collection.FindOne(ctx, payload)

	data := make(map[string]interface{})

	err := result.Decode(data)

	if err != nil {
		return nil, err
	}
	return data, nil
}

//Create creates an object in database
func (d *DatabaseImpl) Create(ctx context.Context, collectionName string, payload interface{}) (interface{}, error) {
	collection := d.Database(d.Name).Collection(collectionName)

	result, err := collection.InsertOne(ctx, payload)
	if err != nil {
		return nil, fmt.Errorf("creation error: %v", err)
	}

	return result, nil
}

//Get retrieves data from the database
func (d *DatabaseImpl) Get(ctx context.Context, collectionName string, filter bson.M) (interface{}, error) {
	collection := d.Database(d.Name).Collection(collectionName)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get error: %v", err)
	}

	defer cursor.Close(ctx)

	var data []interface{}
	for cursor.Next(ctx) {
		episode := make(map[string]interface{})
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		data = append(data, episode)
		fmt.Println(episode)
	}

	return data, nil
}

//Update updates the filtered result using provided data
func (d *DatabaseImpl) Update(ctx context.Context, collectionName string, filter bson.M, payload interface{}) (interface{}, error) {

	collection := d.Database(d.Name).Collection(collectionName)

	updateData := bson.M{"$set": payload}

	result, err := collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return nil, fmt.Errorf("update error: %v", err)
	}

	return result, nil
}

//Delete deletes all records matching the filter inside the collection
func (d *DatabaseImpl) Delete(ctx context.Context, collectionName string, filer bson.M) (int64, error) {

	collection := d.Database(d.Name).Collection(collectionName)

	result, err := collection.DeleteMany(ctx, filer)
	if err != nil {
		return 0, fmt.Errorf("delete error: %v", err)
	}

	return result.DeletedCount, nil
}
