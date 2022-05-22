package main

import (
	"log"
	"strings"
)

/*
Add the ability to detect a port scan, where a single source IP connects to more than 3 host ports in the previous minute.
*/
var currentremoteIp, lastRemoteIp, currentLocalIp, currentPort, lastPort, currentMinute, lastMinute string
var portList []string
var stateCounter int

// this should really return a map that's keyed on remote IP with list of ports as value
func portScanCheck(ipPortList []string) (attackerIp string, localIp string, ports string) {
	for _, line := range ipPortList {
		line = strings.TrimSpace(line)
		splitLine := strings.Split(line, " ")
		if len(splitLine) != 7 {
			log.Fatalln("Invalid input...exiting")
		}
		currentMinute = splitLine[1][3:5]
		currentremoteIp = strings.Split(splitLine[4], ":")[0]
		currentLocalIp = strings.Split(splitLine[6], ":")[0]
		currentPort = strings.Split(splitLine[6], ":")[1]

		if (currentremoteIp == lastRemoteIp || lastRemoteIp == "") && (currentMinute == lastMinute || lastMinute == "") && currentPort != lastPort {
			stateCounter += 1
			portList = append(portList, currentPort)
			attackerIp = currentremoteIp
			localIp = currentLocalIp
		}
		lastRemoteIp = currentremoteIp
		lastPort = currentPort
		lastMinute = currentMinute
	}
	if stateCounter > 3 {
		ports = strings.Join(portList, ",")
	}
	return
}
