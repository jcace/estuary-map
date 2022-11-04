package main

import (
	"fmt"
	"net"

	"github.com/ipinfo/go/v2/ipinfo"
)

func GenerateIpinfoMap(ips []net.IP) string {
	// Seemingly does not need an API key for this!
	client := ipinfo.NewClient(nil, nil, "")
	report, err := client.GetIPMap(ips)
	if err != nil {
		fmt.Println(err)
	}

	return report.ReportURL
}
