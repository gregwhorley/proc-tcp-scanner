package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var testFile string = "testFile.txt"

func createFakeData() {
	content := []byte("sl  local_address rem_address   st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode\n   0: 00000000:1F99 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 30876 1 0000000000000000 100 0 0 10 0\n   1: 00000000:DFF9 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 22564 1 0000000000000000 100 0 0 10 0\n   2: 00000000:BDDB 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 22572 1 0000000000000000 100 0 0 10 0\n   3: 00000000:0801 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 15863 1 0000000000000000 100 0 0 10 0\n")
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

func TestReport(t *testing.T) {
	createFakeData()
	t.Cleanup(cleanUpFakeData)
	err := report(testFile)
	if err != nil {
		t.Error("Failure while calling report() with valid file", err)
	}
}

func TestBadReport(t *testing.T) {
	badFile := "notAFile.txt"
	err := report(badFile)
	if err != nil {
		t.Log("Caught error calling report() with invalid file", err)
	}
}
