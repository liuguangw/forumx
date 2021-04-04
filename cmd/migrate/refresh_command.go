package migrate

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

//回滚所有迁移然后执行所有迁移
func refreshCommand() *cli.Command {
	versionCmd := &cli.Command{
		Name:  "refresh",
		Usage: "Roll back all of your migrations and then execute the migrate command",
		Action: func(c *cli.Context) error {
			fmt.Println("reset migration")
			return nil
		},
	}
	return versionCmd
}
