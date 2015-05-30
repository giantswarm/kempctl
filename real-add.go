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
		ip       string
		port     string
		protocol string
	}
)

func init() {
	realAddCmd.Flags().StringVar(&realAddFlags.ip, "ip", "", "IP address of the real server")
	realAddCmd.Flags().StringVar(&realAddFlags.port, "port", "", "Port of the real server")
}

func realAddRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Virtual service ID is missing.")
		os.Exit(1)
	}
	id := args[0]

	client := createClient()
	err := client.AddRealServerById(id, kemp.RealServer{
		IPAddress: realAddFlags.ip,
		Port:      realAddFlags.port,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	virtualShowRun(cmd, []string{id})

	os.Exit(0)
}
