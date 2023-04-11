package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

func main() {
	//MongoDB connection
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Close the client
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Access database
	collection := client.Database("test").Collection("users")

	// Insert a new user
	user := User{Name: "John Smith", Email: "john.smith@example.com"}
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	//Insert many users
	newuser := []interface{}{
		User{Name: "rob pike", Email: "robby19@example.com"},
		User{Name: "bob jordan", Email: "bob34@example.com"},
		User{Name: "rock ding", Email: "rocky67@example.com"},
	}
	_, err = collection.InsertMany(ctx, newuser)

	// Find a user by name
	filter := bson.M{"name": "John Smith"}
	var result User
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)

	// Update a user's email
	update := bson.M{"$set": bson.M{"email": "john.smith@newexample.com"}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	//Delete a user by name
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
}
