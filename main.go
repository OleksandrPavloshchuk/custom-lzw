package main

import (
	"./config"
	"./header"
	"./lzw"
	"./version"
	"fmt"
	"os"
)

/**
 * This construction is written in order to implement several stages of compressing in future
 */
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
				return
			}
		}
	}
	printError(err)
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	os.Exit(2)
}

func printHeader() {
    if h, err := config.GetHeaderReader()(); err != nil {
		printError(err)
	} else {
		header.Print(&h)
	}
}

func main() {
	config.Acquire()
	switch config.GetMode() {
	case config.Version:
		version.Print()
	case config.Archive:
		call([]func(*[]byte) (*[]byte, error){lzw.Encode}, true)
	case config.Extract:
		call([]func(*[]byte) (*[]byte, error){lzw.Decode}, false)
	case config.PrintHeader:
	  printHeader()
	}
	os.Exit(0)
}
