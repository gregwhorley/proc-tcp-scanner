package main

import "testing"

func TestPortScannerLoad(t *testing.T) {
	testDataA := loadOutput("127.0.0.1", "8080", "209.10.11.12", "33333")
	testDataB := loadOutput("127.0.0.1", "8081", "209.10.11.12", "33333")
	testDataC := loadOutput("127.0.0.1", "8082", "209.10.11.12", "33333")
	testDataD := loadOutput("127.0.0.1", "8083", "209.10.11.12", "33333")
	testData := []string{testDataA, testDataB, testDataC, testDataD}
	attacker, localIp, ports := portScanCheck(testData)
	loadedOutput := loadPortScanOutput(attacker, localIp, ports)
	t.Log(loadedOutput)
}
