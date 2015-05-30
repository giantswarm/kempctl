package main

import (
	"fmt"
	"os"

	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
)

var (
	virtualListCmd = &cobra.Command{
		Use:   "list",
		Short: "List virtual services",
		Long:  `List virtual services`,
		Run:   virtualListRun,
	}
)

const (
	virtualListHeader = "Id | Name | IPAddress | Port | Protocol | Transparent | Status | Backends | Check"
	virtualListScheme = "%s | %s | %s | %s | %s | %s | %s | %s | %s/%s"
)

func virtualListRun(cmd *cobra.Command, args []string) {
	client := createClient()
	result, err := client.ListVirtualServices()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	lines := []string{virtualListHeader}
	for _, v := range result {
		lines = append(lines, fmt.Sprintf(virtualListScheme, v.ID, v.NickName, v.IPAddress, v.Port, v.Protocol, v.Transparent, v.Status, v.NumberOfRSs, v.CheckType, v.CheckPort))
	}
	fmt.Println(columnize.SimpleFormat(lines))
	os.Exit(0)
}
