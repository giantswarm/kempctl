package main

import "github.com/spf13/cobra"

var (
	realCmd = &cobra.Command{
		Use:   "real",
		Short: "Manage real servers",
		Long:  `Manage real servers`,
		Run:   realRun,
	}
)

func realRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
