package main

import (
	"log"
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

	virtualAddFlags = &VirtualAddFlags{}
)

type VirtualAddFlags struct {
	name            string
	ip              string
	port            string
	protocol        string
	transparent     bool
	checkType       string
	checkURL        string
	checkPort       string
	sslAcceleration bool
}

func init() {
	createAddFlags(virtualAddCmd, virtualAddFlags)
}

func createAddFlags(cmd *cobra.Command, flags *VirtualAddFlags) {
	cmd.Flags().StringVar(&flags.name, "name", "", "Name of the virtual service")
	cmd.Flags().StringVar(&flags.ip, "ip", "", "IP address of the virtual service")
	cmd.Flags().StringVar(&flags.port, "port", "", "Port of the virtual service")
	cmd.Flags().StringVar(&flags.protocol, "protocol", "tcp", "Protocol of the virtual service")
	cmd.Flags().BoolVar(&flags.transparent, "transparent", false, "When using Layer 7, when this is enabled the connection arriving at the Real Server appears to come directly from the client. Alternatively, the connection can be non-transparent which means that the connections at the Real Server appear to come from the LoadMaster.")
	cmd.Flags().StringVar(&flags.checkType, "check-type", "http", "Specify which protocol is to be used to check the health of the Real Server (icmp, https, http, tcp, smtp, nntp, ftp, telnet, pop3, imap, rdp, bdata, none.")
	cmd.Flags().StringVar(&flags.checkURL, "check-url", "/", "When the CheckType is set to http or https: By default, the health checker tries to access the URL / to determine if the machine is available. A different URL can be set in the CheckURL parameter.")
	cmd.Flags().StringVar(&flags.checkPort, "check-port", "8080", "The port to be checked. If a port is not specified, the Real Server port is used.")
	cmd.Flags().BoolVar(&flags.sslAcceleration, "ssl-acceleration", false, "Enable SSL handling on this Virtual Service (activate automatically if port 443 is used).")
}

func virtualAddRun(cmd *cobra.Command, args []string) {
	if virtualAddFlags.port == "443" && !virtualAddFlags.sslAcceleration {
		virtualAddFlags.sslAcceleration = true
	}
	if virtualAddFlags.ip == "" {
		log.Fatal("IP address for the virtual service missing")
	}

	client := createClient()
	data, err := client.AddVirtualService(kemp.VirtualServiceParams{
		Name:            virtualAddFlags.name,
		IPAddress:       virtualAddFlags.ip,
		Port:            virtualAddFlags.port,
		Protocol:        virtualAddFlags.protocol,
		Transparent:     virtualAddFlags.transparent,
		CheckType:       virtualAddFlags.checkType,
		CheckURL:        virtualAddFlags.checkURL,
		CheckPort:       virtualAddFlags.checkPort,
		SSLAcceleration: virtualAddFlags.sslAcceleration,
	})

	if err != nil {
		log.Fatal(err)
	}

	virtualShowRun(cmd, []string{string(data.ID)})

	os.Exit(0)
}
