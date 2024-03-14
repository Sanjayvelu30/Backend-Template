package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *DBHelper {
	var dbHelper *DBHelper
	optionsClient := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx := context.Background()
	databaseClient, err := mongo.Connect(ctx, optionsClient)
	if err != nil {
		return dbHelper
	}
	database := databaseClient.Database("WedEase")
	Collection1 := database.Collection("Collection")
	dbHelper = NewDBHelper(database, Collection1)

	return dbHelper
}
