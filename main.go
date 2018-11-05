package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"errors"
)

func call(f func(string,string) error, inputFileName string, outputFileName string) {
	if inputFileName==outputFileName {
		fmt.Printf("input and output files should not coincide\n")
		os.Exit(1)
	} else {
		err:=f(inputFileName, outputFileName)
		if err!=nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(2)
		} else {
			os.Exit(0)
		}
	}
}

func printUsageMessage() {
	fmt.Printf("Usage: HomeMadeArchiver [-a|-e] <input file> <output file>\n")
}

func encode(inputFileName string, outputFileName string) error {
    src,err:=ioutil.ReadFile(inputFileName)
    if err!=nil {
        return err
    }
    codeWriter:=CodeWriter{}
    Encode( src, codeWriter )    
    return codeWriter.Write(outputFileName)
}

func decode(inputFileName string, outputFileName string) error {
    return errors.New("TODO decode")
}


func main() {
	if len(os.Args)<4 {
		printUsageMessage()
		os.Exit(1)
	} else {
		switch strings.ToUpper(os.Args[1]) {
			case "-A":
				call( encode, os.Args[2], os.Args[3])
			case "-E":
				call( decode, os.Args[2], os.Args[3])
			default:
				printUsageMessage()
		}
	}
}
