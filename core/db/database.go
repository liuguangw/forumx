package db

import (
	"github.com/liuguangw/forumx/core/environment"
	"go.mongodb.org/mongo-driver/mongo"
)

//获取数据库Handle
func Database() (*mongo.Database, error) {
	dbName := environment.DatabaseName()
	client, err := Client()
	if err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}
