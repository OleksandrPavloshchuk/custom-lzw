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

func call(transforms []func([]byte,uint64) ([]byte, error)) {
	src, err := config.GetReader()()
	if err == nil {
		res := src
		sourceSize := uint64(len(src))
		for _, t := range transforms {
			res, err = t(res, sourceSize)
			if err != nil {
				break
			}
		}
		if err == nil {
			err = config.GetWriter()(res)
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
		call([]func([]byte,uint64) ([]byte, error){lzw.Encode, static.Encode})
	case config.Extract:
		call([]func([]byte,uint64) ([]byte, error){static.Decode, lzw.Decode})
	case config.PrintHeader:
		if h, err := config.GetHeaderReader()(); err != nil {
			printError(err)
		} else {
			header.Print(h)
		}
	}
	os.Exit(0)
}
