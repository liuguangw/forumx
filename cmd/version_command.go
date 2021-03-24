package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

//此处将会在编译时用ldflags设置变量的值
var (
	appVersion       = "1.0.0"
	appBuildDate     = "1970-01-01"
	appGitCommitHash = "0000000000000000000000000000000000000000"
)

func versionCommand() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show application version info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("forumx %s (%s %s)\n", appVersion,
				appBuildDate, appGitCommitHash[0:7])
		},
	}
	return versionCmd
}
