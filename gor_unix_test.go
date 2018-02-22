// +build linux darwin

package main

import (
	"path/filepath"
	"strings"
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

// Golang stdlib: html package

func TestHtmlDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "html")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value and did not expect non-zero exit code: %d", response.ExitCode)
	}
}

func TestHtmlGorDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "go_stdlib", "html.gor")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value and did not expect non-zero exit code: %d", response.ExitCode)
	}
}

// Argument handling tests

//   Test short boolean style flags

func TestArgBooleanShortDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "arg", "args")

	// test without argument for default 'false' value
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "false" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value and did not expect non-zero exit code: %d", response.ExitCode)
	}

	// test with argument for re-defined true value
	response2 := subprocess.RunShell("", "", testpath, "-t")
	if response2.StdOut != "true" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response2.StdOut)
	}
	if response2.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value and did not expect non-zero exit code: %d", response2.ExitCode)
	}
}

func TestArgBooleanShortGorDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "arg", "args.gor")

	// test without argument for default 'false' value
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "false" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value and did not expect non-zero exit code: %d", response.ExitCode)
	}

	// test with argument for re-defined true value
	response2 := subprocess.RunShell("", "", testpath, "-t")
	if response2.StdOut != "true" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response2.StdOut)
	}
	if response2.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value and did not expect non-zero exit code: %d", response2.ExitCode)
	}
}

//   Test long defintion flags

func TestArgStringLongDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "arg", "args")

	// test without argument for default 'false' value
	response := subprocess.RunShell("", "", testpath, "--ps", "--teststring", "test")
	if response.StdOut != "test" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value and did not expect non-zero exit code: %d", response.ExitCode)
	}
}

func TestArgStringLongGorDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "arg", "args.gor")

	// test without argument for default 'false' value
	response := subprocess.RunShell("", "", testpath, "--ps", "--teststring", "test")
	if response.StdOut != "test" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value and did not expect non-zero exit code: %d", response.ExitCode)
	}
}

// Package import tests

func TestPackageImportDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "pkg", "pkg")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "Hello" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value and did not expect non-zero exit code: %d", response.ExitCode)
	}
}

func TestPackageImportGorDirectOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "pkg", "pkg.gor")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "Hello" {
		t.Errorf("[FAIL] Expected response does not match received response '%s'", response.StdOut)
	}
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Returned non-zero exit code value and did not expect non-zero exit code: %d", response.ExitCode)
	}
}

// Compiler error tests

func TestCompileFailMissingImportOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "compile_errors", "missing_import")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "" {
		t.Errorf("[FAIL] Expected no standard output response and received '%s'", response.StdOut)
	}
	if len(response.StdErr) == 0 {
		t.Errorf("[FAIL] No stderr output when stderr was expected")
	}
	if response.ExitCode == 0 {
		t.Errorf("[FAIL] Returned zero exit code value and did not expect zero exit code.  Returned code: %d", response.ExitCode)
		t.Errorf("%s", response.StdErr)
	}
}

func TestCompileFailMissingImportGorOnUnix(t *testing.T) {
	testpath := filepath.Join("test", "compile_errors", "missing_import.gor")
	response := subprocess.RunShell("", "", testpath)
	if response.StdOut != "" {
		t.Errorf("[FAIL] Expected no standard output response and received '%s'", response.StdOut)
	}
	if len(response.StdErr) == 0 {
		t.Errorf("[FAIL] No stderr output when stderr was expected")
	}
	if response.ExitCode == 0 {
		t.Errorf("[FAIL] Returned zero exit code value and did not expect zero exit code.  Returned code: %d", response.ExitCode)
		t.Errorf("%s", response.StdErr)
	}
}
