package db

import (
    "context"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// Define a variable for the client. It is exposed to be used by other packages.
var Client *mongo.Client

// Init initializes the database connection.
func Init() error {
    var err error

    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

    // Connect to MongoDB
    Client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return err
    }

    // Check the connection
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = Client.Ping(ctx, nil)
    if err != nil {
        return err
    }

    log.Println("Connected to MongoDB!")
    return nil
}

// GetCollection returns a reference to a collection in the database.
func GetCollection(collectionName string) *mongo.Collection {
    return Client.Database("Master").Collection(collectionName)
}