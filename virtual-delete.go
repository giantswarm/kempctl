package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	virtualDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete virtual services",
		Long:  `Delete virtual services`,
		Run:   virtualDeleteRun,
	}
)

func virtualDeleteRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Virtual service ID is missing.")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "Too many parameters.")
		os.Exit(1)
	}
	id := args[0]

	client := createClient()
	err := client.DeleteVirtualServiceById(id)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}
