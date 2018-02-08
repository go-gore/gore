package main

import (
	"testing"
	"github.com/go-rillas/subprocess"
)


// Golang stdlib: bytes package
func TestBytesGor(t *testing.T) {
	response := subprocess.Run("test/bytes.gor")
	if response.StdOut != "true\n" {
		t.Errorf("[FAIL] Expected response of 'true' and received response '%s'", response.StdOut)
	}
}

func TestBytes(t *testing.T) {
	response := subprocess.Run("test/bytes")
	if response.StdOut != "true\n" {
		t.Errorf("[FAIL] Expected response of 'true' and received response '%s'", response.StdOut)
	}
}