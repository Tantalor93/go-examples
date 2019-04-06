package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		os.Exit(1)
	}

	collection := client.Database("local").Collection("numbers")
	collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		os.Exit(1)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(result)
	}

}
