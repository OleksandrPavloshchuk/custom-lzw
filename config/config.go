package config

import (
	"../header"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Mode int

const (
	Version     = iota
	Archive     = iota
	Extract     = iota
	PrintHeader = iota
)

var (
	inputFile     *string
	outputFile    *string
	isVersion     *bool
	isArchive     *bool
	isExtract     *bool
	isPrintHeader *bool
)

func GetReader() func() ([]byte, error) {
	if len(*inputFile) != 0 {
		return func() ([]byte, error) {
			return ioutil.ReadFile(*inputFile)
		}
	} else {
		return func() ([]byte, error) {
			r := make([]byte, 0)
			b := make([]byte, 1)
			for {
				_, err := os.Stdin.Read(b)
				if err == io.EOF {
					break
				}
				r = append(r, b...)
			}
			return r, nil
		}
	}
}

func GetHeaderReader() func() ([]byte, error) {
	return func() ([]byte, error) {
		if f, err := os.Open(*inputFile); err != nil {
			return nil, err
		} else {
			r := make([]byte, header.GetLength())
			if _, err := io.ReadFull(f, r); err != nil {
				return nil, err
			} else {
				return r, nil
			}
		}
	}
}

func GetWriter() func([]byte) error {
	if len(*outputFile) != 0 {
		return func(data []byte) error {
			return ioutil.WriteFile(*outputFile, data, 0644)
		}
	} else {
		return func(data []byte) error {
			_, err := os.Stdout.Write(data)
			return err
		}
	}
}

func GetMode() Mode {
	if *isVersion {
		return Version
	}
	if *isPrintHeader {
		return PrintHeader
	}
	if *isArchive {
		return Archive
	} else {
		return Extract
	}
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func Acquire() {
	isArchive = flag.Bool("a", false, "archive file")
	isExtract = flag.Bool("e", false, "extract file")
	isVersion = flag.Bool("v", false, "print version")
	isPrintHeader = flag.Bool("printHeader", false, "print header of archive")
	inputFile = flag.String("in", "", "input file name")
	outputFile = flag.String("out", "", "output file name")
	flag.Parse()
	if !flag.Parsed() {
		Usage()
	}
	if !*isVersion && !*isPrintHeader && ((!*isArchive && !*isExtract) || (*isArchive && *isExtract)) {
		Usage()
	}
	if *inputFile == *outputFile && len(*inputFile) != 0 {
		fmt.Fprintf(os.Stderr, "input and output files should not coincide\n")
		os.Exit(1)
	}
	if *isPrintHeader && len(*inputFile) == 0 {
		fmt.Fprintf(os.Stderr, "no input file\n")
		os.Exit(1)
	}
}
