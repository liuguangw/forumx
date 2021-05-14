package migrate

import (
	"github.com/liuguangw/forumx/core/service/migration"
	"github.com/urfave/cli/v2"
)

//回滚所有迁移
func resetCommand() *cli.Command {
	resetCmd := &cli.Command{
		Name:  "reset",
		Usage: "Roll back all migrations",
		Action: func(c *cli.Context) error {
			return migration.ExecuteReset()
		},
	}
	return resetCmd
}
