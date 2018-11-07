package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func call(f func(string, string) error, inputFileName string, outputFileName string) {
	if inputFileName == outputFileName {
		fmt.Printf("input and output files should not coincide\n")
		os.Exit(1)
	} else {
		if err := f(inputFileName, outputFileName); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(2)
		} else {
			os.Exit(0)
		}
	}
}

func encode(inputFileName string, outputFileName string) error {
	src, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		return err
	}
	codeWriter := CodeWriter{}
	Encode(src, &codeWriter)
	return codeWriter.Write(outputFileName)
}

func decode(inputFileName string, outputFileName string) error {
	codeReader := CodeReader{}
	if err := codeReader.Read(inputFileName); err != nil {
		return err
	}
	result := Decode(&codeReader)
	return ioutil.WriteFile(outputFileName, result, 0644)
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func printVersion() {
	PrintVersion(os.Stdout)
	os.Exit(0)
}

func main() {
	archiveFlag := flag.Bool("a", false, "archive file")
	extractFlag := flag.Bool("e", false, "extract file")
	versionFlag := flag.Bool("v", false, "print version")
	inputFileName := flag.String("in", "", "input file name")
	outputFileName := flag.String("out", "", "output file name")
	flag.Parse()
	if !flag.Parsed() {
		Usage()
	}
	if *versionFlag {
        printVersion()
	}
	if (!*archiveFlag && !*extractFlag) || (*archiveFlag && *extractFlag) || len(*inputFileName)==0 || len(*outputFileName)==0 {
		Usage()
	} else {
		var handler func(string, string) error
		if *archiveFlag {
			handler = encode
		} else if *extractFlag {
			handler = decode
		}
		call(handler, *inputFileName, *outputFileName)
	}
}
