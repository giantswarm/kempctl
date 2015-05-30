package main

import (
	"fmt"
	"os"

	kemp "github.com/giantswarm/kemp-client"
	"github.com/spf13/cobra"
)

var (
	virtualAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add virtual services",
		Long:  `Add virtual services`,
		Run:   virtualAddRun,
	}

	virtualAddFlags struct {
		ip       string
		port     string
		protocol string
	}
)

func init() {
	virtualAddCmd.Flags().StringVar(&virtualAddFlags.ip, "ip", "", "IP address of the virtual service")
	virtualAddCmd.Flags().StringVar(&virtualAddFlags.port, "port", "", "Port of the virtual service")
	virtualAddCmd.Flags().StringVar(&virtualAddFlags.protocol, "protocol", "tcp", "Protocol of the virtual service")
}

func virtualAddRun(cmd *cobra.Command, args []string) {
	client := createClient()
	data, err := client.AddVirtualService(kemp.VirtualService{
		IPAddress: virtualAddFlags.ip,
		Port:      virtualAddFlags.port,
		Protocol:  virtualAddFlags.protocol,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	virtualShowRun(cmd, []string{data.ID})

	os.Exit(0)
}
