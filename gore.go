// gore.go links Go source code files to the go compiler + runtime to create "executable" Go source code files
package main

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()

	// read the source file at (non-flag) argument slice position 0
	b, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	// remove the shebang line of the gore source file
	preFile := bytes.SplitN(b, []byte("\n"), 2)

	// parse file paths
	gorRelativeFilePath := args[0]
	absPath, _ := filepath.Abs(gorRelativeFilePath)
	dirPath := filepath.Dir(absPath)
	basePath := filepath.Base(absPath)
	baseList := strings.Split(basePath, ".")
	baseName := baseList[0]
	// create temp out file path for the go source code to be executed
	outName := "run_" + baseName + ".go"
	outPath := filepath.Join(dirPath, outName)

	// write the temp go source file to be executed
	err = ioutil.WriteFile(outPath, preFile[1], 0644)
	if err != nil {
		log.Fatal(err)
	}

	// defer temp file removal to follow execution of the source
	// the logic here is probably not necessary but is intended to protect against deleting
	// something that we should not be deleting from the file system in case I am wrong.
	if strings.HasPrefix(filepath.Base(outPath), "run_") && filepath.Ext(outPath) == ".go" {
		defer os.Remove(outPath)
	}

	// create the `go run` command for execution of the source file
	args[0] = outPath
	cmdArgs := []string{"run"}
	for _, x := range args {
		cmdArgs = append(cmdArgs, x)
	}

	// Define the command for execution of the source file
	// & capture the stdout and stderr pipes to push through these streams during execution
	cmd := exec.Command("go", cmdArgs...)
	stdout, stdOutErr := cmd.StdoutPipe()
	stderr, stdErrErr := cmd.StderrPipe()
	if stdOutErr != nil {
		log.Fatal(stdOutErr)
	}
	if stdErrErr != nil {
		log.Fatal(stdErrErr)
	}
	// execute the source code
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	// write to stdout if data present
	if _, err := io.Copy(os.Stdout, stdout); err != nil {
		log.Fatal(err)
	}
	// write to stderr if data present
	if _, err := io.Copy(os.Stderr, stderr); err != nil {
		log.Fatal(err)
	}

}
