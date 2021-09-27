package databaseutils

import (
	"github.com/lysofts/database-utils/mongo"
	"github.com/lysofts/database-utils/postres"
	"github.com/lysofts/database-utils/repository"
)

const (
	POSTGRES = "postgres"
	MONGO    = "mongo"
)

func NewDatabase(choice string) repository.DatabaseUtil {
	switch choice {
	case POSTGRES:
		return postres.New()
	case MONGO:
		return mongo.New()
	default:
		return mongo.New()
	}
}
