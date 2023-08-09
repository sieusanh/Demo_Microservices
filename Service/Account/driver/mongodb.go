package driver

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	mongoConfig "go-module/config/mongodb"
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}
func Connect() *MongoDB {
	connStr := fmt.Sprintf(
		mongoConfig.HOST_URL, 
		mongoConfig.USERNAME,
		mongoConfig.PASSWORD,
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection ok")
	Mongo.Client = client
	return Mongo
}

