package main

import (
	"fmt"
	"testing"
)

func TestConvertIp(t *testing.T) {
	testHex := "7F000001"
	res := convertIp(testHex)
	fmt.Println(res)
}
