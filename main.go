package main

import (
	"./config"
	"./lzw"
	"./version"
	"fmt"
	"os"
)

func call(transform func([]byte) ([]byte, error)) {
	read := config.GetReader()
	write := config.GetWriter()
    src, err := read()
	if err == nil {
	    var res []byte
	    res, err = transform(src)
	    if err == nil {
	        err = write(res)
	        if err == nil {
				os.Exit(0)	            
	        }
	    }
	}
    fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	os.Exit(2)	
}

func main() {
	config.Acquire()
	switch config.GetMode() {
	case config.Version:
		version.Print()
		os.Exit(0)
	case config.Archive:
		call(lzw.Encode)
	case config.Extract:
		call(lzw.Decode)
	}
}
