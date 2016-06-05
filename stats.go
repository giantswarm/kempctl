package main

import (
	"fmt"
	"os"

	kemp "github.com/giantswarm/kemp-client"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
)

var (
	statsCmd = &cobra.Command{
		Use:   "stats",
		Short: "Show statistics",
		Long:  "Show statistics",
		Run:   statsRun,
	}

	statsFlags = &StatsFlags{}

	validKinds = []string{
		"all",
		"totals",
		"virtual",
		"real",
	}
)

type StatsFlags struct {
	kind string
}

func init() {
	createStatsFlags(statsCmd, statsFlags)
}

func createStatsFlags(cmd *cobra.Command, flags *StatsFlags) {
	cmd.Flags().StringVar(&flags.kind, "kind", "all", "What results to show (all, totals, virtual, real)")
}

const (
	totalsHeader = "Conns/s | Bytes/s | Pkts/s"
	totalsScheme = "%v | %v | %v"

	virtualServerHeader = "ID | Enabled | Address | Port | Prot | TotalConns | ActiveConns | Conns/s | TotalPkts | TotalBytes | BytesRead | BytesWritten"
	virtualServerScheme = "%v | %v | %v | %v | %v | %v | %v | %v | %v | %v | %v | %v"

	realServerHeader = "VSID | RSID | Enabled | Address | Port | Weight | ActiveConns | Conns/s | TotalPkts | TotalBytes | BytesRead | BytesWritten"
	realServerScheme = "%v | %v | %v | %v | %v | %v | %v | %v | %v | %v | %v | %v"
)

func createTotalsOutput(lines []string, result kemp.Statistics) []string {
	lines = append(lines, totalsHeader)
	lines = append(lines, fmt.Sprintf(totalsScheme, result.Totals.ConnectionsPerSec, result.Totals.BytesPerSec, result.Totals.PacketsPerSec))

	return lines
}

func createVirtualServerOutput(lines []string, result kemp.Statistics) []string {
	lines = append(lines, virtualServerHeader)
	for _, vs := range result.VirtualServers {
		enabled := "✓"
		if vs.Enabled == 0 {
			enabled = "✗"
		}

		lines = append(lines, fmt.Sprintf(virtualServerScheme, vs.Index, enabled, vs.Address, vs.Port, vs.Protocol, vs.TotalConnections, vs.ActiveConnections, vs.ConnectionsPerSec, vs.TotalPackets, vs.TotalBytes, vs.BytesRead, vs.BytesWritten))
	}

	return lines
}

func createRealServerOutput(lines []string, result kemp.Statistics) []string {
	lines = append(lines, realServerHeader)
	for _, rs := range result.RealServers {
		enabled := "✓"
		if rs.Enabled == 0 {
			enabled = "✗"
		}

		lines = append(lines, fmt.Sprintf(realServerScheme, rs.VSIndex, rs.RSIndex, enabled, rs.Address, rs.Port, rs.Weight, rs.ActiveConnections, rs.ConnectionsPerSec, rs.TotalPackets, rs.TotalBytes, rs.BytesRead, rs.BytesWritten))
	}

	return lines
}

func statsRun(cmd *cobra.Command, args []string) {
	isKindValid := false
	for _, validKind := range validKinds {
		if statsFlags.kind == validKind {
			isKindValid = true
		}
	}

	if !isKindValid {
		fmt.Printf("Invalid kind - must be one of: %v\n", validKinds)
		os.Exit(1)
	}

	client := createClient()
	result, err := client.GetStatistics()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	lines := []string{}

	switch statsFlags.kind {
	case "all":
		lines = createTotalsOutput(lines, result)
		lines = createVirtualServerOutput(lines, result)
		lines = createRealServerOutput(lines, result)
	case "totals":
		lines = createTotalsOutput(lines, result)
	case "virtual":
		lines = createVirtualServerOutput(lines, result)
	case "real":
		lines = createRealServerOutput(lines, result)
	}

	fmt.Println(columnize.SimpleFormat(lines))

	os.Exit(0)
}
