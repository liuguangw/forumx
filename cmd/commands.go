package cmd

import (
	"github.com/joho/godotenv"
	"github.com/liuguangw/forumx/cmd/migrate"
	"github.com/urfave/cli/v2"
)

func prepareMainApp() *cli.App {
	mainApp := &cli.App{
		Name:        "forumx",
		Description: "forumx is an efficient forum service API",
		Commands: []*cli.Command{
			versionCommand(),
			migrate.MainCommand(),
		},
	}
	return mainApp
}

func Execute(args []string) error {
	//加载.env环境变量配置
	if err := godotenv.Load(); err != nil {
		return err
	}
	mainApp := prepareMainApp()
	return mainApp.Run(args)
}
