package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadOutput(t *testing.T) {
	testIp := "127.0.0.1"
	testPort := "8080"
	s := loadOutput(testIp, testPort, testIp, testPort)
	assert.Contains(t, s, "New connection: 127.0.0.1:8080 -> 127.0.0.1:8080")
}
