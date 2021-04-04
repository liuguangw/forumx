package migrate

import (
	"fmt"
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
