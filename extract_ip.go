package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

/*
	Read in the file contents line by line
	Extract the local address (2nd word) and remote address (3rd word) using regexp
	Convert hex to string representation of ip and port
	Pretty print it
*/

type ipAndPort struct {
	localIp    string
	localPort  string
	remoteIp   string
	remotePort string
}

type extracted []ipAndPort

func extractLines(filename string) (ext extracted) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	// Based on output from /proc/net/tcp of the form
	// 0: 00000000:1F99 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 30876 1 0000000000000000 100 0 0 10 0
	// Label capture groups for local/remote Ip and port for later expansion
	re := regexp.MustCompile(`^\s*[\dA-F]+:\s+(?P<localIp>[\dA-F]+):(?P<localPort>[\dA-F]+)\s+(?P<remoteIp>[\dA-F]+):(?P<remotePort>[\dA-F]+).*$`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if re.MatchString(scanner.Text()) {
			matchedLocalIp := fmt.Sprintf("${%s}", re.SubexpNames()[1])
			matchedLocalPort := fmt.Sprintf("${%s}", re.SubexpNames()[2])
			matchedRemoteIp := fmt.Sprintf("${%s}", re.SubexpNames()[3])
			matchedRemotePort := fmt.Sprintf("${%s}", re.SubexpNames()[4])
			ext = append(ext, ipAndPort{
				localIp:    re.ReplaceAllString(scanner.Text(), matchedLocalIp),
				localPort:  re.ReplaceAllString(scanner.Text(), matchedLocalPort),
				remoteIp:   re.ReplaceAllString(scanner.Text(), matchedRemoteIp),
				remotePort: re.ReplaceAllString(scanner.Text(), matchedRemotePort),
			})
		} else {
			log.Print("Incoming data did not match")
		}
	}
	return ext
}
