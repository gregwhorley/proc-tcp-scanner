package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	for {
		extractedLines := extractLines("/proc/net/tcp")
		var outputList []string
		for _, line := range extractedLines {
			convertedLocalIp, err := flipIp(line.localIp)
			errCheck(err)
			convertedLocalIp, err = convertIp(convertedLocalIp)
			errCheck(err)
			convertedRemoteIp, err := flipIp(line.remoteIp)
			errCheck(err)
			convertedRemoteIp, err = convertIp(convertedRemoteIp)
			errCheck(err)
			convertedLocalPort, err := convertPort(line.localPort)
			errCheck(err)
			convertedRemotePort, err := convertPort(line.remotePort)

			out := loadOutput(convertedLocalIp, strconv.FormatInt(convertedLocalPort, 10), convertedRemoteIp, strconv.FormatInt(convertedRemotePort, 10))
			fmt.Println(out)
			outputList = append(outputList, out)
		}
		attackerIp, localIp, ports := portScanCheck(outputList)
		scanOut := loadPortScanOutput(attackerIp, localIp, ports)
		fmt.Println(scanOut)
		time.Sleep(10 * time.Second)
	}
}

func errCheck(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
