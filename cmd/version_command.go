package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"runtime"
	"strings"
)

//此处将会在编译时用ldflags设置变量的值
var (
	appVersion       = "1.0.0"
	appBuildTime     = "1970-01-01 00:00:00"
	appGitCommitHash = "0000000000000000000000000000000000000000"
)

func versionCommand() *cli.Command {
	versionCmd := &cli.Command{
		Name:  "version",
		Usage: "Show application version info",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "full", Usage: "show full version info"},
		},
		Action: func(c *cli.Context) error {
			pos := strings.Index(appBuildTime, " ")
			fmt.Printf("forumx %s (%s %s)\n", appVersion,
				appBuildTime[0:pos], appGitCommitHash[0:7])
			//打印详细信息
			if c.Bool("full") {
				fmt.Printf("version: %s\n", appVersion)
				fmt.Printf("runtime-version: %s\n", runtime.Version())
				fmt.Printf("commit-hash: %s\n", appGitCommitHash)
				fmt.Printf("build-time: %s\n", appBuildTime)
			}
			return nil
		},
	}
	return versionCmd
}
