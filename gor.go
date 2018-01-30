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
	"strings"
)

const (
	version = "0.1.0"

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

//type Session struct {
//	FilePath       string
//	File           *ast.File
//	Fset           *token.FileSet
//	Types          *types.Config
//	TypeInfo       types.Info
//	ExtraFilePaths []string
//	ExtraFiles     []*ast.File
//
//	mainBody         *ast.BlockStmt
//	storedBodyLength int
//}

func main() {
	flag.Parse()

	// parse command line flags and handle them
	switch {
	case *versionShort, *versionLong:
		fmt.Printf("gore v%s\n", version)
		os.Exit(0)
	case *helpShort, *helpLong:
		fmt.Println(help)
		os.Exit(0)
	case *usageLong:
		fmt.Println(usage)
		os.Exit(0)
	}

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

	//create temp out file path for the go source code to be executed
	outName := "run_" + baseName + ".go"
	outPath := filepath.Join(dirPath, outName)

	// write the temp go source file to be executed
	err = ioutil.WriteFile(outPath, preFile[1], 0644)
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
				goSourceList = append(goSourceList, goSourceNeedle)
			}
		}
	}

	// TEMP CODE
	//src, _ := ioutil.ReadFile(outPath)
	//s := &Session{
	//	Fset: token.NewFileSet(),
	//	Types: &types.Config{
	//		Importer: importer.Default(),
	//	},
	//}
	//s.FilePath = outPath
	//
	//astf, err := parser.ParseFile(s.Fset, "", src, parser.Mode(0))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//packages := make([]string, len(astf.Imports))
	//for i, imprt := range astf.Imports {
	//	packages[i] = imprt.Path.Value
	//}
	//
	//var files []string
	//for _, packageName := range packages {
	//	packageName = strings.Replace(packageName, "\"", "", -1)  // remove the double quotes around package names
	//	goPkg, err := build.Import(packageName, ".", 0)
	//	if err != nil {
	//		var err2 error
	//		goPkg, err2 = build.ImportDir(packageName, 0)
	//		if err2 != nil {
	//			log.Fatal(err)
	//		}
	//	}
	//
	//	for _, f := range goPkg.GoFiles {
	//		files = append(files, filepath.Join(goPkg.Dir, f))
	//	}
	//}
	//
	//for _, afile := range files {
	//	fmt.Println(afile)
	//}

	// create the `go run` command for execution of the source file
	cmdArgs := []string{"run"}
	if len(args) > 1 {
		cmdArgs = append(cmdArgs, args[1:]...) // arguments to the executable excluding the path to the .gor file at slice position 0
	}
	cmdArgs = append(cmdArgs, outPath) // the go source file with the main function requested by user
	//cmdArgs = append(cmdArgs, goSourceList...) // additional source file paths in same directory

	//fmt.Println(cmdArgs)

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
