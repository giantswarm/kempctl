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
	virtualListHeader = "ID | Name | IPAddress | Port | Protocol | Type | Via | Transparent | Status | Backends | Check | Cert"
	virtualListScheme = "%d | %s | %s | %s | %s | %s | %s | %s | %s/%s:'%s' | %s"
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
		status := v.Status

		if v.Status == "Up" {
			down := 0
			for _, rs := range v.Rs {
				if rs.Status != "Up" {
					down++
				}
			}
			if down > 0 {
				status = fmt.Sprintf("%s (%d down)", v.Status, down)
			}
		}

		lines = append(lines, fmt.Sprintf(virtualListScheme, v.ID, v.Name, v.IPAddress, v.Port, v.Protocol, v.VStype, v.AddVia, v.Transparent, status, v.NumberOfRSs, v.CheckType, v.CheckPort, v.CheckURL, v.CertFile))
	}
	fmt.Println(columnize.SimpleFormat(lines))

	os.Exit(0)
}
