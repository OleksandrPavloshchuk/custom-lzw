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

	if src, err := read(); err == nil {
		if res, err := transform(src); err == nil {
			if err := write(res); err == nil {
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

func printErrorAndExit(err error) {
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
