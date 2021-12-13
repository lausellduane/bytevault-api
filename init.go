package main

import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
    "fmt"
    "context"
)

var ctx = context.TODO()
var client *mongo.Client
var db *mongo.Database

// Connection URI
const uri = "mongodb://172.18.0.2:27017"
func init() {
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("client err: ", err)
		return
	}
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("ping err: ", err)
		return
	}
	fmt.Println("Successfully connected and pinged.")
	db = client.Database("BYTE_VAULT")
}