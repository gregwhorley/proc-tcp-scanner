package main

import (
	"encoding/hex"
	"errors"
	"log"
	"net"
	"strconv"
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

func convertPort(s string) (int64, error) {
	if len(s) > 4 {
		return 0, errors.New("Invalid length of hex string")
	}
	res, err := strconv.ParseInt(s, 16, 0)
	if err != nil {
		log.Fatalln(err)
	}
	return res, nil
}

/*
Assuming this will only be run on x86 cpu arch.
Better to do an arch check so endianness can be determined instead of blindly flipping hex around.
*/
func flipIp(s string) (string, error) {
	if len(s) > 8 {
		return "", errors.New("Invalid length of hex string")
	}
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}
