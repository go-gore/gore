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

	// create the `go run` command for execution of the source file
	args[0] = outPath
	cmdArgs := []string{"run"}
	for _, x := range args {
		cmdArgs = append(cmdArgs, x)
	}

	// Define the command for execution of the source file
	// & capture the stdout and stderr pipes to push through these streams during execution
	cmd := exec.Command("go", cmdArgs...)
	stdoutPipe, stdOutPipeErr := cmd.StdoutPipe()
	stderrPipe, stdErrPipeErr := cmd.StderrPipe()
	if stdOutPipeErr != nil {
		log.Fatal(stdOutPipeErr)
	}
	if stdErrPipeErr != nil {
		log.Fatal(stdErrPipeErr)
	}
	// execute the source code
	if startErr := cmd.Start(); startErr != nil {
		log.Fatal(startErr)
	}
	// write to stdout if data present
	if _, stdOutErr := io.Copy(os.Stdout, stdoutPipe); stdOutErr != nil {
		log.Fatal(stdOutErr)
	}
	// write to stderr if data present
	_, stdErrErr := io.Copy(os.Stderr, stderrPipe)
	if stdErrErr != nil {
		log.Fatal(stdErrErr)
	}

	returnedErr := cmd.Wait()
	// defer temp file removal to follow execution of the source
	// the logic here is probably not necessary but is intended to protect against deleting
	// something that we should not be deleting from the file system in case I am wrong.
	if strings.HasPrefix(filepath.Base(outPath), "run_") && filepath.Ext(outPath) == ".go" {
		os.Remove(outPath)
	}
	if returnedErr != nil {
		os.Exit(1) // assign a default non-zero fail code value of 1 for all failures
	}

}
