package migrate

import (
	"context"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//迁移记录的集合名称
const migrationLogCollectionName = "migrations"

//已安装的数据迁移记录
type installedMigrationLog struct {
	Id    int    `bson:"id"`    //迁移的ID
	Name  string `bson:"name"`  //迁移名称
	Batch int    `bson:"batch"` //迁移的批次
}

//获取已安装的数据迁移记录
//`sortType` 排序方式, 1正序 -1倒序
func getInstalledMigrationLogs(sortType int) ([]*installedMigrationLog, error) {
	database, err := db.Database()
	if err != nil {
		return nil, err
	}
	//判断迁移记录集合是否存在
	collectionFullName := db.CollectionFullName(migrationLogCollectionName)
	collectionNames, err := database.ListCollectionNames(context.TODO(), bson.M{
		"name": collectionFullName,
	})
	if err != nil {
		return nil, err
	}
	//创建集合
	if len(collectionNames) == 0 {
		if err := createMigrationLogCollection(database, collectionFullName); err != nil {
			return nil, err
		}
	}
	coll, err := db.Collection(migrationLogCollectionName)
	if err != nil {
		return nil, err
	}
	opts := options.Find().SetSort(bson.M{
		"id": sortType,
	})
	cursor, err := coll.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	var migrationLogs []*installedMigrationLog
	if err = cursor.All(context.TODO(), &migrationLogs); err != nil {
		return nil, err
	}
	return migrationLogs, nil
}

//创建迁移记录集合
func createMigrationLogCollection(database *mongo.Database, collectionFullName string) error {
	opts := options.CreateCollection()
	if err := database.CreateCollection(context.TODO(), collectionFullName, opts); err != nil {
		return err
	}
	//创建索引
	coll := database.Collection(migrationLogCollectionName)
	indexView := coll.Indexes()
	indexModels := []mongo.IndexModel{
		{
			Keys: bson.M{
				"id": 1,
			},
			Options: options.Index().SetName("id_uni").SetUnique(true),
		},
		{
			Keys: bson.M{
				"name": 1,
			},
			Options: options.Index().SetName("name_uni").SetUnique(true),
		},
		{
			Keys: bson.M{
				"batch": 1,
			},
			Options: options.Index().SetName("batch_index"),
		},
	}
	indexOpts := options.CreateIndexes().SetMaxTime(2 * time.Second)
	if _, err := indexView.CreateMany(context.TODO(), indexModels, indexOpts); err != nil {
		return err
	}
	return nil
}
