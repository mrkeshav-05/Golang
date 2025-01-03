package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    const uri = "mongodb+srv://Keshav:fZGQvGP7vRzgVeZj@neednear1.bm0dook.mongodb.net"
    const dbName = "learningGo"
    const collectionName = "exampleCollection"

    clientOptions := options.Client().ApplyURI(uri)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    defer func() {
        if err := client.Disconnect(ctx); err != nil {
            log.Fatalf("Failed to disconnect MongoDB: %v", err)
        }
    }()

    db := client.Database(dbName)
    collection := db.Collection(collectionName)

    // Insert a document
    doc := bson.D{
        {Key: "name", Value: "Keshav"},
        {Key: "age", Value: 25},
        {Key: "location", Value: "Lucknow"},
    }
    _, err = collection.InsertOne(ctx, doc)
    if err != nil {
        log.Fatalf("Failed to insert document: %v", err)
    }

    fmt.Printf("Document inserted into collection %q in database %q\n", collectionName, dbName)

    // List collections again
    collections, err := db.ListCollectionNames(ctx, bson.M{})
    if err != nil {
        log.Fatalf("Failed to list collections: %v", err)
    }
    fmt.Printf("Collections in database %q: %v\n", dbName, collections)
}
