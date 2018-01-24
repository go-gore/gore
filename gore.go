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
	// for _, arg := range args {
	// 	fmt.Println(arg)
	// }

	// read the whole file at once
	b, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	preFile := bytes.SplitN(b, []byte("\n"), 2)

	gorRelativeFilePath := args[0]
	absPath, _ := filepath.Abs(gorRelativeFilePath)
	dirPath := filepath.Dir(absPath)
	basePath := filepath.Base(absPath)
	baseList := strings.Split(basePath, ".")
	baseName := baseList[0]
	outName := "run_" + baseName + ".go"
	outPath := filepath.Join(dirPath, outName)

	// write the whole body at once
	err = ioutil.WriteFile(outPath, preFile[1], 0644)
	if err != nil {
		log.Fatal(err)
	}

	// only execute file removal if properly formatted with run_ prefix and .go file extension
	if strings.HasPrefix(filepath.Base(outPath), "run_") && filepath.Ext(outPath) == ".go" {
		defer os.Remove(outPath)
	}

	args[0] = outPath
	cmdArgs := []string{"run"}
	for _, x := range args {
		cmdArgs = append(cmdArgs, x)
	}

	cmd := exec.Command("go", cmdArgs...)
	stdout, stdOutErr := cmd.StdoutPipe()
	stderr, stdErrErr := cmd.StderrPipe()
	if stdOutErr != nil {
		log.Fatal(stdOutErr)
	}
	if stdErrErr != nil {
		log.Fatal(stdErrErr)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, stdout); err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stderr, stderr); err != nil {
		log.Fatal(err)
	}

}
