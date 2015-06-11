package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
)

var (
	virtualShowCmd = &cobra.Command{
		Use:   "show <id>",
		Short: "Show virtual services",
		Long:  `Show virtual services`,
		Run:   virtualShowRun,
	}
)

const (
	virtualShowHeader = "ID | IPAddress | Port | Forward | Status | Weight | Enable"
	virtualShowScheme = "%s | %s | %s | %s | %s | %s | %s"
)

func virtualShowRun(cmd *cobra.Command, args []string) {
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
	result, err := client.ShowVirtualServiceByID(id)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// TODO some grouping to get a better understanding of the config
	lines := []string{}
	lines = append(lines, fmt.Sprintf("%s | %s", "ID:", result.ID))
	lines = append(lines, fmt.Sprintf("%s | %s", "Name:", result.Name))
	lines = append(lines, fmt.Sprintf("%s | %s", "IPAddress:", result.IPAddress))
	lines = append(lines, fmt.Sprintf("%s | %s", "Port:", result.Port))
	lines = append(lines, fmt.Sprintf("%s | %s", "Protocol:", result.Protocol))
	lines = append(lines, fmt.Sprintf("%s | %s", "Status:", result.Status))
	lines = append(lines, fmt.Sprintf("%s | %s", "Enable:", result.Enable))
	lines = append(lines, fmt.Sprintf("%s | %s/%s:'%s'", "CheckType:", result.CheckType, result.CheckPort, result.CheckURL))
	lines = append(lines, fmt.Sprintf("%s | %s", "Transparent:", result.Transparent))
	if result.CertFile != "" {
		lines = append(lines, fmt.Sprintf("%s | %s", "CertFile:", result.CertFile))
	}

	if globalFlags.verbose {
		lines = append(lines, fmt.Sprintf("%s | %s", "SSLReverse:", result.SSLReverse))
		lines = append(lines, fmt.Sprintf("%s | %s", "SSLReencrypt:", result.SSLReencrypt))
		lines = append(lines, fmt.Sprintf("%s | %s", "Intercept:", result.Intercept))
		lines = append(lines, fmt.Sprintf("%s | %s", "InterceptOpts:", strings.Join(result.InterceptOpts, ", ")))
		lines = append(lines, fmt.Sprintf("%s | %s", "AlertThreshold:", result.AlertThreshold))
		lines = append(lines, fmt.Sprintf("%s | %s", "Transactionlimit:", result.Transactionlimit))
		lines = append(lines, fmt.Sprintf("%s | %s", "ServerInit:", result.ServerInit))
		lines = append(lines, fmt.Sprintf("%s | %s", "StartTLSMode:", result.StartTLSMode))
		lines = append(lines, fmt.Sprintf("%s | %s", "Idletime:", result.Idletime))
		lines = append(lines, fmt.Sprintf("%s | %s", "Cache:", result.Cache))
		lines = append(lines, fmt.Sprintf("%s | %s", "Compress:", result.Compress))
		lines = append(lines, fmt.Sprintf("%s | %s", "Verify:", result.Verify))
		lines = append(lines, fmt.Sprintf("%s | %s", "UseforSnat:", result.UseforSnat))
		lines = append(lines, fmt.Sprintf("%s | %s", "ForceL7:", result.ForceL7))
		lines = append(lines, fmt.Sprintf("%s | %s", "ClientCert:", result.ClientCert))
		lines = append(lines, fmt.Sprintf("%s | %s", "ErrorCode:", result.ErrorCode))
		lines = append(lines, fmt.Sprintf("%s | %s", "CheckUse1.1:", result.CheckUse11))
		lines = append(lines, fmt.Sprintf("%s | %s", "MatchLen:", result.MatchLen))
		lines = append(lines, fmt.Sprintf("%s | %s", "CheckUseGet:", result.CheckUseGet))
		lines = append(lines, fmt.Sprintf("%s | %s", "SSLRewrite:", result.SSLRewrite))
		lines = append(lines, fmt.Sprintf("%s | %s", "VStype:", result.VStype))
		lines = append(lines, fmt.Sprintf("%s | %s", "FollowVSID:", result.FollowVSID))
		lines = append(lines, fmt.Sprintf("%s | %s", "Schedule:", result.Schedule))
		lines = append(lines, fmt.Sprintf("%s | %s", "PersistTimeout:", result.PersistTimeout))
		lines = append(lines, fmt.Sprintf("%s | %s", "SSLAcceleration:", result.SSLAcceleration))
		lines = append(lines, fmt.Sprintf("%s | %s", "NRules:", result.NRules))
		lines = append(lines, fmt.Sprintf("%s | %s", "NRequestRules:", result.NRequestRules))
		lines = append(lines, fmt.Sprintf("%s | %s", "NResponseRules:", result.NResponseRules))
		lines = append(lines, fmt.Sprintf("%s | %s", "NPreProcessRules:", result.NPreProcessRules))
		lines = append(lines, fmt.Sprintf("%s | %s", "EspEnabled:", result.EspEnabled))
		lines = append(lines, fmt.Sprintf("%s | %s", "InputAuthMode:", result.InputAuthMode))
		lines = append(lines, fmt.Sprintf("%s | %s", "OutputAuthMode:", result.OutputAuthMode))
		lines = append(lines, fmt.Sprintf("%s | %s", "MasterVS:", result.MasterVS))
		lines = append(lines, fmt.Sprintf("%s | %s", "MasterVSID:", result.MasterVSID))
		lines = append(lines, fmt.Sprintf("%s | %s", "AddVia:", result.AddVia))
		lines = append(lines, fmt.Sprintf("%s | %s", "TlsType:", result.TlsType))
		lines = append(lines, fmt.Sprintf("%s | %s", "NeedHostName:", result.NeedHostName))
		lines = append(lines, fmt.Sprintf("%s | %s", "OCSPVerify:", result.OCSPVerify))
	}
	fmt.Println(columnize.SimpleFormat(lines))
	fmt.Println("")

	if len(result.Rs) > 0 {
		lines = []string{virtualShowHeader}
		for _, server := range result.Rs {
			lines = append(lines, fmt.Sprintf(virtualShowScheme, server.ID, server.IPAddress, server.Port, server.Forward, server.Status, server.Weight, server.Enable))
		}
		fmt.Println(columnize.SimpleFormat(lines))
	}
	os.Exit(0)
}
