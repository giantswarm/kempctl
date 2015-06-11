package main

import (
	"fmt"
	"os"
	"strconv"

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
	if len(args) > 0 {
		fmt.Fprintln(os.Stderr, "Too many parameters.")
		os.Exit(1)
	}
	if realAddFlags.virtualServiceID == 0 {
		fmt.Fprintln(os.Stderr, "Virtual Service ID (--virtual-service-id) is required.")
		os.Exit(1)
	}
	if realAddFlags.ip == "" {
		fmt.Fprintln(os.Stderr, "Real server IP (--ip) address is required.")
		os.Exit(1)
	}
	if realAddFlags.port == "" {
		fmt.Fprintln(os.Stderr, "Real server Port (--port) is required.")
		os.Exit(1)
	}

	client := createClient()
	err := client.AddRealServerByID(realAddFlags.virtualServiceID, kemp.RealServer{
		IPAddress: realAddFlags.ip,
		Port:      realAddFlags.port,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	virtualShowRun(cmd, []string{strconv.Itoa(realAddFlags.virtualServiceID)})

	os.Exit(0)
}
