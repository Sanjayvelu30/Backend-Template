package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DBHelper struct {
	database    *mongo.Database
	Collection1 *mongo.Collection
}

type DBHelperProvider interface {
}

func NewDBHelper(database *mongo.Database, Collection1 *mongo.Collection) *DBHelper {
	return &DBHelper{
		database:    *&database,
		Collection1: Collection1,
	}
}
