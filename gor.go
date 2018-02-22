// gor.go links Go source code files to the go compiler + runtime to create "executable" Go source code files
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	version = "0.4.0"

	usage = `Usage: gor (options) [args]
`

	help = `=================================================
 gor
 Copyright 2018 Christopher Simpkins
 MIT License

 Source: https://github.com/go-rillas/gor
=================================================
 Usage:
   $ gor (options) [args]

 Options:
  -h, --help           Application help
      --usage          Application usage
  -v, --version        Application version
`
)

var versionShort, versionLong, helpShort, helpLong, usageLong *bool

func init() {
	// define available command line flags
	versionShort = flag.Bool("v", false, "Application version")
	versionLong = flag.Bool("version", false, "Application version")
	helpShort = flag.Bool("h", false, "Help")
	helpLong = flag.Bool("help", false, "Help")
	usageLong = flag.Bool("usage", false, "Usage")
}

func main() {
	flag.Parse()

	// parse command line flags and handle them
	switch {
	case *versionShort, *versionLong:
		fmt.Printf("gor v%s\n", version)
		os.Exit(0)
	case *helpShort, *helpLong:
		fmt.Println(help)
		os.Exit(0)
	case *usageLong:
		fmt.Println(usage)
		os.Exit(0)
	}

	args := flag.Args()

	//////////////////////////////
	// COMPILE BINARY FROM SOURCE
	//////////////////////////////

	// read the source file at (non-flag) argument slice position 0
	inBytes, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	// remove the shebang line of the gor source file (if present)
	var outBytes []byte
	if bytes.HasPrefix(inBytes, []byte("#!")) {
		preFileList := bytes.SplitN(inBytes, []byte("\n"), 2)
		outBytes = preFileList[1]
	} else {
		outBytes = inBytes
	}

	// parse file paths
	gorRelativeFilePath := args[0]
	absPath, _ := filepath.Abs(gorRelativeFilePath)
	dirPath := filepath.Dir(absPath)
	basePath := filepath.Base(absPath)
	baseList := strings.Split(basePath, ".")
	baseName := baseList[0]

	//create temp out file path for the go source code to be executed
	outName := "run_" + baseName + ".go"
	outPath := filepath.Join(dirPath, outName)

	// write the temp go source file to be executed
	err = ioutil.WriteFile(outPath, outBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// get other go source files in the same directory as requested file
	workingDir, err := os.Open(dirPath)
	if err != nil {
		log.Fatalf("failed to open directory: %v", err)
	}
	defer workingDir.Close()

	var goSourceList []string
	testSourceList, _ := workingDir.Readdirnames(0)
	for _, goSourceNeedle := range testSourceList {
		if goSourceNeedle != filepath.Base(outPath) {
			if strings.HasSuffix(goSourceNeedle, ".go") && !strings.Contains(goSourceNeedle, "_test.go") {
				goSourceAbsPath, err := filepath.Abs(goSourceNeedle)
				if err != nil {
					log.Fatal(err)
				}
				goSourceList = append(goSourceList, goSourceAbsPath)
			}
		}
	}

	// create the executable outfile path based upon platform
	var tempRunPath string
	if runtime.GOOS == "windows" {
		tempRunPath = "run_" + baseName + ".exe"
	} else {
		tempRunPath = "run_" + baseName
	}

	// create the go build command to compile the executable
	cmdArgs := []string{"build", "-o", tempRunPath}
	cmdArgs = append(cmdArgs, outPath)         // the go source file with the main function requested by user
	cmdArgs = append(cmdArgs, goSourceList...) // additional source file paths in same directory

	// Define the go build command for compile of the source
	// & capture the stderr pipe for error reporting
	cmd := exec.Command("go", cmdArgs...)
	stderrPipe, stdErrPipeErr := cmd.StderrPipe()
	if stdErrPipeErr != nil {
		log.Fatal(stdErrPipeErr)
	}
	// execute the source code
	if startErr := cmd.Start(); startErr != nil {
		log.Fatal(startErr)
	}
	// write to stderr if data present
	_, stdErrErr := io.Copy(os.Stderr, stderrPipe)
	if stdErrErr != nil {
		log.Fatal(stdErrErr)
	}

	returnedErr := cmd.Wait()
	if returnedErr != nil {
		removeTempGoSource(outPath)
		os.Exit(1)
	}

	// remove the temporary Go source file that was created from the Go runner source file
	removeTempGoSource(outPath)


	////////////////////
	// BINARY EXECUTION
	////////////////////

	runPath, pathErr := filepath.Abs(tempRunPath)
	if pathErr != nil {
		log.Fatal(pathErr)
	}

	// define execution command
	var runCmd []string
	var executable string

	executable = runPath

	// define any command line arguments that were requested by user
	if len(args) > 1 {
		runCmd = append(runCmd, args[1:]...) // arguments to the executable excluding the path to the executable file
	}

	defer removeRunFile(runPath) // defer removal of the run file

	cmdRun := exec.Command(executable, runCmd...)

	stdoutPipeRun, stdOutPipeErrRun := cmdRun.StdoutPipe()
	stderrPipeRun, stdErrPipeErrRun := cmdRun.StderrPipe()
	if stdOutPipeErrRun != nil {
		log.Fatal(stdOutPipeErrRun)
	}
	if stdErrPipeErrRun != nil {
		log.Fatal(stdErrPipeErrRun)
	}
	// execute the source code
	if startErrRun := cmdRun.Start(); startErrRun != nil {
		log.Fatal(startErrRun)
	}
	// write to stdout if data present
	if _, stdOutErrRun := io.Copy(os.Stdout, stdoutPipeRun); stdOutErrRun != nil {
		log.Fatal(stdOutErrRun)
	}
	// write to stderr if data present
	_, stdErrErrRun := io.Copy(os.Stderr, stderrPipeRun)
	if stdErrErrRun != nil {
		log.Fatal(stdErrErrRun)
	}

	returnedErrRun := cmdRun.Wait()
	if returnedErrRun != nil {
		removeRunFile(runPath) // remove the temporary binary used to execute code before exit with status code 1
		os.Exit(1)
	}

}

func removeTempGoSource(outPath string) {
	if strings.HasPrefix(filepath.Base(outPath), "run_") && filepath.Ext(outPath) == ".go" {
		os.Remove(outPath)
	}
}

func removeRunFile(runPath string) {
	if _, err := os.Stat(runPath); err == nil {
		// temp runner file exists, let's remove it
		os.Remove(runPath)
	} else {
		fmt.Printf("%v", err)
	}
}
