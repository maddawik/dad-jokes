package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Create a Client to a MongoDB server and use Ping to verify that the
// server is running.
func testDB() {
	mongodb_uri := os.Getenv("MONGODB_URI")
	mongodb_username := os.Getenv("MONGODB_USERNAME")
	mongodb_password := os.Getenv("MONGODB_PASSWORD")

	clientCredential := options.Credential{
		Username: mongodb_username,
		Password: mongodb_password,
	}
	clientOpts := options.Client().ApplyURI(mongodb_uri).SetAuth(clientCredential)

	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// Call Ping to verify that the deployment is up and the Client was
	// configured successfully. As mentioned in the Ping documentation, this
	// reduces application resiliency as the server may be temporarily
	// unavailable when Ping is called.
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatalf("Couldn't connect to the database, %v\n", err)
	} else {
		fmt.Println("Database is up!")
	}

	// Get all databases and list them
	result, err := client.ListDatabaseNames(
		context.TODO(),
		bson.D{{"empty", false}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All Databases")
	for _, db := range result {
		fmt.Println(db)
	}
}
