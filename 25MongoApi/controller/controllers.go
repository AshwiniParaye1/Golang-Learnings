package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://ashwini:ashwini@mycluster01.39rec71.mongodb.net/?retryWrites=true&w=majority&appName=MyCluster01"

const dbName = "netflix"

const colName = "watchlist"

// most imp
var collection *mongo.Collection

//connect with MongoDB

func init() {
	//client options
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success!")

	connection := client.Database(dbName).Collection(colName)

	//collection instance/reference
	fmt.Println("Collection instance is ready!")

}
