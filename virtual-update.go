package main

import (
	"fmt"
	"os"
	"strconv"

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
		fmt.Fprintln(os.Stderr, "Virtual service ID (kempctl virtual update <id>) is missing.")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "Too many parameters.")
		os.Exit(1)
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Virtual service ID (kempctl virtual update <id>) should be a number '%s' (%s).", args[0], err.Error()))
		os.Exit(1)
	}

	client := createClient()
	_, err = client.UpdateVirtualService(id, kemp.VirtualServiceParams{
		Name:            virtualUpdateFlags.name,
		IPAddress:       virtualUpdateFlags.ip,
		Port:            virtualUpdateFlags.port,
		Protocol:        virtualUpdateFlags.protocol,
		Transparent:     virtualUpdateFlags.transparent,
		CheckType:       virtualUpdateFlags.checkType,
		CheckURL:        virtualUpdateFlags.checkURL,
		CheckPort:       virtualUpdateFlags.checkPort,
		SSLAcceleration: virtualUpdateFlags.sslAcceleration,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	virtualShowRun(cmd, []string{strconv.Itoa(id)})

	os.Exit(0)
}
