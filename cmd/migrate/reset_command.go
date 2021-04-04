package migrate

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

//回滚所有迁移
func resetCommand() *cli.Command {
	versionCmd := &cli.Command{
		Name:  "reset",
		Usage: "Roll back all migrations",
		Action: func(c *cli.Context) error {
			fmt.Println("reset migration")
			return nil
		},
	}
	return versionCmd
}
