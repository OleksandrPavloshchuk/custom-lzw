package main

import (
    "io"
    "io/ioutil"
	"flag"
	"fmt"
	"os"
	"./version"
	"./lzw"
)

func readStdIn(_ string) ([]byte,error) {
    r := make([]byte,0)
    b := make([]byte,1)
    for {
        _,err := os.Stdin.Read(b)
        if err==io.EOF {
            break
        }
        r = append( r, b...)
    }   
    return r,nil
}

func writeStdOut(_ string, data []byte, _ os.FileMode) error {
    _,err := os.Stdout.Write(data)
    return err
}


func call(f func([]byte, []byte) ([]byte,error), inputFileName string, outputFileName string) {
	if inputFileName == outputFileName && len(inputFileName)!=0 {
		fmt.Fprintf(os.Stderr, "input and output files should not coincide\n")
		os.Exit(1)
	} else {
	
	    var read func(string)([]byte,error)
	    if len(inputFileName)==0 {
	        read = readStdIn
	    } else {
	        read = ioutil.ReadFile
	    }
	    
	    var write func(string,[]byte,os.FileMode) error
	    if len(outputFileName)==0 {
	        write = writeStdOut
	    } else {
	        write = ioutil.WriteFile
	    }	    
	
	    if src, err := read(inputFileName); err==nil {
	        if res, err := f(src, version.ForHeader()); err==nil {
	            if err:=write(outputFileName, res, 0644); err==nil {
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
    	version.Print()
	    os.Exit(0)
	}
	if (!*archiveFlag && !*extractFlag) || (*archiveFlag && *extractFlag) {
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
