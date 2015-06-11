package main

import (
	"fmt"
	"os"
	"strconv"

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
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Virtual service ID should be a number '%s' (%s).", args[0], err.Error()))
	}

	client := createClient()
	err = client.DeleteVirtualServiceByID(id)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}
