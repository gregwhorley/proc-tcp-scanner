package main

import (
	"fmt"
	"time"
)

func loadOutput(localIp string, localPort string, remoteIp string, remotePort string) string {
	// 2021-04-28 15:28:05: New connection: 192.0.2.56:5973 -> 10.0.0.5:80
	t := time.Now().UTC()
	output := fmt.Sprintf("%s: New connection: %s:%s -> %s:%s", t.Format("2006-01-02 15:04:05"), localIp, localPort, remoteIp, remotePort)
	return output
}
