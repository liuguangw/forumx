package migration

import (
	"context"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/migration"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//迁移记录的集合名称
const migrationLogCollectionName = "migrations"

type installedLogHandler struct {
}

func (h *installedLogHandler) CreateCollection() error {
	database, err := db.Database()
	if err != nil {
		return err
	}
	collectionFullName := db.CollectionFullName(migrationLogCollectionName)
	collectionNames, err := database.ListCollectionNames(context.Background(), bson.M{
		"name": collectionFullName,
	})
	if err != nil {
		return errors.Wrap(err, "list collection names failed")
	}
	var collectionExist bool
	for _, collName := range collectionNames {
		if collName == collectionFullName {
			collectionExist = true
			break
		}
	}
	if !collectionExist {
		return createMigrationLogCollection(database, collectionFullName)
	}
	return nil
}

func (h *installedLogHandler) Insert(name string, batch int) error {
	//下次迁移的ID
	nextMigrationLogID := 1
	lastLog, err := h.lastMigrationLog()
	if err != nil {
		return err
	}
	if lastLog != nil {
		nextMigrationLogID = lastLog.ID + 1
	}
	coll, err := db.Collection(migrationLogCollectionName)
	if err != nil {
		return err
	}
	migrationLog := &migrationLog{
		ID:    nextMigrationLogID,
		Name:  name,
		Batch: batch,
	}
	//插入迁移记录
	if _, err := coll.InsertOne(context.TODO(), migrationLog); err != nil {
		return err
	}
	return nil
}

func (h *installedLogHandler) Delete(name string) error {
	coll, err := db.Collection(migrationLogCollectionName)
	if err != nil {
		return err
	}
	//删除迁移记录
	if _, err := coll.DeleteOne(context.TODO(), bson.M{"name": name}); err != nil {
		return err
	}
	return nil
}

func (h *installedLogHandler) ListLogs(sortType int) ([]*migration.InstalledLog, error) {
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
	var migrationLogs []*migrationLog
	if err = cursor.All(context.TODO(), &migrationLogs); err != nil {
		return nil, err
	}
	//类型转换
	var resultList []*migration.InstalledLog
	for _, item := range migrationLogs {
		resultList = append(resultList, &migration.InstalledLog{
			Name:  item.Name,
			Batch: item.Batch,
		})
	}
	return resultList, nil
}

func (h *installedLogHandler) lastMigrationLog() (*migrationLog, error) {
	coll, err := db.Collection(migrationLogCollectionName)
	if err != nil {
		return nil, err
	}
	opts := options.FindOne().SetSort(bson.M{
		"id": -1,
	})
	lastLog := new(migrationLog)
	if err := coll.FindOne(context.TODO(), bson.M{}, opts).Decode(lastLog); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return lastLog, nil
}

//createMigrationLogCollection 创建迁移记录Collection
func createMigrationLogCollection(database *mongo.Database, collectionFullName string) error {
	opts := options.CreateCollection()
	if err := database.CreateCollection(context.Background(), collectionFullName, opts); err != nil {
		return err
	}
	//创建索引
	coll := database.Collection(collectionFullName)
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
