package main

import (
	"./config"
	"./lzw"
	"./huffman/static"
	"./version"
	"fmt"
	"os"
)

func call(transforms []func([]byte) ([]byte, error)) {
	read := config.GetReader()
	write := config.GetWriter()
	src, err := read()
	if err == nil {	
		res := src
		for _, t := range transforms {
    		res, err = t(res)
    		if err!=nil {
    		    break
    		}
		}	
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
		call([]func([]byte)([]byte, error){lzw.Encode, static.Encode})
	case config.Extract:
		call([]func([]byte)([]byte, error){static.Decode, lzw.Decode})
	}
}
