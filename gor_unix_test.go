// +build linux darwin

package main

import (
	"path/filepath"
	"testing"

	"github.com/go-rillas/subprocess"
	"strings"
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


// Golang stdlib: compress package

func TestCompressDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "compress")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "A long time ago in a galaxy far, far away..." {
		t.Errorf("[FAIL] Expected response of 'A long time ago in a galaxy far, far away...' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

func TestCompressGorDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "compress.gor")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "A long time ago in a galaxy far, far away..." {
		t.Errorf("[FAIL] Expected response of 'A long time ago in a galaxy far, far away...' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

// Golang stdlib: container package

func TestContainerDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "container")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "1 2 3 5 " {
		t.Errorf("[FAIL] Expected response of '1 2 3 5' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

func TestContainerGorDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "container.gor")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "1 2 3 5 " {
		t.Errorf("[FAIL] Expected response of '1 2 3 5' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

// Golang stdlib: crypto package

func TestCryptoDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "crypto")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "59 7f 6a 54 00 10 f9 4c 15 d7 18 06 a9 9a 2c 87 10 e7 47 bd" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

func TestCryptoGorDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "crypto.gor")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "59 7f 6a 54 00 10 f9 4c 15 d7 18 06 a9 9a 2c 87 10 e7 47 bd" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

// Golang stdlib: encoding package

func TestEncodingDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "encoding")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "[{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

func TestEncodingGorDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "encoding.gor")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "[{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}


// Golang stdlib: errors package

func TestErrorsDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "errors")
	response := subprocess.RunShell("", "", testpath)
	if !strings.HasPrefix(response.StdErr, "emit macho dwarf: elf header corrupted") {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdErr)
	}
	if response.ExitCode == 0 {
		t.Errorf("[FAIL] Returned zero exit code value and did not expect zero exit code: %d", response.ExitCode)
	}
}

func TestErrorsGorDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "errors.gor")
	response := subprocess.RunShell("", "", testpath)
	if !strings.HasPrefix(response.StdErr, "emit macho dwarf: elf header corrupted") {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdErr)
	}
	if response.ExitCode == 0 {
		t.Errorf("[FAIL] Returned zero exit code value and did not expect zero exit code: %d", response.ExitCode)
	}
}