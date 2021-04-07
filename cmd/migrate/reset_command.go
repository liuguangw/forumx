package migrate

import (
	"context"
	"errors"
	"fmt"
	"github.com/liuguangw/forumx/core/db"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
)

//回滚所有迁移
func resetCommand() *cli.Command {
	resetCmd := &cli.Command{
		Name:  "reset",
		Usage: "Roll back all migrations",
		Action: func(c *cli.Context) error {
			return resetCommandAction()
		},
	}
	return resetCmd
}

func resetCommandAction() error {
	migrationLogs, err := getInstalledMigrationLogs(-1)
	if err != nil {
		return err
	}
	migrationLogsCount := len(migrationLogs)
	if migrationLogsCount == 0 {
		fmt.Println("nothing to rollback")
		return nil
	}
	//获取migration记录集合的handle
	migrationColl, err := db.Collection(migrationLogCollectionName)
	if err != nil {
		return err
	}
	//回滚执行成功后的操作
	migrationHandler := func(name string) error {
		//删除迁移记录
		if _, err := migrationColl.DeleteOne(context.TODO(), bson.M{"name": name}); err != nil {
			return errors.New("rollback " + name + " error: delete migration log error, " + err.Error())
		}
		fmt.Println("rollback " + name + " success")
		return nil
	}
	migrations := allMigrations()
	rollbackCount, err := processReset(migrationLogs, migrations, migrationHandler)
	if err != nil {
		return err
	}
	fmt.Println("rollback success: ", rollbackCount, " logs")
	return nil
}

//处理回滚所有数据,返回回滚的条数
func processReset(installedMigrationLogs []*migrationLog, migrations []Migration,
	migrationHandler migrationHandlerFunc) (int, error) {
	//本次回滚的条数
	rollbackCount := 0
	//遍历迁移记录列表
	for _, migrationLog := range installedMigrationLogs {
		//查找对应的migration
		migrationName := migrationLog.Name
		var migrationObject Migration
		for _, migrationNode := range migrations {
			if migrationName == migrationNode.Name() {
				migrationObject = migrationNode
				break
			}
		}
		if migrationObject == nil {
			return rollbackCount, errors.New("rollback " + migrationName + " error: migration not found")
		}
		//处理执行出错
		if err := migrationObject.Down(); err != nil {
			return rollbackCount, errors.New("rollback " + migrationName + " error: " + err.Error())
		}
		//执行成功后的处理
		if err := migrationHandler(migrationName); err != nil {
			return rollbackCount, err
		}
		rollbackCount++
	}
	return rollbackCount, nil
}
