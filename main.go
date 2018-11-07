package main

import (
	"flag"
	"fmt"
	"os"
	"./version"
	"./lzw"
)

func call(f func(string, string, []byte) error, inputFileName string, outputFileName string) {
	if inputFileName == outputFileName {
		fmt.Printf("input and output files should not coincide\n")
		os.Exit(1)
	} else {
		if err := f(inputFileName, outputFileName, version.ForHeader()); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(2)
		} else {
			os.Exit(0)
		}
	}
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
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
    	version.Print(os.Stdout)
	    os.Exit(0)
	}
	if (!*archiveFlag && !*extractFlag) || (*archiveFlag && *extractFlag) || len(*inputFileName) == 0 || len(*outputFileName) == 0 {
		Usage()
	} else {
		var handler func(string, string, []byte) error
		if *archiveFlag {
			handler = lzw.Encode
		} else {
			handler = lzw.Decode
		}
		call(handler, *inputFileName, *outputFileName)
	}
}
