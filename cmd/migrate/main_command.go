package migrate

import (
	"errors"
	"fmt"
	"github.com/liuguangw/forumx/core"
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
		Action: func(c *cli.Context) error {
			step := c.Int("step")
			fmt.Println("migrate: ", step)
			migrationLogs, err := getInstalledMigrationLogs(1)
			if err != nil {
				return err
			}
			fmt.Println(migrationLogs)
			return nil
		},
		Subcommands: []*cli.Command{
			rollbackCommand(),
			resetCommand(),
			refreshCommand(),
		},
	}
	return versionCmd
}

//处理迁移,返回执行的迁移记录列表
func processMigrate(installedMigrationLogs []*installedMigrationLog, migrations []core.Migration, step int) ([]*installedMigrationLog, error) {
	//本次迁移的记录
	var migrationLogs []*installedMigrationLog
	//迁移的批次
	migrationBatch := 1
	if len(installedMigrationLogs) > 0 {
		lastInstalledMigrationLog := installedMigrationLogs[len(installedMigrationLogs)-1]
		migrationBatch = lastInstalledMigrationLog.Batch + 1
	}
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
			return migrationLogs, errors.New("execute " + migration.Name() + " error: " + err.Error())
		}
		//构造迁移记录
		currentMigrationLog := &installedMigrationLog{
			Name:  migration.Name(),
			Batch: migrationBatch,
		}
		migrationLogs = append(migrationLogs, currentMigrationLog)
		//step限制
		if step > 0 && len(migrationLogs) >= step {
			break
		}
	}
	return migrationLogs, nil
}
