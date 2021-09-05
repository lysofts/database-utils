package mongo_db

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
	ctx            context.Context
	collectionName string
}

func New(ctx context.Context, collectionName string) *Database {
	return &Database{
		ctx:            ctx,
		collectionName: collectionName,
	}
}

//Connect func connects client to authentication database
func (d *Database) Connect() (*mongo.Client, error) {

	databaseURL := databaseutils.GetEnv("AUTH_DATABASE_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(databaseURL))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return client, nil
}

//GetCollection is a  function makes a connection with a
//collection in the database
func (d *Database) GetCollection() *mongo.Collection {
	databaseName := databaseutils.GetEnv("AUTH_DATABASE_NAME")
	client, _ := d.Connect()
	collection := client.Database(databaseName).Collection(d.collectionName)
	return collection
}

//Create creates an object in database
func (d *Database) Create(ctx context.Context, collectionName string, data interface{}) (*mongo.InsertOneResult, error) {
	collection := d.GetCollection()

	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("creation error: %v", err)
	}

	return result, nil
}

//Get retrieves data from the database
func (d *Database) Get(ctx context.Context, collectionName string, filter bson.M) (*mongo.Cursor, error) {
	collection := d.GetCollection()

	result, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get error: %v", err)
	}

	return result, nil
}

//Update updates the filtered result using provided data
func (d *Database) Update(ctx context.Context, collectionName string, filter bson.M, data interface{}) (*mongo.UpdateResult, error) {
	updateData := bson.M{"$set": data}

	collection := d.GetCollection()

	result, err := collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return nil, fmt.Errorf("update error: %v", err)
	}

	return result, nil
}

//Delete deletes all records matching the filter inside the collection
func (d *Database) Delete(ctx context.Context, collectionName string, filer bson.M) (*mongo.DeleteResult, error) {

	collection := d.GetCollection()

	result, err := collection.DeleteMany(ctx, filer)
	if err != nil {
		return nil, fmt.Errorf("delete error: %v", err)
	}

	return result, nil
}
