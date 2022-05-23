package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
	"time"
)

// prototype code for exposing new connection metrics through Prometheus -- this has not been tested
func recordNewConnections() {
	go func() {
		for {
			// do the extract transform and load operations for connections in /proc/net/tcp
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

			// add the length of tha extracted data array from /proc/net/tcp to the counter
			connectionsProcessed.Add(float64(len(extractedLines)))
			time.Sleep(10 * time.Second)
		}
	}()
}

var (
	connectionsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "new_connections_processed_total",
		Help: "The total number of new connections.",
	})
)

func main() {
	recordNewConnections()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":1234", nil))

}

func errCheck(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
