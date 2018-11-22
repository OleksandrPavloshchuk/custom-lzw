package main

import (
	"./config"
	"./header"
	"./huffman/static"
	"./lzw"
	"./version"
	"fmt"
	"os"
)

func call(transforms []func(*[]byte) (*[]byte, error), doAddHeader bool) {
	src, err := config.GetReader()()
	if err == nil {
		res := &src
		for _, t := range transforms {
			res, err = t(res)
			if err != nil {
				break
			}
		}
		if err == nil {
			if doAddHeader {
				res = header.AddHeader(res)
			}
			err = config.GetWriter()(*res)
			if err == nil {
				os.Exit(0)
			}
		}
	}
	printError(err)
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	os.Exit(2)
}

func main() {
	config.Acquire()
	switch config.GetMode() {
	case config.Version:
		version.Print()
	case config.Archive:
		call([]func(*[]byte) (*[]byte, error){lzw.Encode, static.Encode}, true)
	case config.Extract:
		call([]func(*[]byte) (*[]byte, error){static.Decode, lzw.Decode}, false)
	case config.PrintHeader:
		if h, err := config.GetHeaderReader()(); err != nil {
			printError(err)
		} else {
			header.Print(h)
		}
	}
	os.Exit(0)
}
