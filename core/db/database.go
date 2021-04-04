package db

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

//获取数据库Handle
func Database() (*mongo.Database, error) {
	dbName := os.Getenv(dbNameEnvKey)
	if dbName == "" {
		return nil, errors.New(dbNameEnvKey + " environment variable not found")
	}
	client, err := Client()
	if err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}
