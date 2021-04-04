package core

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

var mongoClient *mongo.Client

//获取数据库Client
func DbClient() (*mongo.Client, error) {
	if mongoClient != nil {
		return mongoClient, nil
	}
	dbUri := os.Getenv("DB_URI")
	if dbUri == "" {
		return nil, errors.New("DB_URI environment variable not found")
	}
	//connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		return nil, err
	}
	//ping
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	//set mongoClient
	mongoClient = client
	return client, nil
}
