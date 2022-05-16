package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var testFile string = "testFile.txt"
var testMe, uhOh []ipAndPort

func createFakeData() {
	content := []byte(`sl  local_address rem_address   st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode
0: 00000000:1F99 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 30876 1 0000000000000000 100 0 0 10 0
1: 00000000:DFF9 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 22564 1 0000000000000000 100 0 0 10 0
2: 00000000:BDDB 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 22572 1 0000000000000000 100 0 0 10 0
3: 00000000:0801 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 15863 1 0000000000000000 100 0 0 10 0`)
	err := ioutil.WriteFile(testFile, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func cleanUpFakeData() {
	err := os.Remove(testFile)
	if err != nil {
		log.Fatal(err)
	}
}

func TestExtractLine(t *testing.T) {
	createFakeData()
	t.Cleanup(cleanUpFakeData)
	testMe = extractLines(testFile)
	for _, item := range testMe {
		t.Log(item.localIp, item.localPort, item.remoteIp, item.remotePort)
	}
}

func TestNoMatch(t *testing.T) {
	badFile := "badFile.txt"
	ioutil.WriteFile(badFile, []byte("jdskjdks sdkfjsdkfjsdkj jfjfjfjf"), 0644)
	defer os.Remove(badFile)
	uhOh = extractLines(badFile)
	for _, item := range uhOh {
		if item.localIp == "" && item.localPort == "" {
			t.Log("Expected empty struct values")
		} else {
			t.Error("Struct values should be empty")
		}
	}
}
