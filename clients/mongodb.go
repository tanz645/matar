package clients

import (
	"context"
	"fmt"
	"matar/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoConnectedClient *mongo.Client

func ConnectToMongoDB(ctx context.Context) *mongo.Client {
	fmt.Println(configs.GetEnvVar("MONGOURI"))
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.GetEnvVar("MONGOURI")))
	if err != nil {
		panic(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB")
	mongoConnectedClient = client
	return mongoConnectedClient
}

func GetMongoCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(fmt.Sprintf(configs.Common.Database.Name)).Collection(collectionName)
	return collection
}

func GetConnectedMongoClient() *mongo.Client {
	return mongoConnectedClient
}
