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
