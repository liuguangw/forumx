package cmd

import (
	"github.com/joho/godotenv"
	"github.com/liuguangw/forumx/cmd/migrate"
	"github.com/urfave/cli/v2"
	"os"
)

func prepareMainApp() *cli.App {
	mainApp := &cli.App{
		Name:        "forumx",
		Description: "forumx is an efficient forum service API",
		Commands: []*cli.Command{
			serveCommand(),
			versionCommand(),
			migrate.MainCommand(),
		},
	}
	return mainApp
}

//加载 `.env`文件, 可以使用 `FORUM_ENV_FILENAME` 环境变量设置自定义的文件名
func loadEnvFile() error {
	//自定义.env文件名
	envFileName := os.Getenv("FORUM_ENV_FILENAME")
	if envFileName != "" {
		return godotenv.Load(envFileName)
	}
	return godotenv.Load()
}

//Execute 执行命令行的入口
func Execute(args []string) error {
	//加载环境变量文件
	if err := loadEnvFile(); err != nil {
		return err
	}
	mainApp := prepareMainApp()
	return mainApp.Run(args)
}
