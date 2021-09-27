package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/lysofts/database-utils/repository"
	"github.com/lysofts/database-utils/utils"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	UserCollectionName = "test_users"
)

type DatabaseImpl struct {
	mongo.Client
	URL  string
	Name string
}

func New() repository.DatabaseUtil {
	url := utils.GetEnv("AUTH_DATABASE_URL")
	name := utils.GetEnv("AUTH_DATABASE_NAME")

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

//Create creates an object in database
func (d *DatabaseImpl) Create(ctx context.Context, collectionName string, payload interface{}) (interface{}, error) {
	collection := d.Database(d.Name).Collection(collectionName)

	result, err := collection.InsertOne(ctx, payload)
	if err != nil {
		return nil, fmt.Errorf("creation error: %v", err)
	}

	return result, nil
}

//ReadOne finds and returns exactly one object
func (d *DatabaseImpl) ReadOne(ctx context.Context, collectionName string, query interface{}) (utils.Map, error) {
	collection := d.Database(d.Name).Collection(collectionName)

	result := collection.FindOne(ctx, query)

	data := make(map[string]interface{})

	err := result.Decode(data)

	if err != nil {
		return nil, err
	}

	p, err := utils.BsonToJson(data)
	if err != nil {
		return nil, err
	}
	return p, nil
}

//Read retrieves data from the database
func (d *DatabaseImpl) Read(ctx context.Context, collectionName string, query interface{}) ([]utils.Map, error) {
	collection := d.Database(d.Name).Collection(collectionName)

	cursor, err := collection.Find(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get error: %v", err)
	}

	defer cursor.Close(ctx)

	var data []utils.Map
	for cursor.Next(ctx) {
		episode := utils.Map{}
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}

		p, err := utils.BsonToJson(episode)
		if err != nil {
			return nil, err
		}
		data = append(data, p)
	}

	return data, nil
}

//Update updates the filtered result using provided data
func (d *DatabaseImpl) Update(ctx context.Context, collectionName string, query interface{}, payload interface{}) (interface{}, error) {

	collection := d.Database(d.Name).Collection(collectionName)

	q, err := utils.JsonToBson(query)
	if err != nil {
		return nil, err
	}

	updateData := bson.M{"$set": payload}

	result, err := collection.UpdateOne(ctx, q, updateData)
	if err != nil {
		return nil, fmt.Errorf("update error: %v", err)
	}

	return result, nil
}

//Delete deletes all records matching the filter inside the collection
func (d *DatabaseImpl) Delete(ctx context.Context, collectionName string, query interface{}) (int64, error) {
	collection := d.Database(d.Name).Collection(collectionName)

	q, err := utils.JsonToBson(query)
	if err != nil {
		return 0, err
	}

	result, err := collection.DeleteMany(ctx, q)
	if err != nil {
		return 0, fmt.Errorf("delete error: %v", err)
	}

	return result.DeletedCount, nil
}
