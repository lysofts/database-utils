package repository

import "context"

//DatabaseUtil blue print to database methods
type DatabaseUtil interface {
	//ReadOne find exactly one item matching the query
	ReadOne(ctx context.Context, collectionName string, query interface{}) (interface{}, error)

	//Create create an item/ add item to database
	Create(ctx context.Context, collectionName string, payload interface{}) (interface{}, error)

	//Read gets a list of items from database matching filter
	Read(ctx context.Context, collectionName string, filter interface{}) (interface{}, error)

	//Updates items in the database matching filter
	Update(ctx context.Context, collectionName string, query interface{}, payload interface{}) (interface{}, error)

	//deletes items in the database that matches the query
	Delete(ctx context.Context, collectionName string, query interface{}) (int64, error)
}
