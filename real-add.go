package main

import (
	"fmt"
	"os"

	kemp "github.com/giantswarm/kemp-client"
	"github.com/spf13/cobra"
)

var (
	realAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add real server",
		Long:  `Add real server`,
		Run:   realAddRun,
	}

	realAddFlags struct {
		virtualServiceID int
		ip               string
		port             string
	}
)

func init() {
	realAddCmd.Flags().IntVar(&realAddFlags.virtualServiceID, "virtual-service-id", 0, "ID of the virtual service")
	realAddCmd.Flags().StringVar(&realAddFlags.ip, "ip", "", "IP address of the real server")
	realAddCmd.Flags().StringVar(&realAddFlags.port, "port", "", "Port of the real server")
}

func realAddRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Virtual service ID is missing.")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "Too many parameters.")
		os.Exit(1)
	}
	if realAddFlags.virtualServiceID == 0 {
		fmt.Fprintln(os.Stderr, "Virtual Service ID is required.")
		os.Exit(1)
	}
	if realAddFlags.ip == "" {
		fmt.Fprintln(os.Stderr, "Real server IP address is required.")
		os.Exit(1)
	}
	if realAddFlags.port == "" {
		fmt.Fprintln(os.Stderr, "Real server Port is required.")
		os.Exit(1)
	}

	client := createClient()
	err := client.AddRealServerById(realAddFlags.virtualServiceID, kemp.RealServer{
		IPAddress: realAddFlags.ip,
		Port:      realAddFlags.port,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	virtualShowRun(cmd, []string{string(realAddFlags.virtualServiceID)})

	os.Exit(0)
}
