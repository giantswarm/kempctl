package main

import (
	"fmt"
	"os"

	kemp "github.com/giantswarm/kemp-client"
	"github.com/spf13/cobra"
)

var (
	virtualUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update virtual services",
		Long:  `Update virtual services`,
		Run:   virtualUpdateRun,
	}

	virtualUpdateFlags = &VirtualAddFlags{}
)

func init() {
	createAddFlags(virtualUpdateCmd, virtualUpdateFlags)
}

func virtualUpdateRun(cmd *cobra.Command, args []string) {
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
	_, err := client.UpdateVirtualService(id, kemp.VirtualService{
		Name:            virtualUpdateFlags.name,
		IPAddress:       virtualUpdateFlags.ip,
		Port:            virtualUpdateFlags.port,
		Protocol:        virtualUpdateFlags.protocol,
		Transparent:     virtualUpdateFlags.transparent,
		CheckType:       virtualUpdateFlags.checkType,
		CheckUrl:        virtualUpdateFlags.checkUrl,
		CheckPort:       virtualUpdateFlags.checkPort,
		SSLAcceleration: virtualUpdateFlags.sslAcceleration,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	virtualShowRun(cmd, []string{id})

	os.Exit(0)
}
