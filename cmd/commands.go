package cmd

import (
	"github.com/joho/godotenv"
	"github.com/liuguangw/forumx/cmd/migrate"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
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
	//环境变量配置文件名称
	envFileName := os.Getenv("FORUM_ENV_FILENAME")
	if envFileName == "" {
		envFileName = ".env" //默认名称
	}
	//获取工作目录
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}
	envFilePath := filepath.Join(workingDir, envFileName)
	//如果工作目录下存在环境变量配置文件,则加载配置文件
	if _, err := os.Stat(envFilePath); os.IsExist(err) {
		return godotenv.Load(envFilePath)
	}
	//使用二进制文件所在目录下的
	exeFilePath, err := os.Executable()
	if err != nil {
		return err
	}
	envFilePath = filepath.Join(filepath.Dir(exeFilePath), envFileName)
	if _, err := os.Stat(envFilePath); os.IsExist(err) {
		return godotenv.Load(envFilePath)
	}
	//如果找不到文件则不加载
	return nil
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
