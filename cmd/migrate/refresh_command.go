package migrate

import (
	"context"
	"errors"
	"fmt"
	"github.com/liuguangw/forumx/core/db"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
)

//回滚所有迁移然后执行所有迁移
func refreshCommand() *cli.Command {
	refreshCmd := &cli.Command{
		Name:  "refresh",
		Usage: "Roll back all of your migrations and then execute the migrate command",
		Action: func(c *cli.Context) error {
			return refreshCommandAction()
		},
	}
	return refreshCmd
}

func refreshCommandAction() error {
	migrationLogs, err := getInstalledMigrationLogs(-1)
	if err != nil {
		return err
	}
	//获取migration记录集合的handle
	migrationColl, err := db.Collection(migrationLogCollectionName)
	if err != nil {
		return err
	}
	//回滚执行成功后的操作
	rollbackHandler := func(name string) error {
		//删除迁移记录
		if _, err := migrationColl.DeleteOne(context.TODO(), bson.M{"name": name}); err != nil {
			return errors.New("rollback " + name + " error: delete migration log error, " + err.Error())
		}
		fmt.Println("rollback " + name + " success")
		return nil
	}
	migrations := allMigrations()
	rollbackCount, err := processReset(migrationLogs, migrations, rollbackHandler)
	if err != nil {
		return err
	}
	//执行迁移
	migrationLogs = nil
	//下次迁移的ID和批次序号
	var (
		nextBatch          = 1
		nextMigrationLogID = 1
	)
	//迁移执行成功后的操作
	migrationHandler := func(name string) error {
		//记录迁移
		currentMigrationLog := &migrationLog{
			ID:    nextMigrationLogID,
			Name:  name,
			Batch: nextBatch,
		}
		//插入迁移记录
		if _, err := migrationColl.InsertOne(context.TODO(), currentMigrationLog); err != nil {
			return errors.New("migrate " + name + " error: save migration log error, " + err.Error())
		}
		nextMigrationLogID++
		fmt.Println("migrate " + name + " success")
		return nil
	}
	migrateCount, err := processMigrate(migrationLogs, migrations, migrationHandler, 0)
	if err != nil {
		return err
	}
	fmt.Printf("complete success, rollback count: %d, migrate count: %d\n",
		rollbackCount, migrateCount)
	return nil
}
