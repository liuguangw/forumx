package db

import (
	"github.com/liuguangw/forumx/core/environment"
	"go.mongodb.org/mongo-driver/mongo"
)

//CollectionFullName 获取集合的完整名称
func CollectionFullName(shortName string) string {
	collectionPrefix := environment.CollectionPrefix()
	return collectionPrefix + shortName
}

//Collection 获取集合handle
func Collection(shortName string) (*mongo.Collection, error) {
	db, err := Database()
	if err != nil {
		return nil, err
	}
	collectionName := CollectionFullName(shortName)
	return db.Collection(collectionName), nil
}
