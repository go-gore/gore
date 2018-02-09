// +build linux darwin

package main

import (
	"path/filepath"
	"testing"

	"github.com/go-rillas/subprocess"
)

// Golang stdlib builtin function

func TestBuiltinDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "builtin")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "2" {
		t.Errorf("[FAIL] Expected response of '2' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

func TestBuiltinGorDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "builtin.gor")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "2" {
		t.Errorf("[FAIL] Expected response of '2' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}


// Golang stdlib: bytes package

func TestBytesDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "bytes")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "true" {
		t.Errorf("[FAIL] Expected response of 'true' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

func TestBytesGorDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "bytes.gor")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "true" {
		t.Errorf("[FAIL] Expected response of 'true' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}
