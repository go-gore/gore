package main

import (
	"path/filepath"
	"testing"

	"github.com/go-rillas/subprocess"
)

// Golang stdlib: bytes package

func TestBytes(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "bytes")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "true\n" {
		t.Errorf("[FAIL] Expected response of 'true' and received response '%s'", response.StdOut)
	}
}

func TestBytesGor(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "bytes.gor")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "true\n" {
		t.Errorf("[FAIL] Expected response of 'true' and received response '%s'", response.StdOut)
	}
}
