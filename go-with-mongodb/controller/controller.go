package controller

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const colName = "watchlist"

var collection = *mongo.Collection

func init() {
	//client Options
	clientOptions := options.Client().ApplyURI(connectionString)
	mongo.Connect(context.TODO(), clientOptions)
}
