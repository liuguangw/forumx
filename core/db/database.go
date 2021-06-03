package db

import (
	"github.com/liuguangw/forumx/app/environment"
	"go.mongodb.org/mongo-driver/mongo"
)

//Database 获取数据库Handle
func Database() (*mongo.Database, error) {
	dbName := environment.DatabaseName()
	client, err := Client()
	if err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}
