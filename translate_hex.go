package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func convertIp(s string) string {
	if len(s) > 8 {
		log.Fatalln("Invalid length of hex string")
	}
	var convertedIp string
	for i := 0; i < len(s); i += 2 {
		decoded, err := hex.DecodeString(s[i : i+2])
		if err != nil {
			log.Fatal(err)
		}
		// hacky way to format ipv4 address without appending a dot at the end
		if i < 6 {
			convertedIp += fmt.Sprintf("%d.", decoded)
		} else {
			convertedIp += fmt.Sprintf("%d", decoded)
		}
	}
	return convertedIp
}

func convertPort(s string) string {
	return ""
}

func flipIp(s string) string {
	return ""
}
