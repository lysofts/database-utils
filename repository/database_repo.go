package repository

import (
	"context"

	"github.com/lysofts/database-utils/utils"
)

//DatabaseUtil blue print to database methods
type DatabaseUtil interface {
	//Create create an item/ add item to database
	Create(ctx context.Context, table utils.DatabaseTable, payload interface{}) (interface{}, error)

	//ReadOne find exactly one item matching the query
	ReadOne(ctx context.Context, table utils.DatabaseTable, query interface{}) (utils.Map, error)

	//Read gets a list of items from database matching filter
	Read(ctx context.Context, table utils.DatabaseTable, filter interface{}) ([]utils.Map, error)

	//Updates items in the database matching filter
	Update(ctx context.Context, table utils.DatabaseTable, query interface{}, payload interface{}) (interface{}, error)

	//deletes items in the database that matches the query
	Delete(ctx context.Context, table utils.DatabaseTable, query interface{}) (int64, error)
}
