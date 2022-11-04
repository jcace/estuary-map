package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/pkg/browser"
	"github.com/schollz/progressbar/v3"
)

func main() {
	var ipList []net.IP
	fmt.Println("finding location of Estuary Providers...")
	allMiners, err := GetAllMiners()
	if err != nil {
		log.Fatal(err)
	}

	bar := progressbar.Default(int64(len(*allMiners)))

	for _, m := range *allMiners {
		bar.Add(1)
		// Don't map suspended miners
		if m.Suspended {
			continue
		}

		mStats, err := GetMinerStats(m.Addr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, mAddr := range mStats.ChainInfo.Addresses {
			maybeIp := strings.Split(mAddr, "/")[2]

			ip := net.ParseIP(maybeIp)

			// Will be nil if it's not a valid IP, likely because it's a DNS name
			if ip == nil {
				resolvedIp, err := net.LookupIP(maybeIp)
				if err != nil {
					continue
				}

				ipList = append(ipList, resolvedIp[0])
			} else {
				ipList = append(ipList, ip)
			}
		}

		// Pause to avoid spamming estuary API
		time.Sleep(150 * time.Millisecond)
	}

	mapResult := GenerateIpinfoMap(ipList)

	fmt.Println(mapResult)
	browser.OpenURL(mapResult)
}
