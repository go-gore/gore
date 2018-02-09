package main

import (
	"path/filepath"
	"testing"

	"github.com/go-rillas/subprocess"
)


// Golang stdlib builtin function

func TestBuiltin(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "builtin")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "2" {
		t.Errorf("[FAIL] Expected response of '2' and received response '%s'", response.StdOut)
	}
}

func TestBuiltinGor(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "builtin.gor")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "2" {
		t.Errorf("[FAIL] Expected response of '2' and received response '%s'", response.StdOut)
	}
}

// Golang stdlib: bytes package

func TestBytes(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "bytes")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "true" {
		t.Errorf("[FAIL] Expected response of 'true' and received response '%s'", response.StdOut)
	}
}

func TestBytesGor(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "bytes.gor")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "true" {
		t.Errorf("[FAIL] Expected response of 'true' and received response '%s'", response.StdOut)
	}
}


