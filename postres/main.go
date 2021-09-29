package postres

import (
	"context"

	"github.com/lysofts/database-utils/repository"
	"github.com/lysofts/database-utils/utils"
)

const (
	Usertable = "test_users"
)

type DatabaseImpl struct {
}

func New() repository.DatabaseUtil {
	return &DatabaseImpl{}
}

//Create creates an object in database
func (d *DatabaseImpl) Create(ctx context.Context, table utils.DatabaseTable, payload interface{}) (interface{}, error) {
	//TODO add implementation
	return nil, nil
}

//ReadOne finds and returns exactly one object
func (d *DatabaseImpl) ReadOne(ctx context.Context, table utils.DatabaseTable, query interface{}) (utils.Map, error) {
	//TODO add implementation
	return nil, nil
}

//Read retrieves data from the database
func (d *DatabaseImpl) Read(ctx context.Context, table utils.DatabaseTable, query interface{}) ([]utils.Map, error) {
	//TODO add implementation
	return nil, nil
}

//Update updates the filtered result using provided data
func (d *DatabaseImpl) Update(ctx context.Context, table utils.DatabaseTable, query interface{}, payload interface{}) (interface{}, error) {
	//TODO add implementation
	return nil, nil
}

//Delete deletes all records matching the filter inside the collection
func (d *DatabaseImpl) Delete(ctx context.Context, table utils.DatabaseTable, query interface{}) (int64, error) {
	//TODO add implementation
	return 0, nil
}
