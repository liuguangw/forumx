package migrate

import (
	"context"
	"errors"
	"fmt"
	"github.com/liuguangw/forumx/core/db"
	"github.com/urfave/cli/v2"
)

//执行数据迁移的主命令
func MainCommand() *cli.Command {
	versionCmd := &cli.Command{
		Name:  "migrate",
		Usage: "Run the database migrations",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "step", Usage: "Force the migrations to be run so they can be rolled back individually"},
		},
		Action: mainCommandAction,
		Subcommands: []*cli.Command{
			rollbackCommand(),
			resetCommand(),
			refreshCommand(),
		},
	}
	return versionCmd
}

func mainCommandAction(c *cli.Context) error {
	step := c.Int("step")
	migrationLogs, err := getInstalledMigrationLogs(1)
	if err != nil {
		return err
	}
	migrationLogsCount := len(migrationLogs)
	//下次迁移的ID和批次序号
	var (
		nextBatch          = 1
		nextMigrationLogId = 1
	)
	if migrationLogsCount > 0 {
		lastMigrationLog := migrationLogs[migrationLogsCount-1]
		nextBatch = lastMigrationLog.Batch + 1
		nextMigrationLogId = lastMigrationLog.Id + 1
	}
	//获取migration记录集合的handle
	migrationColl, err := db.Collection(migrationLogCollectionName)
	if err != nil {
		return err
	}
	//迁移执行成功后的操作
	migrationHandler := func(name string) error {
		//记录迁移
		currentMigrationLog := &migrationLog{
			Id:    nextMigrationLogId,
			Name:  name,
			Batch: nextBatch,
		}
		//插入迁移记录
		if _, err := migrationColl.InsertOne(context.TODO(), currentMigrationLog); err != nil {
			return errors.New("migrate " + name + " error: save migration log error, " + err.Error())
		}
		nextMigrationLogId++
		fmt.Println("migrate " + name + " success")
		return nil
	}
	migrations := allMigrations()
	if err := processMigrate(migrationLogs, migrations, migrationHandler, step); err != nil {
		return err
	}
	fmt.Println("migrate all success")
	return nil
}

//处理数据迁移
func processMigrate(installedMigrationLogs []*migrationLog, migrations []Migration,
	migrationHandler migrationHandlerFunc, step int) error {
	//本次迁移的条数
	migratedCount := 0
	//遍历需要执行的迁移列表
	for _, migration := range migrations {
		//判断是否已经执行了迁移
		var installedMigration bool
		for _, migrationLog := range installedMigrationLogs {
			if migrationLog.Name == migration.Name() {
				installedMigration = true
				break
			}
		}
		//已经执行过了
		if installedMigration {
			continue
		}
		//处理执行出错
		if err := migration.Up(); err != nil {
			return errors.New("migrate " + migration.Name() + " error: " + err.Error())
		}
		//执行成功后的处理
		if err := migrationHandler(migration.Name()); err != nil {
			return err
		}
		//step限制
		migratedCount++
		if step > 0 && migratedCount >= step {
			break
		}
	}
	return nil
}
