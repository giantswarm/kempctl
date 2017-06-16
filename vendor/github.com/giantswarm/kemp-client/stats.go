package kempclient

import (
	"encoding/xml"
	"fmt"
	"sort"

	"github.com/juju/errgo"
)

// StatisticsResponse represents the API response from the `stats` endpoint.
type StatisticsResponse struct {
	Debug   string     `xml:",innerxml"`
	XMLName xml.Name   `xml:"Response"`
	Data    Statistics `xml:"Success>Data"`
}

// Statistics represents the statistics data returned by the API.
type Statistics struct {
	Totals          Totals                  `xml:"VStotals"`
	VirtualServices VirtualServiceStatsList `xml:"Vs"`
	RealServers     RealServerStatsList     `xml:"Rs"`
}

// VirtualServiceStatsList is a list of VirtualServiceStats.
type VirtualServiceStatsList []VirtualServiceStats

func (l VirtualServiceStatsList) Len() int {
	return len(l)
}

func (l VirtualServiceStatsList) Less(i int, j int) bool {
	if l[i].Address == l[j].Address {
		return l[i].Port < l[j].Port
	}

	return l[i].Address < l[j].Address
}

func (l VirtualServiceStatsList) Swap(i int, j int) {
	l[i], l[j] = l[j], l[i]
}

// RealServerStatsList is a list of RealServerStats.
type RealServerStatsList []RealServerStats

func (l RealServerStatsList) Len() int {
	return len(l)
}

func (l RealServerStatsList) Less(i int, j int) bool {
	if l[i].Address == l[j].Address {
		return l[i].Port < l[j].Port
	}

	return l[i].Address < l[j].Address
}

func (l RealServerStatsList) Swap(i int, j int) {
	l[i], l[j] = l[j], l[i]
}

// Totals represents global statistics data.
type Totals struct {
	ConnectionsPerSec int `xml:"ConnsPerSec"`
	BitsPerSec        int
	BytesPerSec       int
	PacketsPerSec     int `xml:"PktsPerSec"`
}

// VirtualServiceStats represents statistics for a Virtual Service.
type VirtualServiceStats struct {
	Index             int
	Address           string `xml:"VSAddress"`
	Port              int    `xml:"VSPort"`
	Protocol          string `xml:"VSProt"`
	TotalConnections  int    `xml:"TotalConns"`
	TotalPackets      int    `xml:"TotalPkts"`
	TotalBytes        int
	TotalBits         int
	ActiveConnections int `xml:"ActiveConns"`
	ConnectionsPerSec int `xml:"ConnsPerSec"`
	BytesRead         int
	BytesWritten      int
	Enabled           int `xml:"Enable"`
	WafEnable         int
	ErrorCode         int
}

// RealServerStats represents statistics for a Real Server.
type RealServerStats struct {
	VSIndex           int
	RSIndex           int
	Address           string `xml:"Addr"`
	Port              int
	TotalConnections  int `xml:"Conns"`
	TotalPackets      int `xml:"Pkts"`
	TotalBytes        int `xml:"Bytes"`
	TotalBits         int `xml:"Bits"`
	ActiveConnections int `xml:"ActivConns"`
	ConnectionsPerSec int `xml:"ConnsPerSec"`
	BytesRead         int
	BytesWritten      int
	Enabled           int `xml:"Enable"`
	Weight            int
	Persist           int
}

// GetStatistics calls the API, and returns a Statistics object.
func (c *Client) GetStatistics() (Statistics, error) {
	parameters := make(map[string]string)

	data := StatisticsResponse{}
	err := c.Request("stats", parameters, &data)
	if err != nil {
		return Statistics{}, errgo.NoteMask(err, "kemp could not return stats", errgo.Any)
	}

	if c.debug {
		fmt.Println("DEBUG:", data.Debug)
	}

	sort.Sort(data.Data.VirtualServices)
	sort.Sort(data.Data.RealServers)

	return data.Data, nil
}
