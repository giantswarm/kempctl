package main

import "github.com/spf13/cobra"

var (
	virtualCmd = &cobra.Command{
		Use:   "virtual",
		Short: "Manage virtual services",
		Long:  `Manage virtual services`,
		Run:   virtualRun,
	}
)

func virtualRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
