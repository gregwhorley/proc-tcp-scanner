package main

import "log"

func main() {
	var filename string = "/proc/net/tcp"
	err := report(filename)
	if err != nil {
		log.Fatalln(err)
	}
}
