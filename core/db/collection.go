package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

//获取集合的完整名称
func CollectionFullName(shortName string) string {
	collectionPrefix := os.Getenv(collectionPrefixEnvKey)
	return collectionPrefix + shortName
}

//获取集合handle
func Collection(shortName string) (*mongo.Collection, error) {
	db, err := Database()
	if err != nil {
		return nil, err
	}
	collectionName := CollectionFullName(shortName)
	return db.Collection(collectionName), nil
}
