package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestConvertIp(t *testing.T) {
	testHex := "7F000001"
	ipAddress, _ := convertIp(testHex)
	v := reflect.ValueOf(ipAddress).Kind()
	assert.Equal(t, v, reflect.String)
}

func TestConvertInvalidLength(t *testing.T) {
	badHex := "1234567890ABCDEF"
	_, expectedError := convertIp(badHex)
	assert.EqualError(t, expectedError, "Invalid length of hex string")
}

func TestConvertPort(t *testing.T) {
	testHex := "01BB"
	port, _ := convertPort(testHex)
	assert.Equal(t, port, int64(443))
}

func TestConvertPortInvalidLength(t *testing.T) {
	badHex := "12354586ABCDE"
	_, expectedError := convertPort(badHex)
	assert.EqualError(t, expectedError, "Invalid length of hex string")
}

func TestFlip(t *testing.T) {
	testHex := "7F000001"
	flipped, _ := flipIp(testHex)
	assert.Equal(t, flipped, "100000F7")
}

func TestFlipInvalidLength(t *testing.T) {
	badHex := "87482743BEEF"
	_, expectedError := flipIp(badHex)
	assert.EqualError(t, expectedError, "Invalid length of hex string")
}
