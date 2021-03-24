package cmd

import "github.com/spf13/cobra"

func prepareRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "forumx",
		Short: "forumx is an efficient forum service API",
	}
	rootCmd.AddCommand(
		versionCommand(),
	)
	return rootCmd
}

func Execute() error {
	rootCommand := prepareRootCommand()
	return rootCommand.Execute()
}
