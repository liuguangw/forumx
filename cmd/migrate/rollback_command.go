package migrate

import (
	"context"
	"errors"
	"fmt"
	"github.com/liuguangw/forumx/core/db"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
)

//回滚数据迁移的命令
func rollbackCommand() *cli.Command {
	rollbackCmd := &cli.Command{
		Name:  "rollback",
		Usage: "Roll back the database migrations",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "step", Usage: "Force the migrations to be run so they can be rolled back individually"},
		},
		Action: func(c *cli.Context) error {
			step := c.Int("step")
			return rollbackCommandAction(step)
		},
	}
	return rollbackCmd
}

func rollbackCommandAction(step int) error {
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
	//迁移执行成功后的操作
	migrationHandler := func(name string) error {
		//删除迁移记录
		if _, err := migrationColl.DeleteOne(context.TODO(), bson.M{"name": name}); err != nil {
			return errors.New("rollback " + name + " error: delete migration log error, " + err.Error())
		}
		fmt.Println("rollback " + name + " success")
		return nil
	}
	migrations := allMigrations()
	rollbackCount, err := processRollback(migrationLogs, migrations, migrationHandler, step)
	if err != nil {
		return err
	}
	fmt.Println("rollback success: ", rollbackCount, " logs")
	return nil
}

//处理数据回滚,返回回滚的条数
func processRollback(installedMigrationLogs []*migrationLog, migrations []Migration,
	migrationHandler migrationHandlerFunc, step int) (int, error) {
	//最后执行的迁移记录
	lastMigrationLog := installedMigrationLogs[0]
	//本次回滚的批次
	rollbackBatch := lastMigrationLog.Batch
	//本次回滚的条数
	rollbackCount := 0
	//遍历迁移记录列表
	for _, migrationLog := range installedMigrationLogs {
		if step == 0 && rollbackBatch != migrationLog.Batch {
			//非回滚的批次
			break
		}
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
		//step限制
		rollbackCount++
		if step > 0 && rollbackCount >= step {
			break
		}
	}
	return rollbackCount, nil
}
