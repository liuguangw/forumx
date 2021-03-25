package cmd

import (
	"github.com/urfave/cli/v2"
)

func prepareMainApp() *cli.App {
	mainApp := &cli.App{
		Name:        "forumx",
		Description: "forumx is an efficient forum service API",
	}
	mainApp.Commands = []*cli.Command{
		versionCommand(),
	}
	return mainApp
}

func Execute(args []string) error {
	mainApp := prepareMainApp()
	return mainApp.Run(args)
}
