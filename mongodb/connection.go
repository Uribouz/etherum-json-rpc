package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

//TODO: Unsafe practices, should be using certificates, or encrypted password
const (
	connection_string = "mongodb+srv://ball-database-mongodb-usr:ball123@cluster0.rsu0js8.mongodb.net/?retryWrites=true&w=majority"
)

func init() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connection_string).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.Background(), bson.D{bson.E{Key:"ping", Value:1}}).Err(); err != nil {
		panic(err)
	}
	dbClient = client
	fmt.Println("MongoDB: Connection established.")
}
func Close() {
	if dbClient != nil {
		err := dbClient.Disconnect(context.Background())
		if err != nil {
			panic(err)
		}
	} 
}
