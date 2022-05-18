package main

import (
	"encoding/hex"
	"errors"
	"log"
	"net"
)

func convertIp(s string) (string, error) {
	// assuming this will always be an 8 char hex string that represents an IPv4 address
	if len(s) > 8 {
		return "", errors.New("Invalid length of hex string")
	}
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatalln(err)
	}
	// assuming decoded will always be a fixed length byte slice representing an IPv4 address
	ipv4 := net.IPv4(decoded[0], decoded[1], decoded[2], decoded[3])
	return ipv4.String(), nil
}

func convertPort(s string) string {
	return ""
}

func flipIp(s string) string {
	return ""
}
