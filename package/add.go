package _package

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	url     string
	db      string
	collect string
}
type User struct {
	Name string `bson:"name"`
	Par  string `bson:"par"`
}

var database Data = Data{
	url:     "mongodb://localhost:27017",
	db:      "OSamidb",
	collect: "user",
}

func add(document bson.M) {
	url := database.url
	collect := database.collect
	db := database.db
	clientOptions := options.Client().ApplyURI(url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)

	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Error pinging MongoDB:", err)

	}

	// Access the database
	database := client.Database(db)

	// Access the collection
	collection := database.Collection(collect)

	// Define a document

	// Insert the document
	result, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		fmt.Println("Error inserting document:", err)
	}

	// Print the ID of the inserted document
	fmt.Println("Inserted document ID:", result.InsertedID)
}
