package migrate

import (
	"context"
	"fmt"
	"github.com/liuguangw/forumx/core/db"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

//MainCommand 执行数据迁移的主命令
func MainCommand() *cli.Command {
	migrateCmd := &cli.Command{
		Name:  "migrate",
		Usage: "Run the database migrations",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "step", Usage: "Force the migrations to be run so they can be rolled back individually"},
		},
		Action: func(c *cli.Context) error {
			step := c.Int("step")
			return mainCommandAction(step)
		},
		Subcommands: []*cli.Command{
			rollbackCommand(),
			resetCommand(),
			refreshCommand(),
		},
	}
	return migrateCmd
}

func mainCommandAction(step int) error {
	migrationLogs, err := getInstalledMigrationLogs(1)
	if err != nil {
		return err
	}
	migrationLogsCount := len(migrationLogs)
	//下次迁移的ID和批次序号
	var (
		nextBatch          = 1
		nextMigrationLogID = 1
	)
	if migrationLogsCount > 0 {
		lastMigrationLog := migrationLogs[migrationLogsCount-1]
		nextBatch = lastMigrationLog.Batch + 1
		nextMigrationLogID = lastMigrationLog.ID + 1
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
			ID:    nextMigrationLogID,
			Name:  name,
			Batch: nextBatch,
		}
		//插入迁移记录
		if _, err := migrationColl.InsertOne(context.TODO(), currentMigrationLog); err != nil {
			return errors.Wrap(err, "migrate "+name+" failed, save migration log error")
		}
		nextMigrationLogID++
		fmt.Println("migrate " + name + " success")
		return nil
	}
	migrations := allMigrations()
	migrateCount, err := processMigrate(migrationLogs, migrations, migrationHandler, step)
	if err != nil {
		return err
	}
	if migrateCount == 0 {
		fmt.Println("nothing to migrate")
	} else {
		fmt.Println("migrate success: ", migrateCount, " logs")
	}
	return nil
}

//处理数据迁移,返回迁移执行的条数
func processMigrate(installedMigrationLogs []*migrationLog, migrations []Migration,
	migrationHandler migrationHandlerFunc, step int) (int, error) {
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
			return migratedCount, errors.Wrap(err, "migrate "+migration.Name()+" failed")
		}
		//执行成功后的处理
		if err := migrationHandler(migration.Name()); err != nil {
			return migratedCount, err
		}
		//step限制
		migratedCount++
		if step > 0 && migratedCount >= step {
			break
		}
	}
	return migratedCount, nil
}
