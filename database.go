package databaseutils

import (
	"context"

	"github.com/lysofts/database-utils/mongo"
	"github.com/lysofts/database-utils/postres"
	"github.com/lysofts/database-utils/repository"
	"github.com/lysofts/database-utils/utils"
)

const (
	POSTGRES = "postgres"
	MONGO    = "mongo"
)

type Database struct {
	Choice repository.DatabaseUtil
}

func NewDatabase(choice string) repository.DatabaseUtil {
	db := Database{}
	switch choice {
	case POSTGRES:
		db.Choice = postres.New()
	case MONGO:
		db.Choice = mongo.New()
	default:
		db.Choice = mongo.New()
	}

	return &db
}

//Create creates an object in database
func (db *Database) Create(ctx context.Context, collectionName string, payload interface{}) (interface{}, error) {
	return db.Choice.Create(ctx, collectionName, payload)
}

//ReadOne finds and returns exactly one object
func (db *Database) ReadOne(ctx context.Context, collectionName string, query interface{}) (utils.Map, error) {
	return db.Choice.ReadOne(ctx, collectionName, query)
}

//Read retrieves data from the database
func (db *Database) Read(ctx context.Context, collectionName string, query interface{}) ([]utils.Map, error) {
	return db.Choice.Read(ctx, collectionName, query)
}

//Update updates the filtered result using provided data
func (db *Database) Update(ctx context.Context, collectionName string, query interface{}, payload interface{}) (interface{}, error) {
	return db.Choice.Update(ctx, collectionName, query, payload)
}

//Delete deletes all records matching the filter inside the collection
func (db *Database) Delete(ctx context.Context, collectionName string, query interface{}) (int64, error) {
	return db.Choice.Delete(ctx, collectionName, query)
}
