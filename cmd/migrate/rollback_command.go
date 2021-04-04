package migrate

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

//回滚数据迁移的命令
func rollbackCommand() *cli.Command {
	versionCmd := &cli.Command{
		Name:  "rollback",
		Usage: "Roll back the database migrations",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "step", Usage: "Force the migrations to be run so they can be rolled back individually"},
		},
		Action: func(c *cli.Context) error {
			step := c.Int("step")
			fmt.Println("rollback: ", step)
			return nil
		},
	}
	return versionCmd
}
