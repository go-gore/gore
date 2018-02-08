package main

import (
	"path/filepath"
	"testing"

	"github.com/go-rillas/subprocess"
)

// Golang stdlib: bytes package
func TestBytesGor(t *testing.T) {
	testpath := filepath.Join("test", "bytes.gor")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "true\n" {
		t.Errorf("[FAIL] Expected response of 'true' and received response '%s'", response.StdOut)
	}
}

func TestBytes(t *testing.T) {
	testpath := filepath.Join("test", "bytes")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "true\n" {
		t.Errorf("[FAIL] Expected response of 'true' and received response '%s'", response.StdOut)
	}
}
