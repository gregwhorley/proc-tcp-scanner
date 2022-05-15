package main

import (
	"fmt"
	"time"
)

func loadOutput() string {
	// 2021-04-28 15:28:05: New connection: 192.0.2.56:5973 -> 10.0.0.5:80
	t := time.Now().UTC()
	res := fmt.Sprintf("%s: New connection: %s:%s -> %s:%s", t.Format("2006-01-02 15:04:05"), "127.0.0.1", "8080", "192.168.0.1", "1234")
	return res
}
