package main

import (
    "io/ioutil"
	"flag"
	"fmt"
	"os"
	"./version"
	"./lzw"
)

func call(f func([]byte, []byte) ([]byte,error), inputFileName string, outputFileName string) {
	if inputFileName == outputFileName {
		fmt.Fprintf(os.Stderr, "input and output files should not coincide\n")
		os.Exit(1)
	} else {
	    if src, err := ioutil.ReadFile(inputFileName); err==nil {
	        if res, err := f(src, version.ForHeader()); err==nil {
	            if err:=ioutil.WriteFile(outputFileName, res, 0644); err==nil {
	                os.Exit(0)
        	    } else {
    	            printErrorAndExit(err)
	            }	            
	        } else {
    	        printErrorAndExit(err)	            
	        }
	    } else {
	        printErrorAndExit(err)
	    }
	}
}

func printErrorAndExit(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
    os.Exit(2)    
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
		var handler func([]byte, []byte) ([]byte, error)
		if *archiveFlag {
			handler = lzw.Encode
		} else {
		    lzw.VersionChecker = version.IsCorrect
			handler = lzw.Decode
		}
		call(handler, *inputFileName, *outputFileName)
	}
}
