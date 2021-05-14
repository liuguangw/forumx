package migrate

import (
	"github.com/liuguangw/forumx/core/service/migration"
	"github.com/urfave/cli/v2"
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
			return migration.ExecuteRollback(step)
		},
	}
	return rollbackCmd
}
