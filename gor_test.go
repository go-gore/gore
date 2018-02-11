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
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

func TestBuiltinGor(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "builtin.gor")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "2" {
		t.Errorf("[FAIL] Expected response of '2' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

// Golang stdlib: bytes package

func TestBytes(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "bytes")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "true" {
		t.Errorf("[FAIL] Expected response of 'true' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

func TestBytesGor(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "bytes.gor")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "true" {
		t.Errorf("[FAIL] Expected response of 'true' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

// Golang stdlib: compress package

func TestCompress(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "compress")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "A long time ago in a galaxy far, far away..." {
		t.Errorf("[FAIL] Expected response of 'A long time ago in a galaxy far, far away...' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

func TestCompressGor(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "compress.gor")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "A long time ago in a galaxy far, far away..." {
		t.Errorf("[FAIL] Expected response of 'A long time ago in a galaxy far, far away...' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

// Golang stdlib: container package

func TestContainer(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "container")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "1 2 3 5 " {
		t.Errorf("[FAIL] Expected response of '1 2 3 5' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

func TestContainerGor(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "container.gor")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "1 2 3 5 " {
		t.Errorf("[FAIL] Expected response of '1 2 3 5' and received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

// Golang stdlib: crypto package

func TestCrypto(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "crypto")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "59 7f 6a 54 00 10 f9 4c 15 d7 18 06 a9 9a 2c 87 10 e7 47 bd" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

func TestCryptoGor(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "crypto.gor")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "59 7f 6a 54 00 10 f9 4c 15 d7 18 06 a9 9a 2c 87 10 e7 47 bd" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

// Golang stdlib: encoding package

func TestEncoding(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "encoding")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "[{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}

func TestEncodingGor(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "encoding.gor")
	response := subprocess.RunShell("", "", "gor", testpath)
	if response.StdOut != "[{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value: %d", response.ExitCode)
	}
}